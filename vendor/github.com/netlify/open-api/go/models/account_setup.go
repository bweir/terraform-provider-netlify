// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountSetup account setup
// swagger:model accountSetup
type AccountSetup struct {

	// extra seats block
	ExtraSeatsBlock int64 `json:"extra_seats_block,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// payment method id
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// period
	Period string `json:"period,omitempty"`

	// type id
	// Required: true
	TypeID *string `json:"type_id"`
}

// Validate validates this account setup
func (m *AccountSetup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePeriod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountSetup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var accountSetupTypePeriodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["monthly","yearly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountSetupTypePeriodPropEnum = append(accountSetupTypePeriodPropEnum, v)
	}
}

const (

	// AccountSetupPeriodMonthly captures enum value "monthly"
	AccountSetupPeriodMonthly string = "monthly"

	// AccountSetupPeriodYearly captures enum value "yearly"
	AccountSetupPeriodYearly string = "yearly"
)

// prop value enum
func (m *AccountSetup) validatePeriodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountSetupTypePeriodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountSetup) validatePeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.Period) { // not required
		return nil
	}

	// value enum
	if err := m.validatePeriodEnum("period", "body", m.Period); err != nil {
		return err
	}

	return nil
}

func (m *AccountSetup) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountSetup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountSetup) UnmarshalBinary(b []byte) error {
	var res AccountSetup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
