// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteLinkedFileParams creates a new DeleteLinkedFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLinkedFileParams() *DeleteLinkedFileParams {
	return &DeleteLinkedFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLinkedFileParamsWithTimeout creates a new DeleteLinkedFileParams object
// with the ability to set a timeout on a request.
func NewDeleteLinkedFileParamsWithTimeout(timeout time.Duration) *DeleteLinkedFileParams {
	return &DeleteLinkedFileParams{
		timeout: timeout,
	}
}

// NewDeleteLinkedFileParamsWithContext creates a new DeleteLinkedFileParams object
// with the ability to set a context for a request.
func NewDeleteLinkedFileParamsWithContext(ctx context.Context) *DeleteLinkedFileParams {
	return &DeleteLinkedFileParams{
		Context: ctx,
	}
}

// NewDeleteLinkedFileParamsWithHTTPClient creates a new DeleteLinkedFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLinkedFileParamsWithHTTPClient(client *http.Client) *DeleteLinkedFileParams {
	return &DeleteLinkedFileParams{
		HTTPClient: client,
	}
}

/* DeleteLinkedFileParams contains all the parameters to send to the API endpoint
   for the delete linked file operation.

   Typically these are written to a http.Request.
*/
type DeleteLinkedFileParams struct {

	/* LinkedFilePublicID.

	   The unique identifier of the linked file.

	   Format: int64
	*/
	LinkedFilePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete linked file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLinkedFileParams) WithDefaults() *DeleteLinkedFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete linked file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLinkedFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete linked file params
func (o *DeleteLinkedFileParams) WithTimeout(timeout time.Duration) *DeleteLinkedFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete linked file params
func (o *DeleteLinkedFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete linked file params
func (o *DeleteLinkedFileParams) WithContext(ctx context.Context) *DeleteLinkedFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete linked file params
func (o *DeleteLinkedFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete linked file params
func (o *DeleteLinkedFileParams) WithHTTPClient(client *http.Client) *DeleteLinkedFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete linked file params
func (o *DeleteLinkedFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLinkedFilePublicID adds the linkedFilePublicID to the delete linked file params
func (o *DeleteLinkedFileParams) WithLinkedFilePublicID(linkedFilePublicID int64) *DeleteLinkedFileParams {
	o.SetLinkedFilePublicID(linkedFilePublicID)
	return o
}

// SetLinkedFilePublicID adds the linkedFilePublicId to the delete linked file params
func (o *DeleteLinkedFileParams) SetLinkedFilePublicID(linkedFilePublicID int64) {
	o.LinkedFilePublicID = linkedFilePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLinkedFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param linked-file-public-id
	if err := r.SetPathParam("linked-file-public-id", swag.FormatInt64(o.LinkedFilePublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
