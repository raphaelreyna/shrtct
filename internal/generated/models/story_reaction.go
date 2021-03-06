// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StoryReaction Emoji reaction on a comment.
//
// swagger:model StoryReaction
type StoryReaction struct {

	// Emoji text of the reaction.
	// Required: true
	Emoji *string `json:"emoji"`

	// Permissions who have reacted with this.
	// Required: true
	PermissionIds []strfmt.UUID `json:"permission_ids"`
}

// Validate validates this story reaction
func (m *StoryReaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmoji(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissionIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoryReaction) validateEmoji(formats strfmt.Registry) error {

	if err := validate.Required("emoji", "body", m.Emoji); err != nil {
		return err
	}

	return nil
}

func (m *StoryReaction) validatePermissionIds(formats strfmt.Registry) error {

	if err := validate.Required("permission_ids", "body", m.PermissionIds); err != nil {
		return err
	}

	for i := 0; i < len(m.PermissionIds); i++ {

		if err := validate.FormatOf("permission_ids"+"."+strconv.Itoa(i), "body", "uuid", m.PermissionIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this story reaction based on context it is used
func (m *StoryReaction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StoryReaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoryReaction) UnmarshalBinary(b []byte) error {
	var res StoryReaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
