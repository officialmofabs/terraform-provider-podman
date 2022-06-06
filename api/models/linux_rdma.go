// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LinuxRdma LinuxRdma for Linux cgroup 'rdma' resource management (Linux 4.11)
//
// swagger:model LinuxRdma
type LinuxRdma struct {

	// Maximum number of HCA handles that can be opened. Default is "no limit".
	HcaHandles uint32 `json:"hcaHandles,omitempty"`

	// Maximum number of HCA objects that can be created. Default is "no limit".
	HcaObjects uint32 `json:"hcaObjects,omitempty"`
}

// Validate validates this linux rdma
func (m *LinuxRdma) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this linux rdma based on context it is used
func (m *LinuxRdma) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinuxRdma) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinuxRdma) UnmarshalBinary(b []byte) error {
	var res LinuxRdma
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
