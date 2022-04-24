// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRuntimeQuickscanConfigHandlerFunc turns a function with the right signature into a get runtime quickscan config handler
type GetRuntimeQuickscanConfigHandlerFunc func(GetRuntimeQuickscanConfigParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeQuickscanConfigHandlerFunc) Handle(params GetRuntimeQuickscanConfigParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeQuickscanConfigHandler interface for that can handle valid get runtime quickscan config params
type GetRuntimeQuickscanConfigHandler interface {
	Handle(GetRuntimeQuickscanConfigParams) middleware.Responder
}

// NewGetRuntimeQuickscanConfig creates a new http.Handler for the get runtime quickscan config operation
func NewGetRuntimeQuickscanConfig(ctx *middleware.Context, handler GetRuntimeQuickscanConfigHandler) *GetRuntimeQuickscanConfig {
	return &GetRuntimeQuickscanConfig{Context: ctx, Handler: handler}
}

/* GetRuntimeQuickscanConfig swagger:route GET /runtime/quickscan/config getRuntimeQuickscanConfig

Get runtime quick scan configuration

*/
type GetRuntimeQuickscanConfig struct {
	Context *middleware.Context
	Handler GetRuntimeQuickscanConfigHandler
}

func (o *GetRuntimeQuickscanConfig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRuntimeQuickscanConfigParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
