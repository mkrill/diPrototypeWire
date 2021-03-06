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

// NewSuggestDeprecatedZipcodesUsingGETParams creates a new SuggestDeprecatedZipcodesUsingGETParams object
// with the default values initialized.
func NewSuggestDeprecatedZipcodesUsingGETParams() *SuggestDeprecatedZipcodesUsingGETParams {
	var ()
	return &SuggestDeprecatedZipcodesUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSuggestDeprecatedZipcodesUsingGETParamsWithTimeout creates a new SuggestDeprecatedZipcodesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSuggestDeprecatedZipcodesUsingGETParamsWithTimeout(timeout time.Duration) *SuggestDeprecatedZipcodesUsingGETParams {
	var ()
	return &SuggestDeprecatedZipcodesUsingGETParams{

		timeout: timeout,
	}
}

// NewSuggestDeprecatedZipcodesUsingGETParamsWithContext creates a new SuggestDeprecatedZipcodesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewSuggestDeprecatedZipcodesUsingGETParamsWithContext(ctx context.Context) *SuggestDeprecatedZipcodesUsingGETParams {
	var ()
	return &SuggestDeprecatedZipcodesUsingGETParams{

		Context: ctx,
	}
}

// NewSuggestDeprecatedZipcodesUsingGETParamsWithHTTPClient creates a new SuggestDeprecatedZipcodesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSuggestDeprecatedZipcodesUsingGETParamsWithHTTPClient(client *http.Client) *SuggestDeprecatedZipcodesUsingGETParams {
	var ()
	return &SuggestDeprecatedZipcodesUsingGETParams{
		HTTPClient: client,
	}
}

/*SuggestDeprecatedZipcodesUsingGETParams contains all the parameters to send to the API endpoint
for the suggest deprecated zipcodes using g e t operation typically these are written to a http.Request
*/
type SuggestDeprecatedZipcodesUsingGETParams struct {

	/*City
	  City to filter zipcodes

	*/
	City *string
	/*Zipcode
	  Token to autocomplete

	*/
	Zipcode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the suggest deprecated zipcodes using g e t params
func (o *SuggestDeprecatedZipcodesUsingGETParams) WithTimeout(timeout time.Duration) *SuggestDeprecatedZipcodesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the suggest deprecated zipcodes using g e t params
func (o *SuggestDeprecatedZipcodesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the suggest deprecated zipcodes using g e t params
func (o *SuggestDeprecatedZipcodesUsingGETParams) WithContext(ctx context.Context) *SuggestDeprecatedZipcodesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the suggest deprecated zipcodes using g e t params
func (o *SuggestDeprecatedZipcodesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the suggest deprecated zipcodes using g e t params
func (o *SuggestDeprecatedZipcodesUsingGETParams) WithHTTPClient(client *http.Client) *SuggestDeprecatedZipcodesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the suggest deprecated zipcodes using g e t params
func (o *SuggestDeprecatedZipcodesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCity adds the city to the suggest deprecated zipcodes using g e t params
func (o *SuggestDeprecatedZipcodesUsingGETParams) WithCity(city *string) *SuggestDeprecatedZipcodesUsingGETParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the suggest deprecated zipcodes using g e t params
func (o *SuggestDeprecatedZipcodesUsingGETParams) SetCity(city *string) {
	o.City = city
}

// WithZipcode adds the zipcode to the suggest deprecated zipcodes using g e t params
func (o *SuggestDeprecatedZipcodesUsingGETParams) WithZipcode(zipcode string) *SuggestDeprecatedZipcodesUsingGETParams {
	o.SetZipcode(zipcode)
	return o
}

// SetZipcode adds the zipcode to the suggest deprecated zipcodes using g e t params
func (o *SuggestDeprecatedZipcodesUsingGETParams) SetZipcode(zipcode string) {
	o.Zipcode = zipcode
}

// WriteToRequest writes these params to a swagger request
func (o *SuggestDeprecatedZipcodesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.City != nil {

		// query param city
		var qrCity string
		if o.City != nil {
			qrCity = *o.City
		}
		qCity := qrCity
		if qCity != "" {
			if err := r.SetQueryParam("city", qCity); err != nil {
				return err
			}
		}

	}

	// query param zipcode
	qrZipcode := o.Zipcode
	qZipcode := qrZipcode
	if qZipcode != "" {
		if err := r.SetQueryParam("zipcode", qZipcode); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
