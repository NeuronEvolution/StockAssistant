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

// NewOauthJumpParams creates a new OauthJumpParams object
// with the default values initialized.
func NewOauthJumpParams() *OauthJumpParams {
	var ()
	return &OauthJumpParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOauthJumpParamsWithTimeout creates a new OauthJumpParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOauthJumpParamsWithTimeout(timeout time.Duration) *OauthJumpParams {
	var ()
	return &OauthJumpParams{

		timeout: timeout,
	}
}

// NewOauthJumpParamsWithContext creates a new OauthJumpParams object
// with the default values initialized, and the ability to set a context for a request
func NewOauthJumpParamsWithContext(ctx context.Context) *OauthJumpParams {
	var ()
	return &OauthJumpParams{

		Context: ctx,
	}
}

// NewOauthJumpParamsWithHTTPClient creates a new OauthJumpParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOauthJumpParamsWithHTTPClient(client *http.Client) *OauthJumpParams {
	var ()
	return &OauthJumpParams{
		HTTPClient: client,
	}
}

/*OauthJumpParams contains all the parameters to send to the API endpoint
for the oauth jump operation typically these are written to a http.Request
*/
type OauthJumpParams struct {

	/*AuthorizationCode*/
	AuthorizationCode string
	/*State*/
	State string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the oauth jump params
func (o *OauthJumpParams) WithTimeout(timeout time.Duration) *OauthJumpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oauth jump params
func (o *OauthJumpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oauth jump params
func (o *OauthJumpParams) WithContext(ctx context.Context) *OauthJumpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oauth jump params
func (o *OauthJumpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oauth jump params
func (o *OauthJumpParams) WithHTTPClient(client *http.Client) *OauthJumpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oauth jump params
func (o *OauthJumpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorizationCode adds the authorizationCode to the oauth jump params
func (o *OauthJumpParams) WithAuthorizationCode(authorizationCode string) *OauthJumpParams {
	o.SetAuthorizationCode(authorizationCode)
	return o
}

// SetAuthorizationCode adds the authorizationCode to the oauth jump params
func (o *OauthJumpParams) SetAuthorizationCode(authorizationCode string) {
	o.AuthorizationCode = authorizationCode
}

// WithState adds the state to the oauth jump params
func (o *OauthJumpParams) WithState(state string) *OauthJumpParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the oauth jump params
func (o *OauthJumpParams) SetState(state string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *OauthJumpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param authorizationCode
	qrAuthorizationCode := o.AuthorizationCode
	qAuthorizationCode := qrAuthorizationCode
	if qAuthorizationCode != "" {
		if err := r.SetQueryParam("authorizationCode", qAuthorizationCode); err != nil {
			return err
		}
	}

	// query param state
	qrState := o.State
	qState := qrState
	if qState != "" {
		if err := r.SetQueryParam("state", qState); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}