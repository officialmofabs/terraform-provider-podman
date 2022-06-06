// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeInspectLibpodNotFoundBody volume inspect libpod not found body
//
// swagger:model volumeInspectLibpodNotFoundBody
type VolumeInspectLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume inspect libpod not found body
func (m *VolumeInspectLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume inspect libpod not found body based on context it is used
func (m *VolumeInspectLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeInspectLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeInspectLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res VolumeInspectLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
