// Code generated by go-swagger; DO NOT EDIT.

package address_suggest_v2

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

// NewGetAddressSuggestionsUsingGETParams creates a new GetAddressSuggestionsUsingGETParams object
// with the default values initialized.
func NewGetAddressSuggestionsUsingGETParams() *GetAddressSuggestionsUsingGETParams {
	var ()
	return &GetAddressSuggestionsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAddressSuggestionsUsingGETParamsWithTimeout creates a new GetAddressSuggestionsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAddressSuggestionsUsingGETParamsWithTimeout(timeout time.Duration) *GetAddressSuggestionsUsingGETParams {
	var ()
	return &GetAddressSuggestionsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetAddressSuggestionsUsingGETParamsWithContext creates a new GetAddressSuggestionsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAddressSuggestionsUsingGETParamsWithContext(ctx context.Context) *GetAddressSuggestionsUsingGETParams {
	var ()
	return &GetAddressSuggestionsUsingGETParams{

		Context: ctx,
	}
}

// NewGetAddressSuggestionsUsingGETParamsWithHTTPClient creates a new GetAddressSuggestionsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAddressSuggestionsUsingGETParamsWithHTTPClient(client *http.Client) *GetAddressSuggestionsUsingGETParams {
	var ()
	return &GetAddressSuggestionsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetAddressSuggestionsUsingGETParams contains all the parameters to send to the API endpoint
for the get address suggestions using g e t operation typically these are written to a http.Request
*/
type GetAddressSuggestionsUsingGETParams struct {

	/*Address
	  Address string, by wich suggestion will be done

	*/
	Address string
	/*Fields
	  Response filtering fields. Logic is not implemented yet, values are ignored

	*/
	Fields []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get address suggestions using g e t params
func (o *GetAddressSuggestionsUsingGETParams) WithTimeout(timeout time.Duration) *GetAddressSuggestionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get address suggestions using g e t params
func (o *GetAddressSuggestionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get address suggestions using g e t params
func (o *GetAddressSuggestionsUsingGETParams) WithContext(ctx context.Context) *GetAddressSuggestionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get address suggestions using g e t params
func (o *GetAddressSuggestionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get address suggestions using g e t params
func (o *GetAddressSuggestionsUsingGETParams) WithHTTPClient(client *http.Client) *GetAddressSuggestionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get address suggestions using g e t params
func (o *GetAddressSuggestionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the get address suggestions using g e t params
func (o *GetAddressSuggestionsUsingGETParams) WithAddress(address string) *GetAddressSuggestionsUsingGETParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the get address suggestions using g e t params
func (o *GetAddressSuggestionsUsingGETParams) SetAddress(address string) {
	o.Address = address
}

// WithFields adds the fields to the get address suggestions using g e t params
func (o *GetAddressSuggestionsUsingGETParams) WithFields(fields []string) *GetAddressSuggestionsUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get address suggestions using g e t params
func (o *GetAddressSuggestionsUsingGETParams) SetFields(fields []string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetAddressSuggestionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param address
	qrAddress := o.Address
	qAddress := qrAddress
	if qAddress != "" {
		if err := r.SetQueryParam("address", qAddress); err != nil {
			return err
		}
	}

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "multi")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
