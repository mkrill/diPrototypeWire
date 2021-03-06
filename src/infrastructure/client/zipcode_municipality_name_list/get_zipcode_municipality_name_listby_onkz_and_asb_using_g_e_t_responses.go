// Code generated by go-swagger; DO NOT EDIT.

package zipcode_municipality_name_list

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mkrill/diPrototypeWire/src/infrastructure/models"
)

// GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETReader is a Reader for the GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGET structure.
type GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK creates a GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK with default headers values
func NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK() *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK {
	return &GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK{}
}

/*GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK handles this case with default header values.

Information was found.
*/
type GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK struct {
	Payload []*models.ZipcodeMunicipalityNameDTO
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/zipcodeMunicipalityNames/getZipcodeMunicipalityNameList/{onkz}][%d] getZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK) GetPayload() []*models.ZipcodeMunicipalityNameDTO {
	return o.Payload
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest creates a GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest with default headers values
func NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest() *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest {
	return &GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest{}
}

/*GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest handles this case with default header values.

Bad request.
*/
type GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest struct {
	Payload *models.ProblemDetails
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/zipcodeMunicipalityNames/getZipcodeMunicipalityNameList/{onkz}][%d] getZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETUnauthorized creates a GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETUnauthorized with default headers values
func NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETUnauthorized() *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETUnauthorized {
	return &GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETUnauthorized{}
}

/*GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETUnauthorized struct {
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/zipcodeMunicipalityNames/getZipcodeMunicipalityNameList/{onkz}][%d] getZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETUnauthorized ", 401)
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETForbidden creates a GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETForbidden with default headers values
func NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETForbidden() *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETForbidden {
	return &GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETForbidden{}
}

/*GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETForbidden struct {
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/zipcodeMunicipalityNames/getZipcodeMunicipalityNameList/{onkz}][%d] getZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETForbidden ", 403)
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETNotFound creates a GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETNotFound with default headers values
func NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETNotFound() *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETNotFound {
	return &GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETNotFound{}
}

/*GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETNotFound struct {
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/zipcodeMunicipalityNames/getZipcodeMunicipalityNameList/{onkz}][%d] getZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETNotFound ", 404)
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError creates a GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError with default headers values
func NewGetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError() *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError {
	return &GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError{}
}

/*GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError handles this case with default header values.

Internal Server Error.
*/
type GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError struct {
	Payload *models.ProblemDetails
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/addressLookup/v1/zipcodeMunicipalityNames/getZipcodeMunicipalityNameList/{onkz}][%d] getZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetZipcodeMunicipalityNameListbyOnkzAndAsbUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
