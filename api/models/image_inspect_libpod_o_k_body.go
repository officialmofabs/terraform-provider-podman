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
	"github.com/go-openapi/validate"
)

// ImageInspectLibpodOKBody image inspect libpod o k body
//
// swagger:model imageInspectLibpodOKBody
type ImageInspectLibpodOKBody struct {

	// annotations
	Annotations map[string]string `json:"Annotations,omitempty"`

	// architecture
	Architecture string `json:"Architecture,omitempty"`

	// author
	Author string `json:"Author,omitempty"`

	// comment
	Comment string `json:"Comment,omitempty"`

	// config
	Config *ImageConfig `json:"Config,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"Created,omitempty"`

	// digest
	Digest Digest `json:"Digest,omitempty"`

	// graph driver
	GraphDriver *DriverData `json:"GraphDriver,omitempty"`

	// healthcheck
	Healthcheck *Schema2HealthConfig `json:"Healthcheck,omitempty"`

	// history
	History []*History `json:"History"`

	// ID
	ID string `json:"Id,omitempty"`

	// labels
	Labels map[string]string `json:"Labels,omitempty"`

	// manifest type
	ManifestType string `json:"ManifestType,omitempty"`

	// names history
	NamesHistory []string `json:"NamesHistory"`

	// os
	Os string `json:"Os,omitempty"`

	// parent
	Parent string `json:"Parent,omitempty"`

	// repo digests
	RepoDigests []string `json:"RepoDigests"`

	// repo tags
	RepoTags []string `json:"RepoTags"`

	// root f s
	RootFS *RootFS `json:"RootFS,omitempty"`

	// size
	Size int64 `json:"Size,omitempty"`

	// user
	User string `json:"User,omitempty"`

	// version
	Version string `json:"Version,omitempty"`

	// virtual size
	VirtualSize int64 `json:"VirtualSize,omitempty"`
}

// Validate validates this image inspect libpod o k body
func (m *ImageInspectLibpodOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDigest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGraphDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthcheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootFS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageInspectLibpodOKBody) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config")
			}
			return err
		}
	}

	return nil
}

func (m *ImageInspectLibpodOKBody) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("Created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageInspectLibpodOKBody) validateDigest(formats strfmt.Registry) error {
	if swag.IsZero(m.Digest) { // not required
		return nil
	}

	if err := m.Digest.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Digest")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Digest")
		}
		return err
	}

	return nil
}

func (m *ImageInspectLibpodOKBody) validateGraphDriver(formats strfmt.Registry) error {
	if swag.IsZero(m.GraphDriver) { // not required
		return nil
	}

	if m.GraphDriver != nil {
		if err := m.GraphDriver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GraphDriver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("GraphDriver")
			}
			return err
		}
	}

	return nil
}

func (m *ImageInspectLibpodOKBody) validateHealthcheck(formats strfmt.Registry) error {
	if swag.IsZero(m.Healthcheck) { // not required
		return nil
	}

	if m.Healthcheck != nil {
		if err := m.Healthcheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Healthcheck")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Healthcheck")
			}
			return err
		}
	}

	return nil
}

func (m *ImageInspectLibpodOKBody) validateHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.History) { // not required
		return nil
	}

	for i := 0; i < len(m.History); i++ {
		if swag.IsZero(m.History[i]) { // not required
			continue
		}

		if m.History[i] != nil {
			if err := m.History[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("History" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("History" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ImageInspectLibpodOKBody) validateRootFS(formats strfmt.Registry) error {
	if swag.IsZero(m.RootFS) { // not required
		return nil
	}

	if m.RootFS != nil {
		if err := m.RootFS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RootFS")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RootFS")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this image inspect libpod o k body based on the context it is used
func (m *ImageInspectLibpodOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDigest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGraphDriver(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealthcheck(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRootFS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageInspectLibpodOKBody) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config")
			}
			return err
		}
	}

	return nil
}

func (m *ImageInspectLibpodOKBody) contextValidateDigest(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Digest.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Digest")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Digest")
		}
		return err
	}

	return nil
}

func (m *ImageInspectLibpodOKBody) contextValidateGraphDriver(ctx context.Context, formats strfmt.Registry) error {

	if m.GraphDriver != nil {
		if err := m.GraphDriver.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GraphDriver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("GraphDriver")
			}
			return err
		}
	}

	return nil
}

func (m *ImageInspectLibpodOKBody) contextValidateHealthcheck(ctx context.Context, formats strfmt.Registry) error {

	if m.Healthcheck != nil {
		if err := m.Healthcheck.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Healthcheck")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Healthcheck")
			}
			return err
		}
	}

	return nil
}

func (m *ImageInspectLibpodOKBody) contextValidateHistory(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.History); i++ {

		if m.History[i] != nil {
			if err := m.History[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("History" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("History" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ImageInspectLibpodOKBody) contextValidateRootFS(ctx context.Context, formats strfmt.Registry) error {

	if m.RootFS != nil {
		if err := m.RootFS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RootFS")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RootFS")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageInspectLibpodOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageInspectLibpodOKBody) UnmarshalBinary(b []byte) error {
	var res ImageInspectLibpodOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
