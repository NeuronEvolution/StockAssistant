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

// UserSettingDeleteReader is a Reader for the UserSettingDelete structure.
type UserSettingDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserSettingDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserSettingDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUserSettingDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserSettingDeleteOK creates a UserSettingDeleteOK with default headers values
func NewUserSettingDeleteOK() *UserSettingDeleteOK {
	return &UserSettingDeleteOK{}
}

/*UserSettingDeleteOK handles this case with default header values.

ok
*/
type UserSettingDeleteOK struct {
}

func (o *UserSettingDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /{userId}/settings/{configKey}][%d] userSettingDeleteOK ", 200)
}

func (o *UserSettingDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserSettingDeleteDefault creates a UserSettingDeleteDefault with default headers values
func NewUserSettingDeleteDefault(code int) *UserSettingDeleteDefault {
	return &UserSettingDeleteDefault{
		_statusCode: code,
	}
}

/*UserSettingDeleteDefault handles this case with default header values.

Error response
*/
type UserSettingDeleteDefault struct {
	_statusCode int

	Payload *models.UserSettingDeleteDefaultBody
}

// Code gets the status code for the user setting delete default response
func (o *UserSettingDeleteDefault) Code() int {
	return o._statusCode
}

func (o *UserSettingDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /{userId}/settings/{configKey}][%d] UserSettingDelete default  %+v", o._statusCode, o.Payload)
}

func (o *UserSettingDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserSettingDeleteDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}