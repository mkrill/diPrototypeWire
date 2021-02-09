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

// NewSuggestDeprecatedStreetsUsingGETParams creates a new SuggestDeprecatedStreetsUsingGETParams object
// with the default values initialized.
func NewSuggestDeprecatedStreetsUsingGETParams() *SuggestDeprecatedStreetsUsingGETParams {
	var ()
	return &SuggestDeprecatedStreetsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSuggestDeprecatedStreetsUsingGETParamsWithTimeout creates a new SuggestDeprecatedStreetsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSuggestDeprecatedStreetsUsingGETParamsWithTimeout(timeout time.Duration) *SuggestDeprecatedStreetsUsingGETParams {
	var ()
	return &SuggestDeprecatedStreetsUsingGETParams{

		timeout: timeout,
	}
}

// NewSuggestDeprecatedStreetsUsingGETParamsWithContext creates a new SuggestDeprecatedStreetsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewSuggestDeprecatedStreetsUsingGETParamsWithContext(ctx context.Context) *SuggestDeprecatedStreetsUsingGETParams {
	var ()
	return &SuggestDeprecatedStreetsUsingGETParams{

		Context: ctx,
	}
}

// NewSuggestDeprecatedStreetsUsingGETParamsWithHTTPClient creates a new SuggestDeprecatedStreetsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSuggestDeprecatedStreetsUsingGETParamsWithHTTPClient(client *http.Client) *SuggestDeprecatedStreetsUsingGETParams {
	var ()
	return &SuggestDeprecatedStreetsUsingGETParams{
		HTTPClient: client,
	}
}

/*SuggestDeprecatedStreetsUsingGETParams contains all the parameters to send to the API endpoint
for the suggest deprecated streets using g e t operation typically these are written to a http.Request
*/
type SuggestDeprecatedStreetsUsingGETParams struct {

	/*City
	  City to filter streets

	*/
	City *string
	/*Street
	  Token to autocomplete

	*/
	Street string
	/*Zipcode
	  Zipcode to filter streets

	*/
	Zipcode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) WithTimeout(timeout time.Duration) *SuggestDeprecatedStreetsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) WithContext(ctx context.Context) *SuggestDeprecatedStreetsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) WithHTTPClient(client *http.Client) *SuggestDeprecatedStreetsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCity adds the city to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) WithCity(city *string) *SuggestDeprecatedStreetsUsingGETParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) SetCity(city *string) {
	o.City = city
}

// WithStreet adds the street to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) WithStreet(street string) *SuggestDeprecatedStreetsUsingGETParams {
	o.SetStreet(street)
	return o
}

// SetStreet adds the street to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) SetStreet(street string) {
	o.Street = street
}

// WithZipcode adds the zipcode to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) WithZipcode(zipcode *string) *SuggestDeprecatedStreetsUsingGETParams {
	o.SetZipcode(zipcode)
	return o
}

// SetZipcode adds the zipcode to the suggest deprecated streets using g e t params
func (o *SuggestDeprecatedStreetsUsingGETParams) SetZipcode(zipcode *string) {
	o.Zipcode = zipcode
}

// WriteToRequest writes these params to a swagger request
func (o *SuggestDeprecatedStreetsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param street
	qrStreet := o.Street
	qStreet := qrStreet
	if qStreet != "" {
		if err := r.SetQueryParam("street", qStreet); err != nil {
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