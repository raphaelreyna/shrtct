// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteEpicReader is a Reader for the DeleteEpic structure.
type DeleteEpicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEpicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteEpicNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteEpicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteEpicNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteEpicUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEpicNoContent creates a DeleteEpicNoContent with default headers values
func NewDeleteEpicNoContent() *DeleteEpicNoContent {
	return &DeleteEpicNoContent{}
}

/* DeleteEpicNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteEpicNoContent struct {
}

func (o *DeleteEpicNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}][%d] deleteEpicNoContent ", 204)
}

func (o *DeleteEpicNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEpicBadRequest creates a DeleteEpicBadRequest with default headers values
func NewDeleteEpicBadRequest() *DeleteEpicBadRequest {
	return &DeleteEpicBadRequest{}
}

/* DeleteEpicBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type DeleteEpicBadRequest struct {
}

func (o *DeleteEpicBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}][%d] deleteEpicBadRequest ", 400)
}

func (o *DeleteEpicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEpicNotFound creates a DeleteEpicNotFound with default headers values
func NewDeleteEpicNotFound() *DeleteEpicNotFound {
	return &DeleteEpicNotFound{}
}

/* DeleteEpicNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type DeleteEpicNotFound struct {
}

func (o *DeleteEpicNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}][%d] deleteEpicNotFound ", 404)
}

func (o *DeleteEpicNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEpicUnprocessableEntity creates a DeleteEpicUnprocessableEntity with default headers values
func NewDeleteEpicUnprocessableEntity() *DeleteEpicUnprocessableEntity {
	return &DeleteEpicUnprocessableEntity{}
}

/* DeleteEpicUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type DeleteEpicUnprocessableEntity struct {
}

func (o *DeleteEpicUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}][%d] deleteEpicUnprocessableEntity ", 422)
}

func (o *DeleteEpicUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}