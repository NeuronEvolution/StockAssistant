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

// NewUserStockIndexRenameParams creates a new UserStockIndexRenameParams object
// with the default values initialized.
func NewUserStockIndexRenameParams() *UserStockIndexRenameParams {
	var ()
	return &UserStockIndexRenameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserStockIndexRenameParamsWithTimeout creates a new UserStockIndexRenameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserStockIndexRenameParamsWithTimeout(timeout time.Duration) *UserStockIndexRenameParams {
	var ()
	return &UserStockIndexRenameParams{

		timeout: timeout,
	}
}

// NewUserStockIndexRenameParamsWithContext creates a new UserStockIndexRenameParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserStockIndexRenameParamsWithContext(ctx context.Context) *UserStockIndexRenameParams {
	var ()
	return &UserStockIndexRenameParams{

		Context: ctx,
	}
}

// NewUserStockIndexRenameParamsWithHTTPClient creates a new UserStockIndexRenameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserStockIndexRenameParamsWithHTTPClient(client *http.Client) *UserStockIndexRenameParams {
	var ()
	return &UserStockIndexRenameParams{
		HTTPClient: client,
	}
}

/*UserStockIndexRenameParams contains all the parameters to send to the API endpoint
for the user stock index rename operation typically these are written to a http.Request
*/
type UserStockIndexRenameParams struct {

	/*NameNew
	  new name

	*/
	NameNew string
	/*NameOld
	  old name

	*/
	NameOld string
	/*UserID
	  User id

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user stock index rename params
func (o *UserStockIndexRenameParams) WithTimeout(timeout time.Duration) *UserStockIndexRenameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user stock index rename params
func (o *UserStockIndexRenameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user stock index rename params
func (o *UserStockIndexRenameParams) WithContext(ctx context.Context) *UserStockIndexRenameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user stock index rename params
func (o *UserStockIndexRenameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user stock index rename params
func (o *UserStockIndexRenameParams) WithHTTPClient(client *http.Client) *UserStockIndexRenameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user stock index rename params
func (o *UserStockIndexRenameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNameNew adds the nameNew to the user stock index rename params
func (o *UserStockIndexRenameParams) WithNameNew(nameNew string) *UserStockIndexRenameParams {
	o.SetNameNew(nameNew)
	return o
}

// SetNameNew adds the nameNew to the user stock index rename params
func (o *UserStockIndexRenameParams) SetNameNew(nameNew string) {
	o.NameNew = nameNew
}

// WithNameOld adds the nameOld to the user stock index rename params
func (o *UserStockIndexRenameParams) WithNameOld(nameOld string) *UserStockIndexRenameParams {
	o.SetNameOld(nameOld)
	return o
}

// SetNameOld adds the nameOld to the user stock index rename params
func (o *UserStockIndexRenameParams) SetNameOld(nameOld string) {
	o.NameOld = nameOld
}

// WithUserID adds the userID to the user stock index rename params
func (o *UserStockIndexRenameParams) WithUserID(userID string) *UserStockIndexRenameParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user stock index rename params
func (o *UserStockIndexRenameParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserStockIndexRenameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param nameNew
	qrNameNew := o.NameNew
	qNameNew := qrNameNew
	if qNameNew != "" {
		if err := r.SetQueryParam("nameNew", qNameNew); err != nil {
			return err
		}
	}

	// query param nameOld
	qrNameOld := o.NameOld
	qNameOld := qrNameOld
	if qNameOld != "" {
		if err := r.SetQueryParam("nameOld", qNameOld); err != nil {
			return err
		}
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