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

// NewCreateEpicParams creates a new CreateEpicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEpicParams() *CreateEpicParams {
	return &CreateEpicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEpicParamsWithTimeout creates a new CreateEpicParams object
// with the ability to set a timeout on a request.
func NewCreateEpicParamsWithTimeout(timeout time.Duration) *CreateEpicParams {
	return &CreateEpicParams{
		timeout: timeout,
	}
}

// NewCreateEpicParamsWithContext creates a new CreateEpicParams object
// with the ability to set a context for a request.
func NewCreateEpicParamsWithContext(ctx context.Context) *CreateEpicParams {
	return &CreateEpicParams{
		Context: ctx,
	}
}

// NewCreateEpicParamsWithHTTPClient creates a new CreateEpicParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEpicParamsWithHTTPClient(client *http.Client) *CreateEpicParams {
	return &CreateEpicParams{
		HTTPClient: client,
	}
}

/* CreateEpicParams contains all the parameters to send to the API endpoint
   for the create epic operation.

   Typically these are written to a http.Request.
*/
type CreateEpicParams struct {

	// CreateEpic.
	CreateEpic *models.CreateEpic

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create epic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEpicParams) WithDefaults() *CreateEpicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create epic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEpicParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create epic params
func (o *CreateEpicParams) WithTimeout(timeout time.Duration) *CreateEpicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create epic params
func (o *CreateEpicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create epic params
func (o *CreateEpicParams) WithContext(ctx context.Context) *CreateEpicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create epic params
func (o *CreateEpicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create epic params
func (o *CreateEpicParams) WithHTTPClient(client *http.Client) *CreateEpicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create epic params
func (o *CreateEpicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateEpic adds the createEpic to the create epic params
func (o *CreateEpicParams) WithCreateEpic(createEpic *models.CreateEpic) *CreateEpicParams {
	o.SetCreateEpic(createEpic)
	return o
}

// SetCreateEpic adds the createEpic to the create epic params
func (o *CreateEpicParams) SetCreateEpic(createEpic *models.CreateEpic) {
	o.CreateEpic = createEpic
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEpicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateEpic != nil {
		if err := r.SetBodyParam(o.CreateEpic); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
