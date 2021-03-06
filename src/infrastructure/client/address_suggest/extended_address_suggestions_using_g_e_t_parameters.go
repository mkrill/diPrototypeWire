// Code generated by go-swagger; DO NOT EDIT.

package address_suggest

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

// NewExtendedAddressSuggestionsUsingGETParams creates a new ExtendedAddressSuggestionsUsingGETParams object
// with the default values initialized.
func NewExtendedAddressSuggestionsUsingGETParams() *ExtendedAddressSuggestionsUsingGETParams {
	var ()
	return &ExtendedAddressSuggestionsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtendedAddressSuggestionsUsingGETParamsWithTimeout creates a new ExtendedAddressSuggestionsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtendedAddressSuggestionsUsingGETParamsWithTimeout(timeout time.Duration) *ExtendedAddressSuggestionsUsingGETParams {
	var ()
	return &ExtendedAddressSuggestionsUsingGETParams{

		timeout: timeout,
	}
}

// NewExtendedAddressSuggestionsUsingGETParamsWithContext creates a new ExtendedAddressSuggestionsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtendedAddressSuggestionsUsingGETParamsWithContext(ctx context.Context) *ExtendedAddressSuggestionsUsingGETParams {
	var ()
	return &ExtendedAddressSuggestionsUsingGETParams{

		Context: ctx,
	}
}

// NewExtendedAddressSuggestionsUsingGETParamsWithHTTPClient creates a new ExtendedAddressSuggestionsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtendedAddressSuggestionsUsingGETParamsWithHTTPClient(client *http.Client) *ExtendedAddressSuggestionsUsingGETParams {
	var ()
	return &ExtendedAddressSuggestionsUsingGETParams{
		HTTPClient: client,
	}
}

/*ExtendedAddressSuggestionsUsingGETParams contains all the parameters to send to the API endpoint
for the extended address suggestions using g e t operation typically these are written to a http.Request
*/
type ExtendedAddressSuggestionsUsingGETParams struct {

	/*HouseNumber
	  House number(should be filled for suggestion type: HOUSE)

	*/
	HouseNumber *int32
	/*HouseNumberSupplement
	  House number supplement(should be filled for suggestion type: HOUSE, or could be empty if houseNumber is filled)

	*/
	HouseNumberSupplement *string
	/*LocalityName
	  Municipality(city) or DistributionArea(district) name(should be filled for suggestion type: LOCALITY, STREET, HOUSE)

	*/
	LocalityName *string
	/*StreetName
	  Street name(should be filled for suggestion type: STREET, HOUSE)

	*/
	StreetName *string
	/*SuggestionType
	  Suggestion Type

	*/
	SuggestionType string
	/*ZipCode
	  Zip code(should be filled for suggestion type: ZIPCODE, STREET, HOUSE)

	*/
	ZipCode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) WithTimeout(timeout time.Duration) *ExtendedAddressSuggestionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) WithContext(ctx context.Context) *ExtendedAddressSuggestionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) WithHTTPClient(client *http.Client) *ExtendedAddressSuggestionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHouseNumber adds the houseNumber to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) WithHouseNumber(houseNumber *int32) *ExtendedAddressSuggestionsUsingGETParams {
	o.SetHouseNumber(houseNumber)
	return o
}

// SetHouseNumber adds the houseNumber to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) SetHouseNumber(houseNumber *int32) {
	o.HouseNumber = houseNumber
}

// WithHouseNumberSupplement adds the houseNumberSupplement to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) WithHouseNumberSupplement(houseNumberSupplement *string) *ExtendedAddressSuggestionsUsingGETParams {
	o.SetHouseNumberSupplement(houseNumberSupplement)
	return o
}

// SetHouseNumberSupplement adds the houseNumberSupplement to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) SetHouseNumberSupplement(houseNumberSupplement *string) {
	o.HouseNumberSupplement = houseNumberSupplement
}

// WithLocalityName adds the localityName to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) WithLocalityName(localityName *string) *ExtendedAddressSuggestionsUsingGETParams {
	o.SetLocalityName(localityName)
	return o
}

// SetLocalityName adds the localityName to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) SetLocalityName(localityName *string) {
	o.LocalityName = localityName
}

// WithStreetName adds the streetName to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) WithStreetName(streetName *string) *ExtendedAddressSuggestionsUsingGETParams {
	o.SetStreetName(streetName)
	return o
}

// SetStreetName adds the streetName to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) SetStreetName(streetName *string) {
	o.StreetName = streetName
}

// WithSuggestionType adds the suggestionType to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) WithSuggestionType(suggestionType string) *ExtendedAddressSuggestionsUsingGETParams {
	o.SetSuggestionType(suggestionType)
	return o
}

// SetSuggestionType adds the suggestionType to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) SetSuggestionType(suggestionType string) {
	o.SuggestionType = suggestionType
}

// WithZipCode adds the zipCode to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) WithZipCode(zipCode *string) *ExtendedAddressSuggestionsUsingGETParams {
	o.SetZipCode(zipCode)
	return o
}

// SetZipCode adds the zipCode to the extended address suggestions using g e t params
func (o *ExtendedAddressSuggestionsUsingGETParams) SetZipCode(zipCode *string) {
	o.ZipCode = zipCode
}

// WriteToRequest writes these params to a swagger request
func (o *ExtendedAddressSuggestionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HouseNumber != nil {

		// query param houseNumber
		var qrHouseNumber int32
		if o.HouseNumber != nil {
			qrHouseNumber = *o.HouseNumber
		}
		qHouseNumber := swag.FormatInt32(qrHouseNumber)
		if qHouseNumber != "" {
			if err := r.SetQueryParam("houseNumber", qHouseNumber); err != nil {
				return err
			}
		}

	}

	if o.HouseNumberSupplement != nil {

		// query param houseNumberSupplement
		var qrHouseNumberSupplement string
		if o.HouseNumberSupplement != nil {
			qrHouseNumberSupplement = *o.HouseNumberSupplement
		}
		qHouseNumberSupplement := qrHouseNumberSupplement
		if qHouseNumberSupplement != "" {
			if err := r.SetQueryParam("houseNumberSupplement", qHouseNumberSupplement); err != nil {
				return err
			}
		}

	}

	if o.LocalityName != nil {

		// query param localityName
		var qrLocalityName string
		if o.LocalityName != nil {
			qrLocalityName = *o.LocalityName
		}
		qLocalityName := qrLocalityName
		if qLocalityName != "" {
			if err := r.SetQueryParam("localityName", qLocalityName); err != nil {
				return err
			}
		}

	}

	if o.StreetName != nil {

		// query param streetName
		var qrStreetName string
		if o.StreetName != nil {
			qrStreetName = *o.StreetName
		}
		qStreetName := qrStreetName
		if qStreetName != "" {
			if err := r.SetQueryParam("streetName", qStreetName); err != nil {
				return err
			}
		}

	}

	// path param suggestionType
	if err := r.SetPathParam("suggestionType", o.SuggestionType); err != nil {
		return err
	}

	if o.ZipCode != nil {

		// query param zipCode
		var qrZipCode string
		if o.ZipCode != nil {
			qrZipCode = *o.ZipCode
		}
		qZipCode := qrZipCode
		if qZipCode != "" {
			if err := r.SetQueryParam("zipCode", qZipCode); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
