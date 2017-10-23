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

// NewUserIndexDeleteParams creates a new UserIndexDeleteParams object
// with the default values initialized.
func NewUserIndexDeleteParams() *UserIndexDeleteParams {
	var ()
	return &UserIndexDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserIndexDeleteParamsWithTimeout creates a new UserIndexDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserIndexDeleteParamsWithTimeout(timeout time.Duration) *UserIndexDeleteParams {
	var ()
	return &UserIndexDeleteParams{

		timeout: timeout,
	}
}

// NewUserIndexDeleteParamsWithContext creates a new UserIndexDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserIndexDeleteParamsWithContext(ctx context.Context) *UserIndexDeleteParams {
	var ()
	return &UserIndexDeleteParams{

		Context: ctx,
	}
}

// NewUserIndexDeleteParamsWithHTTPClient creates a new UserIndexDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserIndexDeleteParamsWithHTTPClient(client *http.Client) *UserIndexDeleteParams {
	var ()
	return &UserIndexDeleteParams{
		HTTPClient: client,
	}
}

/*UserIndexDeleteParams contains all the parameters to send to the API endpoint
for the user index delete operation typically these are written to a http.Request
*/
type UserIndexDeleteParams struct {

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

// WithTimeout adds the timeout to the user index delete params
func (o *UserIndexDeleteParams) WithTimeout(timeout time.Duration) *UserIndexDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user index delete params
func (o *UserIndexDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user index delete params
func (o *UserIndexDeleteParams) WithContext(ctx context.Context) *UserIndexDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user index delete params
func (o *UserIndexDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user index delete params
func (o *UserIndexDeleteParams) WithHTTPClient(client *http.Client) *UserIndexDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user index delete params
func (o *UserIndexDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndexID adds the indexID to the user index delete params
func (o *UserIndexDeleteParams) WithIndexID(indexID string) *UserIndexDeleteParams {
	o.SetIndexID(indexID)
	return o
}

// SetIndexID adds the indexId to the user index delete params
func (o *UserIndexDeleteParams) SetIndexID(indexID string) {
	o.IndexID = indexID
}

// WithUserID adds the userID to the user index delete params
func (o *UserIndexDeleteParams) WithUserID(userID string) *UserIndexDeleteParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user index delete params
func (o *UserIndexDeleteParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserIndexDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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