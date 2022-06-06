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

// NewContainerTopLibpodParams creates a new ContainerTopLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerTopLibpodParams() *ContainerTopLibpodParams {
	return &ContainerTopLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerTopLibpodParamsWithTimeout creates a new ContainerTopLibpodParams object
// with the ability to set a timeout on a request.
func NewContainerTopLibpodParamsWithTimeout(timeout time.Duration) *ContainerTopLibpodParams {
	return &ContainerTopLibpodParams{
		timeout: timeout,
	}
}

// NewContainerTopLibpodParamsWithContext creates a new ContainerTopLibpodParams object
// with the ability to set a context for a request.
func NewContainerTopLibpodParamsWithContext(ctx context.Context) *ContainerTopLibpodParams {
	return &ContainerTopLibpodParams{
		Context: ctx,
	}
}

// NewContainerTopLibpodParamsWithHTTPClient creates a new ContainerTopLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerTopLibpodParamsWithHTTPClient(client *http.Client) *ContainerTopLibpodParams {
	return &ContainerTopLibpodParams{
		HTTPClient: client,
	}
}

/* ContainerTopLibpodParams contains all the parameters to send to the API endpoint
   for the container top libpod operation.

   Typically these are written to a http.Request.
*/
type ContainerTopLibpodParams struct {

	/* Delay.

	   if streaming, delay in seconds between updates. Must be >1. (As of version 4.0)

	   Default: 5
	*/
	Delay *int64

	/* Name.

	   Name of container to query for processes (As of version 1.xx)
	*/
	Name string

	/* PsArgs.

	     arguments to pass to ps such as aux.
	Requires ps(1) to be installed in the container if no ps(1) compatible AIX descriptors are used.


	     Default: "-ef"
	*/
	PsArgs *string

	/* Stream.

	   when true, repeatedly stream the latest output (As of version 4.0)
	*/
	Stream *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container top libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerTopLibpodParams) WithDefaults() *ContainerTopLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container top libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerTopLibpodParams) SetDefaults() {
	var (
		delayDefault = int64(5)

		psArgsDefault = string("-ef")
	)

	val := ContainerTopLibpodParams{
		Delay:  &delayDefault,
		PsArgs: &psArgsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the container top libpod params
func (o *ContainerTopLibpodParams) WithTimeout(timeout time.Duration) *ContainerTopLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container top libpod params
func (o *ContainerTopLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container top libpod params
func (o *ContainerTopLibpodParams) WithContext(ctx context.Context) *ContainerTopLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container top libpod params
func (o *ContainerTopLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container top libpod params
func (o *ContainerTopLibpodParams) WithHTTPClient(client *http.Client) *ContainerTopLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container top libpod params
func (o *ContainerTopLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelay adds the delay to the container top libpod params
func (o *ContainerTopLibpodParams) WithDelay(delay *int64) *ContainerTopLibpodParams {
	o.SetDelay(delay)
	return o
}

// SetDelay adds the delay to the container top libpod params
func (o *ContainerTopLibpodParams) SetDelay(delay *int64) {
	o.Delay = delay
}

// WithName adds the name to the container top libpod params
func (o *ContainerTopLibpodParams) WithName(name string) *ContainerTopLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container top libpod params
func (o *ContainerTopLibpodParams) SetName(name string) {
	o.Name = name
}

// WithPsArgs adds the psArgs to the container top libpod params
func (o *ContainerTopLibpodParams) WithPsArgs(psArgs *string) *ContainerTopLibpodParams {
	o.SetPsArgs(psArgs)
	return o
}

// SetPsArgs adds the psArgs to the container top libpod params
func (o *ContainerTopLibpodParams) SetPsArgs(psArgs *string) {
	o.PsArgs = psArgs
}

// WithStream adds the stream to the container top libpod params
func (o *ContainerTopLibpodParams) WithStream(stream *bool) *ContainerTopLibpodParams {
	o.SetStream(stream)
	return o
}

// SetStream adds the stream to the container top libpod params
func (o *ContainerTopLibpodParams) SetStream(stream *bool) {
	o.Stream = stream
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerTopLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Delay != nil {

		// query param delay
		var qrDelay int64

		if o.Delay != nil {
			qrDelay = *o.Delay
		}
		qDelay := swag.FormatInt64(qrDelay)
		if qDelay != "" {

			if err := r.SetQueryParam("delay", qDelay); err != nil {
				return err
			}
		}
	}

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

	if o.Stream != nil {

		// query param stream
		var qrStream bool

		if o.Stream != nil {
			qrStream = *o.Stream
		}
		qStream := swag.FormatBool(qrStream)
		if qStream != "" {

			if err := r.SetQueryParam("stream", qStream); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
