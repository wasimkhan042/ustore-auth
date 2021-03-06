// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserProduct user product
//
// swagger:model UserProduct
type UserProduct struct {

	// expiry date
	// Format: date-time
	ExpiryDate strfmt.DateTime `json:"expiry_date,omitempty"`

	// item name
	ItemName string `json:"item_name,omitempty"`

	// remaining time in days
	RemainingTimeInDays string `json:"remaining_time_in_days,omitempty"`

	// subs price
	SubsPrice float64 `json:"subs_price,omitempty"`
}

// Validate validates this user product
func (m *UserProduct) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiryDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserProduct) validateExpiryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expiry_date", "body", "date-time", m.ExpiryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user product based on context it is used
func (m *UserProduct) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserProduct) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserProduct) UnmarshalBinary(b []byte) error {
	var res UserProduct
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
