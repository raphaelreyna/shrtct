// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StoryHistoryChangeOldNewInt The estimate value for the Story
//
// swagger:model StoryHistoryChangeOldNewInt
type StoryHistoryChangeOldNewInt struct {

	// The new value.
	New int64 `json:"new,omitempty"`

	// The old value.
	Old int64 `json:"old,omitempty"`
}

// Validate validates this story history change old new int
func (m *StoryHistoryChangeOldNewInt) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this story history change old new int based on context it is used
func (m *StoryHistoryChangeOldNewInt) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StoryHistoryChangeOldNewInt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoryHistoryChangeOldNewInt) UnmarshalBinary(b []byte) error {
	var res StoryHistoryChangeOldNewInt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
