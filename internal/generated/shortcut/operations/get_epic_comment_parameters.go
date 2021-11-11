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

// NewGetEpicCommentParams creates a new GetEpicCommentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEpicCommentParams() *GetEpicCommentParams {
	return &GetEpicCommentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEpicCommentParamsWithTimeout creates a new GetEpicCommentParams object
// with the ability to set a timeout on a request.
func NewGetEpicCommentParamsWithTimeout(timeout time.Duration) *GetEpicCommentParams {
	return &GetEpicCommentParams{
		timeout: timeout,
	}
}

// NewGetEpicCommentParamsWithContext creates a new GetEpicCommentParams object
// with the ability to set a context for a request.
func NewGetEpicCommentParamsWithContext(ctx context.Context) *GetEpicCommentParams {
	return &GetEpicCommentParams{
		Context: ctx,
	}
}

// NewGetEpicCommentParamsWithHTTPClient creates a new GetEpicCommentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEpicCommentParamsWithHTTPClient(client *http.Client) *GetEpicCommentParams {
	return &GetEpicCommentParams{
		HTTPClient: client,
	}
}

/* GetEpicCommentParams contains all the parameters to send to the API endpoint
   for the get epic comment operation.

   Typically these are written to a http.Request.
*/
type GetEpicCommentParams struct {

	/* CommentPublicID.

	   The ID of the Comment.

	   Format: int64
	*/
	CommentPublicID int64

	/* EpicPublicID.

	   The ID of the associated Epic.

	   Format: int64
	*/
	EpicPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get epic comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEpicCommentParams) WithDefaults() *GetEpicCommentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get epic comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEpicCommentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get epic comment params
func (o *GetEpicCommentParams) WithTimeout(timeout time.Duration) *GetEpicCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get epic comment params
func (o *GetEpicCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get epic comment params
func (o *GetEpicCommentParams) WithContext(ctx context.Context) *GetEpicCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get epic comment params
func (o *GetEpicCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get epic comment params
func (o *GetEpicCommentParams) WithHTTPClient(client *http.Client) *GetEpicCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get epic comment params
func (o *GetEpicCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentPublicID adds the commentPublicID to the get epic comment params
func (o *GetEpicCommentParams) WithCommentPublicID(commentPublicID int64) *GetEpicCommentParams {
	o.SetCommentPublicID(commentPublicID)
	return o
}

// SetCommentPublicID adds the commentPublicId to the get epic comment params
func (o *GetEpicCommentParams) SetCommentPublicID(commentPublicID int64) {
	o.CommentPublicID = commentPublicID
}

// WithEpicPublicID adds the epicPublicID to the get epic comment params
func (o *GetEpicCommentParams) WithEpicPublicID(epicPublicID int64) *GetEpicCommentParams {
	o.SetEpicPublicID(epicPublicID)
	return o
}

// SetEpicPublicID adds the epicPublicId to the get epic comment params
func (o *GetEpicCommentParams) SetEpicPublicID(epicPublicID int64) {
	o.EpicPublicID = epicPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEpicCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param comment-public-id
	if err := r.SetPathParam("comment-public-id", swag.FormatInt64(o.CommentPublicID)); err != nil {
		return err
	}

	// path param epic-public-id
	if err := r.SetPathParam("epic-public-id", swag.FormatInt64(o.EpicPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
