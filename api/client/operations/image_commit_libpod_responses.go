// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/project0/terraform-provider-podman/api/models"
)

// ImageCommitLibpodReader is a Reader for the ImageCommitLibpod structure.
type ImageCommitLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageCommitLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewImageCommitLibpodCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImageCommitLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageCommitLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageCommitLibpodCreated creates a ImageCommitLibpodCreated with default headers values
func NewImageCommitLibpodCreated() *ImageCommitLibpodCreated {
	return &ImageCommitLibpodCreated{}
}

/* ImageCommitLibpodCreated describes a response with status code 201, with default header values.

no error
*/
type ImageCommitLibpodCreated struct {
}

func (o *ImageCommitLibpodCreated) Error() string {
	return fmt.Sprintf("[POST /libpod/commit][%d] imageCommitLibpodCreated ", 201)
}

func (o *ImageCommitLibpodCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImageCommitLibpodNotFound creates a ImageCommitLibpodNotFound with default headers values
func NewImageCommitLibpodNotFound() *ImageCommitLibpodNotFound {
	return &ImageCommitLibpodNotFound{}
}

/* ImageCommitLibpodNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImageCommitLibpodNotFound struct {
	Payload *models.ImageCommitLibpodNotFoundBody
}

func (o *ImageCommitLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/commit][%d] imageCommitLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ImageCommitLibpodNotFound) GetPayload() *models.ImageCommitLibpodNotFoundBody {
	return o.Payload
}

func (o *ImageCommitLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageCommitLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageCommitLibpodInternalServerError creates a ImageCommitLibpodInternalServerError with default headers values
func NewImageCommitLibpodInternalServerError() *ImageCommitLibpodInternalServerError {
	return &ImageCommitLibpodInternalServerError{}
}

/* ImageCommitLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageCommitLibpodInternalServerError struct {
	Payload *models.ImageCommitLibpodInternalServerErrorBody
}

func (o *ImageCommitLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/commit][%d] imageCommitLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageCommitLibpodInternalServerError) GetPayload() *models.ImageCommitLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ImageCommitLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageCommitLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
