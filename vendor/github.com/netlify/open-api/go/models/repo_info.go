// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RepoInfo repo info
// swagger:model repoInfo
type RepoInfo struct {

	// allowed branches
	AllowedBranches []string `json:"allowed_branches"`

	// cmd
	Cmd string `json:"cmd,omitempty"`

	// deploy key id
	DeployKeyID string `json:"deploy_key_id,omitempty"`

	// dir
	Dir string `json:"dir,omitempty"`

	// env
	Env map[string]string `json:"env,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// private logs
	PrivateLogs bool `json:"private_logs,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// public repo
	PublicRepo bool `json:"public_repo,omitempty"`

	// repo branch
	RepoBranch string `json:"repo_branch,omitempty"`

	// repo path
	RepoPath string `json:"repo_path,omitempty"`

	// repo url
	RepoURL string `json:"repo_url,omitempty"`
}

// Validate validates this repo info
func (m *RepoInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedBranches(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepoInfo) validateAllowedBranches(formats strfmt.Registry) error {

	if swag.IsZero(m.AllowedBranches) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepoInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepoInfo) UnmarshalBinary(b []byte) error {
	var res RepoInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
