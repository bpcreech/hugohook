// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Event event
// swagger:model Event
type Event struct {

	// before
	Before string `json:"before,omitempty"`

	// commits
	Commits Commits `json:"commits"`

	// distinct size
	DistinctSize int64 `json:"distinct_size,omitempty"`

	// head
	Head string `json:"head,omitempty"`

	// hook
	Hook *HookConfig `json:"hook,omitempty"`

	// hook id
	HookID int64 `json:"hook_id,omitempty"`

	// organization
	Organization *Organization `json:"organization,omitempty"`

	// ref
	Ref string `json:"ref,omitempty"`

	// repository
	Repository *Repository `json:"repository,omitempty"`

	// sender
	Sender *User `json:"sender,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// zen
	Zen string `json:"zen,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHook(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateCommits(formats strfmt.Registry) error {

	if swag.IsZero(m.Commits) { // not required
		return nil
	}

	if err := m.Commits.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("commits")
		}
		return err
	}

	return nil
}

func (m *Event) validateHook(formats strfmt.Registry) error {

	if swag.IsZero(m.Hook) { // not required
		return nil
	}

	if m.Hook != nil {
		if err := m.Hook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hook")
			}
			return err
		}
	}

	return nil
}

func (m *Event) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

func (m *Event) validateRepository(formats strfmt.Registry) error {

	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

func (m *Event) validateSender(formats strfmt.Registry) error {

	if swag.IsZero(m.Sender) { // not required
		return nil
	}

	if m.Sender != nil {
		if err := m.Sender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}