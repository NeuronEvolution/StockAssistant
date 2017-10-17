// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UserIndexDeleteHandlerFunc turns a function with the right signature into a user index delete handler
type UserIndexDeleteHandlerFunc func(UserIndexDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserIndexDeleteHandlerFunc) Handle(params UserIndexDeleteParams) middleware.Responder {
	return fn(params)
}

// UserIndexDeleteHandler interface for that can handle valid user index delete params
type UserIndexDeleteHandler interface {
	Handle(UserIndexDeleteParams) middleware.Responder
}

// NewUserIndexDelete creates a new http.Handler for the user index delete operation
func NewUserIndexDelete(ctx *middleware.Context, handler UserIndexDeleteHandler) *UserIndexDelete {
	return &UserIndexDelete{Context: ctx, Handler: handler}
}

/*UserIndexDelete swagger:route DELETE /{userId}/indices/{indexId} userIndexDelete

UserIndexDelete user index delete API

*/
type UserIndexDelete struct {
	Context *middleware.Context
	Handler UserIndexDeleteHandler
}

func (o *UserIndexDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserIndexDeleteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
