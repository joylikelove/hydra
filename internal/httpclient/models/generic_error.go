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

// GenericError GenericError Error response
//
// Error responses are sent when an error (e.g. unauthorized, bad request, ...) occurred.
//
// swagger:model genericError
type GenericError struct {

	// Debug contains debug information. This is usually not available and has to be enabled.
	Debug string `json:"debug,omitempty"`

	// Name is the error name.
	// Required: true
	Error *string `json:"error"`

	// Description contains further information on the nature of the error.
	ErrorDescription string `json:"error_description,omitempty"`

	// Code represents the error status code (404, 403, 401, ...).
	StatusCode int64 `json:"status_code,omitempty"`
}

// Validate validates this generic error
func (m *GenericError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericError) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenericError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericError) UnmarshalBinary(b []byte) error {
	var res GenericError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
