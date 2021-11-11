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

// NewGetStoryLinkParams creates a new GetStoryLinkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStoryLinkParams() *GetStoryLinkParams {
	return &GetStoryLinkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStoryLinkParamsWithTimeout creates a new GetStoryLinkParams object
// with the ability to set a timeout on a request.
func NewGetStoryLinkParamsWithTimeout(timeout time.Duration) *GetStoryLinkParams {
	return &GetStoryLinkParams{
		timeout: timeout,
	}
}

// NewGetStoryLinkParamsWithContext creates a new GetStoryLinkParams object
// with the ability to set a context for a request.
func NewGetStoryLinkParamsWithContext(ctx context.Context) *GetStoryLinkParams {
	return &GetStoryLinkParams{
		Context: ctx,
	}
}

// NewGetStoryLinkParamsWithHTTPClient creates a new GetStoryLinkParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStoryLinkParamsWithHTTPClient(client *http.Client) *GetStoryLinkParams {
	return &GetStoryLinkParams{
		HTTPClient: client,
	}
}

/* GetStoryLinkParams contains all the parameters to send to the API endpoint
   for the get story link operation.

   Typically these are written to a http.Request.
*/
type GetStoryLinkParams struct {

	/* StoryLinkPublicID.

	   The unique ID of the Story Link.

	   Format: int64
	*/
	StoryLinkPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get story link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStoryLinkParams) WithDefaults() *GetStoryLinkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get story link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStoryLinkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get story link params
func (o *GetStoryLinkParams) WithTimeout(timeout time.Duration) *GetStoryLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get story link params
func (o *GetStoryLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get story link params
func (o *GetStoryLinkParams) WithContext(ctx context.Context) *GetStoryLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get story link params
func (o *GetStoryLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get story link params
func (o *GetStoryLinkParams) WithHTTPClient(client *http.Client) *GetStoryLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get story link params
func (o *GetStoryLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStoryLinkPublicID adds the storyLinkPublicID to the get story link params
func (o *GetStoryLinkParams) WithStoryLinkPublicID(storyLinkPublicID int64) *GetStoryLinkParams {
	o.SetStoryLinkPublicID(storyLinkPublicID)
	return o
}

// SetStoryLinkPublicID adds the storyLinkPublicId to the get story link params
func (o *GetStoryLinkParams) SetStoryLinkPublicID(storyLinkPublicID int64) {
	o.StoryLinkPublicID = storyLinkPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStoryLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param story-link-public-id
	if err := r.SetPathParam("story-link-public-id", swag.FormatInt64(o.StoryLinkPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
