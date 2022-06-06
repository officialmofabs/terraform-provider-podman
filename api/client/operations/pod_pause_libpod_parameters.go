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
)

// NewPodPauseLibpodParams creates a new PodPauseLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPodPauseLibpodParams() *PodPauseLibpodParams {
	return &PodPauseLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPodPauseLibpodParamsWithTimeout creates a new PodPauseLibpodParams object
// with the ability to set a timeout on a request.
func NewPodPauseLibpodParamsWithTimeout(timeout time.Duration) *PodPauseLibpodParams {
	return &PodPauseLibpodParams{
		timeout: timeout,
	}
}

// NewPodPauseLibpodParamsWithContext creates a new PodPauseLibpodParams object
// with the ability to set a context for a request.
func NewPodPauseLibpodParamsWithContext(ctx context.Context) *PodPauseLibpodParams {
	return &PodPauseLibpodParams{
		Context: ctx,
	}
}

// NewPodPauseLibpodParamsWithHTTPClient creates a new PodPauseLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewPodPauseLibpodParamsWithHTTPClient(client *http.Client) *PodPauseLibpodParams {
	return &PodPauseLibpodParams{
		HTTPClient: client,
	}
}

/* PodPauseLibpodParams contains all the parameters to send to the API endpoint
   for the pod pause libpod operation.

   Typically these are written to a http.Request.
*/
type PodPauseLibpodParams struct {

	/* Name.

	   the name or ID of the pod
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pod pause libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodPauseLibpodParams) WithDefaults() *PodPauseLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pod pause libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodPauseLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pod pause libpod params
func (o *PodPauseLibpodParams) WithTimeout(timeout time.Duration) *PodPauseLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pod pause libpod params
func (o *PodPauseLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pod pause libpod params
func (o *PodPauseLibpodParams) WithContext(ctx context.Context) *PodPauseLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pod pause libpod params
func (o *PodPauseLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pod pause libpod params
func (o *PodPauseLibpodParams) WithHTTPClient(client *http.Client) *PodPauseLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pod pause libpod params
func (o *PodPauseLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the pod pause libpod params
func (o *PodPauseLibpodParams) WithName(name string) *PodPauseLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pod pause libpod params
func (o *PodPauseLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PodPauseLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
