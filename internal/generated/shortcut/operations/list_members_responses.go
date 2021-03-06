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

// ListMembersReader is a Reader for the ListMembers structure.
type ListMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListMembersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListMembersOK creates a ListMembersOK with default headers values
func NewListMembersOK() *ListMembersOK {
	return &ListMembersOK{}
}

/* ListMembersOK describes a response with status code 200, with default header values.

Resource
*/
type ListMembersOK struct {
	Payload []*models.Member
}

func (o *ListMembersOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/members][%d] listMembersOK  %+v", 200, o.Payload)
}
func (o *ListMembersOK) GetPayload() []*models.Member {
	return o.Payload
}

func (o *ListMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMembersBadRequest creates a ListMembersBadRequest with default headers values
func NewListMembersBadRequest() *ListMembersBadRequest {
	return &ListMembersBadRequest{}
}

/* ListMembersBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type ListMembersBadRequest struct {
}

func (o *ListMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/members][%d] listMembersBadRequest ", 400)
}

func (o *ListMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMembersNotFound creates a ListMembersNotFound with default headers values
func NewListMembersNotFound() *ListMembersNotFound {
	return &ListMembersNotFound{}
}

/* ListMembersNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type ListMembersNotFound struct {
}

func (o *ListMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/members][%d] listMembersNotFound ", 404)
}

func (o *ListMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMembersUnprocessableEntity creates a ListMembersUnprocessableEntity with default headers values
func NewListMembersUnprocessableEntity() *ListMembersUnprocessableEntity {
	return &ListMembersUnprocessableEntity{}
}

/* ListMembersUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type ListMembersUnprocessableEntity struct {
}

func (o *ListMembersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/members][%d] listMembersUnprocessableEntity ", 422)
}

func (o *ListMembersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
