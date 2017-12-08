// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserStockIndexDeleteDefaultBodyErrorsItems user stock index delete default body errors items
// swagger:model userStockIndexDeleteDefaultBodyErrorsItems

type UserStockIndexDeleteDefaultBodyErrorsItems struct {

	// error code
	Code string `json:"code,omitempty"`

	// field name
	Field string `json:"field,omitempty"`

	// error message
	Message string `json:"message,omitempty"`
}

/* polymorph userStockIndexDeleteDefaultBodyErrorsItems code false */

/* polymorph userStockIndexDeleteDefaultBodyErrorsItems field false */

/* polymorph userStockIndexDeleteDefaultBodyErrorsItems message false */

// Validate validates this user stock index delete default body errors items
func (m *UserStockIndexDeleteDefaultBodyErrorsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserStockIndexDeleteDefaultBodyErrorsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserStockIndexDeleteDefaultBodyErrorsItems) UnmarshalBinary(b []byte) error {
	var res UserStockIndexDeleteDefaultBodyErrorsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}