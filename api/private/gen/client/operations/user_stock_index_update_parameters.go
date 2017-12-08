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

	"github.com/NeuronEvolution/StockAssistant/api/private/gen/models"
)

// NewUserStockIndexUpdateParams creates a new UserStockIndexUpdateParams object
// with the default values initialized.
func NewUserStockIndexUpdateParams() *UserStockIndexUpdateParams {
	var ()
	return &UserStockIndexUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserStockIndexUpdateParamsWithTimeout creates a new UserStockIndexUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserStockIndexUpdateParamsWithTimeout(timeout time.Duration) *UserStockIndexUpdateParams {
	var ()
	return &UserStockIndexUpdateParams{

		timeout: timeout,
	}
}

// NewUserStockIndexUpdateParamsWithContext creates a new UserStockIndexUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserStockIndexUpdateParamsWithContext(ctx context.Context) *UserStockIndexUpdateParams {
	var ()
	return &UserStockIndexUpdateParams{

		Context: ctx,
	}
}

// NewUserStockIndexUpdateParamsWithHTTPClient creates a new UserStockIndexUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserStockIndexUpdateParamsWithHTTPClient(client *http.Client) *UserStockIndexUpdateParams {
	var ()
	return &UserStockIndexUpdateParams{
		HTTPClient: client,
	}
}

/*UserStockIndexUpdateParams contains all the parameters to send to the API endpoint
for the user stock index update operation typically these are written to a http.Request
*/
type UserStockIndexUpdateParams struct {

	/*Index
	  Index

	*/
	Index *models.UserStockIndex
	/*IndexName
	  index id

	*/
	IndexName string
	/*UserID
	  User id

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user stock index update params
func (o *UserStockIndexUpdateParams) WithTimeout(timeout time.Duration) *UserStockIndexUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user stock index update params
func (o *UserStockIndexUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user stock index update params
func (o *UserStockIndexUpdateParams) WithContext(ctx context.Context) *UserStockIndexUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user stock index update params
func (o *UserStockIndexUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user stock index update params
func (o *UserStockIndexUpdateParams) WithHTTPClient(client *http.Client) *UserStockIndexUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user stock index update params
func (o *UserStockIndexUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the user stock index update params
func (o *UserStockIndexUpdateParams) WithIndex(index *models.UserStockIndex) *UserStockIndexUpdateParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the user stock index update params
func (o *UserStockIndexUpdateParams) SetIndex(index *models.UserStockIndex) {
	o.Index = index
}

// WithIndexName adds the indexName to the user stock index update params
func (o *UserStockIndexUpdateParams) WithIndexName(indexName string) *UserStockIndexUpdateParams {
	o.SetIndexName(indexName)
	return o
}

// SetIndexName adds the indexName to the user stock index update params
func (o *UserStockIndexUpdateParams) SetIndexName(indexName string) {
	o.IndexName = indexName
}

// WithUserID adds the userID to the user stock index update params
func (o *UserStockIndexUpdateParams) WithUserID(userID string) *UserStockIndexUpdateParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user stock index update params
func (o *UserStockIndexUpdateParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserStockIndexUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Index != nil {
		if err := r.SetBodyParam(o.Index); err != nil {
			return err
		}
	}

	// path param indexName
	if err := r.SetPathParam("indexName", o.IndexName); err != nil {
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