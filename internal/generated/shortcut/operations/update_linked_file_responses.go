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

// UpdateLinkedFileReader is a Reader for the UpdateLinkedFile structure.
type UpdateLinkedFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLinkedFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLinkedFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateLinkedFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateLinkedFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateLinkedFileUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateLinkedFileOK creates a UpdateLinkedFileOK with default headers values
func NewUpdateLinkedFileOK() *UpdateLinkedFileOK {
	return &UpdateLinkedFileOK{}
}

/* UpdateLinkedFileOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateLinkedFileOK struct {
	Payload *models.LinkedFile
}

func (o *UpdateLinkedFileOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/linked-files/{linked-file-public-id}][%d] updateLinkedFileOK  %+v", 200, o.Payload)
}
func (o *UpdateLinkedFileOK) GetPayload() *models.LinkedFile {
	return o.Payload
}

func (o *UpdateLinkedFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LinkedFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLinkedFileBadRequest creates a UpdateLinkedFileBadRequest with default headers values
func NewUpdateLinkedFileBadRequest() *UpdateLinkedFileBadRequest {
	return &UpdateLinkedFileBadRequest{}
}

/* UpdateLinkedFileBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateLinkedFileBadRequest struct {
}

func (o *UpdateLinkedFileBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/linked-files/{linked-file-public-id}][%d] updateLinkedFileBadRequest ", 400)
}

func (o *UpdateLinkedFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLinkedFileNotFound creates a UpdateLinkedFileNotFound with default headers values
func NewUpdateLinkedFileNotFound() *UpdateLinkedFileNotFound {
	return &UpdateLinkedFileNotFound{}
}

/* UpdateLinkedFileNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateLinkedFileNotFound struct {
}

func (o *UpdateLinkedFileNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/linked-files/{linked-file-public-id}][%d] updateLinkedFileNotFound ", 404)
}

func (o *UpdateLinkedFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLinkedFileUnprocessableEntity creates a UpdateLinkedFileUnprocessableEntity with default headers values
func NewUpdateLinkedFileUnprocessableEntity() *UpdateLinkedFileUnprocessableEntity {
	return &UpdateLinkedFileUnprocessableEntity{}
}

/* UpdateLinkedFileUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateLinkedFileUnprocessableEntity struct {
}

func (o *UpdateLinkedFileUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/linked-files/{linked-file-public-id}][%d] updateLinkedFileUnprocessableEntity ", 422)
}

func (o *UpdateLinkedFileUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
