// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StoryHistoryChangeOldNewBool True if the Story has archived, otherwise false.
//
// swagger:model StoryHistoryChangeOldNewBool
type StoryHistoryChangeOldNewBool struct {

	// The new value.
	New bool `json:"new,omitempty"`

	// The old value.
	Old bool `json:"old,omitempty"`
}

// Validate validates this story history change old new bool
func (m *StoryHistoryChangeOldNewBool) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this story history change old new bool based on context it is used
func (m *StoryHistoryChangeOldNewBool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StoryHistoryChangeOldNewBool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoryHistoryChangeOldNewBool) UnmarshalBinary(b []byte) error {
	var res StoryHistoryChangeOldNewBool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
