// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUserIndexGetParams creates a new UserIndexGetParams object
// with the default values initialized.
func NewUserIndexGetParams() UserIndexGetParams {
	var ()
	return UserIndexGetParams{}
}

// UserIndexGetParams contains all the bound params for the user index get operation
// typically these are obtained from a http.Request
//
// swagger:parameters UserIndexGet
type UserIndexGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*index id
	  Required: true
	  In: path
	*/
	IndexID string
	/*User id
	  Required: true
	  In: path
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UserIndexGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rIndexID, rhkIndexID, _ := route.Params.GetOK("indexId")
	if err := o.bindIndexID(rIndexID, rhkIndexID, route.Formats); err != nil {
		res = append(res, err)
	}

	rUserID, rhkUserID, _ := route.Params.GetOK("userId")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserIndexGetParams) bindIndexID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.IndexID = raw

	return nil
}

func (o *UserIndexGetParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.UserID = raw

	return nil
}
