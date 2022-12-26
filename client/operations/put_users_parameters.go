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

// NewPutUsersParams creates a new PutUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutUsersParams() *PutUsersParams {
	return &PutUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutUsersParamsWithTimeout creates a new PutUsersParams object
// with the ability to set a timeout on a request.
func NewPutUsersParamsWithTimeout(timeout time.Duration) *PutUsersParams {
	return &PutUsersParams{
		timeout: timeout,
	}
}

// NewPutUsersParamsWithContext creates a new PutUsersParams object
// with the ability to set a context for a request.
func NewPutUsersParamsWithContext(ctx context.Context) *PutUsersParams {
	return &PutUsersParams{
		Context: ctx,
	}
}

// NewPutUsersParamsWithHTTPClient creates a new PutUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutUsersParamsWithHTTPClient(client *http.Client) *PutUsersParams {
	return &PutUsersParams{
		HTTPClient: client,
	}
}

/*
PutUsersParams contains all the parameters to send to the API endpoint

	for the put users operation.

	Typically these are written to a http.Request.
*/
type PutUsersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutUsersParams) WithDefaults() *PutUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put users params
func (o *PutUsersParams) WithTimeout(timeout time.Duration) *PutUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put users params
func (o *PutUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put users params
func (o *PutUsersParams) WithContext(ctx context.Context) *PutUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put users params
func (o *PutUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put users params
func (o *PutUsersParams) WithHTTPClient(client *http.Client) *PutUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put users params
func (o *PutUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PutUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
