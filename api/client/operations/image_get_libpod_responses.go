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

// ImageGetLibpodReader is a Reader for the ImageGetLibpod structure.
type ImageGetLibpodReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *ImageGetLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageGetLibpodOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImageGetLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageGetLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageGetLibpodOK creates a ImageGetLibpodOK with default headers values
func NewImageGetLibpodOK(writer io.Writer) *ImageGetLibpodOK {
	return &ImageGetLibpodOK{

		Payload: writer,
	}
}

/* ImageGetLibpodOK describes a response with status code 200, with default header values.

no error
*/
type ImageGetLibpodOK struct {
	Payload io.Writer
}

func (o *ImageGetLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/get][%d] imageGetLibpodOK  %+v", 200, o.Payload)
}
func (o *ImageGetLibpodOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *ImageGetLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageGetLibpodNotFound creates a ImageGetLibpodNotFound with default headers values
func NewImageGetLibpodNotFound() *ImageGetLibpodNotFound {
	return &ImageGetLibpodNotFound{}
}

/* ImageGetLibpodNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImageGetLibpodNotFound struct {
	Payload *models.ImageGetLibpodNotFoundBody
}

func (o *ImageGetLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/get][%d] imageGetLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ImageGetLibpodNotFound) GetPayload() *models.ImageGetLibpodNotFoundBody {
	return o.Payload
}

func (o *ImageGetLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageGetLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageGetLibpodInternalServerError creates a ImageGetLibpodInternalServerError with default headers values
func NewImageGetLibpodInternalServerError() *ImageGetLibpodInternalServerError {
	return &ImageGetLibpodInternalServerError{}
}

/* ImageGetLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageGetLibpodInternalServerError struct {
	Payload *models.ImageGetLibpodInternalServerErrorBody
}

func (o *ImageGetLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/get][%d] imageGetLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageGetLibpodInternalServerError) GetPayload() *models.ImageGetLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ImageGetLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageGetLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
