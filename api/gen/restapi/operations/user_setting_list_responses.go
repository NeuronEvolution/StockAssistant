// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/StockAssistant/api/gen/models"
)

// UserSettingListOKCode is the HTTP code returned for type UserSettingListOK
const UserSettingListOKCode int = 200

/*UserSettingListOK ok

swagger:response userSettingListOK
*/
type UserSettingListOK struct {

	/*
	  In: Body
	*/
	Payload models.UserSettingListOKBody `json:"body,omitempty"`
}

// NewUserSettingListOK creates UserSettingListOK with default headers values
func NewUserSettingListOK() *UserSettingListOK {
	return &UserSettingListOK{}
}

// WithPayload adds the payload to the user setting list o k response
func (o *UserSettingListOK) WithPayload(payload models.UserSettingListOKBody) *UserSettingListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user setting list o k response
func (o *UserSettingListOK) SetPayload(payload models.UserSettingListOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserSettingListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.UserSettingListOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*UserSettingListDefault Error response

swagger:response userSettingListDefault
*/
type UserSettingListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.UserSettingListDefaultBody `json:"body,omitempty"`
}

// NewUserSettingListDefault creates UserSettingListDefault with default headers values
func NewUserSettingListDefault(code int) *UserSettingListDefault {
	if code <= 0 {
		code = 500
	}

	return &UserSettingListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user setting list default response
func (o *UserSettingListDefault) WithStatusCode(code int) *UserSettingListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user setting list default response
func (o *UserSettingListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user setting list default response
func (o *UserSettingListDefault) WithPayload(payload *models.UserSettingListDefaultBody) *UserSettingListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user setting list default response
func (o *UserSettingListDefault) SetPayload(payload *models.UserSettingListDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserSettingListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
