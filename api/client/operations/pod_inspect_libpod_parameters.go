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

// NewPodInspectLibpodParams creates a new PodInspectLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPodInspectLibpodParams() *PodInspectLibpodParams {
	return &PodInspectLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPodInspectLibpodParamsWithTimeout creates a new PodInspectLibpodParams object
// with the ability to set a timeout on a request.
func NewPodInspectLibpodParamsWithTimeout(timeout time.Duration) *PodInspectLibpodParams {
	return &PodInspectLibpodParams{
		timeout: timeout,
	}
}

// NewPodInspectLibpodParamsWithContext creates a new PodInspectLibpodParams object
// with the ability to set a context for a request.
func NewPodInspectLibpodParamsWithContext(ctx context.Context) *PodInspectLibpodParams {
	return &PodInspectLibpodParams{
		Context: ctx,
	}
}

// NewPodInspectLibpodParamsWithHTTPClient creates a new PodInspectLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewPodInspectLibpodParamsWithHTTPClient(client *http.Client) *PodInspectLibpodParams {
	return &PodInspectLibpodParams{
		HTTPClient: client,
	}
}

/* PodInspectLibpodParams contains all the parameters to send to the API endpoint
   for the pod inspect libpod operation.

   Typically these are written to a http.Request.
*/
type PodInspectLibpodParams struct {

	/* Name.

	   the name or ID of the pod
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pod inspect libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodInspectLibpodParams) WithDefaults() *PodInspectLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pod inspect libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodInspectLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pod inspect libpod params
func (o *PodInspectLibpodParams) WithTimeout(timeout time.Duration) *PodInspectLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pod inspect libpod params
func (o *PodInspectLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pod inspect libpod params
func (o *PodInspectLibpodParams) WithContext(ctx context.Context) *PodInspectLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pod inspect libpod params
func (o *PodInspectLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pod inspect libpod params
func (o *PodInspectLibpodParams) WithHTTPClient(client *http.Client) *PodInspectLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pod inspect libpod params
func (o *PodInspectLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the pod inspect libpod params
func (o *PodInspectLibpodParams) WithName(name string) *PodInspectLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pod inspect libpod params
func (o *PodInspectLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PodInspectLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
