// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/raphaelreyna/shrtct/internal/generated/models"
)

// GetProjectReader is a Reader for the GetProject structure.
type GetProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetProjectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectOK creates a GetProjectOK with default headers values
func NewGetProjectOK() *GetProjectOK {
	return &GetProjectOK{}
}

/* GetProjectOK describes a response with status code 200, with default header values.

Resource
*/
type GetProjectOK struct {
	Payload *models.Project
}

func (o *GetProjectOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}][%d] getProjectOK  %+v", 200, o.Payload)
}
func (o *GetProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *GetProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectBadRequest creates a GetProjectBadRequest with default headers values
func NewGetProjectBadRequest() *GetProjectBadRequest {
	return &GetProjectBadRequest{}
}

/* GetProjectBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetProjectBadRequest struct {
}

func (o *GetProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}][%d] getProjectBadRequest ", 400)
}

func (o *GetProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectNotFound creates a GetProjectNotFound with default headers values
func NewGetProjectNotFound() *GetProjectNotFound {
	return &GetProjectNotFound{}
}

/* GetProjectNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetProjectNotFound struct {
}

func (o *GetProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}][%d] getProjectNotFound ", 404)
}

func (o *GetProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectUnprocessableEntity creates a GetProjectUnprocessableEntity with default headers values
func NewGetProjectUnprocessableEntity() *GetProjectUnprocessableEntity {
	return &GetProjectUnprocessableEntity{}
}

/* GetProjectUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetProjectUnprocessableEntity struct {
}

func (o *GetProjectUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}][%d] getProjectUnprocessableEntity ", 422)
}

func (o *GetProjectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}