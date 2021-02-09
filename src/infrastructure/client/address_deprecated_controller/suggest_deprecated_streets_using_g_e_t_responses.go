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

// SuggestDeprecatedStreetsUsingGETReader is a Reader for the SuggestDeprecatedStreetsUsingGET structure.
type SuggestDeprecatedStreetsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuggestDeprecatedStreetsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuggestDeprecatedStreetsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSuggestDeprecatedStreetsUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSuggestDeprecatedStreetsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSuggestDeprecatedStreetsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSuggestDeprecatedStreetsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSuggestDeprecatedStreetsUsingGETOK creates a SuggestDeprecatedStreetsUsingGETOK with default headers values
func NewSuggestDeprecatedStreetsUsingGETOK() *SuggestDeprecatedStreetsUsingGETOK {
	return &SuggestDeprecatedStreetsUsingGETOK{}
}

/*SuggestDeprecatedStreetsUsingGETOK handles this case with default header values.

OK
*/
type SuggestDeprecatedStreetsUsingGETOK struct {
	Payload []*models.Address
}

func (o *SuggestDeprecatedStreetsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/streets][%d] suggestDeprecatedStreetsUsingGETOK  %+v", 200, o.Payload)
}

func (o *SuggestDeprecatedStreetsUsingGETOK) GetPayload() []*models.Address {
	return o.Payload
}

func (o *SuggestDeprecatedStreetsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestDeprecatedStreetsUsingGETBadRequest creates a SuggestDeprecatedStreetsUsingGETBadRequest with default headers values
func NewSuggestDeprecatedStreetsUsingGETBadRequest() *SuggestDeprecatedStreetsUsingGETBadRequest {
	return &SuggestDeprecatedStreetsUsingGETBadRequest{}
}

/*SuggestDeprecatedStreetsUsingGETBadRequest handles this case with default header values.

Filter missing
*/
type SuggestDeprecatedStreetsUsingGETBadRequest struct {
}

func (o *SuggestDeprecatedStreetsUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/streets][%d] suggestDeprecatedStreetsUsingGETBadRequest ", 400)
}

func (o *SuggestDeprecatedStreetsUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSuggestDeprecatedStreetsUsingGETUnauthorized creates a SuggestDeprecatedStreetsUsingGETUnauthorized with default headers values
func NewSuggestDeprecatedStreetsUsingGETUnauthorized() *SuggestDeprecatedStreetsUsingGETUnauthorized {
	return &SuggestDeprecatedStreetsUsingGETUnauthorized{}
}

/*SuggestDeprecatedStreetsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type SuggestDeprecatedStreetsUsingGETUnauthorized struct {
}

func (o *SuggestDeprecatedStreetsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/streets][%d] suggestDeprecatedStreetsUsingGETUnauthorized ", 401)
}

func (o *SuggestDeprecatedStreetsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSuggestDeprecatedStreetsUsingGETForbidden creates a SuggestDeprecatedStreetsUsingGETForbidden with default headers values
func NewSuggestDeprecatedStreetsUsingGETForbidden() *SuggestDeprecatedStreetsUsingGETForbidden {
	return &SuggestDeprecatedStreetsUsingGETForbidden{}
}

/*SuggestDeprecatedStreetsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type SuggestDeprecatedStreetsUsingGETForbidden struct {
}

func (o *SuggestDeprecatedStreetsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/streets][%d] suggestDeprecatedStreetsUsingGETForbidden ", 403)
}

func (o *SuggestDeprecatedStreetsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSuggestDeprecatedStreetsUsingGETNotFound creates a SuggestDeprecatedStreetsUsingGETNotFound with default headers values
func NewSuggestDeprecatedStreetsUsingGETNotFound() *SuggestDeprecatedStreetsUsingGETNotFound {
	return &SuggestDeprecatedStreetsUsingGETNotFound{}
}

/*SuggestDeprecatedStreetsUsingGETNotFound handles this case with default header values.

Not Found
*/
type SuggestDeprecatedStreetsUsingGETNotFound struct {
}

func (o *SuggestDeprecatedStreetsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/address-lookup/streets][%d] suggestDeprecatedStreetsUsingGETNotFound ", 404)
}

func (o *SuggestDeprecatedStreetsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
