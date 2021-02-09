// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProblemDetails ProblemDetails
//
// Error response of 4xx and 5xx http statuses.
//
// swagger:model ProblemDetails
type ProblemDetails struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// embedded
	Embedded *SpecificProblemDetails `json:"embedded,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this problem details
func (m *ProblemDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmbedded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProblemDetails) validateEmbedded(formats strfmt.Registry) error {

	if swag.IsZero(m.Embedded) { // not required
		return nil
	}

	if m.Embedded != nil {
		if err := m.Embedded.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("embedded")
			}
			return err
		}
	}

	return nil
}

func (m *ProblemDetails) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *ProblemDetails) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProblemDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProblemDetails) UnmarshalBinary(b []byte) error {
	var res ProblemDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
