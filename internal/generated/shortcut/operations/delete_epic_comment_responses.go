// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteEpicCommentReader is a Reader for the DeleteEpicComment structure.
type DeleteEpicCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEpicCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteEpicCommentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteEpicCommentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteEpicCommentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteEpicCommentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEpicCommentNoContent creates a DeleteEpicCommentNoContent with default headers values
func NewDeleteEpicCommentNoContent() *DeleteEpicCommentNoContent {
	return &DeleteEpicCommentNoContent{}
}

/* DeleteEpicCommentNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteEpicCommentNoContent struct {
}

func (o *DeleteEpicCommentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentNoContent ", 204)
}

func (o *DeleteEpicCommentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEpicCommentBadRequest creates a DeleteEpicCommentBadRequest with default headers values
func NewDeleteEpicCommentBadRequest() *DeleteEpicCommentBadRequest {
	return &DeleteEpicCommentBadRequest{}
}

/* DeleteEpicCommentBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type DeleteEpicCommentBadRequest struct {
}

func (o *DeleteEpicCommentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentBadRequest ", 400)
}

func (o *DeleteEpicCommentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEpicCommentNotFound creates a DeleteEpicCommentNotFound with default headers values
func NewDeleteEpicCommentNotFound() *DeleteEpicCommentNotFound {
	return &DeleteEpicCommentNotFound{}
}

/* DeleteEpicCommentNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type DeleteEpicCommentNotFound struct {
}

func (o *DeleteEpicCommentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentNotFound ", 404)
}

func (o *DeleteEpicCommentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEpicCommentUnprocessableEntity creates a DeleteEpicCommentUnprocessableEntity with default headers values
func NewDeleteEpicCommentUnprocessableEntity() *DeleteEpicCommentUnprocessableEntity {
	return &DeleteEpicCommentUnprocessableEntity{}
}

/* DeleteEpicCommentUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type DeleteEpicCommentUnprocessableEntity struct {
}

func (o *DeleteEpicCommentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentUnprocessableEntity ", 422)
}

func (o *DeleteEpicCommentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
