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

// GetWorkflowReader is a Reader for the GetWorkflow structure.
type GetWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkflowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetWorkflowUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkflowOK creates a GetWorkflowOK with default headers values
func NewGetWorkflowOK() *GetWorkflowOK {
	return &GetWorkflowOK{}
}

/* GetWorkflowOK describes a response with status code 200, with default header values.

Resource
*/
type GetWorkflowOK struct {
	Payload *models.Workflow
}

func (o *GetWorkflowOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/workflows/{workflow-public-id}][%d] getWorkflowOK  %+v", 200, o.Payload)
}
func (o *GetWorkflowOK) GetPayload() *models.Workflow {
	return o.Payload
}

func (o *GetWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Workflow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowBadRequest creates a GetWorkflowBadRequest with default headers values
func NewGetWorkflowBadRequest() *GetWorkflowBadRequest {
	return &GetWorkflowBadRequest{}
}

/* GetWorkflowBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetWorkflowBadRequest struct {
}

func (o *GetWorkflowBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/workflows/{workflow-public-id}][%d] getWorkflowBadRequest ", 400)
}

func (o *GetWorkflowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowNotFound creates a GetWorkflowNotFound with default headers values
func NewGetWorkflowNotFound() *GetWorkflowNotFound {
	return &GetWorkflowNotFound{}
}

/* GetWorkflowNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetWorkflowNotFound struct {
}

func (o *GetWorkflowNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/workflows/{workflow-public-id}][%d] getWorkflowNotFound ", 404)
}

func (o *GetWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowUnprocessableEntity creates a GetWorkflowUnprocessableEntity with default headers values
func NewGetWorkflowUnprocessableEntity() *GetWorkflowUnprocessableEntity {
	return &GetWorkflowUnprocessableEntity{}
}

/* GetWorkflowUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetWorkflowUnprocessableEntity struct {
}

func (o *GetWorkflowUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/workflows/{workflow-public-id}][%d] getWorkflowUnprocessableEntity ", 422)
}

func (o *GetWorkflowUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
