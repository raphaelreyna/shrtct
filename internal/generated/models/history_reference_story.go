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

// HistoryReferenceStory A reference to a Story.
//
// swagger:model HistoryReferenceStory
type HistoryReferenceStory struct {

	// The application URL of the Story.
	// Required: true
	// Max Length: 2048
	// Pattern: ^https?://.+$
	AppURL *string `json:"app_url"`

	// The type of entity referenced.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the entity referenced.
	// Required: true
	ID *int64 `json:"id"`

	// The name of the entity referenced.
	// Required: true
	Name *string `json:"name"`

	// If the referenced entity is a Story, either "bug", "chore", or "feature".
	// Required: true
	// Enum: [feature chore bug]
	StoryType *string `json:"story_type"`
}

// Validate validates this history reference story
func (m *HistoryReferenceStory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoryType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryReferenceStory) validateAppURL(formats strfmt.Registry) error {

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

func (m *HistoryReferenceStory) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceStory) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceStory) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var historyReferenceStoryTypeStoryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["feature","chore","bug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historyReferenceStoryTypeStoryTypePropEnum = append(historyReferenceStoryTypeStoryTypePropEnum, v)
	}
}

const (

	// HistoryReferenceStoryStoryTypeFeature captures enum value "feature"
	HistoryReferenceStoryStoryTypeFeature string = "feature"

	// HistoryReferenceStoryStoryTypeChore captures enum value "chore"
	HistoryReferenceStoryStoryTypeChore string = "chore"

	// HistoryReferenceStoryStoryTypeBug captures enum value "bug"
	HistoryReferenceStoryStoryTypeBug string = "bug"
)

// prop value enum
func (m *HistoryReferenceStory) validateStoryTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historyReferenceStoryTypeStoryTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoryReferenceStory) validateStoryType(formats strfmt.Registry) error {

	if err := validate.Required("story_type", "body", m.StoryType); err != nil {
		return err
	}

	// value enum
	if err := m.validateStoryTypeEnum("story_type", "body", *m.StoryType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this history reference story based on context it is used
func (m *HistoryReferenceStory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryReferenceStory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryReferenceStory) UnmarshalBinary(b []byte) error {
	var res HistoryReferenceStory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
