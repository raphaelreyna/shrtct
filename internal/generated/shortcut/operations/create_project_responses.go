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

// CreateProjectReader is a Reader for the CreateProject structure.
type CreateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateProjectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProjectCreated creates a CreateProjectCreated with default headers values
func NewCreateProjectCreated() *CreateProjectCreated {
	return &CreateProjectCreated{}
}

/* CreateProjectCreated describes a response with status code 201, with default header values.

Resource
*/
type CreateProjectCreated struct {
	Payload *models.Project
}

func (o *CreateProjectCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/projects][%d] createProjectCreated  %+v", 201, o.Payload)
}
func (o *CreateProjectCreated) GetPayload() *models.Project {
	return o.Payload
}

func (o *CreateProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectBadRequest creates a CreateProjectBadRequest with default headers values
func NewCreateProjectBadRequest() *CreateProjectBadRequest {
	return &CreateProjectBadRequest{}
}

/* CreateProjectBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type CreateProjectBadRequest struct {
}

func (o *CreateProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/projects][%d] createProjectBadRequest ", 400)
}

func (o *CreateProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectNotFound creates a CreateProjectNotFound with default headers values
func NewCreateProjectNotFound() *CreateProjectNotFound {
	return &CreateProjectNotFound{}
}

/* CreateProjectNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type CreateProjectNotFound struct {
}

func (o *CreateProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/projects][%d] createProjectNotFound ", 404)
}

func (o *CreateProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectUnprocessableEntity creates a CreateProjectUnprocessableEntity with default headers values
func NewCreateProjectUnprocessableEntity() *CreateProjectUnprocessableEntity {
	return &CreateProjectUnprocessableEntity{}
}

/* CreateProjectUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type CreateProjectUnprocessableEntity struct {
}

func (o *CreateProjectUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/projects][%d] createProjectUnprocessableEntity ", 422)
}

func (o *CreateProjectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}