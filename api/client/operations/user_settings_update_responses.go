// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/NeuronEvolution/StockAssistant/api/models"
)

// UserSettingsUpdateReader is a Reader for the UserSettingsUpdate structure.
type UserSettingsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserSettingsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserSettingsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserSettingsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUserSettingsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserSettingsUpdateOK creates a UserSettingsUpdateOK with default headers values
func NewUserSettingsUpdateOK() *UserSettingsUpdateOK {
	return &UserSettingsUpdateOK{}
}

/*UserSettingsUpdateOK handles this case with default header values.

ok
*/
type UserSettingsUpdateOK struct {
	Payload *models.Setting
}

func (o *UserSettingsUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /{userId}/settings/{configKey}][%d] userSettingsUpdateOK  %+v", 200, o.Payload)
}

func (o *UserSettingsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Setting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSettingsUpdateBadRequest creates a UserSettingsUpdateBadRequest with default headers values
func NewUserSettingsUpdateBadRequest() *UserSettingsUpdateBadRequest {
	return &UserSettingsUpdateBadRequest{}
}

/*UserSettingsUpdateBadRequest handles this case with default header values.

Bad request
*/
type UserSettingsUpdateBadRequest struct {
	Payload string
}

func (o *UserSettingsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /{userId}/settings/{configKey}][%d] userSettingsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UserSettingsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSettingsUpdateInternalServerError creates a UserSettingsUpdateInternalServerError with default headers values
func NewUserSettingsUpdateInternalServerError() *UserSettingsUpdateInternalServerError {
	return &UserSettingsUpdateInternalServerError{}
}

/*UserSettingsUpdateInternalServerError handles this case with default header values.

Internal error
*/
type UserSettingsUpdateInternalServerError struct {
	Payload string
}

func (o *UserSettingsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /{userId}/settings/{configKey}][%d] userSettingsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *UserSettingsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
