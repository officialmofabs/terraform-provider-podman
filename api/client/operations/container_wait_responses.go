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

// ContainerWaitReader is a Reader for the ContainerWait structure.
type ContainerWaitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerWaitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerWaitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerWaitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerWaitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerWaitOK creates a ContainerWaitOK with default headers values
func NewContainerWaitOK() *ContainerWaitOK {
	return &ContainerWaitOK{}
}

/* ContainerWaitOK describes a response with status code 200, with default header values.

Wait container
*/
type ContainerWaitOK struct {
	Payload *models.ContainerWaitOKBodyOAIGen
}

func (o *ContainerWaitOK) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/wait][%d] containerWaitOK  %+v", 200, o.Payload)
}
func (o *ContainerWaitOK) GetPayload() *models.ContainerWaitOKBodyOAIGen {
	return o.Payload
}

func (o *ContainerWaitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerWaitOKBodyOAIGen)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerWaitNotFound creates a ContainerWaitNotFound with default headers values
func NewContainerWaitNotFound() *ContainerWaitNotFound {
	return &ContainerWaitNotFound{}
}

/* ContainerWaitNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerWaitNotFound struct {
	Payload *models.ContainerWaitNotFoundBody
}

func (o *ContainerWaitNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/wait][%d] containerWaitNotFound  %+v", 404, o.Payload)
}
func (o *ContainerWaitNotFound) GetPayload() *models.ContainerWaitNotFoundBody {
	return o.Payload
}

func (o *ContainerWaitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerWaitNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerWaitInternalServerError creates a ContainerWaitInternalServerError with default headers values
func NewContainerWaitInternalServerError() *ContainerWaitInternalServerError {
	return &ContainerWaitInternalServerError{}
}

/* ContainerWaitInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerWaitInternalServerError struct {
	Payload *models.ContainerWaitInternalServerErrorBody
}

func (o *ContainerWaitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/wait][%d] containerWaitInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerWaitInternalServerError) GetPayload() *models.ContainerWaitInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerWaitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerWaitInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
