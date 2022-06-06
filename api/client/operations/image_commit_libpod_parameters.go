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

// NewImageCommitLibpodParams creates a new ImageCommitLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageCommitLibpodParams() *ImageCommitLibpodParams {
	return &ImageCommitLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageCommitLibpodParamsWithTimeout creates a new ImageCommitLibpodParams object
// with the ability to set a timeout on a request.
func NewImageCommitLibpodParamsWithTimeout(timeout time.Duration) *ImageCommitLibpodParams {
	return &ImageCommitLibpodParams{
		timeout: timeout,
	}
}

// NewImageCommitLibpodParamsWithContext creates a new ImageCommitLibpodParams object
// with the ability to set a context for a request.
func NewImageCommitLibpodParamsWithContext(ctx context.Context) *ImageCommitLibpodParams {
	return &ImageCommitLibpodParams{
		Context: ctx,
	}
}

// NewImageCommitLibpodParamsWithHTTPClient creates a new ImageCommitLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageCommitLibpodParamsWithHTTPClient(client *http.Client) *ImageCommitLibpodParams {
	return &ImageCommitLibpodParams{
		HTTPClient: client,
	}
}

/* ImageCommitLibpodParams contains all the parameters to send to the API endpoint
   for the image commit libpod operation.

   Typically these are written to a http.Request.
*/
type ImageCommitLibpodParams struct {

	/* Author.

	   author of the image
	*/
	Author *string

	/* Changes.

	   instructions to apply while committing in Dockerfile format (i.e. "CMD=/bin/foo")
	*/
	Changes []string

	/* Comment.

	   commit message
	*/
	Comment *string

	/* Container.

	   the name or ID of a container
	*/
	Container string

	/* Format.

	   format of the image manifest and metadata (default "oci")
	*/
	Format *string

	/* Pause.

	   pause the container before committing it
	*/
	Pause *bool

	/* Repo.

	   the repository name for the created image
	*/
	Repo *string

	/* Tag.

	   tag name for the created image
	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image commit libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageCommitLibpodParams) WithDefaults() *ImageCommitLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image commit libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageCommitLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image commit libpod params
func (o *ImageCommitLibpodParams) WithTimeout(timeout time.Duration) *ImageCommitLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image commit libpod params
func (o *ImageCommitLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image commit libpod params
func (o *ImageCommitLibpodParams) WithContext(ctx context.Context) *ImageCommitLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image commit libpod params
func (o *ImageCommitLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image commit libpod params
func (o *ImageCommitLibpodParams) WithHTTPClient(client *http.Client) *ImageCommitLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image commit libpod params
func (o *ImageCommitLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthor adds the author to the image commit libpod params
func (o *ImageCommitLibpodParams) WithAuthor(author *string) *ImageCommitLibpodParams {
	o.SetAuthor(author)
	return o
}

// SetAuthor adds the author to the image commit libpod params
func (o *ImageCommitLibpodParams) SetAuthor(author *string) {
	o.Author = author
}

// WithChanges adds the changes to the image commit libpod params
func (o *ImageCommitLibpodParams) WithChanges(changes []string) *ImageCommitLibpodParams {
	o.SetChanges(changes)
	return o
}

// SetChanges adds the changes to the image commit libpod params
func (o *ImageCommitLibpodParams) SetChanges(changes []string) {
	o.Changes = changes
}

// WithComment adds the comment to the image commit libpod params
func (o *ImageCommitLibpodParams) WithComment(comment *string) *ImageCommitLibpodParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the image commit libpod params
func (o *ImageCommitLibpodParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithContainer adds the container to the image commit libpod params
func (o *ImageCommitLibpodParams) WithContainer(container string) *ImageCommitLibpodParams {
	o.SetContainer(container)
	return o
}

// SetContainer adds the container to the image commit libpod params
func (o *ImageCommitLibpodParams) SetContainer(container string) {
	o.Container = container
}

// WithFormat adds the format to the image commit libpod params
func (o *ImageCommitLibpodParams) WithFormat(format *string) *ImageCommitLibpodParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the image commit libpod params
func (o *ImageCommitLibpodParams) SetFormat(format *string) {
	o.Format = format
}

// WithPause adds the pause to the image commit libpod params
func (o *ImageCommitLibpodParams) WithPause(pause *bool) *ImageCommitLibpodParams {
	o.SetPause(pause)
	return o
}

// SetPause adds the pause to the image commit libpod params
func (o *ImageCommitLibpodParams) SetPause(pause *bool) {
	o.Pause = pause
}

// WithRepo adds the repo to the image commit libpod params
func (o *ImageCommitLibpodParams) WithRepo(repo *string) *ImageCommitLibpodParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the image commit libpod params
func (o *ImageCommitLibpodParams) SetRepo(repo *string) {
	o.Repo = repo
}

// WithTag adds the tag to the image commit libpod params
func (o *ImageCommitLibpodParams) WithTag(tag *string) *ImageCommitLibpodParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the image commit libpod params
func (o *ImageCommitLibpodParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *ImageCommitLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Author != nil {

		// query param author
		var qrAuthor string

		if o.Author != nil {
			qrAuthor = *o.Author
		}
		qAuthor := qrAuthor
		if qAuthor != "" {

			if err := r.SetQueryParam("author", qAuthor); err != nil {
				return err
			}
		}
	}

	if o.Changes != nil {

		// binding items for changes
		joinedChanges := o.bindParamChanges(reg)

		// query array param changes
		if err := r.SetQueryParam("changes", joinedChanges...); err != nil {
			return err
		}
	}

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	// query param container
	qrContainer := o.Container
	qContainer := qrContainer
	if qContainer != "" {

		if err := r.SetQueryParam("container", qContainer); err != nil {
			return err
		}
	}

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	if o.Pause != nil {

		// query param pause
		var qrPause bool

		if o.Pause != nil {
			qrPause = *o.Pause
		}
		qPause := swag.FormatBool(qrPause)
		if qPause != "" {

			if err := r.SetQueryParam("pause", qPause); err != nil {
				return err
			}
		}
	}

	if o.Repo != nil {

		// query param repo
		var qrRepo string

		if o.Repo != nil {
			qrRepo = *o.Repo
		}
		qRepo := qrRepo
		if qRepo != "" {

			if err := r.SetQueryParam("repo", qRepo); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamImageCommitLibpod binds the parameter changes
func (o *ImageCommitLibpodParams) bindParamChanges(formats strfmt.Registry) []string {
	changesIR := o.Changes

	var changesIC []string
	for _, changesIIR := range changesIR { // explode []string

		changesIIV := changesIIR // string as string
		changesIC = append(changesIC, changesIIV)
	}

	// items.CollectionFormat: ""
	changesIS := swag.JoinByFormat(changesIC, "")

	return changesIS
}
