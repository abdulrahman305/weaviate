//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package rest

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-openapi/runtime/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"

	"github.com/weaviate/weaviate/adapters/handlers/rest/raft"
	"github.com/weaviate/weaviate/adapters/handlers/rest/state"
	"github.com/weaviate/weaviate/adapters/handlers/rest/swagger_middleware"
	"github.com/weaviate/weaviate/usecases/config"
	"github.com/weaviate/weaviate/usecases/modules"
	"github.com/weaviate/weaviate/usecases/monitoring"
)

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
//
// we are setting the middlewares from within configureAPI, as we need access
// to some resources which are not exposed
func makeSetupMiddlewares(appState *state.State) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.String() == "/v1/.well-known/openid-configuration" || r.URL.String() == "/v1" {
				handler.ServeHTTP(w, r)
				return
			}
			appState.AnonymousAccess.Middleware(handler).ServeHTTP(w, r)
		})
	}
}

func addHandleRoot(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/" {
			w.Header().Add("Location", "/v1")
			w.WriteHeader(http.StatusMovedPermanently)
			w.Write([]byte(`{"links":{"href":"/v1","name":"api v1","documentationHref":` +
				`"https://weaviate.io/developers/weaviate/current/"}}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// clusterv1Regexp is used to intercept requests and redirect them to a dedicated http server independent of swagger
var clusterv1Regexp = regexp.MustCompile("/v1/cluster/*")

// addClusterHandlerMiddleware will inject a middleware that will catch all requests matching clusterv1Regexp.
// If the request match, it will route it to a dedicated http.Handler and skip the next middleware.
// If the request doesn't match, it will continue to the next handler.
func addClusterHandlerMiddleware(next http.Handler, appState *state.State) http.Handler {
	// Instantiate the router outside the returned lambda to avoid re-allocating everytime a new request comes in
	raftRouter := raft.ClusterRouter(appState.SchemaManager.Handler)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/v1/cluster/statistics":
			next.ServeHTTP(w, r)
		case clusterv1Regexp.MatchString(r.URL.Path):
			raftRouter.ServeHTTP(w, r)
		default:
			next.ServeHTTP(w, r)
		}
	})
}

func makeAddModuleHandlers(modules *modules.Provider) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		mux := http.NewServeMux()

		for _, mod := range modules.GetAllWithHTTPHandlers() {
			prefix := fmt.Sprintf("/v1/modules/%s", mod.Name())
			mux.Handle(fmt.Sprintf("%s/", prefix),
				http.StripPrefix(prefix, mod.RootHandler()))
		}

		prefix := "/v1/modules"
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if url := r.URL.String(); len(url) > len(prefix) && url[:len(prefix)] == prefix {
				mux.ServeHTTP(w, r)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
// Contains "x-api-key", "x-api-token" for legacy reasons, older interfaces might need these headers.
func makeSetupGlobalMiddleware(appState *state.State, context *middleware.Context) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		handleCORS := cors.New(cors.Options{
			OptionsPassthrough: true,
			AllowedMethods:     strings.Split(appState.ServerConfig.Config.CORS.AllowMethods, ","),
			AllowedHeaders:     strings.Split(appState.ServerConfig.Config.CORS.AllowHeaders, ","),
			AllowedOrigins:     strings.Split(appState.ServerConfig.Config.CORS.AllowOrigin, ","),
		}).Handler
		handler = handleCORS(handler)
		handler = swagger_middleware.AddMiddleware([]byte(SwaggerJSON), handler)
		handler = makeAddLogging(appState.Logger)(handler)
		if appState.ServerConfig.Config.Monitoring.Enabled {
			handler = makeAddMonitoring(appState.Metrics)(handler)
		}
		handler = addPreflight(handler, appState.ServerConfig.Config.CORS)
		handler = addLiveAndReadyness(appState, handler)
		handler = addHandleRoot(handler)
		handler = makeAddModuleHandlers(appState.Modules)(handler)
		handler = addInjectHeadersIntoContext(handler)
		handler = makeCatchPanics(appState.Logger, newPanicsRequestsTotal(appState.Metrics, appState.Logger))(handler)
		handler = addSourceIpToContext(handler)
		if appState.ServerConfig.Config.Monitoring.Enabled {
			handler = monitoring.InstrumentHTTP(
				handler,
				staticRoute(context),
				appState.HTTPServerMetrics.InflightRequests,
				appState.HTTPServerMetrics.RequestDuration,
				appState.HTTPServerMetrics.RequestBodySize,
				appState.HTTPServerMetrics.ResponseBodySize,
			)
		}
		// Must be the last middleware as it might skip the next handler
		handler = addClusterHandlerMiddleware(handler, appState)
		if appState.ServerConfig.Config.Sentry.Enabled {
			handler = addSentryHandler(handler)
		}

		return handler
	}
}

// staticRoute is used to convert routes in our main http server into static routes
// by removing all the dynamic variables in the route. Useful for instrumentation
// where "route cardinality" matters.

// Example:
// `/schema/Movies/properties` -> `/schema/{className}`
func staticRoute(context *middleware.Context) monitoring.StaticRouteLabel {
	return func(r *http.Request) (*http.Request, string) {
		route := r.URL.String()
		req := r

		matched, rr, ok := context.RouteInfo(r)
		if ok {
			// convert dynamic route to static route.
			// `/api/v1/schema/Question/tenant1` -> `/api/v1/schema/{class}/{tenant}`
			route = matched.PathPattern
			req = rr
		}
		return req, route
	}
}

func addSentryHandler(next http.Handler) http.Handler {
	return sentryhttp.New(sentryhttp.Options{}).Handle(next)
}

func makeAddLogging(logger logrus.FieldLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.
				WithField("action", "restapi_request").
				WithField("method", r.Method).
				WithField("url", r.URL).
				Debug("received HTTP request")
			next.ServeHTTP(w, r)
		})
	}
}

func makeAddMonitoring(metrics *monitoring.PrometheusMetrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			before := time.Now()
			method := r.Method
			path := r.URL.Path
			next.ServeHTTP(w, r)

			if strings.HasPrefix(path, "/v1/batch/objects") && method == http.MethodPost {
				metrics.BatchTime.With(prometheus.Labels{
					"operation":  "total_api_level",
					"class_name": "n/a",
					"shard_name": "n/a",
				}).
					Observe(float64(time.Since(before) / time.Millisecond))

				metrics.BatchSizeBytes.WithLabelValues("rest").Observe(float64(r.ContentLength))
			}
		})
	}
}

func addPreflight(next http.Handler, cfg config.CORS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", cfg.AllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", cfg.AllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", cfg.AllowHeaders)

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func addInjectHeadersIntoContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		changed := false
		for k, v := range r.Header {
			if strings.HasPrefix(k, "X-") || k == "Authorization" {
				ctx = context.WithValue(ctx, k, v)
				changed = true
			}
		}

		if changed {
			next.ServeHTTP(w, r.Clone(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func addLiveAndReadyness(state *state.State, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/v1/.well-known/live" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.URL.String() == "/v1/.well-known/ready" {
			code := http.StatusOK
			// if this node is in maintenance mode, we want to return live but not ready
			// so that kubernetes will allow this pod to run but not send traffic to it
			if state.Cluster.MaintenanceModeEnabledForLocalhost() {
				code = http.StatusServiceUnavailable
			} else if !state.ClusterService.Ready() || state.Cluster.ClusterHealthScore() != 0 {
				code = http.StatusServiceUnavailable
			} else if state.Modules != nil {
				_, err := state.Modules.GetMeta()
				if err != nil {
					code = http.StatusServiceUnavailable
				}
			}
			w.WriteHeader(code)
			return
		}

		next.ServeHTTP(w, r)
	})
}
