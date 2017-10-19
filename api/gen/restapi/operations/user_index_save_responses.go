// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/StockAssistant/api/gen/models"
)

// UserIndexSaveOKCode is the HTTP code returned for type UserIndexSaveOK
const UserIndexSaveOKCode int = 200

/*UserIndexSaveOK ok

swagger:response userIndexSaveOK
*/
type UserIndexSaveOK struct {

	/*
	  In: Body
	*/
	Payload *models.StockIndex `json:"body,omitempty"`
}

// NewUserIndexSaveOK creates UserIndexSaveOK with default headers values
func NewUserIndexSaveOK() *UserIndexSaveOK {
	return &UserIndexSaveOK{}
}

// WithPayload adds the payload to the user index save o k response
func (o *UserIndexSaveOK) WithPayload(payload *models.StockIndex) *UserIndexSaveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user index save o k response
func (o *UserIndexSaveOK) SetPayload(payload *models.StockIndex) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserIndexSaveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UserIndexSaveDefault Error response

swagger:response userIndexSaveDefault
*/
type UserIndexSaveDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.UserIndexSaveDefaultBody `json:"body,omitempty"`
}

// NewUserIndexSaveDefault creates UserIndexSaveDefault with default headers values
func NewUserIndexSaveDefault(code int) *UserIndexSaveDefault {
	if code <= 0 {
		code = 500
	}

	return &UserIndexSaveDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user index save default response
func (o *UserIndexSaveDefault) WithStatusCode(code int) *UserIndexSaveDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user index save default response
func (o *UserIndexSaveDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user index save default response
func (o *UserIndexSaveDefault) WithPayload(payload *models.UserIndexSaveDefaultBody) *UserIndexSaveDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user index save default response
func (o *UserIndexSaveDefault) SetPayload(payload *models.UserIndexSaveDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserIndexSaveDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
