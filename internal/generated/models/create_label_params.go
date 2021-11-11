// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateLabelParams Request parameters for creating a Label on a Shortcut Story.
//
// swagger:model CreateLabelParams
type CreateLabelParams struct {

	// The hex color to be displayed with the Label (for example, "#ff0000").
	// Min Length: 1
	// Pattern: ^#[a-fA-F0-9]{6}$
	Color string `json:"color,omitempty"`

	// The description of the new Label.
	Description string `json:"description,omitempty"`

	// This field can be set to another unique ID. In the case that the Label has been imported from another tool, the ID in the other tool can be indicated here.
	// Min Length: 1
	ExternalID string `json:"external_id,omitempty"`

	// The name of the new Label.
	// Required: true
	// Max Length: 128
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this create label params
func (m *CreateLabelParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateLabelParams) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MinLength("color", "body", m.Color, 1); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", m.Color, `^#[a-fA-F0-9]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *CreateLabelParams) validateExternalID(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalID) { // not required
		return nil
	}

	if err := validate.MinLength("external_id", "body", m.ExternalID, 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateLabelParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 128); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create label params based on context it is used
func (m *CreateLabelParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateLabelParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateLabelParams) UnmarshalBinary(b []byte) error {
	var res CreateLabelParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}