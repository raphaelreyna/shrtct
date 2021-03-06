// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteCategoryReader is a Reader for the DeleteCategory structure.
type DeleteCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCategoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteCategoryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCategoryNoContent creates a DeleteCategoryNoContent with default headers values
func NewDeleteCategoryNoContent() *DeleteCategoryNoContent {
	return &DeleteCategoryNoContent{}
}

/* DeleteCategoryNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteCategoryNoContent struct {
}

func (o *DeleteCategoryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryNoContent ", 204)
}

func (o *DeleteCategoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCategoryBadRequest creates a DeleteCategoryBadRequest with default headers values
func NewDeleteCategoryBadRequest() *DeleteCategoryBadRequest {
	return &DeleteCategoryBadRequest{}
}

/* DeleteCategoryBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type DeleteCategoryBadRequest struct {
}

func (o *DeleteCategoryBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryBadRequest ", 400)
}

func (o *DeleteCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCategoryNotFound creates a DeleteCategoryNotFound with default headers values
func NewDeleteCategoryNotFound() *DeleteCategoryNotFound {
	return &DeleteCategoryNotFound{}
}

/* DeleteCategoryNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type DeleteCategoryNotFound struct {
}

func (o *DeleteCategoryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryNotFound ", 404)
}

func (o *DeleteCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCategoryUnprocessableEntity creates a DeleteCategoryUnprocessableEntity with default headers values
func NewDeleteCategoryUnprocessableEntity() *DeleteCategoryUnprocessableEntity {
	return &DeleteCategoryUnprocessableEntity{}
}

/* DeleteCategoryUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type DeleteCategoryUnprocessableEntity struct {
}

func (o *DeleteCategoryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryUnprocessableEntity ", 422)
}

func (o *DeleteCategoryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
