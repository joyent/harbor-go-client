// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StartReplicationExecution start replication execution
//
// swagger:model StartReplicationExecution
type StartReplicationExecution struct {

	// The ID of policy that the execution belongs to.
	PolicyID int64 `json:"policy_id,omitempty"`
}

// Validate validates this start replication execution
func (m *StartReplicationExecution) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start replication execution based on context it is used
func (m *StartReplicationExecution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StartReplicationExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StartReplicationExecution) UnmarshalBinary(b []byte) error {
	var res StartReplicationExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}