// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/StockAssistant/api/gen/models"
)

// UserIndexDeleteOKCode is the HTTP code returned for type UserIndexDeleteOK
const UserIndexDeleteOKCode int = 200

/*UserIndexDeleteOK ok

swagger:response userIndexDeleteOK
*/
type UserIndexDeleteOK struct {
}

// NewUserIndexDeleteOK creates UserIndexDeleteOK with default headers values
func NewUserIndexDeleteOK() *UserIndexDeleteOK {
	return &UserIndexDeleteOK{}
}

// WriteResponse to the client
func (o *UserIndexDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*UserIndexDeleteDefault Error response

swagger:response userIndexDeleteDefault
*/
type UserIndexDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.UserIndexDeleteDefaultBody `json:"body,omitempty"`
}

// NewUserIndexDeleteDefault creates UserIndexDeleteDefault with default headers values
func NewUserIndexDeleteDefault(code int) *UserIndexDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &UserIndexDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user index delete default response
func (o *UserIndexDeleteDefault) WithStatusCode(code int) *UserIndexDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user index delete default response
func (o *UserIndexDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user index delete default response
func (o *UserIndexDeleteDefault) WithPayload(payload *models.UserIndexDeleteDefaultBody) *UserIndexDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user index delete default response
func (o *UserIndexDeleteDefault) SetPayload(payload *models.UserIndexDeleteDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserIndexDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
