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

// NewSuggestDeprecatedCitiesUsingGETParams creates a new SuggestDeprecatedCitiesUsingGETParams object
// with the default values initialized.
func NewSuggestDeprecatedCitiesUsingGETParams() *SuggestDeprecatedCitiesUsingGETParams {
	var ()
	return &SuggestDeprecatedCitiesUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSuggestDeprecatedCitiesUsingGETParamsWithTimeout creates a new SuggestDeprecatedCitiesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSuggestDeprecatedCitiesUsingGETParamsWithTimeout(timeout time.Duration) *SuggestDeprecatedCitiesUsingGETParams {
	var ()
	return &SuggestDeprecatedCitiesUsingGETParams{

		timeout: timeout,
	}
}

// NewSuggestDeprecatedCitiesUsingGETParamsWithContext creates a new SuggestDeprecatedCitiesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewSuggestDeprecatedCitiesUsingGETParamsWithContext(ctx context.Context) *SuggestDeprecatedCitiesUsingGETParams {
	var ()
	return &SuggestDeprecatedCitiesUsingGETParams{

		Context: ctx,
	}
}

// NewSuggestDeprecatedCitiesUsingGETParamsWithHTTPClient creates a new SuggestDeprecatedCitiesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSuggestDeprecatedCitiesUsingGETParamsWithHTTPClient(client *http.Client) *SuggestDeprecatedCitiesUsingGETParams {
	var ()
	return &SuggestDeprecatedCitiesUsingGETParams{
		HTTPClient: client,
	}
}

/*SuggestDeprecatedCitiesUsingGETParams contains all the parameters to send to the API endpoint
for the suggest deprecated cities using g e t operation typically these are written to a http.Request
*/
type SuggestDeprecatedCitiesUsingGETParams struct {

	/*City
	  Token to autocomplete

	*/
	City string
	/*Zipcode
	  Zipcode to filter cities

	*/
	Zipcode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the suggest deprecated cities using g e t params
func (o *SuggestDeprecatedCitiesUsingGETParams) WithTimeout(timeout time.Duration) *SuggestDeprecatedCitiesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the suggest deprecated cities using g e t params
func (o *SuggestDeprecatedCitiesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the suggest deprecated cities using g e t params
func (o *SuggestDeprecatedCitiesUsingGETParams) WithContext(ctx context.Context) *SuggestDeprecatedCitiesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the suggest deprecated cities using g e t params
func (o *SuggestDeprecatedCitiesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the suggest deprecated cities using g e t params
func (o *SuggestDeprecatedCitiesUsingGETParams) WithHTTPClient(client *http.Client) *SuggestDeprecatedCitiesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the suggest deprecated cities using g e t params
func (o *SuggestDeprecatedCitiesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCity adds the city to the suggest deprecated cities using g e t params
func (o *SuggestDeprecatedCitiesUsingGETParams) WithCity(city string) *SuggestDeprecatedCitiesUsingGETParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the suggest deprecated cities using g e t params
func (o *SuggestDeprecatedCitiesUsingGETParams) SetCity(city string) {
	o.City = city
}

// WithZipcode adds the zipcode to the suggest deprecated cities using g e t params
func (o *SuggestDeprecatedCitiesUsingGETParams) WithZipcode(zipcode *string) *SuggestDeprecatedCitiesUsingGETParams {
	o.SetZipcode(zipcode)
	return o
}

// SetZipcode adds the zipcode to the suggest deprecated cities using g e t params
func (o *SuggestDeprecatedCitiesUsingGETParams) SetZipcode(zipcode *string) {
	o.Zipcode = zipcode
}

// WriteToRequest writes these params to a swagger request
func (o *SuggestDeprecatedCitiesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param city
	qrCity := o.City
	qCity := qrCity
	if qCity != "" {
		if err := r.SetQueryParam("city", qCity); err != nil {
			return err
		}
	}

	if o.Zipcode != nil {

		// query param zipcode
		var qrZipcode string
		if o.Zipcode != nil {
			qrZipcode = *o.Zipcode
		}
		qZipcode := qrZipcode
		if qZipcode != "" {
			if err := r.SetQueryParam("zipcode", qZipcode); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
