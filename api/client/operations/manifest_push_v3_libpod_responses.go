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

// ManifestPushV3LibpodReader is a Reader for the ManifestPushV3Libpod structure.
type ManifestPushV3LibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManifestPushV3LibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewManifestPushV3LibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewManifestPushV3LibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewManifestPushV3LibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewManifestPushV3LibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewManifestPushV3LibpodOK creates a ManifestPushV3LibpodOK with default headers values
func NewManifestPushV3LibpodOK() *ManifestPushV3LibpodOK {
	return &ManifestPushV3LibpodOK{}
}

/* ManifestPushV3LibpodOK describes a response with status code 200, with default header values.

ManifestPushV3LibpodOK manifest push v3 libpod o k
*/
type ManifestPushV3LibpodOK struct {
	Payload *models.IDResponse
}

func (o *ManifestPushV3LibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/manifests/{name}/push][%d] manifestPushV3LibpodOK  %+v", 200, o.Payload)
}
func (o *ManifestPushV3LibpodOK) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *ManifestPushV3LibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManifestPushV3LibpodBadRequest creates a ManifestPushV3LibpodBadRequest with default headers values
func NewManifestPushV3LibpodBadRequest() *ManifestPushV3LibpodBadRequest {
	return &ManifestPushV3LibpodBadRequest{}
}

/* ManifestPushV3LibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type ManifestPushV3LibpodBadRequest struct {
	Payload *models.ManifestPushV3LibpodBadRequestBody
}

func (o *ManifestPushV3LibpodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/manifests/{name}/push][%d] manifestPushV3LibpodBadRequest  %+v", 400, o.Payload)
}
func (o *ManifestPushV3LibpodBadRequest) GetPayload() *models.ManifestPushV3LibpodBadRequestBody {
	return o.Payload
}

func (o *ManifestPushV3LibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManifestPushV3LibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManifestPushV3LibpodNotFound creates a ManifestPushV3LibpodNotFound with default headers values
func NewManifestPushV3LibpodNotFound() *ManifestPushV3LibpodNotFound {
	return &ManifestPushV3LibpodNotFound{}
}

/* ManifestPushV3LibpodNotFound describes a response with status code 404, with default header values.

No such manifest
*/
type ManifestPushV3LibpodNotFound struct {
	Payload *models.ManifestPushV3LibpodNotFoundBody
}

func (o *ManifestPushV3LibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/manifests/{name}/push][%d] manifestPushV3LibpodNotFound  %+v", 404, o.Payload)
}
func (o *ManifestPushV3LibpodNotFound) GetPayload() *models.ManifestPushV3LibpodNotFoundBody {
	return o.Payload
}

func (o *ManifestPushV3LibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManifestPushV3LibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManifestPushV3LibpodInternalServerError creates a ManifestPushV3LibpodInternalServerError with default headers values
func NewManifestPushV3LibpodInternalServerError() *ManifestPushV3LibpodInternalServerError {
	return &ManifestPushV3LibpodInternalServerError{}
}

/* ManifestPushV3LibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ManifestPushV3LibpodInternalServerError struct {
	Payload *models.ManifestPushV3LibpodInternalServerErrorBody
}

func (o *ManifestPushV3LibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/manifests/{name}/push][%d] manifestPushV3LibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ManifestPushV3LibpodInternalServerError) GetPayload() *models.ManifestPushV3LibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ManifestPushV3LibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManifestPushV3LibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
