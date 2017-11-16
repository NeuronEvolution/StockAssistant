// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserStockEvaluate stock evaluate
// swagger:model UserStockEvaluate

type UserStockEvaluate struct {

	// remark
	EvalRemark string `json:"evalRemark,omitempty"`

	// Exchange id
	ExchangeID string `json:"exchangeId,omitempty"`

	// Exchange name
	ExchangeName string `json:"exchangeName,omitempty"`

	// index count
	IndexCount int32 `json:"indexCount,omitempty"`

	// Industry name
	IndustryName string `json:"industryName,omitempty"`

	// Launch date
	LaunchDate strfmt.DateTime `json:"launchDate,omitempty"`

	// Stock code
	StockCode string `json:"stockCode,omitempty"`

	// stock id
	StockID string `json:"stockId,omitempty"`

	// Stock name cn
	StockNameCN string `json:"stockNameCN,omitempty"`

	// score
	TotalScore float64 `json:"totalScore,omitempty"`
}

/* polymorph UserStockEvaluate evalRemark false */

/* polymorph UserStockEvaluate exchangeId false */

/* polymorph UserStockEvaluate exchangeName false */

/* polymorph UserStockEvaluate indexCount false */

/* polymorph UserStockEvaluate industryName false */

/* polymorph UserStockEvaluate launchDate false */

/* polymorph UserStockEvaluate stockCode false */

/* polymorph UserStockEvaluate stockId false */

/* polymorph UserStockEvaluate stockNameCN false */

/* polymorph UserStockEvaluate totalScore false */

// Validate validates this user stock evaluate
func (m *UserStockEvaluate) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserStockEvaluate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserStockEvaluate) UnmarshalBinary(b []byte) error {
	var res UserStockEvaluate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
