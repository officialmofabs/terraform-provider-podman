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

// ImageCommitReader is a Reader for the ImageCommit structure.
type ImageCommitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageCommitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewImageCommitCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImageCommitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageCommitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageCommitCreated creates a ImageCommitCreated with default headers values
func NewImageCommitCreated() *ImageCommitCreated {
	return &ImageCommitCreated{}
}

/* ImageCommitCreated describes a response with status code 201, with default header values.

no error
*/
type ImageCommitCreated struct {
}

func (o *ImageCommitCreated) Error() string {
	return fmt.Sprintf("[POST /commit][%d] imageCommitCreated ", 201)
}

func (o *ImageCommitCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImageCommitNotFound creates a ImageCommitNotFound with default headers values
func NewImageCommitNotFound() *ImageCommitNotFound {
	return &ImageCommitNotFound{}
}

/* ImageCommitNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImageCommitNotFound struct {
	Payload *models.ImageCommitNotFoundBody
}

func (o *ImageCommitNotFound) Error() string {
	return fmt.Sprintf("[POST /commit][%d] imageCommitNotFound  %+v", 404, o.Payload)
}
func (o *ImageCommitNotFound) GetPayload() *models.ImageCommitNotFoundBody {
	return o.Payload
}

func (o *ImageCommitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageCommitNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageCommitInternalServerError creates a ImageCommitInternalServerError with default headers values
func NewImageCommitInternalServerError() *ImageCommitInternalServerError {
	return &ImageCommitInternalServerError{}
}

/* ImageCommitInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageCommitInternalServerError struct {
	Payload *models.ImageCommitInternalServerErrorBody
}

func (o *ImageCommitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /commit][%d] imageCommitInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageCommitInternalServerError) GetPayload() *models.ImageCommitInternalServerErrorBody {
	return o.Payload
}

func (o *ImageCommitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageCommitInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
