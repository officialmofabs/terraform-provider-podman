// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewVolumeDeleteLibpodParams creates a new VolumeDeleteLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeDeleteLibpodParams() *VolumeDeleteLibpodParams {
	return &VolumeDeleteLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeDeleteLibpodParamsWithTimeout creates a new VolumeDeleteLibpodParams object
// with the ability to set a timeout on a request.
func NewVolumeDeleteLibpodParamsWithTimeout(timeout time.Duration) *VolumeDeleteLibpodParams {
	return &VolumeDeleteLibpodParams{
		timeout: timeout,
	}
}

// NewVolumeDeleteLibpodParamsWithContext creates a new VolumeDeleteLibpodParams object
// with the ability to set a context for a request.
func NewVolumeDeleteLibpodParamsWithContext(ctx context.Context) *VolumeDeleteLibpodParams {
	return &VolumeDeleteLibpodParams{
		Context: ctx,
	}
}

// NewVolumeDeleteLibpodParamsWithHTTPClient creates a new VolumeDeleteLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeDeleteLibpodParamsWithHTTPClient(client *http.Client) *VolumeDeleteLibpodParams {
	return &VolumeDeleteLibpodParams{
		HTTPClient: client,
	}
}

/* VolumeDeleteLibpodParams contains all the parameters to send to the API endpoint
   for the volume delete libpod operation.

   Typically these are written to a http.Request.
*/
type VolumeDeleteLibpodParams struct {

	/* Force.

	   force removal
	*/
	Force *bool

	/* Name.

	   the name or ID of the volume
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume delete libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeDeleteLibpodParams) WithDefaults() *VolumeDeleteLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume delete libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeDeleteLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume delete libpod params
func (o *VolumeDeleteLibpodParams) WithTimeout(timeout time.Duration) *VolumeDeleteLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume delete libpod params
func (o *VolumeDeleteLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume delete libpod params
func (o *VolumeDeleteLibpodParams) WithContext(ctx context.Context) *VolumeDeleteLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume delete libpod params
func (o *VolumeDeleteLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume delete libpod params
func (o *VolumeDeleteLibpodParams) WithHTTPClient(client *http.Client) *VolumeDeleteLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume delete libpod params
func (o *VolumeDeleteLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the volume delete libpod params
func (o *VolumeDeleteLibpodParams) WithForce(force *bool) *VolumeDeleteLibpodParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the volume delete libpod params
func (o *VolumeDeleteLibpodParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the volume delete libpod params
func (o *VolumeDeleteLibpodParams) WithName(name string) *VolumeDeleteLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the volume delete libpod params
func (o *VolumeDeleteLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeDeleteLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
