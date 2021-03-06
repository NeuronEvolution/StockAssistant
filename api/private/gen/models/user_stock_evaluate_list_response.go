// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserStockEvaluateListResponse user stock evaluate list response
// swagger:model UserStockEvaluateListResponse
type UserStockEvaluateListResponse struct {

	// items
	Items UserStockEvaluateListResponseItems `json:"items"`

	// next page token
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// Validate validates this user stock evaluate list response
func (m *UserStockEvaluateListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserStockEvaluateListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserStockEvaluateListResponse) UnmarshalBinary(b []byte) error {
	var res UserStockEvaluateListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
