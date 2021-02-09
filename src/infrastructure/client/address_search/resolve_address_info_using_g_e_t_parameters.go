// Code generated by go-swagger; DO NOT EDIT.

package address_search

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

// NewResolveAddressInfoUsingGETParams creates a new ResolveAddressInfoUsingGETParams object
// with the default values initialized.
func NewResolveAddressInfoUsingGETParams() *ResolveAddressInfoUsingGETParams {
	var ()
	return &ResolveAddressInfoUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResolveAddressInfoUsingGETParamsWithTimeout creates a new ResolveAddressInfoUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResolveAddressInfoUsingGETParamsWithTimeout(timeout time.Duration) *ResolveAddressInfoUsingGETParams {
	var ()
	return &ResolveAddressInfoUsingGETParams{

		timeout: timeout,
	}
}

// NewResolveAddressInfoUsingGETParamsWithContext creates a new ResolveAddressInfoUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewResolveAddressInfoUsingGETParamsWithContext(ctx context.Context) *ResolveAddressInfoUsingGETParams {
	var ()
	return &ResolveAddressInfoUsingGETParams{

		Context: ctx,
	}
}

// NewResolveAddressInfoUsingGETParamsWithHTTPClient creates a new ResolveAddressInfoUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResolveAddressInfoUsingGETParamsWithHTTPClient(client *http.Client) *ResolveAddressInfoUsingGETParams {
	var ()
	return &ResolveAddressInfoUsingGETParams{
		HTTPClient: client,
	}
}

/*ResolveAddressInfoUsingGETParams contains all the parameters to send to the API endpoint
for the resolve address info using g e t operation typically these are written to a http.Request
*/
type ResolveAddressInfoUsingGETParams struct {

	/*HouseQualifier
	  Full house number

	*/
	HouseQualifier string
	/*MunicipalityName
	  Municipality (city) name

	*/
	MunicipalityName string
	/*StreetName
	  Street name

	*/
	StreetName string
	/*ZipCode
	  Zip code

	*/
	ZipCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) WithTimeout(timeout time.Duration) *ResolveAddressInfoUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) WithContext(ctx context.Context) *ResolveAddressInfoUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) WithHTTPClient(client *http.Client) *ResolveAddressInfoUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHouseQualifier adds the houseQualifier to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) WithHouseQualifier(houseQualifier string) *ResolveAddressInfoUsingGETParams {
	o.SetHouseQualifier(houseQualifier)
	return o
}

// SetHouseQualifier adds the houseQualifier to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) SetHouseQualifier(houseQualifier string) {
	o.HouseQualifier = houseQualifier
}

// WithMunicipalityName adds the municipalityName to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) WithMunicipalityName(municipalityName string) *ResolveAddressInfoUsingGETParams {
	o.SetMunicipalityName(municipalityName)
	return o
}

// SetMunicipalityName adds the municipalityName to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) SetMunicipalityName(municipalityName string) {
	o.MunicipalityName = municipalityName
}

// WithStreetName adds the streetName to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) WithStreetName(streetName string) *ResolveAddressInfoUsingGETParams {
	o.SetStreetName(streetName)
	return o
}

// SetStreetName adds the streetName to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) SetStreetName(streetName string) {
	o.StreetName = streetName
}

// WithZipCode adds the zipCode to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) WithZipCode(zipCode string) *ResolveAddressInfoUsingGETParams {
	o.SetZipCode(zipCode)
	return o
}

// SetZipCode adds the zipCode to the resolve address info using g e t params
func (o *ResolveAddressInfoUsingGETParams) SetZipCode(zipCode string) {
	o.ZipCode = zipCode
}

// WriteToRequest writes these params to a swagger request
func (o *ResolveAddressInfoUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param houseQualifier
	qrHouseQualifier := o.HouseQualifier
	qHouseQualifier := qrHouseQualifier
	if qHouseQualifier != "" {
		if err := r.SetQueryParam("houseQualifier", qHouseQualifier); err != nil {
			return err
		}
	}

	// query param municipalityName
	qrMunicipalityName := o.MunicipalityName
	qMunicipalityName := qrMunicipalityName
	if qMunicipalityName != "" {
		if err := r.SetQueryParam("municipalityName", qMunicipalityName); err != nil {
			return err
		}
	}

	// query param streetName
	qrStreetName := o.StreetName
	qStreetName := qrStreetName
	if qStreetName != "" {
		if err := r.SetQueryParam("streetName", qStreetName); err != nil {
			return err
		}
	}

	// query param zipCode
	qrZipCode := o.ZipCode
	qZipCode := qrZipCode
	if qZipCode != "" {
		if err := r.SetQueryParam("zipCode", qZipCode); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
