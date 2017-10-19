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

// UserIndexSaveReader is a Reader for the UserIndexSave structure.
type UserIndexSaveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserIndexSaveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserIndexSaveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserIndexSaveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUserIndexSaveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserIndexSaveOK creates a UserIndexSaveOK with default headers values
func NewUserIndexSaveOK() *UserIndexSaveOK {
	return &UserIndexSaveOK{}
}

/*UserIndexSaveOK handles this case with default header values.

ok
*/
type UserIndexSaveOK struct {
	Payload *models.StockIndex
}

func (o *UserIndexSaveOK) Error() string {
	return fmt.Sprintf("[POST /{userId}/indices][%d] userIndexSaveOK  %+v", 200, o.Payload)
}

func (o *UserIndexSaveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StockIndex)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserIndexSaveBadRequest creates a UserIndexSaveBadRequest with default headers values
func NewUserIndexSaveBadRequest() *UserIndexSaveBadRequest {
	return &UserIndexSaveBadRequest{}
}

/*UserIndexSaveBadRequest handles this case with default header values.

Bad request
*/
type UserIndexSaveBadRequest struct {
	Payload string
}

func (o *UserIndexSaveBadRequest) Error() string {
	return fmt.Sprintf("[POST /{userId}/indices][%d] userIndexSaveBadRequest  %+v", 400, o.Payload)
}

func (o *UserIndexSaveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserIndexSaveInternalServerError creates a UserIndexSaveInternalServerError with default headers values
func NewUserIndexSaveInternalServerError() *UserIndexSaveInternalServerError {
	return &UserIndexSaveInternalServerError{}
}

/*UserIndexSaveInternalServerError handles this case with default header values.

Internal server error
*/
type UserIndexSaveInternalServerError struct {
	Payload string
}

func (o *UserIndexSaveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /{userId}/indices][%d] userIndexSaveInternalServerError  %+v", 500, o.Payload)
}

func (o *UserIndexSaveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}