// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserSettingDeleteDefaultBodyErrorsItems user setting delete default body errors items
// swagger:model userSettingDeleteDefaultBodyErrorsItems

type UserSettingDeleteDefaultBodyErrorsItems struct {

	// error code
	Code string `json:"code,omitempty"`

	// field name
	Field string `json:"field,omitempty"`

	// error message
	Message string `json:"message,omitempty"`
}

/* polymorph userSettingDeleteDefaultBodyErrorsItems code false */

/* polymorph userSettingDeleteDefaultBodyErrorsItems field false */

/* polymorph userSettingDeleteDefaultBodyErrorsItems message false */

// Validate validates this user setting delete default body errors items
func (m *UserSettingDeleteDefaultBodyErrorsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserSettingDeleteDefaultBodyErrorsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSettingDeleteDefaultBodyErrorsItems) UnmarshalBinary(b []byte) error {
	var res UserSettingDeleteDefaultBodyErrorsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}