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

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VectorConfig vector config
//
// swagger:model VectorConfig
type VectorConfig struct {

	// Vector-index config, that is specific to the type of index selected in vectorIndexType
	VectorIndexConfig interface{} `json:"vectorIndexConfig,omitempty"`

	// Name of the vector index to use, eg. (HNSW)
	VectorIndexType string `json:"vectorIndexType,omitempty"`

	// Configuration of a specific vectorizer used by this vector
	Vectorizer interface{} `json:"vectorizer,omitempty"`
}

// Validate validates this vector config
func (m *VectorConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vector config based on context it is used
func (m *VectorConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VectorConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VectorConfig) UnmarshalBinary(b []byte) error {
	var res VectorConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
