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

// NewStoryHistoryParams creates a new StoryHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStoryHistoryParams() *StoryHistoryParams {
	return &StoryHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStoryHistoryParamsWithTimeout creates a new StoryHistoryParams object
// with the ability to set a timeout on a request.
func NewStoryHistoryParamsWithTimeout(timeout time.Duration) *StoryHistoryParams {
	return &StoryHistoryParams{
		timeout: timeout,
	}
}

// NewStoryHistoryParamsWithContext creates a new StoryHistoryParams object
// with the ability to set a context for a request.
func NewStoryHistoryParamsWithContext(ctx context.Context) *StoryHistoryParams {
	return &StoryHistoryParams{
		Context: ctx,
	}
}

// NewStoryHistoryParamsWithHTTPClient creates a new StoryHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewStoryHistoryParamsWithHTTPClient(client *http.Client) *StoryHistoryParams {
	return &StoryHistoryParams{
		HTTPClient: client,
	}
}

/* StoryHistoryParams contains all the parameters to send to the API endpoint
   for the story history operation.

   Typically these are written to a http.Request.
*/
type StoryHistoryParams struct {

	/* StoryPublicID.

	   The ID of the Story.

	   Format: int64
	*/
	StoryPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the story history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StoryHistoryParams) WithDefaults() *StoryHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the story history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StoryHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the story history params
func (o *StoryHistoryParams) WithTimeout(timeout time.Duration) *StoryHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the story history params
func (o *StoryHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the story history params
func (o *StoryHistoryParams) WithContext(ctx context.Context) *StoryHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the story history params
func (o *StoryHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the story history params
func (o *StoryHistoryParams) WithHTTPClient(client *http.Client) *StoryHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the story history params
func (o *StoryHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStoryPublicID adds the storyPublicID to the story history params
func (o *StoryHistoryParams) WithStoryPublicID(storyPublicID int64) *StoryHistoryParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the story history params
func (o *StoryHistoryParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *StoryHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param story-public-id
	if err := r.SetPathParam("story-public-id", swag.FormatInt64(o.StoryPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
