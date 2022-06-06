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

// ContainerLogsReader is a Reader for the ContainerLogs structure.
type ContainerLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerLogsOK creates a ContainerLogsOK with default headers values
func NewContainerLogsOK() *ContainerLogsOK {
	return &ContainerLogsOK{}
}

/* ContainerLogsOK describes a response with status code 200, with default header values.

logs returned as a stream in response body.
*/
type ContainerLogsOK struct {
}

func (o *ContainerLogsOK) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/logs][%d] containerLogsOK ", 200)
}

func (o *ContainerLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerLogsNotFound creates a ContainerLogsNotFound with default headers values
func NewContainerLogsNotFound() *ContainerLogsNotFound {
	return &ContainerLogsNotFound{}
}

/* ContainerLogsNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerLogsNotFound struct {
	Payload *models.ContainerLogsNotFoundBody
}

func (o *ContainerLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/logs][%d] containerLogsNotFound  %+v", 404, o.Payload)
}
func (o *ContainerLogsNotFound) GetPayload() *models.ContainerLogsNotFoundBody {
	return o.Payload
}

func (o *ContainerLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerLogsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerLogsInternalServerError creates a ContainerLogsInternalServerError with default headers values
func NewContainerLogsInternalServerError() *ContainerLogsInternalServerError {
	return &ContainerLogsInternalServerError{}
}

/* ContainerLogsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerLogsInternalServerError struct {
	Payload *models.ContainerLogsInternalServerErrorBody
}

func (o *ContainerLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/logs][%d] containerLogsInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerLogsInternalServerError) GetPayload() *models.ContainerLogsInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerLogsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
