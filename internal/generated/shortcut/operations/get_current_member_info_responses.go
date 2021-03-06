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

// GetCurrentMemberInfoReader is a Reader for the GetCurrentMemberInfo structure.
type GetCurrentMemberInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentMemberInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentMemberInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCurrentMemberInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCurrentMemberInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetCurrentMemberInfoUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCurrentMemberInfoOK creates a GetCurrentMemberInfoOK with default headers values
func NewGetCurrentMemberInfoOK() *GetCurrentMemberInfoOK {
	return &GetCurrentMemberInfoOK{}
}

/* GetCurrentMemberInfoOK describes a response with status code 200, with default header values.

Resource
*/
type GetCurrentMemberInfoOK struct {
	Payload *models.MemberInfo
}

func (o *GetCurrentMemberInfoOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/member][%d] getCurrentMemberInfoOK  %+v", 200, o.Payload)
}
func (o *GetCurrentMemberInfoOK) GetPayload() *models.MemberInfo {
	return o.Payload
}

func (o *GetCurrentMemberInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentMemberInfoBadRequest creates a GetCurrentMemberInfoBadRequest with default headers values
func NewGetCurrentMemberInfoBadRequest() *GetCurrentMemberInfoBadRequest {
	return &GetCurrentMemberInfoBadRequest{}
}

/* GetCurrentMemberInfoBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetCurrentMemberInfoBadRequest struct {
}

func (o *GetCurrentMemberInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/member][%d] getCurrentMemberInfoBadRequest ", 400)
}

func (o *GetCurrentMemberInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrentMemberInfoNotFound creates a GetCurrentMemberInfoNotFound with default headers values
func NewGetCurrentMemberInfoNotFound() *GetCurrentMemberInfoNotFound {
	return &GetCurrentMemberInfoNotFound{}
}

/* GetCurrentMemberInfoNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetCurrentMemberInfoNotFound struct {
}

func (o *GetCurrentMemberInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/member][%d] getCurrentMemberInfoNotFound ", 404)
}

func (o *GetCurrentMemberInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrentMemberInfoUnprocessableEntity creates a GetCurrentMemberInfoUnprocessableEntity with default headers values
func NewGetCurrentMemberInfoUnprocessableEntity() *GetCurrentMemberInfoUnprocessableEntity {
	return &GetCurrentMemberInfoUnprocessableEntity{}
}

/* GetCurrentMemberInfoUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetCurrentMemberInfoUnprocessableEntity struct {
}

func (o *GetCurrentMemberInfoUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/member][%d] getCurrentMemberInfoUnprocessableEntity ", 422)
}

func (o *GetCurrentMemberInfoUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
