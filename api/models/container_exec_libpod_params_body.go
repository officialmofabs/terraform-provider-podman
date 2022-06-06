// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerExecLibpodParamsBody container exec libpod params body
//
// swagger:model containerExecLibpodParamsBody
type ContainerExecLibpodParamsBody struct {

	// Attach to stderr of the exec command
	AttachStderr bool `json:"AttachStderr,omitempty"`

	// Attach to stdin of the exec command
	AttachStdin bool `json:"AttachStdin,omitempty"`

	// Attach to stdout of the exec command
	AttachStdout bool `json:"AttachStdout,omitempty"`

	// Command to run, as a string or array of strings.
	Cmd []string `json:"Cmd"`

	// "Override the key sequence for detaching a container. Format is a single character [a-Z] or ctrl-<value> where <value> is one of: a-z, @, ^, [, , or _."
	//
	DetachKeys string `json:"DetachKeys,omitempty"`

	// A list of environment variables in the form ["VAR=value", ...]
	Env []string `json:"Env"`

	// Runs the exec process with extended privileges
	Privileged *bool `json:"Privileged,omitempty"`

	// Allocate a pseudo-TTY
	Tty bool `json:"Tty,omitempty"`

	// "The user, and optionally, group to run the exec process inside the container. Format is one of: user, user:group, uid, or uid:gid."
	//
	User string `json:"User,omitempty"`

	// The working directory for the exec process inside the container.
	WorkingDir string `json:"WorkingDir,omitempty"`
}

// Validate validates this container exec libpod params body
func (m *ContainerExecLibpodParamsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container exec libpod params body based on context it is used
func (m *ContainerExecLibpodParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerExecLibpodParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerExecLibpodParamsBody) UnmarshalBinary(b []byte) error {
	var res ContainerExecLibpodParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
