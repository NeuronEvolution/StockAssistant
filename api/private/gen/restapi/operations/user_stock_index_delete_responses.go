// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UserStockIndexDeleteOKCode is the HTTP code returned for type UserStockIndexDeleteOK
const UserStockIndexDeleteOKCode int = 200

/*UserStockIndexDeleteOK ok

swagger:response userStockIndexDeleteOK
*/
type UserStockIndexDeleteOK struct {
}

// NewUserStockIndexDeleteOK creates UserStockIndexDeleteOK with default headers values
func NewUserStockIndexDeleteOK() *UserStockIndexDeleteOK {
	return &UserStockIndexDeleteOK{}
}

// WriteResponse to the client
func (o *UserStockIndexDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
