// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HistoryActionStoryCommentCreate An action representing a Story Comment being created.
//
// swagger:model HistoryActionStoryCommentCreate
type HistoryActionStoryCommentCreate struct {

	// The action of the entity referenced.
	// Required: true
	// Enum: [create]
	Action *string `json:"action"`

	// The application URL of the Story Comment.
	// Required: true
	// Max Length: 2048
	// Pattern: ^https?://.+$
	AppURL *string `json:"app_url"`

	// The Member ID of who created the Story Comment.
	// Required: true
	// Format: uuid
	AuthorID *strfmt.UUID `json:"author_id"`

	// The type of entity referenced.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the entity referenced.
	// Required: true
	ID *int64 `json:"id"`

	// The text of the Story Comment.
	// Required: true
	Text *string `json:"text"`
}

// Validate validates this history action story comment create
func (m *HistoryActionStoryCommentCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var historyActionStoryCommentCreateTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["create"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historyActionStoryCommentCreateTypeActionPropEnum = append(historyActionStoryCommentCreateTypeActionPropEnum, v)
	}
}

const (

	// HistoryActionStoryCommentCreateActionCreate captures enum value "create"
	HistoryActionStoryCommentCreateActionCreate string = "create"
)

// prop value enum
func (m *HistoryActionStoryCommentCreate) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historyActionStoryCommentCreateTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoryActionStoryCommentCreate) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionStoryCommentCreate) validateAppURL(formats strfmt.Registry) error {

	if err := validate.Required("app_url", "body", m.AppURL); err != nil {
		return err
	}

	if err := validate.MaxLength("app_url", "body", *m.AppURL, 2048); err != nil {
		return err
	}

	if err := validate.Pattern("app_url", "body", *m.AppURL, `^https?://.+$`); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionStoryCommentCreate) validateAuthorID(formats strfmt.Registry) error {

	if err := validate.Required("author_id", "body", m.AuthorID); err != nil {
		return err
	}

	if err := validate.FormatOf("author_id", "body", "uuid", m.AuthorID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionStoryCommentCreate) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionStoryCommentCreate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionStoryCommentCreate) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this history action story comment create based on context it is used
func (m *HistoryActionStoryCommentCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryActionStoryCommentCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryActionStoryCommentCreate) UnmarshalBinary(b []byte) error {
	var res HistoryActionStoryCommentCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}