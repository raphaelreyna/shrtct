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

// PullRequest Corresponds to a VCS Pull Request attached to a Shortcut story.
//
// swagger:model PullRequest
type PullRequest struct {

	// The ID of the branch for the particular pull request.
	// Required: true
	BranchID *int64 `json:"branch_id"`

	// The name of the branch for the particular pull request.
	// Required: true
	BranchName *string `json:"branch_name"`

	// The status of the Continuous Integration workflow for the pull request.
	BuildStatus string `json:"build_status,omitempty"`

	// True/False boolean indicating whether the VCS pull request has been closed.
	// Required: true
	Closed *bool `json:"closed"`

	// The time/date the pull request was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// True/False boolean indicating whether the VCS pull request is in the draft state.
	// Required: true
	Draft *bool `json:"draft"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The unique ID associated with the pull request in Shortcut.
	// Required: true
	ID *int64 `json:"id"`

	// True/False boolean indicating whether the VCS pull request has been merged.
	// Required: true
	Merged *bool `json:"merged"`

	// Number of lines added in the pull request, according to VCS.
	// Required: true
	NumAdded *int64 `json:"num_added"`

	// The number of commits on the pull request.
	// Required: true
	NumCommits *int64 `json:"num_commits"`

	// Number of lines modified in the pull request, according to VCS.
	// Required: true
	NumModified *int64 `json:"num_modified"`

	// Number of lines removed in the pull request, according to VCS.
	// Required: true
	NumRemoved *int64 `json:"num_removed"`

	// The pull request's unique number ID in VCS.
	// Required: true
	Number *int64 `json:"number"`

	// The ID of the repository for the particular pull request.
	// Required: true
	RepositoryID *int64 `json:"repository_id"`

	// The status of the review for the pull request.
	ReviewStatus string `json:"review_status,omitempty"`

	// The ID of the target branch for the particular pull request.
	// Required: true
	TargetBranchID *int64 `json:"target_branch_id"`

	// The name of the target branch for the particular pull request.
	// Required: true
	TargetBranchName *string `json:"target_branch_name"`

	// The title of the pull request.
	// Required: true
	Title *string `json:"title"`

	// The time/date the pull request was created.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// The URL for the pull request.
	// Required: true
	URL *string `json:"url"`

	// An array of PullRequestLabels attached to the PullRequest.
	VcsLabels []*PullRequestLabel `json:"vcs_labels"`
}

// Validate validates this pull request
func (m *PullRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranchID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBranchName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClosed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDraft(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMerged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumCommits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumRemoved(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetBranchID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetBranchName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PullRequest) validateBranchID(formats strfmt.Registry) error {

	if err := validate.Required("branch_id", "body", m.BranchID); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateBranchName(formats strfmt.Registry) error {

	if err := validate.Required("branch_name", "body", m.BranchName); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateClosed(formats strfmt.Registry) error {

	if err := validate.Required("closed", "body", m.Closed); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateDraft(formats strfmt.Registry) error {

	if err := validate.Required("draft", "body", m.Draft); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateMerged(formats strfmt.Registry) error {

	if err := validate.Required("merged", "body", m.Merged); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateNumAdded(formats strfmt.Registry) error {

	if err := validate.Required("num_added", "body", m.NumAdded); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateNumCommits(formats strfmt.Registry) error {

	if err := validate.Required("num_commits", "body", m.NumCommits); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateNumModified(formats strfmt.Registry) error {

	if err := validate.Required("num_modified", "body", m.NumModified); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateNumRemoved(formats strfmt.Registry) error {

	if err := validate.Required("num_removed", "body", m.NumRemoved); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateRepositoryID(formats strfmt.Registry) error {

	if err := validate.Required("repository_id", "body", m.RepositoryID); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateTargetBranchID(formats strfmt.Registry) error {

	if err := validate.Required("target_branch_id", "body", m.TargetBranchID); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateTargetBranchName(formats strfmt.Registry) error {

	if err := validate.Required("target_branch_name", "body", m.TargetBranchName); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *PullRequest) validateVcsLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.VcsLabels) { // not required
		return nil
	}

	for i := 0; i < len(m.VcsLabels); i++ {
		if swag.IsZero(m.VcsLabels[i]) { // not required
			continue
		}

		if m.VcsLabels[i] != nil {
			if err := m.VcsLabels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vcs_labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vcs_labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this pull request based on the context it is used
func (m *PullRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVcsLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PullRequest) contextValidateVcsLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VcsLabels); i++ {

		if m.VcsLabels[i] != nil {
			if err := m.VcsLabels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vcs_labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vcs_labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PullRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PullRequest) UnmarshalBinary(b []byte) error {
	var res PullRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
