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

// NewContainerListParams creates a new ContainerListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerListParams() *ContainerListParams {
	return &ContainerListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerListParamsWithTimeout creates a new ContainerListParams object
// with the ability to set a timeout on a request.
func NewContainerListParamsWithTimeout(timeout time.Duration) *ContainerListParams {
	return &ContainerListParams{
		timeout: timeout,
	}
}

// NewContainerListParamsWithContext creates a new ContainerListParams object
// with the ability to set a context for a request.
func NewContainerListParamsWithContext(ctx context.Context) *ContainerListParams {
	return &ContainerListParams{
		Context: ctx,
	}
}

// NewContainerListParamsWithHTTPClient creates a new ContainerListParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerListParamsWithHTTPClient(client *http.Client) *ContainerListParams {
	return &ContainerListParams{
		HTTPClient: client,
	}
}

/* ContainerListParams contains all the parameters to send to the API endpoint
   for the container list operation.

   Typically these are written to a http.Request.
*/
type ContainerListParams struct {

	/* All.

	   Return all containers. By default, only running containers are shown
	*/
	All *bool

	/* External.

	   Return containers in storage not controlled by Podman
	*/
	External *bool

	/* Filters.

	    Returns a list of containers.
	- ancestor=(<image-name>[:<tag>], <image id>, or <image@digest>)
	- before=(<container id> or <container name>)
	- expose=(<port>[/<proto>]|<startport-endport>/[<proto>])
	- exited=<int> containers with exit code of <int>
	- health=(starting|healthy|unhealthy|none)
	- id=<ID> a container's ID
	- is-task=(true|false)
	- label=key or label="key=value" of a container label
	- name=<name> a container's name
	- network=(<network id> or <network name>)
	- publish=(<port>[/<proto>]|<startport-endport>/[<proto>])
	- since=(<container id> or <container name>)
	- status=(created|restarting|running|removing|paused|exited|dead)
	- volume=(<volume name> or <mount point destination>)

	*/
	Filters *string

	/* Limit.

	   Return this number of most recently created containers, including non-running ones.
	*/
	Limit *int64

	/* Size.

	   Return the size of container as fields SizeRw and SizeRootFs.
	*/
	Size *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerListParams) WithDefaults() *ContainerListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerListParams) SetDefaults() {
	var (
		allDefault = bool(false)

		externalDefault = bool(false)

		sizeDefault = bool(false)
	)

	val := ContainerListParams{
		All:      &allDefault,
		External: &externalDefault,
		Size:     &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the container list params
func (o *ContainerListParams) WithTimeout(timeout time.Duration) *ContainerListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container list params
func (o *ContainerListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container list params
func (o *ContainerListParams) WithContext(ctx context.Context) *ContainerListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container list params
func (o *ContainerListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container list params
func (o *ContainerListParams) WithHTTPClient(client *http.Client) *ContainerListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container list params
func (o *ContainerListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAll adds the all to the container list params
func (o *ContainerListParams) WithAll(all *bool) *ContainerListParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the container list params
func (o *ContainerListParams) SetAll(all *bool) {
	o.All = all
}

// WithExternal adds the external to the container list params
func (o *ContainerListParams) WithExternal(external *bool) *ContainerListParams {
	o.SetExternal(external)
	return o
}

// SetExternal adds the external to the container list params
func (o *ContainerListParams) SetExternal(external *bool) {
	o.External = external
}

// WithFilters adds the filters to the container list params
func (o *ContainerListParams) WithFilters(filters *string) *ContainerListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the container list params
func (o *ContainerListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithLimit adds the limit to the container list params
func (o *ContainerListParams) WithLimit(limit *int64) *ContainerListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the container list params
func (o *ContainerListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithSize adds the size to the container list params
func (o *ContainerListParams) WithSize(size *bool) *ContainerListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the container list params
func (o *ContainerListParams) SetSize(size *bool) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.External != nil {

		// query param external
		var qrExternal bool

		if o.External != nil {
			qrExternal = *o.External
		}
		qExternal := swag.FormatBool(qrExternal)
		if qExternal != "" {

			if err := r.SetQueryParam("external", qExternal); err != nil {
				return err
			}
		}
	}

	if o.Filters != nil {

		// query param filters
		var qrFilters string

		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {

			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize bool

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatBool(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
