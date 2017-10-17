// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UserSettingsUpdateHandlerFunc turns a function with the right signature into a user settings update handler
type UserSettingsUpdateHandlerFunc func(UserSettingsUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserSettingsUpdateHandlerFunc) Handle(params UserSettingsUpdateParams) middleware.Responder {
	return fn(params)
}

// UserSettingsUpdateHandler interface for that can handle valid user settings update params
type UserSettingsUpdateHandler interface {
	Handle(UserSettingsUpdateParams) middleware.Responder
}

// NewUserSettingsUpdate creates a new http.Handler for the user settings update operation
func NewUserSettingsUpdate(ctx *middleware.Context, handler UserSettingsUpdateHandler) *UserSettingsUpdate {
	return &UserSettingsUpdate{Context: ctx, Handler: handler}
}

/*UserSettingsUpdate swagger:route PATCH /{userId}/settings/{configKey} userSettingsUpdate

UserSettingsUpdate user settings update API

*/
type UserSettingsUpdate struct {
	Context *middleware.Context
	Handler UserSettingsUpdateHandler
}

func (o *UserSettingsUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserSettingsUpdateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}