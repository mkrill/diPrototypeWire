// Code generated by go-swagger; DO NOT EDIT.

package address_deprecated_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mkrill/diPrototypeWire/src/infrastructure/models"
)

// SuggestDeprecatedZipcodesUsingGETReader is a Reader for the SuggestDeprecatedZipcodesUsingGET structure.
type SuggestDeprecatedZipcodesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuggestDeprecatedZipcodesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuggestDeprecatedZipcodesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSuggestDeprecatedZipcodesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSuggestDeprecatedZipcodesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSuggestDeprecatedZipcodesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSuggestDeprecatedZipcodesUsingGETOK creates a SuggestDeprecatedZipcodesUsingGETOK with default headers values
func NewSuggestDeprecatedZipcodesUsingGETOK() *SuggestDeprecatedZipcodesUsingGETOK {
	return &SuggestDeprecatedZipcodesUsingGETOK{}
}

/*SuggestDeprecatedZipcodesUsingGETOK handles this case with default header values.

OK
*/
type SuggestDeprecatedZipcodesUsingGETOK struct {
	Payload []*models.Address
}

func (o *SuggestDeprecatedZipcodesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/zipcodes][%d] suggestDeprecatedZipcodesUsingGETOK  %+v", 200, o.Payload)
}

func (o *SuggestDeprecatedZipcodesUsingGETOK) GetPayload() []*models.Address {
	return o.Payload
}

func (o *SuggestDeprecatedZipcodesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestDeprecatedZipcodesUsingGETUnauthorized creates a SuggestDeprecatedZipcodesUsingGETUnauthorized with default headers values
func NewSuggestDeprecatedZipcodesUsingGETUnauthorized() *SuggestDeprecatedZipcodesUsingGETUnauthorized {
	return &SuggestDeprecatedZipcodesUsingGETUnauthorized{}
}

/*SuggestDeprecatedZipcodesUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type SuggestDeprecatedZipcodesUsingGETUnauthorized struct {
}

func (o *SuggestDeprecatedZipcodesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/zipcodes][%d] suggestDeprecatedZipcodesUsingGETUnauthorized ", 401)
}

func (o *SuggestDeprecatedZipcodesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSuggestDeprecatedZipcodesUsingGETForbidden creates a SuggestDeprecatedZipcodesUsingGETForbidden with default headers values
func NewSuggestDeprecatedZipcodesUsingGETForbidden() *SuggestDeprecatedZipcodesUsingGETForbidden {
	return &SuggestDeprecatedZipcodesUsingGETForbidden{}
}

/*SuggestDeprecatedZipcodesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type SuggestDeprecatedZipcodesUsingGETForbidden struct {
}

func (o *SuggestDeprecatedZipcodesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/zipcodes][%d] suggestDeprecatedZipcodesUsingGETForbidden ", 403)
}

func (o *SuggestDeprecatedZipcodesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSuggestDeprecatedZipcodesUsingGETNotFound creates a SuggestDeprecatedZipcodesUsingGETNotFound with default headers values
func NewSuggestDeprecatedZipcodesUsingGETNotFound() *SuggestDeprecatedZipcodesUsingGETNotFound {
	return &SuggestDeprecatedZipcodesUsingGETNotFound{}
}

/*SuggestDeprecatedZipcodesUsingGETNotFound handles this case with default header values.

Not Found
*/
type SuggestDeprecatedZipcodesUsingGETNotFound struct {
}

func (o *SuggestDeprecatedZipcodesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/zipcodes][%d] suggestDeprecatedZipcodesUsingGETNotFound ", 404)
}

func (o *SuggestDeprecatedZipcodesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
