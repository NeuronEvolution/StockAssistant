// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StockIndexAdviceListDefaultBodyErrorsItems stock index advice list default body errors items
// swagger:model stockIndexAdviceListDefaultBodyErrorsItems

type StockIndexAdviceListDefaultBodyErrorsItems struct {

	// error code
	Code string `json:"code,omitempty"`

	// field name
	Field string `json:"field,omitempty"`

	// error message
	Message string `json:"message,omitempty"`
}

/* polymorph stockIndexAdviceListDefaultBodyErrorsItems code false */

/* polymorph stockIndexAdviceListDefaultBodyErrorsItems field false */

/* polymorph stockIndexAdviceListDefaultBodyErrorsItems message false */

// Validate validates this stock index advice list default body errors items
func (m *StockIndexAdviceListDefaultBodyErrorsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StockIndexAdviceListDefaultBodyErrorsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StockIndexAdviceListDefaultBodyErrorsItems) UnmarshalBinary(b []byte) error {
	var res StockIndexAdviceListDefaultBodyErrorsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
