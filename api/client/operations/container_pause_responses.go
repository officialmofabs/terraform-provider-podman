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

// ContainerPauseReader is a Reader for the ContainerPause structure.
type ContainerPauseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerPauseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerPauseNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerPauseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerPauseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerPauseNoContent creates a ContainerPauseNoContent with default headers values
func NewContainerPauseNoContent() *ContainerPauseNoContent {
	return &ContainerPauseNoContent{}
}

/* ContainerPauseNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerPauseNoContent struct {
}

func (o *ContainerPauseNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/pause][%d] containerPauseNoContent ", 204)
}

func (o *ContainerPauseNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerPauseNotFound creates a ContainerPauseNotFound with default headers values
func NewContainerPauseNotFound() *ContainerPauseNotFound {
	return &ContainerPauseNotFound{}
}

/* ContainerPauseNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerPauseNotFound struct {
	Payload *models.ContainerPauseNotFoundBody
}

func (o *ContainerPauseNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/pause][%d] containerPauseNotFound  %+v", 404, o.Payload)
}
func (o *ContainerPauseNotFound) GetPayload() *models.ContainerPauseNotFoundBody {
	return o.Payload
}

func (o *ContainerPauseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerPauseNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerPauseInternalServerError creates a ContainerPauseInternalServerError with default headers values
func NewContainerPauseInternalServerError() *ContainerPauseInternalServerError {
	return &ContainerPauseInternalServerError{}
}

/* ContainerPauseInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerPauseInternalServerError struct {
	Payload *models.ContainerPauseInternalServerErrorBody
}

func (o *ContainerPauseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/pause][%d] containerPauseInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerPauseInternalServerError) GetPayload() *models.ContainerPauseInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerPauseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerPauseInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
