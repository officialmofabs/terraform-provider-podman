// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodRestartReport pod restart report
//
// swagger:model PodRestartReport
type PodRestartReport struct {

	// errs
	Errs []string `json:"Errs"`

	// Id
	ID string `json:"Id,omitempty"`
}

// Validate validates this pod restart report
func (m *PodRestartReport) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod restart report based on context it is used
func (m *PodRestartReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodRestartReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodRestartReport) UnmarshalBinary(b []byte) error {
	var res PodRestartReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
