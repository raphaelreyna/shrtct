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

// NewListCategoriesParams creates a new ListCategoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListCategoriesParams() *ListCategoriesParams {
	return &ListCategoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListCategoriesParamsWithTimeout creates a new ListCategoriesParams object
// with the ability to set a timeout on a request.
func NewListCategoriesParamsWithTimeout(timeout time.Duration) *ListCategoriesParams {
	return &ListCategoriesParams{
		timeout: timeout,
	}
}

// NewListCategoriesParamsWithContext creates a new ListCategoriesParams object
// with the ability to set a context for a request.
func NewListCategoriesParamsWithContext(ctx context.Context) *ListCategoriesParams {
	return &ListCategoriesParams{
		Context: ctx,
	}
}

// NewListCategoriesParamsWithHTTPClient creates a new ListCategoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListCategoriesParamsWithHTTPClient(client *http.Client) *ListCategoriesParams {
	return &ListCategoriesParams{
		HTTPClient: client,
	}
}

/* ListCategoriesParams contains all the parameters to send to the API endpoint
   for the list categories operation.

   Typically these are written to a http.Request.
*/
type ListCategoriesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list categories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCategoriesParams) WithDefaults() *ListCategoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list categories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCategoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list categories params
func (o *ListCategoriesParams) WithTimeout(timeout time.Duration) *ListCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list categories params
func (o *ListCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list categories params
func (o *ListCategoriesParams) WithContext(ctx context.Context) *ListCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list categories params
func (o *ListCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list categories params
func (o *ListCategoriesParams) WithHTTPClient(client *http.Client) *ListCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list categories params
func (o *ListCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
