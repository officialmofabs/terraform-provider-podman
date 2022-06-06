// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutContainerArchiveBadRequestBody put container archive bad request body
//
// swagger:model putContainerArchiveBadRequestBody
type PutContainerArchiveBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this put container archive bad request body
func (m *PutContainerArchiveBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put container archive bad request body based on context it is used
func (m *PutContainerArchiveBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutContainerArchiveBadRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutContainerArchiveBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PutContainerArchiveBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
