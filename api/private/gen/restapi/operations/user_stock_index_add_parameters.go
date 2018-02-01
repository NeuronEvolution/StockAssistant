// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/NeuronEvolution/StockAssistant/api/private/gen/models"
)

// NewUserStockIndexAddParams creates a new UserStockIndexAddParams object
// with the default values initialized.
func NewUserStockIndexAddParams() UserStockIndexAddParams {
	var ()
	return UserStockIndexAddParams{}
}

// UserStockIndexAddParams contains all the bound params for the user stock index add operation
// typically these are obtained from a http.Request
//
// swagger:parameters UserStockIndexAdd
type UserStockIndexAddParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Index
	  Required: true
	  In: body
	*/
	Index *models.UserStockIndex
	/*User id
	  Required: true
	  In: path
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UserStockIndexAddParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UserStockIndex
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("index", "body"))
			} else {
				res = append(res, errors.NewParseError("index", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Index = &body
			}
		}

	} else {
		res = append(res, errors.Required("index", "body"))
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

func (o *UserStockIndexAddParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.UserID = raw

	return nil
}
