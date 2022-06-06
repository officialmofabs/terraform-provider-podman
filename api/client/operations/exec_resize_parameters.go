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

// NewExecResizeParams creates a new ExecResizeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecResizeParams() *ExecResizeParams {
	return &ExecResizeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecResizeParamsWithTimeout creates a new ExecResizeParams object
// with the ability to set a timeout on a request.
func NewExecResizeParamsWithTimeout(timeout time.Duration) *ExecResizeParams {
	return &ExecResizeParams{
		timeout: timeout,
	}
}

// NewExecResizeParamsWithContext creates a new ExecResizeParams object
// with the ability to set a context for a request.
func NewExecResizeParamsWithContext(ctx context.Context) *ExecResizeParams {
	return &ExecResizeParams{
		Context: ctx,
	}
}

// NewExecResizeParamsWithHTTPClient creates a new ExecResizeParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecResizeParamsWithHTTPClient(client *http.Client) *ExecResizeParams {
	return &ExecResizeParams{
		HTTPClient: client,
	}
}

/* ExecResizeParams contains all the parameters to send to the API endpoint
   for the exec resize operation.

   Typically these are written to a http.Request.
*/
type ExecResizeParams struct {

	/* H.

	   Height of the TTY session in characters
	*/
	H *int64

	/* ID.

	   Exec instance ID
	*/
	ID string

	/* Running.

	   Ignore containers not running errors
	*/
	Running *bool

	/* W.

	   Width of the TTY session in characters
	*/
	W *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the exec resize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecResizeParams) WithDefaults() *ExecResizeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the exec resize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecResizeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the exec resize params
func (o *ExecResizeParams) WithTimeout(timeout time.Duration) *ExecResizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the exec resize params
func (o *ExecResizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the exec resize params
func (o *ExecResizeParams) WithContext(ctx context.Context) *ExecResizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the exec resize params
func (o *ExecResizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the exec resize params
func (o *ExecResizeParams) WithHTTPClient(client *http.Client) *ExecResizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the exec resize params
func (o *ExecResizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithH adds the h to the exec resize params
func (o *ExecResizeParams) WithH(h *int64) *ExecResizeParams {
	o.SetH(h)
	return o
}

// SetH adds the h to the exec resize params
func (o *ExecResizeParams) SetH(h *int64) {
	o.H = h
}

// WithID adds the id to the exec resize params
func (o *ExecResizeParams) WithID(id string) *ExecResizeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the exec resize params
func (o *ExecResizeParams) SetID(id string) {
	o.ID = id
}

// WithRunning adds the running to the exec resize params
func (o *ExecResizeParams) WithRunning(running *bool) *ExecResizeParams {
	o.SetRunning(running)
	return o
}

// SetRunning adds the running to the exec resize params
func (o *ExecResizeParams) SetRunning(running *bool) {
	o.Running = running
}

// WithW adds the w to the exec resize params
func (o *ExecResizeParams) WithW(w *int64) *ExecResizeParams {
	o.SetW(w)
	return o
}

// SetW adds the w to the exec resize params
func (o *ExecResizeParams) SetW(w *int64) {
	o.W = w
}

// WriteToRequest writes these params to a swagger request
func (o *ExecResizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.H != nil {

		// query param h
		var qrH int64

		if o.H != nil {
			qrH = *o.H
		}
		qH := swag.FormatInt64(qrH)
		if qH != "" {

			if err := r.SetQueryParam("h", qH); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Running != nil {

		// query param running
		var qrRunning bool

		if o.Running != nil {
			qrRunning = *o.Running
		}
		qRunning := swag.FormatBool(qrRunning)
		if qRunning != "" {

			if err := r.SetQueryParam("running", qRunning); err != nil {
				return err
			}
		}
	}

	if o.W != nil {

		// query param w
		var qrW int64

		if o.W != nil {
			qrW = *o.W
		}
		qW := swag.FormatInt64(qrW)
		if qW != "" {

			if err := r.SetQueryParam("w", qW); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
