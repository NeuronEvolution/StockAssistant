// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserIndexRenameDefaultBody user index rename default body
// swagger:model userIndexRenameDefaultBody

type UserIndexRenameDefaultBody struct {

	// Error code
	Code string `json:"code,omitempty"`

	// errors
	Errors UserIndexEvaluateGetDefaultBodyErrors `json:"errors"`

	// Error message
	Message string `json:"message,omitempty"`

	// status
	Status *int32 `json:"status,omitempty"`
}

/* polymorph userIndexRenameDefaultBody code false */

/* polymorph userIndexRenameDefaultBody errors false */

/* polymorph userIndexRenameDefaultBody message false */

/* polymorph userIndexRenameDefaultBody status false */

// Validate validates this user index rename default body
func (m *UserIndexRenameDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserIndexRenameDefaultBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserIndexRenameDefaultBody) UnmarshalBinary(b []byte) error {
	var res UserIndexRenameDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
