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

// NewUserSettingListParams creates a new UserSettingListParams object
// with the default values initialized.
func NewUserSettingListParams() *UserSettingListParams {
	var ()
	return &UserSettingListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserSettingListParamsWithTimeout creates a new UserSettingListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserSettingListParamsWithTimeout(timeout time.Duration) *UserSettingListParams {
	var ()
	return &UserSettingListParams{

		timeout: timeout,
	}
}

// NewUserSettingListParamsWithContext creates a new UserSettingListParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserSettingListParamsWithContext(ctx context.Context) *UserSettingListParams {
	var ()
	return &UserSettingListParams{

		Context: ctx,
	}
}

// NewUserSettingListParamsWithHTTPClient creates a new UserSettingListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserSettingListParamsWithHTTPClient(client *http.Client) *UserSettingListParams {
	var ()
	return &UserSettingListParams{
		HTTPClient: client,
	}
}

/*UserSettingListParams contains all the parameters to send to the API endpoint
for the user setting list operation typically these are written to a http.Request
*/
type UserSettingListParams struct {

	/*UserID
	  User id

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user setting list params
func (o *UserSettingListParams) WithTimeout(timeout time.Duration) *UserSettingListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user setting list params
func (o *UserSettingListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user setting list params
func (o *UserSettingListParams) WithContext(ctx context.Context) *UserSettingListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user setting list params
func (o *UserSettingListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user setting list params
func (o *UserSettingListParams) WithHTTPClient(client *http.Client) *UserSettingListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user setting list params
func (o *UserSettingListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the user setting list params
func (o *UserSettingListParams) WithUserID(userID string) *UserSettingListParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user setting list params
func (o *UserSettingListParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserSettingListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}