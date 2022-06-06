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

// PodPruneLibpodReader is a Reader for the PodPruneLibpod structure.
type PodPruneLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PodPruneLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPodPruneLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPodPruneLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPodPruneLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPodPruneLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPodPruneLibpodOK creates a PodPruneLibpodOK with default headers values
func NewPodPruneLibpodOK() *PodPruneLibpodOK {
	return &PodPruneLibpodOK{}
}

/* PodPruneLibpodOK describes a response with status code 200, with default header values.

Prune pod
*/
type PodPruneLibpodOK struct {
	Payload *models.PodPruneReport
}

func (o *PodPruneLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/prune][%d] podPruneLibpodOK  %+v", 200, o.Payload)
}
func (o *PodPruneLibpodOK) GetPayload() *models.PodPruneReport {
	return o.Payload
}

func (o *PodPruneLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodPruneReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodPruneLibpodBadRequest creates a PodPruneLibpodBadRequest with default headers values
func NewPodPruneLibpodBadRequest() *PodPruneLibpodBadRequest {
	return &PodPruneLibpodBadRequest{}
}

/* PodPruneLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type PodPruneLibpodBadRequest struct {
	Payload *models.PodPruneLibpodBadRequestBody
}

func (o *PodPruneLibpodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/prune][%d] podPruneLibpodBadRequest  %+v", 400, o.Payload)
}
func (o *PodPruneLibpodBadRequest) GetPayload() *models.PodPruneLibpodBadRequestBody {
	return o.Payload
}

func (o *PodPruneLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodPruneLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodPruneLibpodConflict creates a PodPruneLibpodConflict with default headers values
func NewPodPruneLibpodConflict() *PodPruneLibpodConflict {
	return &PodPruneLibpodConflict{}
}

/* PodPruneLibpodConflict describes a response with status code 409, with default header values.

pod already exists
*/
type PodPruneLibpodConflict struct {
}

func (o *PodPruneLibpodConflict) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/prune][%d] podPruneLibpodConflict ", 409)
}

func (o *PodPruneLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPodPruneLibpodInternalServerError creates a PodPruneLibpodInternalServerError with default headers values
func NewPodPruneLibpodInternalServerError() *PodPruneLibpodInternalServerError {
	return &PodPruneLibpodInternalServerError{}
}

/* PodPruneLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PodPruneLibpodInternalServerError struct {
	Payload *models.PodPruneLibpodInternalServerErrorBody
}

func (o *PodPruneLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/prune][%d] podPruneLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *PodPruneLibpodInternalServerError) GetPayload() *models.PodPruneLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *PodPruneLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodPruneLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
