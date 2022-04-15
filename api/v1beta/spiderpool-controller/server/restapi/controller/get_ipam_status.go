// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetIpamStatusHandlerFunc turns a function with the right signature into a get ipam status handler
type GetIpamStatusHandlerFunc func(GetIpamStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIpamStatusHandlerFunc) Handle(params GetIpamStatusParams) middleware.Responder {
	return fn(params)
}

// GetIpamStatusHandler interface for that can handle valid get ipam status params
type GetIpamStatusHandler interface {
	Handle(GetIpamStatusParams) middleware.Responder
}

// NewGetIpamStatus creates a new http.Handler for the get ipam status operation
func NewGetIpamStatus(ctx *middleware.Context, handler GetIpamStatusHandler) *GetIpamStatus {
	return &GetIpamStatus{Context: ctx, Handler: handler}
}

/* GetIpamStatus swagger:route GET /ipam/status controller getIpamStatus

Get status

Get ipam status for spiderpool controller cli debug usage


*/
type GetIpamStatus struct {
	Context *middleware.Context
	Handler GetIpamStatusHandler
}

func (o *GetIpamStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetIpamStatusParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
