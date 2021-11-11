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

// NewDeleteTaskParams creates a new DeleteTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTaskParams() *DeleteTaskParams {
	return &DeleteTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTaskParamsWithTimeout creates a new DeleteTaskParams object
// with the ability to set a timeout on a request.
func NewDeleteTaskParamsWithTimeout(timeout time.Duration) *DeleteTaskParams {
	return &DeleteTaskParams{
		timeout: timeout,
	}
}

// NewDeleteTaskParamsWithContext creates a new DeleteTaskParams object
// with the ability to set a context for a request.
func NewDeleteTaskParamsWithContext(ctx context.Context) *DeleteTaskParams {
	return &DeleteTaskParams{
		Context: ctx,
	}
}

// NewDeleteTaskParamsWithHTTPClient creates a new DeleteTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTaskParamsWithHTTPClient(client *http.Client) *DeleteTaskParams {
	return &DeleteTaskParams{
		HTTPClient: client,
	}
}

/* DeleteTaskParams contains all the parameters to send to the API endpoint
   for the delete task operation.

   Typically these are written to a http.Request.
*/
type DeleteTaskParams struct {

	/* StoryPublicID.

	   The unique ID of the Story this Task is associated with.

	   Format: int64
	*/
	StoryPublicID int64

	/* TaskPublicID.

	   The unique ID of the Task.

	   Format: int64
	*/
	TaskPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTaskParams) WithDefaults() *DeleteTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete task params
func (o *DeleteTaskParams) WithTimeout(timeout time.Duration) *DeleteTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete task params
func (o *DeleteTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete task params
func (o *DeleteTaskParams) WithContext(ctx context.Context) *DeleteTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete task params
func (o *DeleteTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete task params
func (o *DeleteTaskParams) WithHTTPClient(client *http.Client) *DeleteTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete task params
func (o *DeleteTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStoryPublicID adds the storyPublicID to the delete task params
func (o *DeleteTaskParams) WithStoryPublicID(storyPublicID int64) *DeleteTaskParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the delete task params
func (o *DeleteTaskParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WithTaskPublicID adds the taskPublicID to the delete task params
func (o *DeleteTaskParams) WithTaskPublicID(taskPublicID int64) *DeleteTaskParams {
	o.SetTaskPublicID(taskPublicID)
	return o
}

// SetTaskPublicID adds the taskPublicId to the delete task params
func (o *DeleteTaskParams) SetTaskPublicID(taskPublicID int64) {
	o.TaskPublicID = taskPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param story-public-id
	if err := r.SetPathParam("story-public-id", swag.FormatInt64(o.StoryPublicID)); err != nil {
		return err
	}

	// path param task-public-id
	if err := r.SetPathParam("task-public-id", swag.FormatInt64(o.TaskPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
