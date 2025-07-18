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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// ListAllUsersHandlerFunc turns a function with the right signature into a list all users handler
type ListAllUsersHandlerFunc func(ListAllUsersParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllUsersHandlerFunc) Handle(params ListAllUsersParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ListAllUsersHandler interface for that can handle valid list all users params
type ListAllUsersHandler interface {
	Handle(ListAllUsersParams, *models.Principal) middleware.Responder
}

// NewListAllUsers creates a new http.Handler for the list all users operation
func NewListAllUsers(ctx *middleware.Context, handler ListAllUsersHandler) *ListAllUsers {
	return &ListAllUsers{Context: ctx, Handler: handler}
}

/*
	ListAllUsers swagger:route GET /users/db users listAllUsers

list all db users
*/
type ListAllUsers struct {
	Context *middleware.Context
	Handler ListAllUsersHandler
}

func (o *ListAllUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListAllUsersParams()
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
