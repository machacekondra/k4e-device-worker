// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HardwareProfileConfiguration hardware profile configuration
//
// swagger:model hardware-profile-configuration
type HardwareProfileConfiguration struct {

	// include
	Include bool `json:"include,omitempty"`

	// scope
	// Enum: [full delta]
	Scope string `json:"scope,omitempty"`
}

// Validate validates this hardware profile configuration
func (m *HardwareProfileConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hardwareProfileConfigurationTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["full","delta"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hardwareProfileConfigurationTypeScopePropEnum = append(hardwareProfileConfigurationTypeScopePropEnum, v)
	}
}

const (

	// HardwareProfileConfigurationScopeFull captures enum value "full"
	HardwareProfileConfigurationScopeFull string = "full"

	// HardwareProfileConfigurationScopeDelta captures enum value "delta"
	HardwareProfileConfigurationScopeDelta string = "delta"
)

// prop value enum
func (m *HardwareProfileConfiguration) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hardwareProfileConfigurationTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HardwareProfileConfiguration) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HardwareProfileConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HardwareProfileConfiguration) UnmarshalBinary(b []byte) error {
	var res HardwareProfileConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
