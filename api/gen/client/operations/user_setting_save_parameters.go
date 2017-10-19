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

	"github.com/NeuronEvolution/StockAssistant/api/gen/models"
)

// NewUserSettingSaveParams creates a new UserSettingSaveParams object
// with the default values initialized.
func NewUserSettingSaveParams() *UserSettingSaveParams {
	var ()
	return &UserSettingSaveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserSettingSaveParamsWithTimeout creates a new UserSettingSaveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserSettingSaveParamsWithTimeout(timeout time.Duration) *UserSettingSaveParams {
	var ()
	return &UserSettingSaveParams{

		timeout: timeout,
	}
}

// NewUserSettingSaveParamsWithContext creates a new UserSettingSaveParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserSettingSaveParamsWithContext(ctx context.Context) *UserSettingSaveParams {
	var ()
	return &UserSettingSaveParams{

		Context: ctx,
	}
}

// NewUserSettingSaveParamsWithHTTPClient creates a new UserSettingSaveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserSettingSaveParamsWithHTTPClient(client *http.Client) *UserSettingSaveParams {
	var ()
	return &UserSettingSaveParams{
		HTTPClient: client,
	}
}

/*UserSettingSaveParams contains all the parameters to send to the API endpoint
for the user setting save operation typically these are written to a http.Request
*/
type UserSettingSaveParams struct {

	/*Setting
	  setting

	*/
	Setting *models.Setting
	/*UserID
	  User id

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user setting save params
func (o *UserSettingSaveParams) WithTimeout(timeout time.Duration) *UserSettingSaveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user setting save params
func (o *UserSettingSaveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user setting save params
func (o *UserSettingSaveParams) WithContext(ctx context.Context) *UserSettingSaveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user setting save params
func (o *UserSettingSaveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user setting save params
func (o *UserSettingSaveParams) WithHTTPClient(client *http.Client) *UserSettingSaveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user setting save params
func (o *UserSettingSaveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSetting adds the setting to the user setting save params
func (o *UserSettingSaveParams) WithSetting(setting *models.Setting) *UserSettingSaveParams {
	o.SetSetting(setting)
	return o
}

// SetSetting adds the setting to the user setting save params
func (o *UserSettingSaveParams) SetSetting(setting *models.Setting) {
	o.Setting = setting
}

// WithUserID adds the userID to the user setting save params
func (o *UserSettingSaveParams) WithUserID(userID string) *UserSettingSaveParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user setting save params
func (o *UserSettingSaveParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserSettingSaveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Setting == nil {
		o.Setting = new(models.Setting)
	}

	if err := r.SetBodyParam(o.Setting); err != nil {
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