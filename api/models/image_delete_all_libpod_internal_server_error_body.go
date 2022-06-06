// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImageDeleteAllLibpodInternalServerErrorBody image delete all libpod internal server error body
//
// swagger:model imageDeleteAllLibpodInternalServerErrorBody
type ImageDeleteAllLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image delete all libpod internal server error body
func (m *ImageDeleteAllLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image delete all libpod internal server error body based on context it is used
func (m *ImageDeleteAllLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageDeleteAllLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageDeleteAllLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImageDeleteAllLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
