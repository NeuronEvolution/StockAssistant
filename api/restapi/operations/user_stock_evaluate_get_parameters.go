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

// NewUserStockEvaluateGetParams creates a new UserStockEvaluateGetParams object
// with the default values initialized.
func NewUserStockEvaluateGetParams() UserStockEvaluateGetParams {
	var ()
	return UserStockEvaluateGetParams{}
}

// UserStockEvaluateGetParams contains all the bound params for the user stock evaluate get operation
// typically these are obtained from a http.Request
//
// swagger:parameters UserStockEvaluateGet
type UserStockEvaluateGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*stock id
	  Required: true
	  In: path
	*/
	StockID string
	/*User id
	  Required: true
	  In: path
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UserStockEvaluateGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rStockID, rhkStockID, _ := route.Params.GetOK("stockId")
	if err := o.bindStockID(rStockID, rhkStockID, route.Formats); err != nil {
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

func (o *UserStockEvaluateGetParams) bindStockID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.StockID = raw

	return nil
}

func (o *UserStockEvaluateGetParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.UserID = raw

	return nil
}
