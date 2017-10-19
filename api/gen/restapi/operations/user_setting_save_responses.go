// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/StockAssistant/api/gen/models"
)

// UserSettingSaveOKCode is the HTTP code returned for type UserSettingSaveOK
const UserSettingSaveOKCode int = 200

/*UserSettingSaveOK ok

swagger:response userSettingSaveOK
*/
type UserSettingSaveOK struct {

	/*
	  In: Body
	*/
	Payload *models.Setting `json:"body,omitempty"`
}

// NewUserSettingSaveOK creates UserSettingSaveOK with default headers values
func NewUserSettingSaveOK() *UserSettingSaveOK {
	return &UserSettingSaveOK{}
}

// WithPayload adds the payload to the user setting save o k response
func (o *UserSettingSaveOK) WithPayload(payload *models.Setting) *UserSettingSaveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user setting save o k response
func (o *UserSettingSaveOK) SetPayload(payload *models.Setting) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserSettingSaveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserSettingSaveBadRequestCode is the HTTP code returned for type UserSettingSaveBadRequest
const UserSettingSaveBadRequestCode int = 400

/*UserSettingSaveBadRequest Bad request

swagger:response userSettingSaveBadRequest
*/
type UserSettingSaveBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUserSettingSaveBadRequest creates UserSettingSaveBadRequest with default headers values
func NewUserSettingSaveBadRequest() *UserSettingSaveBadRequest {
	return &UserSettingSaveBadRequest{}
}

// WithPayload adds the payload to the user setting save bad request response
func (o *UserSettingSaveBadRequest) WithPayload(payload string) *UserSettingSaveBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user setting save bad request response
func (o *UserSettingSaveBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserSettingSaveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// UserSettingSaveInternalServerErrorCode is the HTTP code returned for type UserSettingSaveInternalServerError
const UserSettingSaveInternalServerErrorCode int = 500

/*UserSettingSaveInternalServerError Internal server error

swagger:response userSettingSaveInternalServerError
*/
type UserSettingSaveInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUserSettingSaveInternalServerError creates UserSettingSaveInternalServerError with default headers values
func NewUserSettingSaveInternalServerError() *UserSettingSaveInternalServerError {
	return &UserSettingSaveInternalServerError{}
}

// WithPayload adds the payload to the user setting save internal server error response
func (o *UserSettingSaveInternalServerError) WithPayload(payload string) *UserSettingSaveInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user setting save internal server error response
func (o *UserSettingSaveInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserSettingSaveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
