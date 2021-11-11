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

// Member Details about individual Shortcut user within the Shortcut organization that has issued the token.
//
// swagger:model Member
type Member struct {

	// The time/date the Member was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// Whether this member was created as a placeholder entity.
	// Required: true
	CreatedWithoutInvite *bool `json:"created_without_invite"`

	// True/false boolean indicating whether the Member has been disabled within this Organization.
	// Required: true
	Disabled *bool `json:"disabled"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The Member's group ids
	// Required: true
	GroupIds []strfmt.UUID `json:"group_ids"`

	// The Member's ID in Shortcut.
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// profile
	// Required: true
	Profile *Profile `json:"profile"`

	// The id of the member that replaces this one when merged.
	// Format: uuid
	ReplacedBy strfmt.UUID `json:"replaced_by,omitempty"`

	// The Member's role in the Shortcut organization.
	// Required: true
	Role *string `json:"role"`

	// The user state, one of partial, full, disabled, or imported.  A partial
	//            user is disabled, has no means to log in, and is not an import user.  A full
	//            user is enabled and has a means to log in.  A disabled user is disabled and has
	//            a means to log in.  An import user is disabled, has no means to log in, and is
	//            marked as an import user.
	// Required: true
	// Enum: [partial full disabled imported]
	State *string `json:"state"`

	// The time/date the Member was last updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this member
func (m *Member) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedWithoutInvite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplacedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *Member) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateCreatedWithoutInvite(formats strfmt.Registry) error {

	if err := validate.Required("created_without_invite", "body", m.CreatedWithoutInvite); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateDisabled(formats strfmt.Registry) error {

	if err := validate.Required("disabled", "body", m.Disabled); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateGroupIds(formats strfmt.Registry) error {

	if err := validate.Required("group_ids", "body", m.GroupIds); err != nil {
		return err
	}

	for i := 0; i < len(m.GroupIds); i++ {

		if err := validate.FormatOf("group_ids"+"."+strconv.Itoa(i), "body", "uuid", m.GroupIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Member) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateProfile(formats strfmt.Registry) error {

	if err := validate.Required("profile", "body", m.Profile); err != nil {
		return err
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *Member) validateReplacedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplacedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("replaced_by", "body", "uuid", m.ReplacedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

var memberTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["partial","full","disabled","imported"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		memberTypeStatePropEnum = append(memberTypeStatePropEnum, v)
	}
}

const (

	// MemberStatePartial captures enum value "partial"
	MemberStatePartial string = "partial"

	// MemberStateFull captures enum value "full"
	MemberStateFull string = "full"

	// MemberStateDisabled captures enum value "disabled"
	MemberStateDisabled string = "disabled"

	// MemberStateImported captures enum value "imported"
	MemberStateImported string = "imported"
)

// prop value enum
func (m *Member) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, memberTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Member) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this member based on the context it is used
func (m *Member) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Member) contextValidateProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.Profile != nil {
		if err := m.Profile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Member) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Member) UnmarshalBinary(b []byte) error {
	var res Member
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
