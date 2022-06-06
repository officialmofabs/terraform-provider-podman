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

// ImagePushLibpodReader is a Reader for the ImagePushLibpod structure.
type ImagePushLibpodReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *ImagePushLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagePushLibpodOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImagePushLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagePushLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagePushLibpodOK creates a ImagePushLibpodOK with default headers values
func NewImagePushLibpodOK(writer io.Writer) *ImagePushLibpodOK {
	return &ImagePushLibpodOK{

		Payload: writer,
	}
}

/* ImagePushLibpodOK describes a response with status code 200, with default header values.

no error
*/
type ImagePushLibpodOK struct {
	Payload io.Writer
}

func (o *ImagePushLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/push][%d] imagePushLibpodOK  %+v", 200, o.Payload)
}
func (o *ImagePushLibpodOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *ImagePushLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagePushLibpodNotFound creates a ImagePushLibpodNotFound with default headers values
func NewImagePushLibpodNotFound() *ImagePushLibpodNotFound {
	return &ImagePushLibpodNotFound{}
}

/* ImagePushLibpodNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImagePushLibpodNotFound struct {
	Payload *models.ImagePushLibpodNotFoundBody
}

func (o *ImagePushLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/push][%d] imagePushLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ImagePushLibpodNotFound) GetPayload() *models.ImagePushLibpodNotFoundBody {
	return o.Payload
}

func (o *ImagePushLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImagePushLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagePushLibpodInternalServerError creates a ImagePushLibpodInternalServerError with default headers values
func NewImagePushLibpodInternalServerError() *ImagePushLibpodInternalServerError {
	return &ImagePushLibpodInternalServerError{}
}

/* ImagePushLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImagePushLibpodInternalServerError struct {
	Payload *models.ImagePushLibpodInternalServerErrorBody
}

func (o *ImagePushLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/push][%d] imagePushLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ImagePushLibpodInternalServerError) GetPayload() *models.ImagePushLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ImagePushLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImagePushLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
