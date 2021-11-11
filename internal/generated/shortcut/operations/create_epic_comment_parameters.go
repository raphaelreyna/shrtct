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

// NewCreateEpicCommentParams creates a new CreateEpicCommentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEpicCommentParams() *CreateEpicCommentParams {
	return &CreateEpicCommentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEpicCommentParamsWithTimeout creates a new CreateEpicCommentParams object
// with the ability to set a timeout on a request.
func NewCreateEpicCommentParamsWithTimeout(timeout time.Duration) *CreateEpicCommentParams {
	return &CreateEpicCommentParams{
		timeout: timeout,
	}
}

// NewCreateEpicCommentParamsWithContext creates a new CreateEpicCommentParams object
// with the ability to set a context for a request.
func NewCreateEpicCommentParamsWithContext(ctx context.Context) *CreateEpicCommentParams {
	return &CreateEpicCommentParams{
		Context: ctx,
	}
}

// NewCreateEpicCommentParamsWithHTTPClient creates a new CreateEpicCommentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEpicCommentParamsWithHTTPClient(client *http.Client) *CreateEpicCommentParams {
	return &CreateEpicCommentParams{
		HTTPClient: client,
	}
}

/* CreateEpicCommentParams contains all the parameters to send to the API endpoint
   for the create epic comment operation.

   Typically these are written to a http.Request.
*/
type CreateEpicCommentParams struct {

	// CreateEpicComment.
	CreateEpicComment *models.CreateEpicComment

	/* EpicPublicID.

	   The ID of the associated Epic.

	   Format: int64
	*/
	EpicPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create epic comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEpicCommentParams) WithDefaults() *CreateEpicCommentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create epic comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEpicCommentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create epic comment params
func (o *CreateEpicCommentParams) WithTimeout(timeout time.Duration) *CreateEpicCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create epic comment params
func (o *CreateEpicCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create epic comment params
func (o *CreateEpicCommentParams) WithContext(ctx context.Context) *CreateEpicCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create epic comment params
func (o *CreateEpicCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create epic comment params
func (o *CreateEpicCommentParams) WithHTTPClient(client *http.Client) *CreateEpicCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create epic comment params
func (o *CreateEpicCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateEpicComment adds the createEpicComment to the create epic comment params
func (o *CreateEpicCommentParams) WithCreateEpicComment(createEpicComment *models.CreateEpicComment) *CreateEpicCommentParams {
	o.SetCreateEpicComment(createEpicComment)
	return o
}

// SetCreateEpicComment adds the createEpicComment to the create epic comment params
func (o *CreateEpicCommentParams) SetCreateEpicComment(createEpicComment *models.CreateEpicComment) {
	o.CreateEpicComment = createEpicComment
}

// WithEpicPublicID adds the epicPublicID to the create epic comment params
func (o *CreateEpicCommentParams) WithEpicPublicID(epicPublicID int64) *CreateEpicCommentParams {
	o.SetEpicPublicID(epicPublicID)
	return o
}

// SetEpicPublicID adds the epicPublicId to the create epic comment params
func (o *CreateEpicCommentParams) SetEpicPublicID(epicPublicID int64) {
	o.EpicPublicID = epicPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEpicCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateEpicComment != nil {
		if err := r.SetBodyParam(o.CreateEpicComment); err != nil {
			return err
		}
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
