// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// OauthJumpHandlerFunc turns a function with the right signature into a oauth jump handler
type OauthJumpHandlerFunc func(OauthJumpParams) middleware.Responder

// Handle executing the request and returning a response
func (fn OauthJumpHandlerFunc) Handle(params OauthJumpParams) middleware.Responder {
	return fn(params)
}

// OauthJumpHandler interface for that can handle valid oauth jump params
type OauthJumpHandler interface {
	Handle(OauthJumpParams) middleware.Responder
}

// NewOauthJump creates a new http.Handler for the oauth jump operation
func NewOauthJump(ctx *middleware.Context, handler OauthJumpHandler) *OauthJump {
	return &OauthJump{Context: ctx, Handler: handler}
}

/*OauthJump swagger:route POST /oauthJump oauthJump

OauthJump oauth jump API

*/
type OauthJump struct {
	Context *middleware.Context
	Handler OauthJumpHandler
}

func (o *OauthJump) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewOauthJumpParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}