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

// NewPodStatsAllLibpodParams creates a new PodStatsAllLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPodStatsAllLibpodParams() *PodStatsAllLibpodParams {
	return &PodStatsAllLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPodStatsAllLibpodParamsWithTimeout creates a new PodStatsAllLibpodParams object
// with the ability to set a timeout on a request.
func NewPodStatsAllLibpodParamsWithTimeout(timeout time.Duration) *PodStatsAllLibpodParams {
	return &PodStatsAllLibpodParams{
		timeout: timeout,
	}
}

// NewPodStatsAllLibpodParamsWithContext creates a new PodStatsAllLibpodParams object
// with the ability to set a context for a request.
func NewPodStatsAllLibpodParamsWithContext(ctx context.Context) *PodStatsAllLibpodParams {
	return &PodStatsAllLibpodParams{
		Context: ctx,
	}
}

// NewPodStatsAllLibpodParamsWithHTTPClient creates a new PodStatsAllLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewPodStatsAllLibpodParamsWithHTTPClient(client *http.Client) *PodStatsAllLibpodParams {
	return &PodStatsAllLibpodParams{
		HTTPClient: client,
	}
}

/* PodStatsAllLibpodParams contains all the parameters to send to the API endpoint
   for the pod stats all libpod operation.

   Typically these are written to a http.Request.
*/
type PodStatsAllLibpodParams struct {

	/* All.

	   Provide statistics for all running pods.
	*/
	All *bool

	/* NamesOrIDs.

	   Names or IDs of pods.
	*/
	NamesOrIDs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pod stats all libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodStatsAllLibpodParams) WithDefaults() *PodStatsAllLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pod stats all libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodStatsAllLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pod stats all libpod params
func (o *PodStatsAllLibpodParams) WithTimeout(timeout time.Duration) *PodStatsAllLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pod stats all libpod params
func (o *PodStatsAllLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pod stats all libpod params
func (o *PodStatsAllLibpodParams) WithContext(ctx context.Context) *PodStatsAllLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pod stats all libpod params
func (o *PodStatsAllLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pod stats all libpod params
func (o *PodStatsAllLibpodParams) WithHTTPClient(client *http.Client) *PodStatsAllLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pod stats all libpod params
func (o *PodStatsAllLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAll adds the all to the pod stats all libpod params
func (o *PodStatsAllLibpodParams) WithAll(all *bool) *PodStatsAllLibpodParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the pod stats all libpod params
func (o *PodStatsAllLibpodParams) SetAll(all *bool) {
	o.All = all
}

// WithNamesOrIDs adds the namesOrIDs to the pod stats all libpod params
func (o *PodStatsAllLibpodParams) WithNamesOrIDs(namesOrIDs []string) *PodStatsAllLibpodParams {
	o.SetNamesOrIDs(namesOrIDs)
	return o
}

// SetNamesOrIDs adds the namesOrIDs to the pod stats all libpod params
func (o *PodStatsAllLibpodParams) SetNamesOrIDs(namesOrIDs []string) {
	o.NamesOrIDs = namesOrIDs
}

// WriteToRequest writes these params to a swagger request
func (o *PodStatsAllLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.All != nil {

		// query param all
		var qrAll bool

		if o.All != nil {
			qrAll = *o.All
		}
		qAll := swag.FormatBool(qrAll)
		if qAll != "" {

			if err := r.SetQueryParam("all", qAll); err != nil {
				return err
			}
		}
	}

	if o.NamesOrIDs != nil {

		// binding items for namesOrIDs
		joinedNamesOrIDs := o.bindParamNamesOrIDs(reg)

		// query array param namesOrIDs
		if err := r.SetQueryParam("namesOrIDs", joinedNamesOrIDs...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPodStatsAllLibpod binds the parameter namesOrIDs
func (o *PodStatsAllLibpodParams) bindParamNamesOrIDs(formats strfmt.Registry) []string {
	namesOrIDsIR := o.NamesOrIDs

	var namesOrIDsIC []string
	for _, namesOrIDsIIR := range namesOrIDsIR { // explode []string

		namesOrIDsIIV := namesOrIDsIIR // string as string
		namesOrIDsIC = append(namesOrIDsIC, namesOrIDsIIV)
	}

	// items.CollectionFormat: ""
	namesOrIDsIS := swag.JoinByFormat(namesOrIDsIC, "")

	return namesOrIDsIS
}
