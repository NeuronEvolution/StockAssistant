// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUserIndexGetParams creates a new UserIndexGetParams object
// with the default values initialized.
func NewUserIndexGetParams() *UserIndexGetParams {
	var ()
	return &UserIndexGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserIndexGetParamsWithTimeout creates a new UserIndexGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserIndexGetParamsWithTimeout(timeout time.Duration) *UserIndexGetParams {
	var ()
	return &UserIndexGetParams{

		timeout: timeout,
	}
}

// NewUserIndexGetParamsWithContext creates a new UserIndexGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserIndexGetParamsWithContext(ctx context.Context) *UserIndexGetParams {
	var ()
	return &UserIndexGetParams{

		Context: ctx,
	}
}

// NewUserIndexGetParamsWithHTTPClient creates a new UserIndexGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserIndexGetParamsWithHTTPClient(client *http.Client) *UserIndexGetParams {
	var ()
	return &UserIndexGetParams{
		HTTPClient: client,
	}
}

/*UserIndexGetParams contains all the parameters to send to the API endpoint
for the user index get operation typically these are written to a http.Request
*/
type UserIndexGetParams struct {

	/*IndexID
	  index id

	*/
	IndexID string
	/*UserID
	  User id

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user index get params
func (o *UserIndexGetParams) WithTimeout(timeout time.Duration) *UserIndexGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user index get params
func (o *UserIndexGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user index get params
func (o *UserIndexGetParams) WithContext(ctx context.Context) *UserIndexGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user index get params
func (o *UserIndexGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user index get params
func (o *UserIndexGetParams) WithHTTPClient(client *http.Client) *UserIndexGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user index get params
func (o *UserIndexGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndexID adds the indexID to the user index get params
func (o *UserIndexGetParams) WithIndexID(indexID string) *UserIndexGetParams {
	o.SetIndexID(indexID)
	return o
}

// SetIndexID adds the indexId to the user index get params
func (o *UserIndexGetParams) SetIndexID(indexID string) {
	o.IndexID = indexID
}

// WithUserID adds the userID to the user index get params
func (o *UserIndexGetParams) WithUserID(userID string) *UserIndexGetParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user index get params
func (o *UserIndexGetParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserIndexGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param indexId
	if err := r.SetPathParam("indexId", o.IndexID); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
