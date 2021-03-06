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

// GetFileReader is a Reader for the GetFile structure.
type GetFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetFileUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFileOK creates a GetFileOK with default headers values
func NewGetFileOK() *GetFileOK {
	return &GetFileOK{}
}

/* GetFileOK describes a response with status code 200, with default header values.

Resource
*/
type GetFileOK struct {
	Payload *models.UploadedFile
}

func (o *GetFileOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileOK  %+v", 200, o.Payload)
}
func (o *GetFileOK) GetPayload() *models.UploadedFile {
	return o.Payload
}

func (o *GetFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UploadedFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileBadRequest creates a GetFileBadRequest with default headers values
func NewGetFileBadRequest() *GetFileBadRequest {
	return &GetFileBadRequest{}
}

/* GetFileBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetFileBadRequest struct {
}

func (o *GetFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileBadRequest ", 400)
}

func (o *GetFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFileNotFound creates a GetFileNotFound with default headers values
func NewGetFileNotFound() *GetFileNotFound {
	return &GetFileNotFound{}
}

/* GetFileNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetFileNotFound struct {
}

func (o *GetFileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileNotFound ", 404)
}

func (o *GetFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFileUnprocessableEntity creates a GetFileUnprocessableEntity with default headers values
func NewGetFileUnprocessableEntity() *GetFileUnprocessableEntity {
	return &GetFileUnprocessableEntity{}
}

/* GetFileUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetFileUnprocessableEntity struct {
}

func (o *GetFileUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileUnprocessableEntity ", 422)
}

func (o *GetFileUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
