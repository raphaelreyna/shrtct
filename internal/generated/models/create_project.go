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

// CreateProject create project
//
// swagger:model CreateProject
type CreateProject struct {

	// The Project abbreviation used in Story summaries. Should be kept to 3 characters at most.
	Abbreviation string `json:"abbreviation,omitempty"`

	// The color you wish to use for the Project in the system.
	// Min Length: 1
	// Pattern: ^#[a-fA-F0-9]{6}$
	Color string `json:"color,omitempty"`

	// Defaults to the time/date it is created but can be set to reflect another date.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The Project description.
	// Max Length: 100000
	Description string `json:"description,omitempty"`

	// This field can be set to another unique ID. In the case that the Project has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`

	// An array of UUIDs for any members you want to add as Owners on this new Epic.
	FollowerIds []strfmt.UUID `json:"follower_ids"`

	// The number of weeks per iteration in this Project.
	IterationLength int64 `json:"iteration_length,omitempty"`

	// The name of the Project.
	// Required: true
	// Max Length: 128
	// Min Length: 1
	Name *string `json:"name"`

	// The date at which the Project was started.
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`

	// The ID of the team the project belongs to.
	// Required: true
	TeamID *int64 `json:"team_id"`

	// Defaults to the time/date it is created but can be set to reflect another date.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this create project
func (m *CreateProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFollowerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamID(formats); err != nil {
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

func (m *CreateProject) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MinLength("color", "body", m.Color, 1); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", m.Color, `^#[a-fA-F0-9]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *CreateProject) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateProject) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 100000); err != nil {
		return err
	}

	return nil
}

func (m *CreateProject) validateFollowerIds(formats strfmt.Registry) error {
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

func (m *CreateProject) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 128); err != nil {
		return err
	}

	return nil
}

func (m *CreateProject) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateProject) validateTeamID(formats strfmt.Registry) error {

	if err := validate.Required("team_id", "body", m.TeamID); err != nil {
		return err
	}

	return nil
}

func (m *CreateProject) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create project based on context it is used
func (m *CreateProject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateProject) UnmarshalBinary(b []byte) error {
	var res CreateProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
