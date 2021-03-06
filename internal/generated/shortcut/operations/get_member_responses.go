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

// GetMemberReader is a Reader for the GetMember structure.
type GetMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetMemberUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMemberOK creates a GetMemberOK with default headers values
func NewGetMemberOK() *GetMemberOK {
	return &GetMemberOK{}
}

/* GetMemberOK describes a response with status code 200, with default header values.

Resource
*/
type GetMemberOK struct {
	Payload *models.Member
}

func (o *GetMemberOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/members/{member-public-id}][%d] getMemberOK  %+v", 200, o.Payload)
}
func (o *GetMemberOK) GetPayload() *models.Member {
	return o.Payload
}

func (o *GetMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Member)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMemberBadRequest creates a GetMemberBadRequest with default headers values
func NewGetMemberBadRequest() *GetMemberBadRequest {
	return &GetMemberBadRequest{}
}

/* GetMemberBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetMemberBadRequest struct {
}

func (o *GetMemberBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/members/{member-public-id}][%d] getMemberBadRequest ", 400)
}

func (o *GetMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMemberNotFound creates a GetMemberNotFound with default headers values
func NewGetMemberNotFound() *GetMemberNotFound {
	return &GetMemberNotFound{}
}

/* GetMemberNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetMemberNotFound struct {
}

func (o *GetMemberNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/members/{member-public-id}][%d] getMemberNotFound ", 404)
}

func (o *GetMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMemberUnprocessableEntity creates a GetMemberUnprocessableEntity with default headers values
func NewGetMemberUnprocessableEntity() *GetMemberUnprocessableEntity {
	return &GetMemberUnprocessableEntity{}
}

/* GetMemberUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetMemberUnprocessableEntity struct {
}

func (o *GetMemberUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/members/{member-public-id}][%d] getMemberUnprocessableEntity ", 422)
}

func (o *GetMemberUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
