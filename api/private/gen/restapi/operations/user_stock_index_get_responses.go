// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/StockAssistant/api/private/gen/models"
)

// UserStockIndexGetOKCode is the HTTP code returned for type UserStockIndexGetOK
const UserStockIndexGetOKCode int = 200

/*UserStockIndexGetOK ok

swagger:response userStockIndexGetOK
*/
type UserStockIndexGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserStockIndex `json:"body,omitempty"`
}

// NewUserStockIndexGetOK creates UserStockIndexGetOK with default headers values
func NewUserStockIndexGetOK() *UserStockIndexGetOK {
	return &UserStockIndexGetOK{}
}

// WithPayload adds the payload to the user stock index get o k response
func (o *UserStockIndexGetOK) WithPayload(payload *models.UserStockIndex) *UserStockIndexGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user stock index get o k response
func (o *UserStockIndexGetOK) SetPayload(payload *models.UserStockIndex) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserStockIndexGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
