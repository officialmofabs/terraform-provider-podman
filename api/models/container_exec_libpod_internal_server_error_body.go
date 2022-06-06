// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerExecLibpodInternalServerErrorBody container exec libpod internal server error body
//
// swagger:model containerExecLibpodInternalServerErrorBody
type ContainerExecLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container exec libpod internal server error body
func (m *ContainerExecLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container exec libpod internal server error body based on context it is used
func (m *ContainerExecLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerExecLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerExecLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerExecLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
