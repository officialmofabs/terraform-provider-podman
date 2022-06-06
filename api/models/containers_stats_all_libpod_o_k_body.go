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

// ContainersStatsAllLibpodOKBody containers stats all libpod o k body
//
// swagger:model containersStatsAllLibpodOKBody
type ContainersStatsAllLibpodOKBody struct {

	// avg CPU
	AvgCPU float64 `json:"AvgCPU,omitempty"`

	// block input
	BlockInput uint64 `json:"BlockInput,omitempty"`

	// block output
	BlockOutput uint64 `json:"BlockOutput,omitempty"`

	// CPU
	CPU float64 `json:"CPU,omitempty"`

	// CPU nano
	CPUNano uint64 `json:"CPUNano,omitempty"`

	// CPU system nano
	CPUSystemNano uint64 `json:"CPUSystemNano,omitempty"`

	// container ID
	ContainerID string `json:"ContainerID,omitempty"`

	// duration
	Duration uint64 `json:"Duration,omitempty"`

	// mem limit
	MemLimit uint64 `json:"MemLimit,omitempty"`

	// mem perc
	MemPerc float64 `json:"MemPerc,omitempty"`

	// mem usage
	MemUsage uint64 `json:"MemUsage,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// net input
	NetInput uint64 `json:"NetInput,omitempty"`

	// net output
	NetOutput uint64 `json:"NetOutput,omitempty"`

	// p i ds
	PIDs uint64 `json:"PIDs,omitempty"`

	// per CPU
	PerCPU []uint64 `json:"PerCPU"`

	// system nano
	SystemNano uint64 `json:"SystemNano,omitempty"`

	// up time
	UpTime Duration `json:"UpTime,omitempty"`
}

// Validate validates this containers stats all libpod o k body
func (m *ContainersStatsAllLibpodOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainersStatsAllLibpodOKBody) validateUpTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpTime) { // not required
		return nil
	}

	if err := m.UpTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("UpTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("UpTime")
		}
		return err
	}

	return nil
}

// ContextValidate validate this containers stats all libpod o k body based on the context it is used
func (m *ContainersStatsAllLibpodOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUpTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainersStatsAllLibpodOKBody) contextValidateUpTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.UpTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("UpTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("UpTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainersStatsAllLibpodOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainersStatsAllLibpodOKBody) UnmarshalBinary(b []byte) error {
	var res ContainersStatsAllLibpodOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
