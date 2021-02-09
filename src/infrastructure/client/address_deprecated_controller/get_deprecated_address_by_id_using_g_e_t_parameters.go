// Code generated by go-swagger; DO NOT EDIT.

package address_deprecated_controller

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

// NewGetDeprecatedAddressByIDUsingGETParams creates a new GetDeprecatedAddressByIDUsingGETParams object
// with the default values initialized.
func NewGetDeprecatedAddressByIDUsingGETParams() *GetDeprecatedAddressByIDUsingGETParams {
	var ()
	return &GetDeprecatedAddressByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeprecatedAddressByIDUsingGETParamsWithTimeout creates a new GetDeprecatedAddressByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeprecatedAddressByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetDeprecatedAddressByIDUsingGETParams {
	var ()
	return &GetDeprecatedAddressByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetDeprecatedAddressByIDUsingGETParamsWithContext creates a new GetDeprecatedAddressByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeprecatedAddressByIDUsingGETParamsWithContext(ctx context.Context) *GetDeprecatedAddressByIDUsingGETParams {
	var ()
	return &GetDeprecatedAddressByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetDeprecatedAddressByIDUsingGETParamsWithHTTPClient creates a new GetDeprecatedAddressByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeprecatedAddressByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetDeprecatedAddressByIDUsingGETParams {
	var ()
	return &GetDeprecatedAddressByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetDeprecatedAddressByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get deprecated address by Id using g e t operation typically these are written to a http.Request
*/
type GetDeprecatedAddressByIDUsingGETParams struct {

	/*KlsID
	  klsId

	*/
	KlsID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deprecated address by Id using g e t params
func (o *GetDeprecatedAddressByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetDeprecatedAddressByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deprecated address by Id using g e t params
func (o *GetDeprecatedAddressByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deprecated address by Id using g e t params
func (o *GetDeprecatedAddressByIDUsingGETParams) WithContext(ctx context.Context) *GetDeprecatedAddressByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deprecated address by Id using g e t params
func (o *GetDeprecatedAddressByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deprecated address by Id using g e t params
func (o *GetDeprecatedAddressByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetDeprecatedAddressByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deprecated address by Id using g e t params
func (o *GetDeprecatedAddressByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKlsID adds the klsID to the get deprecated address by Id using g e t params
func (o *GetDeprecatedAddressByIDUsingGETParams) WithKlsID(klsID string) *GetDeprecatedAddressByIDUsingGETParams {
	o.SetKlsID(klsID)
	return o
}

// SetKlsID adds the klsId to the get deprecated address by Id using g e t params
func (o *GetDeprecatedAddressByIDUsingGETParams) SetKlsID(klsID string) {
	o.KlsID = klsID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeprecatedAddressByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param klsId
	if err := r.SetPathParam("klsId", o.KlsID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
