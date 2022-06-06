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

// SecretInfoReport secret info report
//
// swagger:model SecretInfoReport
type SecretInfoReport struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"CreatedAt,omitempty"`

	// ID
	ID string `json:"ID,omitempty"`

	// spec
	Spec *SecretSpec `json:"Spec,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"UpdatedAt,omitempty"`
}

// Validate validates this secret info report
func (m *SecretInfoReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretInfoReport) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SecretInfoReport) validateSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Spec")
			}
			return err
		}
	}

	return nil
}

func (m *SecretInfoReport) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this secret info report based on the context it is used
func (m *SecretInfoReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretInfoReport) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.Spec != nil {
		if err := m.Spec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecretInfoReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretInfoReport) UnmarshalBinary(b []byte) error {
	var res SecretInfoReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
