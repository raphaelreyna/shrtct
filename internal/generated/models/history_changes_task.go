// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HistoryChangesTask The changes that have occurred as a result of the action.
//
// swagger:model HistoryChangesTask
type HistoryChangesTask struct {

	// complete
	Complete *StoryHistoryChangeOldNewBool `json:"complete,omitempty"`

	// description
	Description *StoryHistoryChangeOldNewStr `json:"description,omitempty"`

	// mention ids
	MentionIds *StoryHistoryChangeAddsRemovesUUID `json:"mention_ids,omitempty"`

	// owner ids
	OwnerIds *StoryHistoryChangeAddsRemovesUUID `json:"owner_ids,omitempty"`
}

// Validate validates this history changes task
func (m *HistoryChangesTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComplete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryChangesTask) validateComplete(formats strfmt.Registry) error {
	if swag.IsZero(m.Complete) { // not required
		return nil
	}

	if m.Complete != nil {
		if err := m.Complete.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("complete")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("complete")
			}
			return err
		}
	}

	return nil
}

func (m *HistoryChangesTask) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if m.Description != nil {
		if err := m.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

func (m *HistoryChangesTask) validateMentionIds(formats strfmt.Registry) error {
	if swag.IsZero(m.MentionIds) { // not required
		return nil
	}

	if m.MentionIds != nil {
		if err := m.MentionIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mention_ids")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mention_ids")
			}
			return err
		}
	}

	return nil
}

func (m *HistoryChangesTask) validateOwnerIds(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerIds) { // not required
		return nil
	}

	if m.OwnerIds != nil {
		if err := m.OwnerIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner_ids")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner_ids")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this history changes task based on the context it is used
func (m *HistoryChangesTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComplete(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMentionIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwnerIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryChangesTask) contextValidateComplete(ctx context.Context, formats strfmt.Registry) error {

	if m.Complete != nil {
		if err := m.Complete.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("complete")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("complete")
			}
			return err
		}
	}

	return nil
}

func (m *HistoryChangesTask) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if m.Description != nil {
		if err := m.Description.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

func (m *HistoryChangesTask) contextValidateMentionIds(ctx context.Context, formats strfmt.Registry) error {

	if m.MentionIds != nil {
		if err := m.MentionIds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mention_ids")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mention_ids")
			}
			return err
		}
	}

	return nil
}

func (m *HistoryChangesTask) contextValidateOwnerIds(ctx context.Context, formats strfmt.Registry) error {

	if m.OwnerIds != nil {
		if err := m.OwnerIds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner_ids")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner_ids")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistoryChangesTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryChangesTask) UnmarshalBinary(b []byte) error {
	var res HistoryChangesTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
