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

// GetGroupReader is a Reader for the GetGroup structure.
type GetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetGroupUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupOK creates a GetGroupOK with default headers values
func NewGetGroupOK() *GetGroupOK {
	return &GetGroupOK{}
}

/* GetGroupOK describes a response with status code 200, with default header values.

Resource
*/
type GetGroupOK struct {
	Payload *models.Group
}

func (o *GetGroupOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupOK  %+v", 200, o.Payload)
}
func (o *GetGroupOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *GetGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupBadRequest creates a GetGroupBadRequest with default headers values
func NewGetGroupBadRequest() *GetGroupBadRequest {
	return &GetGroupBadRequest{}
}

/* GetGroupBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetGroupBadRequest struct {
}

func (o *GetGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupBadRequest ", 400)
}

func (o *GetGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGroupNotFound creates a GetGroupNotFound with default headers values
func NewGetGroupNotFound() *GetGroupNotFound {
	return &GetGroupNotFound{}
}

/* GetGroupNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetGroupNotFound struct {
}

func (o *GetGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupNotFound ", 404)
}

func (o *GetGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGroupUnprocessableEntity creates a GetGroupUnprocessableEntity with default headers values
func NewGetGroupUnprocessableEntity() *GetGroupUnprocessableEntity {
	return &GetGroupUnprocessableEntity{}
}

/* GetGroupUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetGroupUnprocessableEntity struct {
}

func (o *GetGroupUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupUnprocessableEntity ", 422)
}

func (o *GetGroupUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
