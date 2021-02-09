// Code generated by go-swagger; DO NOT EDIT.

package address_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mkrill/diPrototypeWire/src/infrastructure/models"
)

// ResolveAddressInfoUsingGETReader is a Reader for the ResolveAddressInfoUsingGET structure.
type ResolveAddressInfoUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResolveAddressInfoUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResolveAddressInfoUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResolveAddressInfoUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewResolveAddressInfoUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResolveAddressInfoUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResolveAddressInfoUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResolveAddressInfoUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResolveAddressInfoUsingGETOK creates a ResolveAddressInfoUsingGETOK with default headers values
func NewResolveAddressInfoUsingGETOK() *ResolveAddressInfoUsingGETOK {
	return &ResolveAddressInfoUsingGETOK{}
}

/*ResolveAddressInfoUsingGETOK handles this case with default header values.

Information was found.
*/
type ResolveAddressInfoUsingGETOK struct {
	Payload *models.AddressInfo
}

func (o *ResolveAddressInfoUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/addressSearch/resolveAddressInfo][%d] resolveAddressInfoUsingGETOK  %+v", 200, o.Payload)
}

func (o *ResolveAddressInfoUsingGETOK) GetPayload() *models.AddressInfo {
	return o.Payload
}

func (o *ResolveAddressInfoUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddressInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResolveAddressInfoUsingGETBadRequest creates a ResolveAddressInfoUsingGETBadRequest with default headers values
func NewResolveAddressInfoUsingGETBadRequest() *ResolveAddressInfoUsingGETBadRequest {
	return &ResolveAddressInfoUsingGETBadRequest{}
}

/*ResolveAddressInfoUsingGETBadRequest handles this case with default header values.

Bad request.
*/
type ResolveAddressInfoUsingGETBadRequest struct {
	Payload *models.ProblemDetails
}

func (o *ResolveAddressInfoUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/addressSearch/resolveAddressInfo][%d] resolveAddressInfoUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *ResolveAddressInfoUsingGETBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ResolveAddressInfoUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResolveAddressInfoUsingGETUnauthorized creates a ResolveAddressInfoUsingGETUnauthorized with default headers values
func NewResolveAddressInfoUsingGETUnauthorized() *ResolveAddressInfoUsingGETUnauthorized {
	return &ResolveAddressInfoUsingGETUnauthorized{}
}

/*ResolveAddressInfoUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type ResolveAddressInfoUsingGETUnauthorized struct {
}

func (o *ResolveAddressInfoUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/addressSearch/resolveAddressInfo][%d] resolveAddressInfoUsingGETUnauthorized ", 401)
}

func (o *ResolveAddressInfoUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResolveAddressInfoUsingGETForbidden creates a ResolveAddressInfoUsingGETForbidden with default headers values
func NewResolveAddressInfoUsingGETForbidden() *ResolveAddressInfoUsingGETForbidden {
	return &ResolveAddressInfoUsingGETForbidden{}
}

/*ResolveAddressInfoUsingGETForbidden handles this case with default header values.

Forbidden
*/
type ResolveAddressInfoUsingGETForbidden struct {
}

func (o *ResolveAddressInfoUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/addressSearch/resolveAddressInfo][%d] resolveAddressInfoUsingGETForbidden ", 403)
}

func (o *ResolveAddressInfoUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResolveAddressInfoUsingGETNotFound creates a ResolveAddressInfoUsingGETNotFound with default headers values
func NewResolveAddressInfoUsingGETNotFound() *ResolveAddressInfoUsingGETNotFound {
	return &ResolveAddressInfoUsingGETNotFound{}
}

/*ResolveAddressInfoUsingGETNotFound handles this case with default header values.

KlsId was not found for specified address.
*/
type ResolveAddressInfoUsingGETNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ResolveAddressInfoUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/addressSearch/resolveAddressInfo][%d] resolveAddressInfoUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *ResolveAddressInfoUsingGETNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ResolveAddressInfoUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResolveAddressInfoUsingGETInternalServerError creates a ResolveAddressInfoUsingGETInternalServerError with default headers values
func NewResolveAddressInfoUsingGETInternalServerError() *ResolveAddressInfoUsingGETInternalServerError {
	return &ResolveAddressInfoUsingGETInternalServerError{}
}

/*ResolveAddressInfoUsingGETInternalServerError handles this case with default header values.

Internal Server Error.
*/
type ResolveAddressInfoUsingGETInternalServerError struct {
	Payload *models.ProblemDetails
}

func (o *ResolveAddressInfoUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/addressSearch/resolveAddressInfo][%d] resolveAddressInfoUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *ResolveAddressInfoUsingGETInternalServerError) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ResolveAddressInfoUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}