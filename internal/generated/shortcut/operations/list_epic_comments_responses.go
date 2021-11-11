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

// ListEpicCommentsReader is a Reader for the ListEpicComments structure.
type ListEpicCommentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEpicCommentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEpicCommentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListEpicCommentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListEpicCommentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListEpicCommentsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListEpicCommentsOK creates a ListEpicCommentsOK with default headers values
func NewListEpicCommentsOK() *ListEpicCommentsOK {
	return &ListEpicCommentsOK{}
}

/* ListEpicCommentsOK describes a response with status code 200, with default header values.

Resource
*/
type ListEpicCommentsOK struct {
	Payload []*models.ThreadedComment
}

func (o *ListEpicCommentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics/{epic-public-id}/comments][%d] listEpicCommentsOK  %+v", 200, o.Payload)
}
func (o *ListEpicCommentsOK) GetPayload() []*models.ThreadedComment {
	return o.Payload
}

func (o *ListEpicCommentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEpicCommentsBadRequest creates a ListEpicCommentsBadRequest with default headers values
func NewListEpicCommentsBadRequest() *ListEpicCommentsBadRequest {
	return &ListEpicCommentsBadRequest{}
}

/* ListEpicCommentsBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type ListEpicCommentsBadRequest struct {
}

func (o *ListEpicCommentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics/{epic-public-id}/comments][%d] listEpicCommentsBadRequest ", 400)
}

func (o *ListEpicCommentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEpicCommentsNotFound creates a ListEpicCommentsNotFound with default headers values
func NewListEpicCommentsNotFound() *ListEpicCommentsNotFound {
	return &ListEpicCommentsNotFound{}
}

/* ListEpicCommentsNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type ListEpicCommentsNotFound struct {
}

func (o *ListEpicCommentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics/{epic-public-id}/comments][%d] listEpicCommentsNotFound ", 404)
}

func (o *ListEpicCommentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEpicCommentsUnprocessableEntity creates a ListEpicCommentsUnprocessableEntity with default headers values
func NewListEpicCommentsUnprocessableEntity() *ListEpicCommentsUnprocessableEntity {
	return &ListEpicCommentsUnprocessableEntity{}
}

/* ListEpicCommentsUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type ListEpicCommentsUnprocessableEntity struct {
}

func (o *ListEpicCommentsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics/{epic-public-id}/comments][%d] listEpicCommentsUnprocessableEntity ", 422)
}

func (o *ListEpicCommentsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
