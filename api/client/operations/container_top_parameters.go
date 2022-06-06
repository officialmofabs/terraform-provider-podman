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

// NewContainerTopParams creates a new ContainerTopParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerTopParams() *ContainerTopParams {
	return &ContainerTopParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerTopParamsWithTimeout creates a new ContainerTopParams object
// with the ability to set a timeout on a request.
func NewContainerTopParamsWithTimeout(timeout time.Duration) *ContainerTopParams {
	return &ContainerTopParams{
		timeout: timeout,
	}
}

// NewContainerTopParamsWithContext creates a new ContainerTopParams object
// with the ability to set a context for a request.
func NewContainerTopParamsWithContext(ctx context.Context) *ContainerTopParams {
	return &ContainerTopParams{
		Context: ctx,
	}
}

// NewContainerTopParamsWithHTTPClient creates a new ContainerTopParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerTopParamsWithHTTPClient(client *http.Client) *ContainerTopParams {
	return &ContainerTopParams{
		HTTPClient: client,
	}
}

/* ContainerTopParams contains all the parameters to send to the API endpoint
   for the container top operation.

   Typically these are written to a http.Request.
*/
type ContainerTopParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	/* PsArgs.

	   arguments to pass to ps such as aux. Requires ps(1) to be installed in the container if no ps(1) compatible AIX descriptors are used.

	   Default: "-ef"
	*/
	PsArgs *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container top params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerTopParams) WithDefaults() *ContainerTopParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container top params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerTopParams) SetDefaults() {
	var (
		psArgsDefault = string("-ef")
	)

	val := ContainerTopParams{
		PsArgs: &psArgsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the container top params
func (o *ContainerTopParams) WithTimeout(timeout time.Duration) *ContainerTopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container top params
func (o *ContainerTopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container top params
func (o *ContainerTopParams) WithContext(ctx context.Context) *ContainerTopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container top params
func (o *ContainerTopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container top params
func (o *ContainerTopParams) WithHTTPClient(client *http.Client) *ContainerTopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container top params
func (o *ContainerTopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the container top params
func (o *ContainerTopParams) WithName(name string) *ContainerTopParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container top params
func (o *ContainerTopParams) SetName(name string) {
	o.Name = name
}

// WithPsArgs adds the psArgs to the container top params
func (o *ContainerTopParams) WithPsArgs(psArgs *string) *ContainerTopParams {
	o.SetPsArgs(psArgs)
	return o
}

// SetPsArgs adds the psArgs to the container top params
func (o *ContainerTopParams) SetPsArgs(psArgs *string) {
	o.PsArgs = psArgs
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerTopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.PsArgs != nil {

		// query param ps_args
		var qrPsArgs string

		if o.PsArgs != nil {
			qrPsArgs = *o.PsArgs
		}
		qPsArgs := qrPsArgs
		if qPsArgs != "" {

			if err := r.SetQueryParam("ps_args", qPsArgs); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
