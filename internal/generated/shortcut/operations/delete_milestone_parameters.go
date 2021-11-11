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

// NewDeleteMilestoneParams creates a new DeleteMilestoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMilestoneParams() *DeleteMilestoneParams {
	return &DeleteMilestoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMilestoneParamsWithTimeout creates a new DeleteMilestoneParams object
// with the ability to set a timeout on a request.
func NewDeleteMilestoneParamsWithTimeout(timeout time.Duration) *DeleteMilestoneParams {
	return &DeleteMilestoneParams{
		timeout: timeout,
	}
}

// NewDeleteMilestoneParamsWithContext creates a new DeleteMilestoneParams object
// with the ability to set a context for a request.
func NewDeleteMilestoneParamsWithContext(ctx context.Context) *DeleteMilestoneParams {
	return &DeleteMilestoneParams{
		Context: ctx,
	}
}

// NewDeleteMilestoneParamsWithHTTPClient creates a new DeleteMilestoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMilestoneParamsWithHTTPClient(client *http.Client) *DeleteMilestoneParams {
	return &DeleteMilestoneParams{
		HTTPClient: client,
	}
}

/* DeleteMilestoneParams contains all the parameters to send to the API endpoint
   for the delete milestone operation.

   Typically these are written to a http.Request.
*/
type DeleteMilestoneParams struct {

	/* MilestonePublicID.

	   The ID of the Milestone.

	   Format: int64
	*/
	MilestonePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMilestoneParams) WithDefaults() *DeleteMilestoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMilestoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete milestone params
func (o *DeleteMilestoneParams) WithTimeout(timeout time.Duration) *DeleteMilestoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete milestone params
func (o *DeleteMilestoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete milestone params
func (o *DeleteMilestoneParams) WithContext(ctx context.Context) *DeleteMilestoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete milestone params
func (o *DeleteMilestoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete milestone params
func (o *DeleteMilestoneParams) WithHTTPClient(client *http.Client) *DeleteMilestoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete milestone params
func (o *DeleteMilestoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMilestonePublicID adds the milestonePublicID to the delete milestone params
func (o *DeleteMilestoneParams) WithMilestonePublicID(milestonePublicID int64) *DeleteMilestoneParams {
	o.SetMilestonePublicID(milestonePublicID)
	return o
}

// SetMilestonePublicID adds the milestonePublicId to the delete milestone params
func (o *DeleteMilestoneParams) SetMilestonePublicID(milestonePublicID int64) {
	o.MilestonePublicID = milestonePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMilestoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param milestone-public-id
	if err := r.SetPathParam("milestone-public-id", swag.FormatInt64(o.MilestonePublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
