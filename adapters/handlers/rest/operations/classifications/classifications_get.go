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

// Code generated by go-swagger; DO NOT EDIT.

package classifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// ClassificationsGetHandlerFunc turns a function with the right signature into a classifications get handler
type ClassificationsGetHandlerFunc func(ClassificationsGetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ClassificationsGetHandlerFunc) Handle(params ClassificationsGetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ClassificationsGetHandler interface for that can handle valid classifications get params
type ClassificationsGetHandler interface {
	Handle(ClassificationsGetParams, *models.Principal) middleware.Responder
}

// NewClassificationsGet creates a new http.Handler for the classifications get operation
func NewClassificationsGet(ctx *middleware.Context, handler ClassificationsGetHandler) *ClassificationsGet {
	return &ClassificationsGet{Context: ctx, Handler: handler}
}

/*
	ClassificationsGet swagger:route GET /classifications/{id} classifications classificationsGet

# View previously created classification

Get status, results and metadata of a previously created classification
*/
type ClassificationsGet struct {
	Context *middleware.Context
	Handler ClassificationsGetHandler
}

func (o *ClassificationsGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewClassificationsGetParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
