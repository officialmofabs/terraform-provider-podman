// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemVersionOKBody system version o k body
//
// swagger:model systemVersionOKBody
type SystemVersionOKBody struct {

	// API version
	APIVersion string `json:"ApiVersion,omitempty"`

	// arch
	Arch string `json:"Arch,omitempty"`

	// build time
	BuildTime string `json:"BuildTime,omitempty"`

	// components
	Components []*ComponentVersion `json:"Components"`

	// experimental
	Experimental bool `json:"Experimental,omitempty"`

	// git commit
	GitCommit string `json:"GitCommit,omitempty"`

	// go version
	GoVersion string `json:"GoVersion,omitempty"`

	// kernel version
	KernelVersion string `json:"KernelVersion,omitempty"`

	// min API version
	MinAPIVersion string `json:"MinAPIVersion,omitempty"`

	// os
	Os string `json:"Os,omitempty"`

	// platform
	Platform interface{} `json:"Platform,omitempty"`

	// version
	Version string `json:"Version,omitempty"`
}

// Validate validates this system version o k body
func (m *SystemVersionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemVersionOKBody) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this system version o k body based on the context it is used
func (m *SystemVersionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemVersionOKBody) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Components); i++ {

		if m.Components[i] != nil {
			if err := m.Components[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemVersionOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemVersionOKBody) UnmarshalBinary(b []byte) error {
	var res SystemVersionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
