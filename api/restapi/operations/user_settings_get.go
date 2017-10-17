// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UserSettingsGetHandlerFunc turns a function with the right signature into a user settings get handler
type UserSettingsGetHandlerFunc func(UserSettingsGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserSettingsGetHandlerFunc) Handle(params UserSettingsGetParams) middleware.Responder {
	return fn(params)
}

// UserSettingsGetHandler interface for that can handle valid user settings get params
type UserSettingsGetHandler interface {
	Handle(UserSettingsGetParams) middleware.Responder
}

// NewUserSettingsGet creates a new http.Handler for the user settings get operation
func NewUserSettingsGet(ctx *middleware.Context, handler UserSettingsGetHandler) *UserSettingsGet {
	return &UserSettingsGet{Context: ctx, Handler: handler}
}

/*UserSettingsGet swagger:route GET /{userId}/settings/{configKey} userSettingsGet

UserSettingsGet user settings get API

*/
type UserSettingsGet struct {
	Context *middleware.Context
	Handler UserSettingsGetHandler
}

func (o *UserSettingsGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserSettingsGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}