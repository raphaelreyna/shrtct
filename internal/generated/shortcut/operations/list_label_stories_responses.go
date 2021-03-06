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

// ListLabelStoriesReader is a Reader for the ListLabelStories structure.
type ListLabelStoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLabelStoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLabelStoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListLabelStoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListLabelStoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListLabelStoriesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListLabelStoriesOK creates a ListLabelStoriesOK with default headers values
func NewListLabelStoriesOK() *ListLabelStoriesOK {
	return &ListLabelStoriesOK{}
}

/* ListLabelStoriesOK describes a response with status code 200, with default header values.

Resource
*/
type ListLabelStoriesOK struct {
	Payload []*models.StorySlim
}

func (o *ListLabelStoriesOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}/stories][%d] listLabelStoriesOK  %+v", 200, o.Payload)
}
func (o *ListLabelStoriesOK) GetPayload() []*models.StorySlim {
	return o.Payload
}

func (o *ListLabelStoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLabelStoriesBadRequest creates a ListLabelStoriesBadRequest with default headers values
func NewListLabelStoriesBadRequest() *ListLabelStoriesBadRequest {
	return &ListLabelStoriesBadRequest{}
}

/* ListLabelStoriesBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type ListLabelStoriesBadRequest struct {
}

func (o *ListLabelStoriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}/stories][%d] listLabelStoriesBadRequest ", 400)
}

func (o *ListLabelStoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListLabelStoriesNotFound creates a ListLabelStoriesNotFound with default headers values
func NewListLabelStoriesNotFound() *ListLabelStoriesNotFound {
	return &ListLabelStoriesNotFound{}
}

/* ListLabelStoriesNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type ListLabelStoriesNotFound struct {
}

func (o *ListLabelStoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}/stories][%d] listLabelStoriesNotFound ", 404)
}

func (o *ListLabelStoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListLabelStoriesUnprocessableEntity creates a ListLabelStoriesUnprocessableEntity with default headers values
func NewListLabelStoriesUnprocessableEntity() *ListLabelStoriesUnprocessableEntity {
	return &ListLabelStoriesUnprocessableEntity{}
}

/* ListLabelStoriesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type ListLabelStoriesUnprocessableEntity struct {
}

func (o *ListLabelStoriesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}/stories][%d] listLabelStoriesUnprocessableEntity ", 422)
}

func (o *ListLabelStoriesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
