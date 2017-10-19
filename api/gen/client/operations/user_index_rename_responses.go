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

// UserIndexRenameReader is a Reader for the UserIndexRename structure.
type UserIndexRenameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserIndexRenameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserIndexRenameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserIndexRenameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUserIndexRenameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserIndexRenameOK creates a UserIndexRenameOK with default headers values
func NewUserIndexRenameOK() *UserIndexRenameOK {
	return &UserIndexRenameOK{}
}

/*UserIndexRenameOK handles this case with default header values.

ok
*/
type UserIndexRenameOK struct {
	Payload *models.StockIndex
}

func (o *UserIndexRenameOK) Error() string {
	return fmt.Sprintf("[POST /{userId}/indices/rename][%d] userIndexRenameOK  %+v", 200, o.Payload)
}

func (o *UserIndexRenameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StockIndex)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserIndexRenameBadRequest creates a UserIndexRenameBadRequest with default headers values
func NewUserIndexRenameBadRequest() *UserIndexRenameBadRequest {
	return &UserIndexRenameBadRequest{}
}

/*UserIndexRenameBadRequest handles this case with default header values.

Bad request
*/
type UserIndexRenameBadRequest struct {
	Payload string
}

func (o *UserIndexRenameBadRequest) Error() string {
	return fmt.Sprintf("[POST /{userId}/indices/rename][%d] userIndexRenameBadRequest  %+v", 400, o.Payload)
}

func (o *UserIndexRenameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserIndexRenameInternalServerError creates a UserIndexRenameInternalServerError with default headers values
func NewUserIndexRenameInternalServerError() *UserIndexRenameInternalServerError {
	return &UserIndexRenameInternalServerError{}
}

/*UserIndexRenameInternalServerError handles this case with default header values.

Internal server error
*/
type UserIndexRenameInternalServerError struct {
	Payload string
}

func (o *UserIndexRenameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /{userId}/indices/rename][%d] userIndexRenameInternalServerError  %+v", 500, o.Payload)
}

func (o *UserIndexRenameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}