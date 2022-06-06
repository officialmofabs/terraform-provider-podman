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

// SystemEventsLibpodReader is a Reader for the SystemEventsLibpod structure.
type SystemEventsLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemEventsLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemEventsLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSystemEventsLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemEventsLibpodOK creates a SystemEventsLibpodOK with default headers values
func NewSystemEventsLibpodOK() *SystemEventsLibpodOK {
	return &SystemEventsLibpodOK{}
}

/* SystemEventsLibpodOK describes a response with status code 200, with default header values.

returns a string of json data describing an event
*/
type SystemEventsLibpodOK struct {
}

func (o *SystemEventsLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/events][%d] systemEventsLibpodOK ", 200)
}

func (o *SystemEventsLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSystemEventsLibpodInternalServerError creates a SystemEventsLibpodInternalServerError with default headers values
func NewSystemEventsLibpodInternalServerError() *SystemEventsLibpodInternalServerError {
	return &SystemEventsLibpodInternalServerError{}
}

/* SystemEventsLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SystemEventsLibpodInternalServerError struct {
	Payload *models.SystemEventsLibpodInternalServerErrorBody
}

func (o *SystemEventsLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/events][%d] systemEventsLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *SystemEventsLibpodInternalServerError) GetPayload() *models.SystemEventsLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *SystemEventsLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemEventsLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
