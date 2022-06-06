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

// NewImageDeleteAllLibpodParams creates a new ImageDeleteAllLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageDeleteAllLibpodParams() *ImageDeleteAllLibpodParams {
	return &ImageDeleteAllLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageDeleteAllLibpodParamsWithTimeout creates a new ImageDeleteAllLibpodParams object
// with the ability to set a timeout on a request.
func NewImageDeleteAllLibpodParamsWithTimeout(timeout time.Duration) *ImageDeleteAllLibpodParams {
	return &ImageDeleteAllLibpodParams{
		timeout: timeout,
	}
}

// NewImageDeleteAllLibpodParamsWithContext creates a new ImageDeleteAllLibpodParams object
// with the ability to set a context for a request.
func NewImageDeleteAllLibpodParamsWithContext(ctx context.Context) *ImageDeleteAllLibpodParams {
	return &ImageDeleteAllLibpodParams{
		Context: ctx,
	}
}

// NewImageDeleteAllLibpodParamsWithHTTPClient creates a new ImageDeleteAllLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageDeleteAllLibpodParamsWithHTTPClient(client *http.Client) *ImageDeleteAllLibpodParams {
	return &ImageDeleteAllLibpodParams{
		HTTPClient: client,
	}
}

/* ImageDeleteAllLibpodParams contains all the parameters to send to the API endpoint
   for the image delete all libpod operation.

   Typically these are written to a http.Request.
*/
type ImageDeleteAllLibpodParams struct {

	/* All.

	   Remove all images.

	   Default: true
	*/
	All *bool

	/* Force.

	   Force image removal (including containers using the images).
	*/
	Force *bool

	/* Ignore.

	   Ignore if a specified image does not exist and do not throw an error.
	*/
	Ignore *bool

	/* Images.

	   Images IDs or names to remove.
	*/
	Images []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image delete all libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageDeleteAllLibpodParams) WithDefaults() *ImageDeleteAllLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image delete all libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageDeleteAllLibpodParams) SetDefaults() {
	var (
		allDefault = bool(true)
	)

	val := ImageDeleteAllLibpodParams{
		All: &allDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) WithTimeout(timeout time.Duration) *ImageDeleteAllLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) WithContext(ctx context.Context) *ImageDeleteAllLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) WithHTTPClient(client *http.Client) *ImageDeleteAllLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAll adds the all to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) WithAll(all *bool) *ImageDeleteAllLibpodParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) SetAll(all *bool) {
	o.All = all
}

// WithForce adds the force to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) WithForce(force *bool) *ImageDeleteAllLibpodParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) SetForce(force *bool) {
	o.Force = force
}

// WithIgnore adds the ignore to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) WithIgnore(ignore *bool) *ImageDeleteAllLibpodParams {
	o.SetIgnore(ignore)
	return o
}

// SetIgnore adds the ignore to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) SetIgnore(ignore *bool) {
	o.Ignore = ignore
}

// WithImages adds the images to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) WithImages(images []string) *ImageDeleteAllLibpodParams {
	o.SetImages(images)
	return o
}

// SetImages adds the images to the image delete all libpod params
func (o *ImageDeleteAllLibpodParams) SetImages(images []string) {
	o.Images = images
}

// WriteToRequest writes these params to a swagger request
func (o *ImageDeleteAllLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Ignore != nil {

		// query param ignore
		var qrIgnore bool

		if o.Ignore != nil {
			qrIgnore = *o.Ignore
		}
		qIgnore := swag.FormatBool(qrIgnore)
		if qIgnore != "" {

			if err := r.SetQueryParam("ignore", qIgnore); err != nil {
				return err
			}
		}
	}

	if o.Images != nil {

		// binding items for images
		joinedImages := o.bindParamImages(reg)

		// query array param images
		if err := r.SetQueryParam("images", joinedImages...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamImageDeleteAllLibpod binds the parameter images
func (o *ImageDeleteAllLibpodParams) bindParamImages(formats strfmt.Registry) []string {
	imagesIR := o.Images

	var imagesIC []string
	for _, imagesIIR := range imagesIR { // explode []string

		imagesIIV := imagesIIR // string as string
		imagesIC = append(imagesIC, imagesIIV)
	}

	// items.CollectionFormat: ""
	imagesIS := swag.JoinByFormat(imagesIC, "")

	return imagesIS
}
