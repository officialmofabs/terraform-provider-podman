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

// ImageLoadReader is a Reader for the ImageLoad structure.
type ImageLoadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageLoadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageLoadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewImageLoadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageLoadOK creates a ImageLoadOK with default headers values
func NewImageLoadOK() *ImageLoadOK {
	return &ImageLoadOK{}
}

/* ImageLoadOK describes a response with status code 200, with default header values.

no error
*/
type ImageLoadOK struct {
}

func (o *ImageLoadOK) Error() string {
	return fmt.Sprintf("[POST /images/load][%d] imageLoadOK ", 200)
}

func (o *ImageLoadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImageLoadInternalServerError creates a ImageLoadInternalServerError with default headers values
func NewImageLoadInternalServerError() *ImageLoadInternalServerError {
	return &ImageLoadInternalServerError{}
}

/* ImageLoadInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageLoadInternalServerError struct {
	Payload *models.ImageLoadInternalServerErrorBody
}

func (o *ImageLoadInternalServerError) Error() string {
	return fmt.Sprintf("[POST /images/load][%d] imageLoadInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageLoadInternalServerError) GetPayload() *models.ImageLoadInternalServerErrorBody {
	return o.Payload
}

func (o *ImageLoadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageLoadInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
