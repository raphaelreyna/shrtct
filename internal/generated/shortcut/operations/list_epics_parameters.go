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

	"github.com/raphaelreyna/shrtct/internal/generated/models"
)

// NewListEpicsParams creates a new ListEpicsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListEpicsParams() *ListEpicsParams {
	return &ListEpicsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListEpicsParamsWithTimeout creates a new ListEpicsParams object
// with the ability to set a timeout on a request.
func NewListEpicsParamsWithTimeout(timeout time.Duration) *ListEpicsParams {
	return &ListEpicsParams{
		timeout: timeout,
	}
}

// NewListEpicsParamsWithContext creates a new ListEpicsParams object
// with the ability to set a context for a request.
func NewListEpicsParamsWithContext(ctx context.Context) *ListEpicsParams {
	return &ListEpicsParams{
		Context: ctx,
	}
}

// NewListEpicsParamsWithHTTPClient creates a new ListEpicsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListEpicsParamsWithHTTPClient(client *http.Client) *ListEpicsParams {
	return &ListEpicsParams{
		HTTPClient: client,
	}
}

/* ListEpicsParams contains all the parameters to send to the API endpoint
   for the list epics operation.

   Typically these are written to a http.Request.
*/
type ListEpicsParams struct {

	// ListEpics.
	ListEpics *models.ListEpics

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list epics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEpicsParams) WithDefaults() *ListEpicsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list epics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEpicsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list epics params
func (o *ListEpicsParams) WithTimeout(timeout time.Duration) *ListEpicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list epics params
func (o *ListEpicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list epics params
func (o *ListEpicsParams) WithContext(ctx context.Context) *ListEpicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list epics params
func (o *ListEpicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list epics params
func (o *ListEpicsParams) WithHTTPClient(client *http.Client) *ListEpicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list epics params
func (o *ListEpicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithListEpics adds the listEpics to the list epics params
func (o *ListEpicsParams) WithListEpics(listEpics *models.ListEpics) *ListEpicsParams {
	o.SetListEpics(listEpics)
	return o
}

// SetListEpics adds the listEpics to the list epics params
func (o *ListEpicsParams) SetListEpics(listEpics *models.ListEpics) {
	o.ListEpics = listEpics
}

// WriteToRequest writes these params to a swagger request
func (o *ListEpicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ListEpics != nil {
		if err := r.SetBodyParam(o.ListEpics); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
