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

// GetMilestoneReader is a Reader for the GetMilestone structure.
type GetMilestoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMilestoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMilestoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMilestoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMilestoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetMilestoneUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMilestoneOK creates a GetMilestoneOK with default headers values
func NewGetMilestoneOK() *GetMilestoneOK {
	return &GetMilestoneOK{}
}

/* GetMilestoneOK describes a response with status code 200, with default header values.

Resource
*/
type GetMilestoneOK struct {
	Payload *models.Milestone
}

func (o *GetMilestoneOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneOK  %+v", 200, o.Payload)
}
func (o *GetMilestoneOK) GetPayload() *models.Milestone {
	return o.Payload
}

func (o *GetMilestoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Milestone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMilestoneBadRequest creates a GetMilestoneBadRequest with default headers values
func NewGetMilestoneBadRequest() *GetMilestoneBadRequest {
	return &GetMilestoneBadRequest{}
}

/* GetMilestoneBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetMilestoneBadRequest struct {
}

func (o *GetMilestoneBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneBadRequest ", 400)
}

func (o *GetMilestoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMilestoneNotFound creates a GetMilestoneNotFound with default headers values
func NewGetMilestoneNotFound() *GetMilestoneNotFound {
	return &GetMilestoneNotFound{}
}

/* GetMilestoneNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetMilestoneNotFound struct {
}

func (o *GetMilestoneNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneNotFound ", 404)
}

func (o *GetMilestoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMilestoneUnprocessableEntity creates a GetMilestoneUnprocessableEntity with default headers values
func NewGetMilestoneUnprocessableEntity() *GetMilestoneUnprocessableEntity {
	return &GetMilestoneUnprocessableEntity{}
}

/* GetMilestoneUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetMilestoneUnprocessableEntity struct {
}

func (o *GetMilestoneUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneUnprocessableEntity ", 422)
}

func (o *GetMilestoneUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}