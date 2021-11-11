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

// UpdateEpicReader is a Reader for the UpdateEpic structure.
type UpdateEpicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEpicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEpicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateEpicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEpicNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateEpicUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateEpicOK creates a UpdateEpicOK with default headers values
func NewUpdateEpicOK() *UpdateEpicOK {
	return &UpdateEpicOK{}
}

/* UpdateEpicOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateEpicOK struct {
	Payload *models.Epic
}

func (o *UpdateEpicOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}][%d] updateEpicOK  %+v", 200, o.Payload)
}
func (o *UpdateEpicOK) GetPayload() *models.Epic {
	return o.Payload
}

func (o *UpdateEpicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Epic)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEpicBadRequest creates a UpdateEpicBadRequest with default headers values
func NewUpdateEpicBadRequest() *UpdateEpicBadRequest {
	return &UpdateEpicBadRequest{}
}

/* UpdateEpicBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateEpicBadRequest struct {
}

func (o *UpdateEpicBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}][%d] updateEpicBadRequest ", 400)
}

func (o *UpdateEpicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEpicNotFound creates a UpdateEpicNotFound with default headers values
func NewUpdateEpicNotFound() *UpdateEpicNotFound {
	return &UpdateEpicNotFound{}
}

/* UpdateEpicNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateEpicNotFound struct {
}

func (o *UpdateEpicNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}][%d] updateEpicNotFound ", 404)
}

func (o *UpdateEpicNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEpicUnprocessableEntity creates a UpdateEpicUnprocessableEntity with default headers values
func NewUpdateEpicUnprocessableEntity() *UpdateEpicUnprocessableEntity {
	return &UpdateEpicUnprocessableEntity{}
}

/* UpdateEpicUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateEpicUnprocessableEntity struct {
}

func (o *UpdateEpicUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}][%d] updateEpicUnprocessableEntity ", 422)
}

func (o *UpdateEpicUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
