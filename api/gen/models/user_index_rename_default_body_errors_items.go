// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserIndexRenameDefaultBodyErrorsItems user index rename default body errors items
// swagger:model userIndexRenameDefaultBodyErrorsItems

type UserIndexRenameDefaultBodyErrorsItems struct {

	// error code
	Code string `json:"code,omitempty"`

	// field name
	Field string `json:"field,omitempty"`

	// error message
	Message string `json:"message,omitempty"`
}

/* polymorph userIndexRenameDefaultBodyErrorsItems code false */

/* polymorph userIndexRenameDefaultBodyErrorsItems field false */

/* polymorph userIndexRenameDefaultBodyErrorsItems message false */

// Validate validates this user index rename default body errors items
func (m *UserIndexRenameDefaultBodyErrorsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserIndexRenameDefaultBodyErrorsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserIndexRenameDefaultBodyErrorsItems) UnmarshalBinary(b []byte) error {
	var res UserIndexRenameDefaultBodyErrorsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
