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

// GetIterationReader is a Reader for the GetIteration structure.
type GetIterationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIterationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIterationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIterationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIterationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetIterationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIterationOK creates a GetIterationOK with default headers values
func NewGetIterationOK() *GetIterationOK {
	return &GetIterationOK{}
}

/* GetIterationOK describes a response with status code 200, with default header values.

Resource
*/
type GetIterationOK struct {
	Payload *models.Iteration
}

func (o *GetIterationOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/iterations/{iteration-public-id}][%d] getIterationOK  %+v", 200, o.Payload)
}
func (o *GetIterationOK) GetPayload() *models.Iteration {
	return o.Payload
}

func (o *GetIterationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Iteration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIterationBadRequest creates a GetIterationBadRequest with default headers values
func NewGetIterationBadRequest() *GetIterationBadRequest {
	return &GetIterationBadRequest{}
}

/* GetIterationBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetIterationBadRequest struct {
}

func (o *GetIterationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/iterations/{iteration-public-id}][%d] getIterationBadRequest ", 400)
}

func (o *GetIterationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIterationNotFound creates a GetIterationNotFound with default headers values
func NewGetIterationNotFound() *GetIterationNotFound {
	return &GetIterationNotFound{}
}

/* GetIterationNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetIterationNotFound struct {
}

func (o *GetIterationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/iterations/{iteration-public-id}][%d] getIterationNotFound ", 404)
}

func (o *GetIterationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIterationUnprocessableEntity creates a GetIterationUnprocessableEntity with default headers values
func NewGetIterationUnprocessableEntity() *GetIterationUnprocessableEntity {
	return &GetIterationUnprocessableEntity{}
}

/* GetIterationUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetIterationUnprocessableEntity struct {
}

func (o *GetIterationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/iterations/{iteration-public-id}][%d] getIterationUnprocessableEntity ", 422)
}

func (o *GetIterationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
