// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeDeleteLibpodInternalServerErrorBody volume delete libpod internal server error body
//
// swagger:model volumeDeleteLibpodInternalServerErrorBody
type VolumeDeleteLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume delete libpod internal server error body
func (m *VolumeDeleteLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume delete libpod internal server error body based on context it is used
func (m *VolumeDeleteLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeDeleteLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeDeleteLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res VolumeDeleteLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
