// Code generated by go-swagger; DO NOT EDIT.

package basic_error_controller

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

// NewErrorUsingPOSTParams creates a new ErrorUsingPOSTParams object
// with the default values initialized.
func NewErrorUsingPOSTParams() *ErrorUsingPOSTParams {

	return &ErrorUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewErrorUsingPOSTParamsWithTimeout creates a new ErrorUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewErrorUsingPOSTParamsWithTimeout(timeout time.Duration) *ErrorUsingPOSTParams {

	return &ErrorUsingPOSTParams{

		timeout: timeout,
	}
}

// NewErrorUsingPOSTParamsWithContext creates a new ErrorUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewErrorUsingPOSTParamsWithContext(ctx context.Context) *ErrorUsingPOSTParams {

	return &ErrorUsingPOSTParams{

		Context: ctx,
	}
}

// NewErrorUsingPOSTParamsWithHTTPClient creates a new ErrorUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewErrorUsingPOSTParamsWithHTTPClient(client *http.Client) *ErrorUsingPOSTParams {

	return &ErrorUsingPOSTParams{
		HTTPClient: client,
	}
}

/*ErrorUsingPOSTParams contains all the parameters to send to the API endpoint
for the error using p o s t operation typically these are written to a http.Request
*/
type ErrorUsingPOSTParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the error using p o s t params
func (o *ErrorUsingPOSTParams) WithTimeout(timeout time.Duration) *ErrorUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the error using p o s t params
func (o *ErrorUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the error using p o s t params
func (o *ErrorUsingPOSTParams) WithContext(ctx context.Context) *ErrorUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the error using p o s t params
func (o *ErrorUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the error using p o s t params
func (o *ErrorUsingPOSTParams) WithHTTPClient(client *http.Client) *ErrorUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the error using p o s t params
func (o *ErrorUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ErrorUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
