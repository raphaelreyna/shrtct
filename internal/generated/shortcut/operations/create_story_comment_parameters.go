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

	"github.com/raphaelreyna/shrtct/internal/generated/models"
)

// NewCreateStoryCommentParams creates a new CreateStoryCommentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateStoryCommentParams() *CreateStoryCommentParams {
	return &CreateStoryCommentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStoryCommentParamsWithTimeout creates a new CreateStoryCommentParams object
// with the ability to set a timeout on a request.
func NewCreateStoryCommentParamsWithTimeout(timeout time.Duration) *CreateStoryCommentParams {
	return &CreateStoryCommentParams{
		timeout: timeout,
	}
}

// NewCreateStoryCommentParamsWithContext creates a new CreateStoryCommentParams object
// with the ability to set a context for a request.
func NewCreateStoryCommentParamsWithContext(ctx context.Context) *CreateStoryCommentParams {
	return &CreateStoryCommentParams{
		Context: ctx,
	}
}

// NewCreateStoryCommentParamsWithHTTPClient creates a new CreateStoryCommentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateStoryCommentParamsWithHTTPClient(client *http.Client) *CreateStoryCommentParams {
	return &CreateStoryCommentParams{
		HTTPClient: client,
	}
}

/* CreateStoryCommentParams contains all the parameters to send to the API endpoint
   for the create story comment operation.

   Typically these are written to a http.Request.
*/
type CreateStoryCommentParams struct {

	// CreateStoryComment.
	CreateStoryComment *models.CreateStoryComment

	/* StoryPublicID.

	   The ID of the Story that the Comment is in.

	   Format: int64
	*/
	StoryPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create story comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStoryCommentParams) WithDefaults() *CreateStoryCommentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create story comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStoryCommentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create story comment params
func (o *CreateStoryCommentParams) WithTimeout(timeout time.Duration) *CreateStoryCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create story comment params
func (o *CreateStoryCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create story comment params
func (o *CreateStoryCommentParams) WithContext(ctx context.Context) *CreateStoryCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create story comment params
func (o *CreateStoryCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create story comment params
func (o *CreateStoryCommentParams) WithHTTPClient(client *http.Client) *CreateStoryCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create story comment params
func (o *CreateStoryCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateStoryComment adds the createStoryComment to the create story comment params
func (o *CreateStoryCommentParams) WithCreateStoryComment(createStoryComment *models.CreateStoryComment) *CreateStoryCommentParams {
	o.SetCreateStoryComment(createStoryComment)
	return o
}

// SetCreateStoryComment adds the createStoryComment to the create story comment params
func (o *CreateStoryCommentParams) SetCreateStoryComment(createStoryComment *models.CreateStoryComment) {
	o.CreateStoryComment = createStoryComment
}

// WithStoryPublicID adds the storyPublicID to the create story comment params
func (o *CreateStoryCommentParams) WithStoryPublicID(storyPublicID int64) *CreateStoryCommentParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the create story comment params
func (o *CreateStoryCommentParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStoryCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateStoryComment != nil {
		if err := r.SetBodyParam(o.CreateStoryComment); err != nil {
			return err
		}
	}

	// path param story-public-id
	if err := r.SetPathParam("story-public-id", swag.FormatInt64(o.StoryPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
