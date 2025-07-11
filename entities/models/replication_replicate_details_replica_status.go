//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

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

// ReplicationReplicateDetailsReplicaStatus Represents the current or historical status of a shard replica involved in a replication operation, including its operational state and any associated errors.
//
// swagger:model ReplicationReplicateDetailsReplicaStatus
type ReplicationReplicateDetailsReplicaStatus struct {

	// A list of error messages encountered by this replica during the replication operation, if any.
	Errors []*ReplicationReplicateDetailsReplicaStatusError `json:"errors"`

	// The current operational state of the replica during the replication process.
	// Enum: [REGISTERED HYDRATING FINALIZING DEHYDRATING READY CANCELLED]
	State string `json:"state,omitempty"`

	// The UNIX timestamp in ms when this state was first entered. This is an approximate time and so should not be used for precise timing.
	WhenStartedUnixMs int64 `json:"whenStartedUnixMs,omitempty"`
}

// Validate validates this replication replicate details replica status
func (m *ReplicationReplicateDetailsReplicaStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
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

func (m *ReplicationReplicateDetailsReplicaStatus) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var replicationReplicateDetailsReplicaStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REGISTERED","HYDRATING","FINALIZING","DEHYDRATING","READY","CANCELLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		replicationReplicateDetailsReplicaStatusTypeStatePropEnum = append(replicationReplicateDetailsReplicaStatusTypeStatePropEnum, v)
	}
}

const (

	// ReplicationReplicateDetailsReplicaStatusStateREGISTERED captures enum value "REGISTERED"
	ReplicationReplicateDetailsReplicaStatusStateREGISTERED string = "REGISTERED"

	// ReplicationReplicateDetailsReplicaStatusStateHYDRATING captures enum value "HYDRATING"
	ReplicationReplicateDetailsReplicaStatusStateHYDRATING string = "HYDRATING"

	// ReplicationReplicateDetailsReplicaStatusStateFINALIZING captures enum value "FINALIZING"
	ReplicationReplicateDetailsReplicaStatusStateFINALIZING string = "FINALIZING"

	// ReplicationReplicateDetailsReplicaStatusStateDEHYDRATING captures enum value "DEHYDRATING"
	ReplicationReplicateDetailsReplicaStatusStateDEHYDRATING string = "DEHYDRATING"

	// ReplicationReplicateDetailsReplicaStatusStateREADY captures enum value "READY"
	ReplicationReplicateDetailsReplicaStatusStateREADY string = "READY"

	// ReplicationReplicateDetailsReplicaStatusStateCANCELLED captures enum value "CANCELLED"
	ReplicationReplicateDetailsReplicaStatusStateCANCELLED string = "CANCELLED"
)

// prop value enum
func (m *ReplicationReplicateDetailsReplicaStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, replicationReplicateDetailsReplicaStatusTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReplicationReplicateDetailsReplicaStatus) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this replication replicate details replica status based on the context it is used
func (m *ReplicationReplicateDetailsReplicaStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplicationReplicateDetailsReplicaStatus) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationReplicateDetailsReplicaStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationReplicateDetailsReplicaStatus) UnmarshalBinary(b []byte) error {
	var res ReplicationReplicateDetailsReplicaStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
