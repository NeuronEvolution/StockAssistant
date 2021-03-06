// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUserStockIndexRenameParams creates a new UserStockIndexRenameParams object
// with the default values initialized.
func NewUserStockIndexRenameParams() UserStockIndexRenameParams {
	var ()
	return UserStockIndexRenameParams{}
}

// UserStockIndexRenameParams contains all the bound params for the user stock index rename operation
// typically these are obtained from a http.Request
//
// swagger:parameters UserStockIndexRename
type UserStockIndexRenameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*new name
	  Required: true
	  In: query
	*/
	NameNew string
	/*old name
	  Required: true
	  In: query
	*/
	NameOld string
	/*User id
	  Required: true
	  In: path
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UserStockIndexRenameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qNameNew, qhkNameNew, _ := qs.GetOK("nameNew")
	if err := o.bindNameNew(qNameNew, qhkNameNew, route.Formats); err != nil {
		res = append(res, err)
	}

	qNameOld, qhkNameOld, _ := qs.GetOK("nameOld")
	if err := o.bindNameOld(qNameOld, qhkNameOld, route.Formats); err != nil {
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

func (o *UserStockIndexRenameParams) bindNameNew(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("nameNew", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("nameNew", "query", raw); err != nil {
		return err
	}

	o.NameNew = raw

	return nil
}

func (o *UserStockIndexRenameParams) bindNameOld(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("nameOld", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("nameOld", "query", raw); err != nil {
		return err
	}

	o.NameOld = raw

	return nil
}

func (o *UserStockIndexRenameParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.UserID = raw

	return nil
}
