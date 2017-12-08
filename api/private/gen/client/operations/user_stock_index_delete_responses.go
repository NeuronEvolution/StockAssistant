// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/NeuronEvolution/StockAssistant/api/private/gen/models"
)

// UserStockIndexDeleteReader is a Reader for the UserStockIndexDelete structure.
type UserStockIndexDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserStockIndexDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserStockIndexDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUserStockIndexDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserStockIndexDeleteOK creates a UserStockIndexDeleteOK with default headers values
func NewUserStockIndexDeleteOK() *UserStockIndexDeleteOK {
	return &UserStockIndexDeleteOK{}
}

/*UserStockIndexDeleteOK handles this case with default header values.

ok
*/
type UserStockIndexDeleteOK struct {
}

func (o *UserStockIndexDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /{userId}/stockIndices/{indexName}][%d] userStockIndexDeleteOK ", 200)
}

func (o *UserStockIndexDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserStockIndexDeleteDefault creates a UserStockIndexDeleteDefault with default headers values
func NewUserStockIndexDeleteDefault(code int) *UserStockIndexDeleteDefault {
	return &UserStockIndexDeleteDefault{
		_statusCode: code,
	}
}

/*UserStockIndexDeleteDefault handles this case with default header values.

Error response
*/
type UserStockIndexDeleteDefault struct {
	_statusCode int

	Payload *models.UserStockIndexDeleteDefaultBody
}

// Code gets the status code for the user stock index delete default response
func (o *UserStockIndexDeleteDefault) Code() int {
	return o._statusCode
}

func (o *UserStockIndexDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /{userId}/stockIndices/{indexName}][%d] UserStockIndexDelete default  %+v", o._statusCode, o.Payload)
}

func (o *UserStockIndexDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserStockIndexDeleteDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}