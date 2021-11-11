// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateEpic update epic
//
// swagger:model UpdateEpic
type UpdateEpic struct {

	// The ID of the Epic we want to move this Epic after.
	AfterID int64 `json:"after_id,omitempty"`

	// A true/false boolean indicating whether the Epic is in archived state.
	Archived bool `json:"archived,omitempty"`

	// The ID of the Epic we want to move this Epic before.
	BeforeID int64 `json:"before_id,omitempty"`

	// A manual override for the time/date the Epic was completed.
	// Format: date-time
	CompletedAtOverride *strfmt.DateTime `json:"completed_at_override,omitempty"`

	// The Epic's deadline.
	// Format: date-time
	Deadline *strfmt.DateTime `json:"deadline,omitempty"`

	// The Epic's description.
	// Max Length: 100000
	Description string `json:"description,omitempty"`

	// The ID of the Epic State.
	EpicStateID int64 `json:"epic_state_id,omitempty"`

	// An array of UUIDs for any Members you want to add as Followers on this Epic.
	FollowerIds []strfmt.UUID `json:"follower_ids"`

	// The ID of the group to associate with the epic.
	// Format: uuid
	GroupID *strfmt.UUID `json:"group_id,omitempty"`

	// An array of Labels attached to the Epic.
	Labels []*CreateLabelParams `json:"labels"`

	// The ID of the Milestone this Epic is related to.
	MilestoneID *int64 `json:"milestone_id,omitempty"`

	// The Epic's name.
	// Max Length: 256
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// An array of UUIDs for any members you want to add as Owners on this Epic.
	OwnerIds []strfmt.UUID `json:"owner_ids"`

	// The Epic's planned start date.
	// Format: date-time
	PlannedStartDate *strfmt.DateTime `json:"planned_start_date,omitempty"`

	// The ID of the member that requested the epic.
	// Format: uuid
	RequestedByID strfmt.UUID `json:"requested_by_id,omitempty"`

	// A manual override for the time/date the Epic was started.
	// Format: date-time
	StartedAtOverride *strfmt.DateTime `json:"started_at_override,omitempty"`

	// `Deprecated` The Epic's state (to do, in progress, or done); will be ignored when `epic_state_id` is set.
	// Enum: [in progress to do done]
	State string `json:"state,omitempty"`
}

// Validate validates this update epic
func (m *UpdateEpic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedAtOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeadline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFollowerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlannedStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedByID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAtOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateEpic) validateCompletedAtOverride(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletedAtOverride) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at_override", "body", "date-time", m.CompletedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEpic) validateDeadline(formats strfmt.Registry) error {
	if swag.IsZero(m.Deadline) { // not required
		return nil
	}

	if err := validate.FormatOf("deadline", "body", "date-time", m.Deadline.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEpic) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 100000); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEpic) validateFollowerIds(formats strfmt.Registry) error {
	if swag.IsZero(m.FollowerIds) { // not required
		return nil
	}

	for i := 0; i < len(m.FollowerIds); i++ {

		if err := validate.FormatOf("follower_ids"+"."+strconv.Itoa(i), "body", "uuid", m.FollowerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *UpdateEpic) validateGroupID(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupID) { // not required
		return nil
	}

	if err := validate.FormatOf("group_id", "body", "uuid", m.GroupID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEpic) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateEpic) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 256); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEpic) validateOwnerIds(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerIds) { // not required
		return nil
	}

	for i := 0; i < len(m.OwnerIds); i++ {

		if err := validate.FormatOf("owner_ids"+"."+strconv.Itoa(i), "body", "uuid", m.OwnerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *UpdateEpic) validatePlannedStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PlannedStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("planned_start_date", "body", "date-time", m.PlannedStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEpic) validateRequestedByID(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedByID) { // not required
		return nil
	}

	if err := validate.FormatOf("requested_by_id", "body", "uuid", m.RequestedByID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEpic) validateStartedAtOverride(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedAtOverride) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at_override", "body", "date-time", m.StartedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

var updateEpicTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["in progress","to do","done"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateEpicTypeStatePropEnum = append(updateEpicTypeStatePropEnum, v)
	}
}

const (

	// UpdateEpicStateInProgress captures enum value "in progress"
	UpdateEpicStateInProgress string = "in progress"

	// UpdateEpicStateToDo captures enum value "to do"
	UpdateEpicStateToDo string = "to do"

	// UpdateEpicStateDone captures enum value "done"
	UpdateEpicStateDone string = "done"
)

// prop value enum
func (m *UpdateEpic) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateEpicTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateEpic) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update epic based on the context it is used
func (m *UpdateEpic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateEpic) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateEpic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateEpic) UnmarshalBinary(b []byte) error {
	var res UpdateEpic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
