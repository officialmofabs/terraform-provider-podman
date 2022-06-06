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

// VolumePruneLibpodReader is a Reader for the VolumePruneLibpod structure.
type VolumePruneLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumePruneLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumePruneLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewVolumePruneLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumePruneLibpodOK creates a VolumePruneLibpodOK with default headers values
func NewVolumePruneLibpodOK() *VolumePruneLibpodOK {
	return &VolumePruneLibpodOK{}
}

/* VolumePruneLibpodOK describes a response with status code 200, with default header values.

Volume prune response
*/
type VolumePruneLibpodOK struct {
	Payload []*models.PruneReport
}

func (o *VolumePruneLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/volumes/prune][%d] volumePruneLibpodOK  %+v", 200, o.Payload)
}
func (o *VolumePruneLibpodOK) GetPayload() []*models.PruneReport {
	return o.Payload
}

func (o *VolumePruneLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumePruneLibpodInternalServerError creates a VolumePruneLibpodInternalServerError with default headers values
func NewVolumePruneLibpodInternalServerError() *VolumePruneLibpodInternalServerError {
	return &VolumePruneLibpodInternalServerError{}
}

/* VolumePruneLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type VolumePruneLibpodInternalServerError struct {
	Payload *models.VolumePruneLibpodInternalServerErrorBody
}

func (o *VolumePruneLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/volumes/prune][%d] volumePruneLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *VolumePruneLibpodInternalServerError) GetPayload() *models.VolumePruneLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *VolumePruneLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumePruneLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
