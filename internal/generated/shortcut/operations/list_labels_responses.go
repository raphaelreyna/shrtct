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

// ListLabelsReader is a Reader for the ListLabels structure.
type ListLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListLabelsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListLabelsOK creates a ListLabelsOK with default headers values
func NewListLabelsOK() *ListLabelsOK {
	return &ListLabelsOK{}
}

/* ListLabelsOK describes a response with status code 200, with default header values.

Resource
*/
type ListLabelsOK struct {
	Payload []*models.Label
}

func (o *ListLabelsOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels][%d] listLabelsOK  %+v", 200, o.Payload)
}
func (o *ListLabelsOK) GetPayload() []*models.Label {
	return o.Payload
}

func (o *ListLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLabelsBadRequest creates a ListLabelsBadRequest with default headers values
func NewListLabelsBadRequest() *ListLabelsBadRequest {
	return &ListLabelsBadRequest{}
}

/* ListLabelsBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type ListLabelsBadRequest struct {
}

func (o *ListLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels][%d] listLabelsBadRequest ", 400)
}

func (o *ListLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListLabelsNotFound creates a ListLabelsNotFound with default headers values
func NewListLabelsNotFound() *ListLabelsNotFound {
	return &ListLabelsNotFound{}
}

/* ListLabelsNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type ListLabelsNotFound struct {
}

func (o *ListLabelsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels][%d] listLabelsNotFound ", 404)
}

func (o *ListLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListLabelsUnprocessableEntity creates a ListLabelsUnprocessableEntity with default headers values
func NewListLabelsUnprocessableEntity() *ListLabelsUnprocessableEntity {
	return &ListLabelsUnprocessableEntity{}
}

/* ListLabelsUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type ListLabelsUnprocessableEntity struct {
}

func (o *ListLabelsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels][%d] listLabelsUnprocessableEntity ", 422)
}

func (o *ListLabelsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
