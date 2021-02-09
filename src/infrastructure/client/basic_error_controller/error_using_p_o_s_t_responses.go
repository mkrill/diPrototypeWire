// Code generated by go-swagger; DO NOT EDIT.

package basic_error_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ErrorUsingPOSTReader is a Reader for the ErrorUsingPOST structure.
type ErrorUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ErrorUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewErrorUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewErrorUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewErrorUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewErrorUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewErrorUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewErrorUsingPOSTOK creates a ErrorUsingPOSTOK with default headers values
func NewErrorUsingPOSTOK() *ErrorUsingPOSTOK {
	return &ErrorUsingPOSTOK{}
}

/*ErrorUsingPOSTOK handles this case with default header values.

OK
*/
type ErrorUsingPOSTOK struct {
	Payload map[string]interface{}
}

func (o *ErrorUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /error][%d] errorUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ErrorUsingPOSTOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *ErrorUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewErrorUsingPOSTCreated creates a ErrorUsingPOSTCreated with default headers values
func NewErrorUsingPOSTCreated() *ErrorUsingPOSTCreated {
	return &ErrorUsingPOSTCreated{}
}

/*ErrorUsingPOSTCreated handles this case with default header values.

Created
*/
type ErrorUsingPOSTCreated struct {
}

func (o *ErrorUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /error][%d] errorUsingPOSTCreated ", 201)
}

func (o *ErrorUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorUsingPOSTUnauthorized creates a ErrorUsingPOSTUnauthorized with default headers values
func NewErrorUsingPOSTUnauthorized() *ErrorUsingPOSTUnauthorized {
	return &ErrorUsingPOSTUnauthorized{}
}

/*ErrorUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type ErrorUsingPOSTUnauthorized struct {
}

func (o *ErrorUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /error][%d] errorUsingPOSTUnauthorized ", 401)
}

func (o *ErrorUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorUsingPOSTForbidden creates a ErrorUsingPOSTForbidden with default headers values
func NewErrorUsingPOSTForbidden() *ErrorUsingPOSTForbidden {
	return &ErrorUsingPOSTForbidden{}
}

/*ErrorUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type ErrorUsingPOSTForbidden struct {
}

func (o *ErrorUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /error][%d] errorUsingPOSTForbidden ", 403)
}

func (o *ErrorUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorUsingPOSTNotFound creates a ErrorUsingPOSTNotFound with default headers values
func NewErrorUsingPOSTNotFound() *ErrorUsingPOSTNotFound {
	return &ErrorUsingPOSTNotFound{}
}

/*ErrorUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type ErrorUsingPOSTNotFound struct {
}

func (o *ErrorUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /error][%d] errorUsingPOSTNotFound ", 404)
}

func (o *ErrorUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
