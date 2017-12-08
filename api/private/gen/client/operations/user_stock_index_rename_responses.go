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

// UserStockIndexRenameReader is a Reader for the UserStockIndexRename structure.
type UserStockIndexRenameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserStockIndexRenameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserStockIndexRenameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUserStockIndexRenameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserStockIndexRenameOK creates a UserStockIndexRenameOK with default headers values
func NewUserStockIndexRenameOK() *UserStockIndexRenameOK {
	return &UserStockIndexRenameOK{}
}

/*UserStockIndexRenameOK handles this case with default header values.

ok
*/
type UserStockIndexRenameOK struct {
	Payload *models.UserStockIndex
}

func (o *UserStockIndexRenameOK) Error() string {
	return fmt.Sprintf("[POST /{userId}/stockIndices/rename][%d] userStockIndexRenameOK  %+v", 200, o.Payload)
}

func (o *UserStockIndexRenameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserStockIndex)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserStockIndexRenameDefault creates a UserStockIndexRenameDefault with default headers values
func NewUserStockIndexRenameDefault(code int) *UserStockIndexRenameDefault {
	return &UserStockIndexRenameDefault{
		_statusCode: code,
	}
}

/*UserStockIndexRenameDefault handles this case with default header values.

Error response
*/
type UserStockIndexRenameDefault struct {
	_statusCode int

	Payload *models.UserStockIndexRenameDefaultBody
}

// Code gets the status code for the user stock index rename default response
func (o *UserStockIndexRenameDefault) Code() int {
	return o._statusCode
}

func (o *UserStockIndexRenameDefault) Error() string {
	return fmt.Sprintf("[POST /{userId}/stockIndices/rename][%d] UserStockIndexRename default  %+v", o._statusCode, o.Payload)
}

func (o *UserStockIndexRenameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserStockIndexRenameDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}