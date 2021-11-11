// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EnableGroupsReader is a Reader for the EnableGroups structure.
type EnableGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEnableGroupsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnableGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnableGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewEnableGroupsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnableGroupsNoContent creates a EnableGroupsNoContent with default headers values
func NewEnableGroupsNoContent() *EnableGroupsNoContent {
	return &EnableGroupsNoContent{}
}

/* EnableGroupsNoContent describes a response with status code 204, with default header values.

No Content
*/
type EnableGroupsNoContent struct {
}

func (o *EnableGroupsNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/enable][%d] enableGroupsNoContent ", 204)
}

func (o *EnableGroupsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableGroupsBadRequest creates a EnableGroupsBadRequest with default headers values
func NewEnableGroupsBadRequest() *EnableGroupsBadRequest {
	return &EnableGroupsBadRequest{}
}

/* EnableGroupsBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type EnableGroupsBadRequest struct {
}

func (o *EnableGroupsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/enable][%d] enableGroupsBadRequest ", 400)
}

func (o *EnableGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableGroupsNotFound creates a EnableGroupsNotFound with default headers values
func NewEnableGroupsNotFound() *EnableGroupsNotFound {
	return &EnableGroupsNotFound{}
}

/* EnableGroupsNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type EnableGroupsNotFound struct {
}

func (o *EnableGroupsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/enable][%d] enableGroupsNotFound ", 404)
}

func (o *EnableGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableGroupsUnprocessableEntity creates a EnableGroupsUnprocessableEntity with default headers values
func NewEnableGroupsUnprocessableEntity() *EnableGroupsUnprocessableEntity {
	return &EnableGroupsUnprocessableEntity{}
}

/* EnableGroupsUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type EnableGroupsUnprocessableEntity struct {
}

func (o *EnableGroupsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/enable][%d] enableGroupsUnprocessableEntity ", 422)
}

func (o *EnableGroupsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
