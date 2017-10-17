// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/StockAssistant/api/models"
)

// UserStockEvaluateGetOKCode is the HTTP code returned for type UserStockEvaluateGetOK
const UserStockEvaluateGetOKCode int = 200

/*UserStockEvaluateGetOK ok

swagger:response userStockEvaluateGetOK
*/
type UserStockEvaluateGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.StockEvaluate `json:"body,omitempty"`
}

// NewUserStockEvaluateGetOK creates UserStockEvaluateGetOK with default headers values
func NewUserStockEvaluateGetOK() *UserStockEvaluateGetOK {
	return &UserStockEvaluateGetOK{}
}

// WithPayload adds the payload to the user stock evaluate get o k response
func (o *UserStockEvaluateGetOK) WithPayload(payload *models.StockEvaluate) *UserStockEvaluateGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user stock evaluate get o k response
func (o *UserStockEvaluateGetOK) SetPayload(payload *models.StockEvaluate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserStockEvaluateGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserStockEvaluateGetBadRequestCode is the HTTP code returned for type UserStockEvaluateGetBadRequest
const UserStockEvaluateGetBadRequestCode int = 400

/*UserStockEvaluateGetBadRequest Bad request

swagger:response userStockEvaluateGetBadRequest
*/
type UserStockEvaluateGetBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUserStockEvaluateGetBadRequest creates UserStockEvaluateGetBadRequest with default headers values
func NewUserStockEvaluateGetBadRequest() *UserStockEvaluateGetBadRequest {
	return &UserStockEvaluateGetBadRequest{}
}

// WithPayload adds the payload to the user stock evaluate get bad request response
func (o *UserStockEvaluateGetBadRequest) WithPayload(payload string) *UserStockEvaluateGetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user stock evaluate get bad request response
func (o *UserStockEvaluateGetBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserStockEvaluateGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// UserStockEvaluateGetInternalServerErrorCode is the HTTP code returned for type UserStockEvaluateGetInternalServerError
const UserStockEvaluateGetInternalServerErrorCode int = 500

/*UserStockEvaluateGetInternalServerError Internal error

swagger:response userStockEvaluateGetInternalServerError
*/
type UserStockEvaluateGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUserStockEvaluateGetInternalServerError creates UserStockEvaluateGetInternalServerError with default headers values
func NewUserStockEvaluateGetInternalServerError() *UserStockEvaluateGetInternalServerError {
	return &UserStockEvaluateGetInternalServerError{}
}

// WithPayload adds the payload to the user stock evaluate get internal server error response
func (o *UserStockEvaluateGetInternalServerError) WithPayload(payload string) *UserStockEvaluateGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user stock evaluate get internal server error response
func (o *UserStockEvaluateGetInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserStockEvaluateGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}