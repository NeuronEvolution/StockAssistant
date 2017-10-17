// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/StockAssistant/api/models"
)

// UserStockEvaluateAddOrUpdateOKCode is the HTTP code returned for type UserStockEvaluateAddOrUpdateOK
const UserStockEvaluateAddOrUpdateOKCode int = 200

/*UserStockEvaluateAddOrUpdateOK ok

swagger:response userStockEvaluateAddOrUpdateOK
*/
type UserStockEvaluateAddOrUpdateOK struct {

	/*
	  In: Body
	*/
	Payload *models.StockEvaluate `json:"body,omitempty"`
}

// NewUserStockEvaluateAddOrUpdateOK creates UserStockEvaluateAddOrUpdateOK with default headers values
func NewUserStockEvaluateAddOrUpdateOK() *UserStockEvaluateAddOrUpdateOK {
	return &UserStockEvaluateAddOrUpdateOK{}
}

// WithPayload adds the payload to the user stock evaluate add or update o k response
func (o *UserStockEvaluateAddOrUpdateOK) WithPayload(payload *models.StockEvaluate) *UserStockEvaluateAddOrUpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user stock evaluate add or update o k response
func (o *UserStockEvaluateAddOrUpdateOK) SetPayload(payload *models.StockEvaluate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserStockEvaluateAddOrUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserStockEvaluateAddOrUpdateBadRequestCode is the HTTP code returned for type UserStockEvaluateAddOrUpdateBadRequest
const UserStockEvaluateAddOrUpdateBadRequestCode int = 400

/*UserStockEvaluateAddOrUpdateBadRequest Bad request

swagger:response userStockEvaluateAddOrUpdateBadRequest
*/
type UserStockEvaluateAddOrUpdateBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUserStockEvaluateAddOrUpdateBadRequest creates UserStockEvaluateAddOrUpdateBadRequest with default headers values
func NewUserStockEvaluateAddOrUpdateBadRequest() *UserStockEvaluateAddOrUpdateBadRequest {
	return &UserStockEvaluateAddOrUpdateBadRequest{}
}

// WithPayload adds the payload to the user stock evaluate add or update bad request response
func (o *UserStockEvaluateAddOrUpdateBadRequest) WithPayload(payload string) *UserStockEvaluateAddOrUpdateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user stock evaluate add or update bad request response
func (o *UserStockEvaluateAddOrUpdateBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserStockEvaluateAddOrUpdateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// UserStockEvaluateAddOrUpdateInternalServerErrorCode is the HTTP code returned for type UserStockEvaluateAddOrUpdateInternalServerError
const UserStockEvaluateAddOrUpdateInternalServerErrorCode int = 500

/*UserStockEvaluateAddOrUpdateInternalServerError Internal error

swagger:response userStockEvaluateAddOrUpdateInternalServerError
*/
type UserStockEvaluateAddOrUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUserStockEvaluateAddOrUpdateInternalServerError creates UserStockEvaluateAddOrUpdateInternalServerError with default headers values
func NewUserStockEvaluateAddOrUpdateInternalServerError() *UserStockEvaluateAddOrUpdateInternalServerError {
	return &UserStockEvaluateAddOrUpdateInternalServerError{}
}

// WithPayload adds the payload to the user stock evaluate add or update internal server error response
func (o *UserStockEvaluateAddOrUpdateInternalServerError) WithPayload(payload string) *UserStockEvaluateAddOrUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user stock evaluate add or update internal server error response
func (o *UserStockEvaluateAddOrUpdateInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserStockEvaluateAddOrUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
