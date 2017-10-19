// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserIndexSaveDefaultBodyErrorsItems user index save default body errors items
// swagger:model userIndexSaveDefaultBodyErrorsItems

type UserIndexSaveDefaultBodyErrorsItems struct {

	// error code
	Code string `json:"code,omitempty"`

	// field name
	Field string `json:"field,omitempty"`

	// error message
	Message string `json:"message,omitempty"`
}

/* polymorph userIndexSaveDefaultBodyErrorsItems code false */

/* polymorph userIndexSaveDefaultBodyErrorsItems field false */

/* polymorph userIndexSaveDefaultBodyErrorsItems message false */

// Validate validates this user index save default body errors items
func (m *UserIndexSaveDefaultBodyErrorsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserIndexSaveDefaultBodyErrorsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserIndexSaveDefaultBodyErrorsItems) UnmarshalBinary(b []byte) error {
	var res UserIndexSaveDefaultBodyErrorsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
