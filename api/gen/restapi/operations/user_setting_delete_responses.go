// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UserSettingDeleteOKCode is the HTTP code returned for type UserSettingDeleteOK
const UserSettingDeleteOKCode int = 200

/*UserSettingDeleteOK ok

swagger:response userSettingDeleteOK
*/
type UserSettingDeleteOK struct {
}

// NewUserSettingDeleteOK creates UserSettingDeleteOK with default headers values
func NewUserSettingDeleteOK() *UserSettingDeleteOK {
	return &UserSettingDeleteOK{}
}

// WriteResponse to the client
func (o *UserSettingDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// UserSettingDeleteBadRequestCode is the HTTP code returned for type UserSettingDeleteBadRequest
const UserSettingDeleteBadRequestCode int = 400

/*UserSettingDeleteBadRequest Bad request

swagger:response userSettingDeleteBadRequest
*/
type UserSettingDeleteBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUserSettingDeleteBadRequest creates UserSettingDeleteBadRequest with default headers values
func NewUserSettingDeleteBadRequest() *UserSettingDeleteBadRequest {
	return &UserSettingDeleteBadRequest{}
}

// WithPayload adds the payload to the user setting delete bad request response
func (o *UserSettingDeleteBadRequest) WithPayload(payload string) *UserSettingDeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user setting delete bad request response
func (o *UserSettingDeleteBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserSettingDeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// UserSettingDeleteInternalServerErrorCode is the HTTP code returned for type UserSettingDeleteInternalServerError
const UserSettingDeleteInternalServerErrorCode int = 500

/*UserSettingDeleteInternalServerError Internal server error

swagger:response userSettingDeleteInternalServerError
*/
type UserSettingDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUserSettingDeleteInternalServerError creates UserSettingDeleteInternalServerError with default headers values
func NewUserSettingDeleteInternalServerError() *UserSettingDeleteInternalServerError {
	return &UserSettingDeleteInternalServerError{}
}

// WithPayload adds the payload to the user setting delete internal server error response
func (o *UserSettingDeleteInternalServerError) WithPayload(payload string) *UserSettingDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user setting delete internal server error response
func (o *UserSettingDeleteInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserSettingDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}