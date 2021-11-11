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

// HistoryActionTaskUpdate An action representing a Task being updated.
//
// swagger:model HistoryActionTaskUpdate
type HistoryActionTaskUpdate struct {

	// The action of the entity referenced.
	// Required: true
	// Enum: [update]
	Action *string `json:"action"`

	// changes
	// Required: true
	Changes *HistoryChangesTask `json:"changes"`

	// Whether or not the Task is complete.
	Complete bool `json:"complete,omitempty"`

	// The description of the Task.
	// Required: true
	Description *string `json:"description"`

	// The type of entity referenced.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the entity referenced.
	// Required: true
	ID *int64 `json:"id"`

	// The Story ID that contains the Task.
	// Required: true
	StoryID *int64 `json:"story_id"`
}

// Validate validates this history action task update
func (m *HistoryActionTaskUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoryID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var historyActionTaskUpdateTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["update"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historyActionTaskUpdateTypeActionPropEnum = append(historyActionTaskUpdateTypeActionPropEnum, v)
	}
}

const (

	// HistoryActionTaskUpdateActionUpdate captures enum value "update"
	HistoryActionTaskUpdateActionUpdate string = "update"
)

// prop value enum
func (m *HistoryActionTaskUpdate) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historyActionTaskUpdateTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoryActionTaskUpdate) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionTaskUpdate) validateChanges(formats strfmt.Registry) error {

	if err := validate.Required("changes", "body", m.Changes); err != nil {
		return err
	}

	if m.Changes != nil {
		if err := m.Changes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changes")
			}
			return err
		}
	}

	return nil
}

func (m *HistoryActionTaskUpdate) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionTaskUpdate) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionTaskUpdate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionTaskUpdate) validateStoryID(formats strfmt.Registry) error {

	if err := validate.Required("story_id", "body", m.StoryID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this history action task update based on the context it is used
func (m *HistoryActionTaskUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryActionTaskUpdate) contextValidateChanges(ctx context.Context, formats strfmt.Registry) error {

	if m.Changes != nil {
		if err := m.Changes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistoryActionTaskUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryActionTaskUpdate) UnmarshalBinary(b []byte) error {
	var res HistoryActionTaskUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}