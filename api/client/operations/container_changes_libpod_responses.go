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

// ContainerChangesLibpodReader is a Reader for the ContainerChangesLibpod structure.
type ContainerChangesLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerChangesLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerChangesLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerChangesLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerChangesLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerChangesLibpodOK creates a ContainerChangesLibpodOK with default headers values
func NewContainerChangesLibpodOK() *ContainerChangesLibpodOK {
	return &ContainerChangesLibpodOK{}
}

/* ContainerChangesLibpodOK describes a response with status code 200, with default header values.

Array of Changes
*/
type ContainerChangesLibpodOK struct {
}

func (o *ContainerChangesLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/changes][%d] containerChangesLibpodOK ", 200)
}

func (o *ContainerChangesLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerChangesLibpodNotFound creates a ContainerChangesLibpodNotFound with default headers values
func NewContainerChangesLibpodNotFound() *ContainerChangesLibpodNotFound {
	return &ContainerChangesLibpodNotFound{}
}

/* ContainerChangesLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerChangesLibpodNotFound struct {
	Payload *models.ContainerChangesLibpodNotFoundBody
}

func (o *ContainerChangesLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/changes][%d] containerChangesLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ContainerChangesLibpodNotFound) GetPayload() *models.ContainerChangesLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerChangesLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerChangesLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerChangesLibpodInternalServerError creates a ContainerChangesLibpodInternalServerError with default headers values
func NewContainerChangesLibpodInternalServerError() *ContainerChangesLibpodInternalServerError {
	return &ContainerChangesLibpodInternalServerError{}
}

/* ContainerChangesLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerChangesLibpodInternalServerError struct {
	Payload *models.ContainerChangesLibpodInternalServerErrorBody
}

func (o *ContainerChangesLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/changes][%d] containerChangesLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerChangesLibpodInternalServerError) GetPayload() *models.ContainerChangesLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerChangesLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerChangesLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
