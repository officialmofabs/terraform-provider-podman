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

// ContainerWaitLibpodReader is a Reader for the ContainerWaitLibpod structure.
type ContainerWaitLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerWaitLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerWaitLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerWaitLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerWaitLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerWaitLibpodOK creates a ContainerWaitLibpodOK with default headers values
func NewContainerWaitLibpodOK() *ContainerWaitLibpodOK {
	return &ContainerWaitLibpodOK{}
}

/* ContainerWaitLibpodOK describes a response with status code 200, with default header values.

Status code
*/
type ContainerWaitLibpodOK struct {
	Payload int32
}

func (o *ContainerWaitLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/wait][%d] containerWaitLibpodOK  %+v", 200, o.Payload)
}
func (o *ContainerWaitLibpodOK) GetPayload() int32 {
	return o.Payload
}

func (o *ContainerWaitLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerWaitLibpodNotFound creates a ContainerWaitLibpodNotFound with default headers values
func NewContainerWaitLibpodNotFound() *ContainerWaitLibpodNotFound {
	return &ContainerWaitLibpodNotFound{}
}

/* ContainerWaitLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerWaitLibpodNotFound struct {
	Payload *models.ContainerWaitLibpodNotFoundBody
}

func (o *ContainerWaitLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/wait][%d] containerWaitLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ContainerWaitLibpodNotFound) GetPayload() *models.ContainerWaitLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerWaitLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerWaitLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerWaitLibpodInternalServerError creates a ContainerWaitLibpodInternalServerError with default headers values
func NewContainerWaitLibpodInternalServerError() *ContainerWaitLibpodInternalServerError {
	return &ContainerWaitLibpodInternalServerError{}
}

/* ContainerWaitLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerWaitLibpodInternalServerError struct {
	Payload *models.ContainerWaitLibpodInternalServerErrorBody
}

func (o *ContainerWaitLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/wait][%d] containerWaitLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerWaitLibpodInternalServerError) GetPayload() *models.ContainerWaitLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerWaitLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerWaitLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
