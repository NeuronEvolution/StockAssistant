// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserIndexEvaluateGetDefaultBody user index evaluate get default body
// swagger:model userIndexEvaluateGetDefaultBody

type UserIndexEvaluateGetDefaultBody struct {

	// Error code
	Code string `json:"code,omitempty"`

	// errors
	Errors UserIndexEvaluateGetDefaultBodyErrors `json:"errors"`

	// Error message
	Message string `json:"message,omitempty"`

	// status
	Status *int32 `json:"status,omitempty"`
}

/* polymorph userIndexEvaluateGetDefaultBody code false */

/* polymorph userIndexEvaluateGetDefaultBody errors false */

/* polymorph userIndexEvaluateGetDefaultBody message false */

/* polymorph userIndexEvaluateGetDefaultBody status false */

// Validate validates this user index evaluate get default body
func (m *UserIndexEvaluateGetDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserIndexEvaluateGetDefaultBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserIndexEvaluateGetDefaultBody) UnmarshalBinary(b []byte) error {
	var res UserIndexEvaluateGetDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
