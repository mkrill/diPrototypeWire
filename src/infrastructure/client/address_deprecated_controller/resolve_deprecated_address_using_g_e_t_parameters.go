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

// NewResolveDeprecatedAddressUsingGETParams creates a new ResolveDeprecatedAddressUsingGETParams object
// with the default values initialized.
func NewResolveDeprecatedAddressUsingGETParams() *ResolveDeprecatedAddressUsingGETParams {
	var ()
	return &ResolveDeprecatedAddressUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResolveDeprecatedAddressUsingGETParamsWithTimeout creates a new ResolveDeprecatedAddressUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResolveDeprecatedAddressUsingGETParamsWithTimeout(timeout time.Duration) *ResolveDeprecatedAddressUsingGETParams {
	var ()
	return &ResolveDeprecatedAddressUsingGETParams{

		timeout: timeout,
	}
}

// NewResolveDeprecatedAddressUsingGETParamsWithContext creates a new ResolveDeprecatedAddressUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewResolveDeprecatedAddressUsingGETParamsWithContext(ctx context.Context) *ResolveDeprecatedAddressUsingGETParams {
	var ()
	return &ResolveDeprecatedAddressUsingGETParams{

		Context: ctx,
	}
}

// NewResolveDeprecatedAddressUsingGETParamsWithHTTPClient creates a new ResolveDeprecatedAddressUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResolveDeprecatedAddressUsingGETParamsWithHTTPClient(client *http.Client) *ResolveDeprecatedAddressUsingGETParams {
	var ()
	return &ResolveDeprecatedAddressUsingGETParams{
		HTTPClient: client,
	}
}

/*ResolveDeprecatedAddressUsingGETParams contains all the parameters to send to the API endpoint
for the resolve deprecated address using g e t operation typically these are written to a http.Request
*/
type ResolveDeprecatedAddressUsingGETParams struct {

	/*City
	  city

	*/
	City string
	/*Street
	  street

	*/
	Street string
	/*Streetnumber
	  streetnumber

	*/
	Streetnumber string
	/*Zipcode
	  zipcode

	*/
	Zipcode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) WithTimeout(timeout time.Duration) *ResolveDeprecatedAddressUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) WithContext(ctx context.Context) *ResolveDeprecatedAddressUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) WithHTTPClient(client *http.Client) *ResolveDeprecatedAddressUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCity adds the city to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) WithCity(city string) *ResolveDeprecatedAddressUsingGETParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) SetCity(city string) {
	o.City = city
}

// WithStreet adds the street to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) WithStreet(street string) *ResolveDeprecatedAddressUsingGETParams {
	o.SetStreet(street)
	return o
}

// SetStreet adds the street to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) SetStreet(street string) {
	o.Street = street
}

// WithStreetnumber adds the streetnumber to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) WithStreetnumber(streetnumber string) *ResolveDeprecatedAddressUsingGETParams {
	o.SetStreetnumber(streetnumber)
	return o
}

// SetStreetnumber adds the streetnumber to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) SetStreetnumber(streetnumber string) {
	o.Streetnumber = streetnumber
}

// WithZipcode adds the zipcode to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) WithZipcode(zipcode string) *ResolveDeprecatedAddressUsingGETParams {
	o.SetZipcode(zipcode)
	return o
}

// SetZipcode adds the zipcode to the resolve deprecated address using g e t params
func (o *ResolveDeprecatedAddressUsingGETParams) SetZipcode(zipcode string) {
	o.Zipcode = zipcode
}

// WriteToRequest writes these params to a swagger request
func (o *ResolveDeprecatedAddressUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param street
	qrStreet := o.Street
	qStreet := qrStreet
	if qStreet != "" {
		if err := r.SetQueryParam("street", qStreet); err != nil {
			return err
		}
	}

	// query param streetnumber
	qrStreetnumber := o.Streetnumber
	qStreetnumber := qrStreetnumber
	if qStreetnumber != "" {
		if err := r.SetQueryParam("streetnumber", qStreetnumber); err != nil {
			return err
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
