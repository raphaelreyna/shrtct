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

// CreateCommentComment create comment comment
//
// swagger:model CreateCommentComment
type CreateCommentComment struct {

	// The Member ID of the Comment's author. Defaults to the user identified by the API token.
	// Format: uuid
	AuthorID strfmt.UUID `json:"author_id,omitempty"`

	// Defaults to the time/date the comment is created, but can be set to reflect another date.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// This field can be set to another unique ID. In the case that the comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`

	// The comment text.
	// Required: true
	// Max Length: 100000
	// Min Length: 1
	Text *string `json:"text"`

	// Defaults to the time/date the comment is last updated, but can be set to reflect another date.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this create comment comment
func (m *CreateCommentComment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCommentComment) validateAuthorID(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthorID) { // not required
		return nil
	}

	if err := validate.FormatOf("author_id", "body", "uuid", m.AuthorID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateCommentComment) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateCommentComment) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	if err := validate.MinLength("text", "body", *m.Text, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("text", "body", *m.Text, 100000); err != nil {
		return err
	}

	return nil
}

func (m *CreateCommentComment) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create comment comment based on context it is used
func (m *CreateCommentComment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateCommentComment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCommentComment) UnmarshalBinary(b []byte) error {
	var res CreateCommentComment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
