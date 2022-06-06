// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerHealthCheckConfig ContainerHealthCheckConfig describes a container healthcheck with attributes
// like command, retries, interval, start period, and timeout.
//
// swagger:model ContainerHealthCheckConfig
type ContainerHealthCheckConfig struct {

	// healthconfig
	Healthconfig *Schema2HealthConfig `json:"healthconfig,omitempty"`
}

// Validate validates this container health check config
func (m *ContainerHealthCheckConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthconfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerHealthCheckConfig) validateHealthconfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Healthconfig) { // not required
		return nil
	}

	if m.Healthconfig != nil {
		if err := m.Healthconfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("healthconfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("healthconfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this container health check config based on the context it is used
func (m *ContainerHealthCheckConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHealthconfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerHealthCheckConfig) contextValidateHealthconfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Healthconfig != nil {
		if err := m.Healthconfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("healthconfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("healthconfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerHealthCheckConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerHealthCheckConfig) UnmarshalBinary(b []byte) error {
	var res ContainerHealthCheckConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
