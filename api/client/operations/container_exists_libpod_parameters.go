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

// NewContainerExistsLibpodParams creates a new ContainerExistsLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerExistsLibpodParams() *ContainerExistsLibpodParams {
	return &ContainerExistsLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerExistsLibpodParamsWithTimeout creates a new ContainerExistsLibpodParams object
// with the ability to set a timeout on a request.
func NewContainerExistsLibpodParamsWithTimeout(timeout time.Duration) *ContainerExistsLibpodParams {
	return &ContainerExistsLibpodParams{
		timeout: timeout,
	}
}

// NewContainerExistsLibpodParamsWithContext creates a new ContainerExistsLibpodParams object
// with the ability to set a context for a request.
func NewContainerExistsLibpodParamsWithContext(ctx context.Context) *ContainerExistsLibpodParams {
	return &ContainerExistsLibpodParams{
		Context: ctx,
	}
}

// NewContainerExistsLibpodParamsWithHTTPClient creates a new ContainerExistsLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerExistsLibpodParamsWithHTTPClient(client *http.Client) *ContainerExistsLibpodParams {
	return &ContainerExistsLibpodParams{
		HTTPClient: client,
	}
}

/* ContainerExistsLibpodParams contains all the parameters to send to the API endpoint
   for the container exists libpod operation.

   Typically these are written to a http.Request.
*/
type ContainerExistsLibpodParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container exists libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerExistsLibpodParams) WithDefaults() *ContainerExistsLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container exists libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerExistsLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container exists libpod params
func (o *ContainerExistsLibpodParams) WithTimeout(timeout time.Duration) *ContainerExistsLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container exists libpod params
func (o *ContainerExistsLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container exists libpod params
func (o *ContainerExistsLibpodParams) WithContext(ctx context.Context) *ContainerExistsLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container exists libpod params
func (o *ContainerExistsLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container exists libpod params
func (o *ContainerExistsLibpodParams) WithHTTPClient(client *http.Client) *ContainerExistsLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container exists libpod params
func (o *ContainerExistsLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the container exists libpod params
func (o *ContainerExistsLibpodParams) WithName(name string) *ContainerExistsLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container exists libpod params
func (o *ContainerExistsLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerExistsLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
