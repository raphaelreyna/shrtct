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

// NewDisableGroupsParams creates a new DisableGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisableGroupsParams() *DisableGroupsParams {
	return &DisableGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisableGroupsParamsWithTimeout creates a new DisableGroupsParams object
// with the ability to set a timeout on a request.
func NewDisableGroupsParamsWithTimeout(timeout time.Duration) *DisableGroupsParams {
	return &DisableGroupsParams{
		timeout: timeout,
	}
}

// NewDisableGroupsParamsWithContext creates a new DisableGroupsParams object
// with the ability to set a context for a request.
func NewDisableGroupsParamsWithContext(ctx context.Context) *DisableGroupsParams {
	return &DisableGroupsParams{
		Context: ctx,
	}
}

// NewDisableGroupsParamsWithHTTPClient creates a new DisableGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisableGroupsParamsWithHTTPClient(client *http.Client) *DisableGroupsParams {
	return &DisableGroupsParams{
		HTTPClient: client,
	}
}

/* DisableGroupsParams contains all the parameters to send to the API endpoint
   for the disable groups operation.

   Typically these are written to a http.Request.
*/
type DisableGroupsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disable groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableGroupsParams) WithDefaults() *DisableGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disable groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disable groups params
func (o *DisableGroupsParams) WithTimeout(timeout time.Duration) *DisableGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable groups params
func (o *DisableGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable groups params
func (o *DisableGroupsParams) WithContext(ctx context.Context) *DisableGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable groups params
func (o *DisableGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable groups params
func (o *DisableGroupsParams) WithHTTPClient(client *http.Client) *DisableGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable groups params
func (o *DisableGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DisableGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
