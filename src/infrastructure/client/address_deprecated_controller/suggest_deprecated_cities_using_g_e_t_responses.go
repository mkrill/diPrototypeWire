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

// SuggestDeprecatedCitiesUsingGETReader is a Reader for the SuggestDeprecatedCitiesUsingGET structure.
type SuggestDeprecatedCitiesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuggestDeprecatedCitiesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuggestDeprecatedCitiesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSuggestDeprecatedCitiesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSuggestDeprecatedCitiesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSuggestDeprecatedCitiesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSuggestDeprecatedCitiesUsingGETOK creates a SuggestDeprecatedCitiesUsingGETOK with default headers values
func NewSuggestDeprecatedCitiesUsingGETOK() *SuggestDeprecatedCitiesUsingGETOK {
	return &SuggestDeprecatedCitiesUsingGETOK{}
}

/*SuggestDeprecatedCitiesUsingGETOK handles this case with default header values.

OK
*/
type SuggestDeprecatedCitiesUsingGETOK struct {
	Payload []*models.Address
}

func (o *SuggestDeprecatedCitiesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/cities][%d] suggestDeprecatedCitiesUsingGETOK  %+v", 200, o.Payload)
}

func (o *SuggestDeprecatedCitiesUsingGETOK) GetPayload() []*models.Address {
	return o.Payload
}

func (o *SuggestDeprecatedCitiesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestDeprecatedCitiesUsingGETUnauthorized creates a SuggestDeprecatedCitiesUsingGETUnauthorized with default headers values
func NewSuggestDeprecatedCitiesUsingGETUnauthorized() *SuggestDeprecatedCitiesUsingGETUnauthorized {
	return &SuggestDeprecatedCitiesUsingGETUnauthorized{}
}

/*SuggestDeprecatedCitiesUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type SuggestDeprecatedCitiesUsingGETUnauthorized struct {
}

func (o *SuggestDeprecatedCitiesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/cities][%d] suggestDeprecatedCitiesUsingGETUnauthorized ", 401)
}

func (o *SuggestDeprecatedCitiesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSuggestDeprecatedCitiesUsingGETForbidden creates a SuggestDeprecatedCitiesUsingGETForbidden with default headers values
func NewSuggestDeprecatedCitiesUsingGETForbidden() *SuggestDeprecatedCitiesUsingGETForbidden {
	return &SuggestDeprecatedCitiesUsingGETForbidden{}
}

/*SuggestDeprecatedCitiesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type SuggestDeprecatedCitiesUsingGETForbidden struct {
}

func (o *SuggestDeprecatedCitiesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/cities][%d] suggestDeprecatedCitiesUsingGETForbidden ", 403)
}

func (o *SuggestDeprecatedCitiesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSuggestDeprecatedCitiesUsingGETNotFound creates a SuggestDeprecatedCitiesUsingGETNotFound with default headers values
func NewSuggestDeprecatedCitiesUsingGETNotFound() *SuggestDeprecatedCitiesUsingGETNotFound {
	return &SuggestDeprecatedCitiesUsingGETNotFound{}
}

/*SuggestDeprecatedCitiesUsingGETNotFound handles this case with default header values.

Not Found
*/
type SuggestDeprecatedCitiesUsingGETNotFound struct {
}

func (o *SuggestDeprecatedCitiesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/cities][%d] suggestDeprecatedCitiesUsingGETNotFound ", 404)
}

func (o *SuggestDeprecatedCitiesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
