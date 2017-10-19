// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserIndexDeleteDefaultBody user index delete default body
// swagger:model userIndexDeleteDefaultBody

type UserIndexDeleteDefaultBody struct {

	// Error code
	Code string `json:"code,omitempty"`

	// errors
	Errors UserIndexEvaluateGetDefaultBodyErrors `json:"errors"`

	// Error message
	Message string `json:"message,omitempty"`

	// status
	Status *int32 `json:"status,omitempty"`
}

/* polymorph userIndexDeleteDefaultBody code false */

/* polymorph userIndexDeleteDefaultBody errors false */

/* polymorph userIndexDeleteDefaultBody message false */

/* polymorph userIndexDeleteDefaultBody status false */

// Validate validates this user index delete default body
func (m *UserIndexDeleteDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserIndexDeleteDefaultBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserIndexDeleteDefaultBody) UnmarshalBinary(b []byte) error {
	var res UserIndexDeleteDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}