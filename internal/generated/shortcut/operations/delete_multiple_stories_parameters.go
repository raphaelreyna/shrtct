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

	"github.com/raphaelreyna/shrtct/internal/generated/models"
)

// NewDeleteMultipleStoriesParams creates a new DeleteMultipleStoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMultipleStoriesParams() *DeleteMultipleStoriesParams {
	return &DeleteMultipleStoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMultipleStoriesParamsWithTimeout creates a new DeleteMultipleStoriesParams object
// with the ability to set a timeout on a request.
func NewDeleteMultipleStoriesParamsWithTimeout(timeout time.Duration) *DeleteMultipleStoriesParams {
	return &DeleteMultipleStoriesParams{
		timeout: timeout,
	}
}

// NewDeleteMultipleStoriesParamsWithContext creates a new DeleteMultipleStoriesParams object
// with the ability to set a context for a request.
func NewDeleteMultipleStoriesParamsWithContext(ctx context.Context) *DeleteMultipleStoriesParams {
	return &DeleteMultipleStoriesParams{
		Context: ctx,
	}
}

// NewDeleteMultipleStoriesParamsWithHTTPClient creates a new DeleteMultipleStoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMultipleStoriesParamsWithHTTPClient(client *http.Client) *DeleteMultipleStoriesParams {
	return &DeleteMultipleStoriesParams{
		HTTPClient: client,
	}
}

/* DeleteMultipleStoriesParams contains all the parameters to send to the API endpoint
   for the delete multiple stories operation.

   Typically these are written to a http.Request.
*/
type DeleteMultipleStoriesParams struct {

	// DeleteStories.
	DeleteStories *models.DeleteStories

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete multiple stories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMultipleStoriesParams) WithDefaults() *DeleteMultipleStoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete multiple stories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMultipleStoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete multiple stories params
func (o *DeleteMultipleStoriesParams) WithTimeout(timeout time.Duration) *DeleteMultipleStoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete multiple stories params
func (o *DeleteMultipleStoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete multiple stories params
func (o *DeleteMultipleStoriesParams) WithContext(ctx context.Context) *DeleteMultipleStoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete multiple stories params
func (o *DeleteMultipleStoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete multiple stories params
func (o *DeleteMultipleStoriesParams) WithHTTPClient(client *http.Client) *DeleteMultipleStoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete multiple stories params
func (o *DeleteMultipleStoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteStories adds the deleteStories to the delete multiple stories params
func (o *DeleteMultipleStoriesParams) WithDeleteStories(deleteStories *models.DeleteStories) *DeleteMultipleStoriesParams {
	o.SetDeleteStories(deleteStories)
	return o
}

// SetDeleteStories adds the deleteStories to the delete multiple stories params
func (o *DeleteMultipleStoriesParams) SetDeleteStories(deleteStories *models.DeleteStories) {
	o.DeleteStories = deleteStories
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMultipleStoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.DeleteStories != nil {
		if err := r.SetBodyParam(o.DeleteStories); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
