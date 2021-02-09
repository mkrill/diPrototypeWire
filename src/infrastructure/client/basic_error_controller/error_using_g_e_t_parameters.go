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

// NewErrorUsingGETParams creates a new ErrorUsingGETParams object
// with the default values initialized.
func NewErrorUsingGETParams() *ErrorUsingGETParams {

	return &ErrorUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewErrorUsingGETParamsWithTimeout creates a new ErrorUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewErrorUsingGETParamsWithTimeout(timeout time.Duration) *ErrorUsingGETParams {

	return &ErrorUsingGETParams{

		timeout: timeout,
	}
}

// NewErrorUsingGETParamsWithContext creates a new ErrorUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewErrorUsingGETParamsWithContext(ctx context.Context) *ErrorUsingGETParams {

	return &ErrorUsingGETParams{

		Context: ctx,
	}
}

// NewErrorUsingGETParamsWithHTTPClient creates a new ErrorUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewErrorUsingGETParamsWithHTTPClient(client *http.Client) *ErrorUsingGETParams {

	return &ErrorUsingGETParams{
		HTTPClient: client,
	}
}

/*ErrorUsingGETParams contains all the parameters to send to the API endpoint
for the error using g e t operation typically these are written to a http.Request
*/
type ErrorUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the error using g e t params
func (o *ErrorUsingGETParams) WithTimeout(timeout time.Duration) *ErrorUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the error using g e t params
func (o *ErrorUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the error using g e t params
func (o *ErrorUsingGETParams) WithContext(ctx context.Context) *ErrorUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the error using g e t params
func (o *ErrorUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the error using g e t params
func (o *ErrorUsingGETParams) WithHTTPClient(client *http.Client) *ErrorUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the error using g e t params
func (o *ErrorUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ErrorUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
