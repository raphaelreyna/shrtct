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

// NewListEpicCommentsParams creates a new ListEpicCommentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListEpicCommentsParams() *ListEpicCommentsParams {
	return &ListEpicCommentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListEpicCommentsParamsWithTimeout creates a new ListEpicCommentsParams object
// with the ability to set a timeout on a request.
func NewListEpicCommentsParamsWithTimeout(timeout time.Duration) *ListEpicCommentsParams {
	return &ListEpicCommentsParams{
		timeout: timeout,
	}
}

// NewListEpicCommentsParamsWithContext creates a new ListEpicCommentsParams object
// with the ability to set a context for a request.
func NewListEpicCommentsParamsWithContext(ctx context.Context) *ListEpicCommentsParams {
	return &ListEpicCommentsParams{
		Context: ctx,
	}
}

// NewListEpicCommentsParamsWithHTTPClient creates a new ListEpicCommentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListEpicCommentsParamsWithHTTPClient(client *http.Client) *ListEpicCommentsParams {
	return &ListEpicCommentsParams{
		HTTPClient: client,
	}
}

/* ListEpicCommentsParams contains all the parameters to send to the API endpoint
   for the list epic comments operation.

   Typically these are written to a http.Request.
*/
type ListEpicCommentsParams struct {

	/* EpicPublicID.

	   The unique ID of the Epic.

	   Format: int64
	*/
	EpicPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list epic comments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEpicCommentsParams) WithDefaults() *ListEpicCommentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list epic comments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEpicCommentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list epic comments params
func (o *ListEpicCommentsParams) WithTimeout(timeout time.Duration) *ListEpicCommentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list epic comments params
func (o *ListEpicCommentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list epic comments params
func (o *ListEpicCommentsParams) WithContext(ctx context.Context) *ListEpicCommentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list epic comments params
func (o *ListEpicCommentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list epic comments params
func (o *ListEpicCommentsParams) WithHTTPClient(client *http.Client) *ListEpicCommentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list epic comments params
func (o *ListEpicCommentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEpicPublicID adds the epicPublicID to the list epic comments params
func (o *ListEpicCommentsParams) WithEpicPublicID(epicPublicID int64) *ListEpicCommentsParams {
	o.SetEpicPublicID(epicPublicID)
	return o
}

// SetEpicPublicID adds the epicPublicId to the list epic comments params
func (o *ListEpicCommentsParams) SetEpicPublicID(epicPublicID int64) {
	o.EpicPublicID = epicPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *ListEpicCommentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param epic-public-id
	if err := r.SetPathParam("epic-public-id", swag.FormatInt64(o.EpicPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}