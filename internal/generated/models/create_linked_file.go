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

// CreateLinkedFile create linked file
//
// swagger:model CreateLinkedFile
type CreateLinkedFile struct {

	// The content type of the image (e.g. txt/plain).
	ContentType string `json:"content_type,omitempty"`

	// The description of the file.
	Description string `json:"description,omitempty"`

	// The name of the file.
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// The filesize, if the integration provided it.
	Size int64 `json:"size,omitempty"`

	// The ID of the linked story.
	StoryID int64 `json:"story_id,omitempty"`

	// The URL of the thumbnail, if the integration provided it.
	// Max Length: 2048
	// Pattern: ^https?://.+$
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// The integration type of the file (e.g. google, dropbox, box).
	// Required: true
	// Enum: [google url dropbox box onedrive]
	Type *string `json:"type"`

	// The UUID of the member that uploaded the file.
	// Format: uuid
	UploaderID strfmt.UUID `json:"uploader_id,omitempty"`

	// The URL of linked file.
	// Required: true
	// Max Length: 2048
	// Pattern: ^https?://.+$
	URL *string `json:"url"`
}

// Validate validates this create linked file
func (m *CreateLinkedFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThumbnailURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploaderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateLinkedFile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateLinkedFile) validateThumbnailURL(formats strfmt.Registry) error {
	if swag.IsZero(m.ThumbnailURL) { // not required
		return nil
	}

	if err := validate.MaxLength("thumbnail_url", "body", m.ThumbnailURL, 2048); err != nil {
		return err
	}

	if err := validate.Pattern("thumbnail_url", "body", m.ThumbnailURL, `^https?://.+$`); err != nil {
		return err
	}

	return nil
}

var createLinkedFileTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["google","url","dropbox","box","onedrive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createLinkedFileTypeTypePropEnum = append(createLinkedFileTypeTypePropEnum, v)
	}
}

const (

	// CreateLinkedFileTypeGoogle captures enum value "google"
	CreateLinkedFileTypeGoogle string = "google"

	// CreateLinkedFileTypeURL captures enum value "url"
	CreateLinkedFileTypeURL string = "url"

	// CreateLinkedFileTypeDropbox captures enum value "dropbox"
	CreateLinkedFileTypeDropbox string = "dropbox"

	// CreateLinkedFileTypeBox captures enum value "box"
	CreateLinkedFileTypeBox string = "box"

	// CreateLinkedFileTypeOnedrive captures enum value "onedrive"
	CreateLinkedFileTypeOnedrive string = "onedrive"
)

// prop value enum
func (m *CreateLinkedFile) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createLinkedFileTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateLinkedFile) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CreateLinkedFile) validateUploaderID(formats strfmt.Registry) error {
	if swag.IsZero(m.UploaderID) { // not required
		return nil
	}

	if err := validate.FormatOf("uploader_id", "body", "uuid", m.UploaderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateLinkedFile) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.MaxLength("url", "body", *m.URL, 2048); err != nil {
		return err
	}

	if err := validate.Pattern("url", "body", *m.URL, `^https?://.+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create linked file based on context it is used
func (m *CreateLinkedFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateLinkedFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateLinkedFile) UnmarshalBinary(b []byte) error {
	var res CreateLinkedFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
