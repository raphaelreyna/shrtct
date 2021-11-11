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

// UpdateEntityTemplate Request parameters for changing either a template's name or any of
//   the attributes it is designed to pre-populate.
//
// swagger:model UpdateEntityTemplate
type UpdateEntityTemplate struct {

	// The updated template name.
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// story contents
	StoryContents *UpdateStoryContents `json:"story_contents,omitempty"`
}

// Validate validates this update entity template
func (m *UpdateEntityTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoryContents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateEntityTemplate) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEntityTemplate) validateStoryContents(formats strfmt.Registry) error {
	if swag.IsZero(m.StoryContents) { // not required
		return nil
	}

	if m.StoryContents != nil {
		if err := m.StoryContents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("story_contents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("story_contents")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update entity template based on the context it is used
func (m *UpdateEntityTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStoryContents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateEntityTemplate) contextValidateStoryContents(ctx context.Context, formats strfmt.Registry) error {

	if m.StoryContents != nil {
		if err := m.StoryContents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("story_contents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("story_contents")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateEntityTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateEntityTemplate) UnmarshalBinary(b []byte) error {
	var res UpdateEntityTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
