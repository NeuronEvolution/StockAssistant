// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/NeuronEvolution/StockAssistant/api/gen/models"
)

// UserStockEvaluateSaveReader is a Reader for the UserStockEvaluateSave structure.
type UserStockEvaluateSaveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserStockEvaluateSaveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserStockEvaluateSaveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUserStockEvaluateSaveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserStockEvaluateSaveOK creates a UserStockEvaluateSaveOK with default headers values
func NewUserStockEvaluateSaveOK() *UserStockEvaluateSaveOK {
	return &UserStockEvaluateSaveOK{}
}

/*UserStockEvaluateSaveOK handles this case with default header values.

ok
*/
type UserStockEvaluateSaveOK struct {
	Payload *models.StockEvaluate
}

func (o *UserStockEvaluateSaveOK) Error() string {
	return fmt.Sprintf("[POST /{userId}/stockEvaluates][%d] userStockEvaluateSaveOK  %+v", 200, o.Payload)
}

func (o *UserStockEvaluateSaveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StockEvaluate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserStockEvaluateSaveDefault creates a UserStockEvaluateSaveDefault with default headers values
func NewUserStockEvaluateSaveDefault(code int) *UserStockEvaluateSaveDefault {
	return &UserStockEvaluateSaveDefault{
		_statusCode: code,
	}
}

/*UserStockEvaluateSaveDefault handles this case with default header values.

Error response
*/
type UserStockEvaluateSaveDefault struct {
	_statusCode int

	Payload *models.UserStockEvaluateSaveDefaultBody
}

// Code gets the status code for the user stock evaluate save default response
func (o *UserStockEvaluateSaveDefault) Code() int {
	return o._statusCode
}

func (o *UserStockEvaluateSaveDefault) Error() string {
	return fmt.Sprintf("[POST /{userId}/stockEvaluates][%d] UserStockEvaluateSave default  %+v", o._statusCode, o.Payload)
}

func (o *UserStockEvaluateSaveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserStockEvaluateSaveDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
