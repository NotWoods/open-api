// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeployFiles deploy files
//
// swagger:model deployFiles
type DeployFiles struct {

	// async
	Async bool `json:"async,omitempty"`

	// branch
	Branch string `json:"branch,omitempty"`

	// draft
	Draft bool `json:"draft,omitempty"`

	// files
	Files interface{} `json:"files,omitempty"`

	// framework
	Framework string `json:"framework,omitempty"`

	// functions
	Functions interface{} `json:"functions,omitempty"`
}

// Validate validates this deploy files
func (m *DeployFiles) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeployFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeployFiles) UnmarshalBinary(b []byte) error {
	var res DeployFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
