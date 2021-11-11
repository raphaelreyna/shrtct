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

// NewEnableGroupsParams creates a new EnableGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableGroupsParams() *EnableGroupsParams {
	return &EnableGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableGroupsParamsWithTimeout creates a new EnableGroupsParams object
// with the ability to set a timeout on a request.
func NewEnableGroupsParamsWithTimeout(timeout time.Duration) *EnableGroupsParams {
	return &EnableGroupsParams{
		timeout: timeout,
	}
}

// NewEnableGroupsParamsWithContext creates a new EnableGroupsParams object
// with the ability to set a context for a request.
func NewEnableGroupsParamsWithContext(ctx context.Context) *EnableGroupsParams {
	return &EnableGroupsParams{
		Context: ctx,
	}
}

// NewEnableGroupsParamsWithHTTPClient creates a new EnableGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableGroupsParamsWithHTTPClient(client *http.Client) *EnableGroupsParams {
	return &EnableGroupsParams{
		HTTPClient: client,
	}
}

/* EnableGroupsParams contains all the parameters to send to the API endpoint
   for the enable groups operation.

   Typically these are written to a http.Request.
*/
type EnableGroupsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enable groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableGroupsParams) WithDefaults() *EnableGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enable groups params
func (o *EnableGroupsParams) WithTimeout(timeout time.Duration) *EnableGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable groups params
func (o *EnableGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable groups params
func (o *EnableGroupsParams) WithContext(ctx context.Context) *EnableGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable groups params
func (o *EnableGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable groups params
func (o *EnableGroupsParams) WithHTTPClient(client *http.Client) *EnableGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable groups params
func (o *EnableGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EnableGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}