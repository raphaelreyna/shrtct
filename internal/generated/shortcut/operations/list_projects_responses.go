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

// ListProjectsReader is a Reader for the ListProjects structure.
type ListProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListProjectsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListProjectsOK creates a ListProjectsOK with default headers values
func NewListProjectsOK() *ListProjectsOK {
	return &ListProjectsOK{}
}

/* ListProjectsOK describes a response with status code 200, with default header values.

Resource
*/
type ListProjectsOK struct {
	Payload []*models.Project
}

func (o *ListProjectsOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/projects][%d] listProjectsOK  %+v", 200, o.Payload)
}
func (o *ListProjectsOK) GetPayload() []*models.Project {
	return o.Payload
}

func (o *ListProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsBadRequest creates a ListProjectsBadRequest with default headers values
func NewListProjectsBadRequest() *ListProjectsBadRequest {
	return &ListProjectsBadRequest{}
}

/* ListProjectsBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type ListProjectsBadRequest struct {
}

func (o *ListProjectsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/projects][%d] listProjectsBadRequest ", 400)
}

func (o *ListProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectsNotFound creates a ListProjectsNotFound with default headers values
func NewListProjectsNotFound() *ListProjectsNotFound {
	return &ListProjectsNotFound{}
}

/* ListProjectsNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type ListProjectsNotFound struct {
}

func (o *ListProjectsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/projects][%d] listProjectsNotFound ", 404)
}

func (o *ListProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectsUnprocessableEntity creates a ListProjectsUnprocessableEntity with default headers values
func NewListProjectsUnprocessableEntity() *ListProjectsUnprocessableEntity {
	return &ListProjectsUnprocessableEntity{}
}

/* ListProjectsUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type ListProjectsUnprocessableEntity struct {
}

func (o *ListProjectsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/projects][%d] listProjectsUnprocessableEntity ", 422)
}

func (o *ListProjectsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
