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

// HistoryActionProjectUpdate An action representing a Project being updated.
//
// swagger:model HistoryActionProjectUpdate
type HistoryActionProjectUpdate struct {

	// The action of the entity referenced.
	// Required: true
	// Enum: [update]
	Action *string `json:"action"`

	// The application URL of the Project.
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

	// The name of the Project.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this history action project update
func (m *HistoryActionProjectUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var historyActionProjectUpdateTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["update"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historyActionProjectUpdateTypeActionPropEnum = append(historyActionProjectUpdateTypeActionPropEnum, v)
	}
}

const (

	// HistoryActionProjectUpdateActionUpdate captures enum value "update"
	HistoryActionProjectUpdateActionUpdate string = "update"
)

// prop value enum
func (m *HistoryActionProjectUpdate) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historyActionProjectUpdateTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoryActionProjectUpdate) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionProjectUpdate) validateAppURL(formats strfmt.Registry) error {

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

func (m *HistoryActionProjectUpdate) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionProjectUpdate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionProjectUpdate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this history action project update based on context it is used
func (m *HistoryActionProjectUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryActionProjectUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryActionProjectUpdate) UnmarshalBinary(b []byte) error {
	var res HistoryActionProjectUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
