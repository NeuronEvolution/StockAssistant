// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserStockIndex User stock index
// swagger:model UserStockIndex

type UserStockIndex struct {

	// ai weight
	AiWeight int32 `json:"aiWeight,omitempty"`

	// desc
	Desc string `json:"desc,omitempty"`

	// Eval weight
	EvalWeight int32 `json:"evalWeight,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

/* polymorph UserStockIndex aiWeight false */

/* polymorph UserStockIndex desc false */

/* polymorph UserStockIndex evalWeight false */

/* polymorph UserStockIndex name false */

// Validate validates this user stock index
func (m *UserStockIndex) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserStockIndex) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserStockIndex) UnmarshalBinary(b []byte) error {
	var res UserStockIndex
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}