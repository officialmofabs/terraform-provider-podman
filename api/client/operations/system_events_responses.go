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

// SystemEventsReader is a Reader for the SystemEvents structure.
type SystemEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSystemEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemEventsOK creates a SystemEventsOK with default headers values
func NewSystemEventsOK() *SystemEventsOK {
	return &SystemEventsOK{}
}

/* SystemEventsOK describes a response with status code 200, with default header values.

returns a string of json data describing an event
*/
type SystemEventsOK struct {
}

func (o *SystemEventsOK) Error() string {
	return fmt.Sprintf("[GET /events][%d] systemEventsOK ", 200)
}

func (o *SystemEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSystemEventsInternalServerError creates a SystemEventsInternalServerError with default headers values
func NewSystemEventsInternalServerError() *SystemEventsInternalServerError {
	return &SystemEventsInternalServerError{}
}

/* SystemEventsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SystemEventsInternalServerError struct {
	Payload *models.SystemEventsInternalServerErrorBody
}

func (o *SystemEventsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /events][%d] systemEventsInternalServerError  %+v", 500, o.Payload)
}
func (o *SystemEventsInternalServerError) GetPayload() *models.SystemEventsInternalServerErrorBody {
	return o.Payload
}

func (o *SystemEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemEventsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
