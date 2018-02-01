// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/StockAssistant/api/private/gen/models"
)

// StockGetOKCode is the HTTP code returned for type StockGetOK
const StockGetOKCode int = 200

/*StockGetOK ok

swagger:response stockGetOK
*/
type StockGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Stock `json:"body,omitempty"`
}

// NewStockGetOK creates StockGetOK with default headers values
func NewStockGetOK() *StockGetOK {
	return &StockGetOK{}
}

// WithPayload adds the payload to the stock get o k response
func (o *StockGetOK) WithPayload(payload *models.Stock) *StockGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stock get o k response
func (o *StockGetOK) SetPayload(payload *models.Stock) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StockGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
