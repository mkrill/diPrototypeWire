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
)

// NewAddressSuggestionsUsingGETParams creates a new AddressSuggestionsUsingGETParams object
// with the default values initialized.
func NewAddressSuggestionsUsingGETParams() *AddressSuggestionsUsingGETParams {
	var ()
	return &AddressSuggestionsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddressSuggestionsUsingGETParamsWithTimeout creates a new AddressSuggestionsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddressSuggestionsUsingGETParamsWithTimeout(timeout time.Duration) *AddressSuggestionsUsingGETParams {
	var ()
	return &AddressSuggestionsUsingGETParams{

		timeout: timeout,
	}
}

// NewAddressSuggestionsUsingGETParamsWithContext creates a new AddressSuggestionsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddressSuggestionsUsingGETParamsWithContext(ctx context.Context) *AddressSuggestionsUsingGETParams {
	var ()
	return &AddressSuggestionsUsingGETParams{

		Context: ctx,
	}
}

// NewAddressSuggestionsUsingGETParamsWithHTTPClient creates a new AddressSuggestionsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddressSuggestionsUsingGETParamsWithHTTPClient(client *http.Client) *AddressSuggestionsUsingGETParams {
	var ()
	return &AddressSuggestionsUsingGETParams{
		HTTPClient: client,
	}
}

/*AddressSuggestionsUsingGETParams contains all the parameters to send to the API endpoint
for the address suggestions using g e t operation typically these are written to a http.Request
*/
type AddressSuggestionsUsingGETParams struct {

	/*HouseQualifier
	  Full house number(should be filled for suggestion type: HOUSE)

	*/
	HouseQualifier *string
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

// WithTimeout adds the timeout to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) WithTimeout(timeout time.Duration) *AddressSuggestionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) WithContext(ctx context.Context) *AddressSuggestionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) WithHTTPClient(client *http.Client) *AddressSuggestionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHouseQualifier adds the houseQualifier to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) WithHouseQualifier(houseQualifier *string) *AddressSuggestionsUsingGETParams {
	o.SetHouseQualifier(houseQualifier)
	return o
}

// SetHouseQualifier adds the houseQualifier to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) SetHouseQualifier(houseQualifier *string) {
	o.HouseQualifier = houseQualifier
}

// WithLocalityName adds the localityName to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) WithLocalityName(localityName *string) *AddressSuggestionsUsingGETParams {
	o.SetLocalityName(localityName)
	return o
}

// SetLocalityName adds the localityName to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) SetLocalityName(localityName *string) {
	o.LocalityName = localityName
}

// WithStreetName adds the streetName to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) WithStreetName(streetName *string) *AddressSuggestionsUsingGETParams {
	o.SetStreetName(streetName)
	return o
}

// SetStreetName adds the streetName to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) SetStreetName(streetName *string) {
	o.StreetName = streetName
}

// WithSuggestionType adds the suggestionType to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) WithSuggestionType(suggestionType string) *AddressSuggestionsUsingGETParams {
	o.SetSuggestionType(suggestionType)
	return o
}

// SetSuggestionType adds the suggestionType to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) SetSuggestionType(suggestionType string) {
	o.SuggestionType = suggestionType
}

// WithZipCode adds the zipCode to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) WithZipCode(zipCode *string) *AddressSuggestionsUsingGETParams {
	o.SetZipCode(zipCode)
	return o
}

// SetZipCode adds the zipCode to the address suggestions using g e t params
func (o *AddressSuggestionsUsingGETParams) SetZipCode(zipCode *string) {
	o.ZipCode = zipCode
}

// WriteToRequest writes these params to a swagger request
func (o *AddressSuggestionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HouseQualifier != nil {

		// query param houseQualifier
		var qrHouseQualifier string
		if o.HouseQualifier != nil {
			qrHouseQualifier = *o.HouseQualifier
		}
		qHouseQualifier := qrHouseQualifier
		if qHouseQualifier != "" {
			if err := r.SetQueryParam("houseQualifier", qHouseQualifier); err != nil {
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