// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VolumeCreateLibpodCreatedBody volume create libpod created body
//
// swagger:model volumeCreateLibpodCreatedBody
type VolumeCreateLibpodCreatedBody struct {

	// Anonymous indicates that the volume was created as an anonymous
	// volume for a specific container, and will be be removed when any
	// container using it is removed.
	Anonymous bool `json:"Anonymous,omitempty"`

	// CreatedAt is the date and time the volume was created at. This is not
	// stored for older Libpod volumes; if so, it will be omitted.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"CreatedAt,omitempty"`

	// Driver is the driver used to create the volume.
	// If set to "local" or "", the Local driver (Podman built-in code) is
	// used to service the volume; otherwise, a volume plugin with the given
	// name is used to mount and manage the volume.
	Driver string `json:"Driver,omitempty"`

	// GID is the GID that the volume was created with.
	GID int64 `json:"GID,omitempty"`

	// Labels includes the volume's configured labels, key:value pairs that
	// can be passed during volume creation to provide information for third
	// party tools.
	Labels map[string]string `json:"Labels,omitempty"`

	// MountCount is the number of times this volume has been mounted.
	MountCount uint64 `json:"MountCount,omitempty"`

	// Mountpoint is the path on the host where the volume is mounted.
	Mountpoint string `json:"Mountpoint,omitempty"`

	// Name is the name of the volume.
	Name string `json:"Name,omitempty"`

	// NeedsChown indicates that the next time the volume is mounted into
	// a container, the container will chown the volume to the container process
	// UID/GID.
	NeedsChown bool `json:"NeedsChown,omitempty"`

	// NeedsCopyUp indicates that the next time the volume is mounted into
	NeedsCopyUp bool `json:"NeedsCopyUp,omitempty"`

	// Options is a set of options that were used when creating the volume.
	// For the Local driver, these are mount options that will be used to
	// determine how a local filesystem is mounted; they are handled as
	// parameters to Mount in a manner described in the volume create
	// manpage.
	// For non-local drivers, these are passed as-is to the volume plugin.
	Options map[string]string `json:"Options,omitempty"`

	// Scope is unused and provided solely for Docker compatibility. It is
	// unconditionally set to "local".
	Scope string `json:"Scope,omitempty"`

	// Status is used to return information on the volume's current state,
	// if the volume was created using a volume plugin (uses a Driver that
	// is not the local driver).
	// Status is provided to us by an external program, so no guarantees are
	// made about its format or contents. Further, it is an optional field,
	// so it may not be set even in cases where a volume plugin is in use.
	Status map[string]interface{} `json:"Status,omitempty"`

	// UID is the UID that the volume was created with.
	UID int64 `json:"UID,omitempty"`
}

// Validate validates this volume create libpod created body
func (m *VolumeCreateLibpodCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeCreateLibpodCreatedBody) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this volume create libpod created body based on context it is used
func (m *VolumeCreateLibpodCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeCreateLibpodCreatedBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeCreateLibpodCreatedBody) UnmarshalBinary(b []byte) error {
	var res VolumeCreateLibpodCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
