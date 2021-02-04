// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSpoeConfigurationVersionHandlerFunc turns a function with the right signature into a get spoe configuration version handler
type GetSpoeConfigurationVersionHandlerFunc func(GetSpoeConfigurationVersionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSpoeConfigurationVersionHandlerFunc) Handle(params GetSpoeConfigurationVersionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSpoeConfigurationVersionHandler interface for that can handle valid get spoe configuration version params
type GetSpoeConfigurationVersionHandler interface {
	Handle(GetSpoeConfigurationVersionParams, interface{}) middleware.Responder
}

// NewGetSpoeConfigurationVersion creates a new http.Handler for the get spoe configuration version operation
func NewGetSpoeConfigurationVersion(ctx *middleware.Context, handler GetSpoeConfigurationVersionHandler) *GetSpoeConfigurationVersion {
	return &GetSpoeConfigurationVersion{Context: ctx, Handler: handler}
}

/*GetSpoeConfigurationVersion swagger:route GET /services/haproxy/spoe/version Spoe getSpoeConfigurationVersion

Return a SPOE configuration version

Returns SPOE configuration version.

*/
type GetSpoeConfigurationVersion struct {
	Context *middleware.Context
	Handler GetSpoeConfigurationVersionHandler
}

func (o *GetSpoeConfigurationVersion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSpoeConfigurationVersionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}