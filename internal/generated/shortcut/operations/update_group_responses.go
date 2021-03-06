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

// UpdateGroupReader is a Reader for the UpdateGroup structure.
type UpdateGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateGroupUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateGroupOK creates a UpdateGroupOK with default headers values
func NewUpdateGroupOK() *UpdateGroupOK {
	return &UpdateGroupOK{}
}

/* UpdateGroupOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateGroupOK struct {
	Payload *models.Group
}

func (o *UpdateGroupOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupOK  %+v", 200, o.Payload)
}
func (o *UpdateGroupOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *UpdateGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupBadRequest creates a UpdateGroupBadRequest with default headers values
func NewUpdateGroupBadRequest() *UpdateGroupBadRequest {
	return &UpdateGroupBadRequest{}
}

/* UpdateGroupBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateGroupBadRequest struct {
}

func (o *UpdateGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupBadRequest ", 400)
}

func (o *UpdateGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGroupForbidden creates a UpdateGroupForbidden with default headers values
func NewUpdateGroupForbidden() *UpdateGroupForbidden {
	return &UpdateGroupForbidden{}
}

/* UpdateGroupForbidden describes a response with status code 403, with default header values.

UpdateGroupForbidden update group forbidden
*/
type UpdateGroupForbidden struct {
	Payload *models.UnusableEntitlementError
}

func (o *UpdateGroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupForbidden  %+v", 403, o.Payload)
}
func (o *UpdateGroupForbidden) GetPayload() *models.UnusableEntitlementError {
	return o.Payload
}

func (o *UpdateGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnusableEntitlementError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupNotFound creates a UpdateGroupNotFound with default headers values
func NewUpdateGroupNotFound() *UpdateGroupNotFound {
	return &UpdateGroupNotFound{}
}

/* UpdateGroupNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateGroupNotFound struct {
}

func (o *UpdateGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupNotFound ", 404)
}

func (o *UpdateGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGroupUnprocessableEntity creates a UpdateGroupUnprocessableEntity with default headers values
func NewUpdateGroupUnprocessableEntity() *UpdateGroupUnprocessableEntity {
	return &UpdateGroupUnprocessableEntity{}
}

/* UpdateGroupUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateGroupUnprocessableEntity struct {
}

func (o *UpdateGroupUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupUnprocessableEntity ", 422)
}

func (o *UpdateGroupUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
