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

// ListCategoryMilestonesReader is a Reader for the ListCategoryMilestones structure.
type ListCategoryMilestonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCategoryMilestonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCategoryMilestonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCategoryMilestonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCategoryMilestonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListCategoryMilestonesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCategoryMilestonesOK creates a ListCategoryMilestonesOK with default headers values
func NewListCategoryMilestonesOK() *ListCategoryMilestonesOK {
	return &ListCategoryMilestonesOK{}
}

/* ListCategoryMilestonesOK describes a response with status code 200, with default header values.

Resource
*/
type ListCategoryMilestonesOK struct {
	Payload []*models.Milestone
}

func (o *ListCategoryMilestonesOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesOK  %+v", 200, o.Payload)
}
func (o *ListCategoryMilestonesOK) GetPayload() []*models.Milestone {
	return o.Payload
}

func (o *ListCategoryMilestonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCategoryMilestonesBadRequest creates a ListCategoryMilestonesBadRequest with default headers values
func NewListCategoryMilestonesBadRequest() *ListCategoryMilestonesBadRequest {
	return &ListCategoryMilestonesBadRequest{}
}

/* ListCategoryMilestonesBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type ListCategoryMilestonesBadRequest struct {
}

func (o *ListCategoryMilestonesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesBadRequest ", 400)
}

func (o *ListCategoryMilestonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCategoryMilestonesNotFound creates a ListCategoryMilestonesNotFound with default headers values
func NewListCategoryMilestonesNotFound() *ListCategoryMilestonesNotFound {
	return &ListCategoryMilestonesNotFound{}
}

/* ListCategoryMilestonesNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type ListCategoryMilestonesNotFound struct {
}

func (o *ListCategoryMilestonesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesNotFound ", 404)
}

func (o *ListCategoryMilestonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCategoryMilestonesUnprocessableEntity creates a ListCategoryMilestonesUnprocessableEntity with default headers values
func NewListCategoryMilestonesUnprocessableEntity() *ListCategoryMilestonesUnprocessableEntity {
	return &ListCategoryMilestonesUnprocessableEntity{}
}

/* ListCategoryMilestonesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type ListCategoryMilestonesUnprocessableEntity struct {
}

func (o *ListCategoryMilestonesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesUnprocessableEntity ", 422)
}

func (o *ListCategoryMilestonesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
