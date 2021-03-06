// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/StockAssistant/api/private/gen/models"
)

// UserStockIndexRenameOKCode is the HTTP code returned for type UserStockIndexRenameOK
const UserStockIndexRenameOKCode int = 200

/*UserStockIndexRenameOK ok

swagger:response userStockIndexRenameOK
*/
type UserStockIndexRenameOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserStockIndex `json:"body,omitempty"`
}

// NewUserStockIndexRenameOK creates UserStockIndexRenameOK with default headers values
func NewUserStockIndexRenameOK() *UserStockIndexRenameOK {
	return &UserStockIndexRenameOK{}
}

// WithPayload adds the payload to the user stock index rename o k response
func (o *UserStockIndexRenameOK) WithPayload(payload *models.UserStockIndex) *UserStockIndexRenameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user stock index rename o k response
func (o *UserStockIndexRenameOK) SetPayload(payload *models.UserStockIndex) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserStockIndexRenameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
