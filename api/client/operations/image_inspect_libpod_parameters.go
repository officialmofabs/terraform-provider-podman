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

// NewImageInspectLibpodParams creates a new ImageInspectLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageInspectLibpodParams() *ImageInspectLibpodParams {
	return &ImageInspectLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageInspectLibpodParamsWithTimeout creates a new ImageInspectLibpodParams object
// with the ability to set a timeout on a request.
func NewImageInspectLibpodParamsWithTimeout(timeout time.Duration) *ImageInspectLibpodParams {
	return &ImageInspectLibpodParams{
		timeout: timeout,
	}
}

// NewImageInspectLibpodParamsWithContext creates a new ImageInspectLibpodParams object
// with the ability to set a context for a request.
func NewImageInspectLibpodParamsWithContext(ctx context.Context) *ImageInspectLibpodParams {
	return &ImageInspectLibpodParams{
		Context: ctx,
	}
}

// NewImageInspectLibpodParamsWithHTTPClient creates a new ImageInspectLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageInspectLibpodParamsWithHTTPClient(client *http.Client) *ImageInspectLibpodParams {
	return &ImageInspectLibpodParams{
		HTTPClient: client,
	}
}

/* ImageInspectLibpodParams contains all the parameters to send to the API endpoint
   for the image inspect libpod operation.

   Typically these are written to a http.Request.
*/
type ImageInspectLibpodParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image inspect libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageInspectLibpodParams) WithDefaults() *ImageInspectLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image inspect libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageInspectLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image inspect libpod params
func (o *ImageInspectLibpodParams) WithTimeout(timeout time.Duration) *ImageInspectLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image inspect libpod params
func (o *ImageInspectLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image inspect libpod params
func (o *ImageInspectLibpodParams) WithContext(ctx context.Context) *ImageInspectLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image inspect libpod params
func (o *ImageInspectLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image inspect libpod params
func (o *ImageInspectLibpodParams) WithHTTPClient(client *http.Client) *ImageInspectLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image inspect libpod params
func (o *ImageInspectLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the image inspect libpod params
func (o *ImageInspectLibpodParams) WithName(name string) *ImageInspectLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the image inspect libpod params
func (o *ImageInspectLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ImageInspectLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
