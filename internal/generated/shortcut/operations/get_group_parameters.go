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
)

// NewGetGroupParams creates a new GetGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGroupParams() *GetGroupParams {
	return &GetGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupParamsWithTimeout creates a new GetGroupParams object
// with the ability to set a timeout on a request.
func NewGetGroupParamsWithTimeout(timeout time.Duration) *GetGroupParams {
	return &GetGroupParams{
		timeout: timeout,
	}
}

// NewGetGroupParamsWithContext creates a new GetGroupParams object
// with the ability to set a context for a request.
func NewGetGroupParamsWithContext(ctx context.Context) *GetGroupParams {
	return &GetGroupParams{
		Context: ctx,
	}
}

// NewGetGroupParamsWithHTTPClient creates a new GetGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGroupParamsWithHTTPClient(client *http.Client) *GetGroupParams {
	return &GetGroupParams{
		HTTPClient: client,
	}
}

/* GetGroupParams contains all the parameters to send to the API endpoint
   for the get group operation.

   Typically these are written to a http.Request.
*/
type GetGroupParams struct {

	/* GroupPublicID.

	   The unique ID of the Group.

	   Format: uuid
	*/
	GroupPublicID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGroupParams) WithDefaults() *GetGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get group params
func (o *GetGroupParams) WithTimeout(timeout time.Duration) *GetGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get group params
func (o *GetGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get group params
func (o *GetGroupParams) WithContext(ctx context.Context) *GetGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get group params
func (o *GetGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get group params
func (o *GetGroupParams) WithHTTPClient(client *http.Client) *GetGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get group params
func (o *GetGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupPublicID adds the groupPublicID to the get group params
func (o *GetGroupParams) WithGroupPublicID(groupPublicID strfmt.UUID) *GetGroupParams {
	o.SetGroupPublicID(groupPublicID)
	return o
}

// SetGroupPublicID adds the groupPublicId to the get group params
func (o *GetGroupParams) SetGroupPublicID(groupPublicID strfmt.UUID) {
	o.GroupPublicID = groupPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param group-public-id
	if err := r.SetPathParam("group-public-id", o.GroupPublicID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}