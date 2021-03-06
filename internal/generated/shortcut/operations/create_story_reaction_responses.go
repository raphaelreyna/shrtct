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

// CreateStoryReactionReader is a Reader for the CreateStoryReaction structure.
type CreateStoryReactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStoryReactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateStoryReactionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateStoryReactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateStoryReactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateStoryReactionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateStoryReactionCreated creates a CreateStoryReactionCreated with default headers values
func NewCreateStoryReactionCreated() *CreateStoryReactionCreated {
	return &CreateStoryReactionCreated{}
}

/* CreateStoryReactionCreated describes a response with status code 201, with default header values.

Resource
*/
type CreateStoryReactionCreated struct {
	Payload []*models.StoryReaction
}

func (o *CreateStoryReactionCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/reactions][%d] createStoryReactionCreated  %+v", 201, o.Payload)
}
func (o *CreateStoryReactionCreated) GetPayload() []*models.StoryReaction {
	return o.Payload
}

func (o *CreateStoryReactionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStoryReactionBadRequest creates a CreateStoryReactionBadRequest with default headers values
func NewCreateStoryReactionBadRequest() *CreateStoryReactionBadRequest {
	return &CreateStoryReactionBadRequest{}
}

/* CreateStoryReactionBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type CreateStoryReactionBadRequest struct {
}

func (o *CreateStoryReactionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/reactions][%d] createStoryReactionBadRequest ", 400)
}

func (o *CreateStoryReactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateStoryReactionNotFound creates a CreateStoryReactionNotFound with default headers values
func NewCreateStoryReactionNotFound() *CreateStoryReactionNotFound {
	return &CreateStoryReactionNotFound{}
}

/* CreateStoryReactionNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type CreateStoryReactionNotFound struct {
}

func (o *CreateStoryReactionNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/reactions][%d] createStoryReactionNotFound ", 404)
}

func (o *CreateStoryReactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateStoryReactionUnprocessableEntity creates a CreateStoryReactionUnprocessableEntity with default headers values
func NewCreateStoryReactionUnprocessableEntity() *CreateStoryReactionUnprocessableEntity {
	return &CreateStoryReactionUnprocessableEntity{}
}

/* CreateStoryReactionUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type CreateStoryReactionUnprocessableEntity struct {
}

func (o *CreateStoryReactionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/reactions][%d] createStoryReactionUnprocessableEntity ", 422)
}

func (o *CreateStoryReactionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
