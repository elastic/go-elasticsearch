// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// MappingLimitSettingsNestedObjects type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L445-L452
type MappingLimitSettingsNestedObjects struct {
	// Limit The maximum number of nested JSON objects that a single document can contain
	// across all nested types. This limit helps
	// to prevent out of memory errors when a document contains too many nested
	// objects.
	Limit *int `json:"limit,omitempty"`
}

// MappingLimitSettingsNestedObjectsBuilder holds MappingLimitSettingsNestedObjects struct and provides a builder API.
type MappingLimitSettingsNestedObjectsBuilder struct {
	v *MappingLimitSettingsNestedObjects
}

// NewMappingLimitSettingsNestedObjects provides a builder for the MappingLimitSettingsNestedObjects struct.
func NewMappingLimitSettingsNestedObjectsBuilder() *MappingLimitSettingsNestedObjectsBuilder {
	r := MappingLimitSettingsNestedObjectsBuilder{
		&MappingLimitSettingsNestedObjects{},
	}

	return &r
}

// Build finalize the chain and returns the MappingLimitSettingsNestedObjects struct
func (rb *MappingLimitSettingsNestedObjectsBuilder) Build() MappingLimitSettingsNestedObjects {
	return *rb.v
}

// Limit The maximum number of nested JSON objects that a single document can contain
// across all nested types. This limit helps
// to prevent out of memory errors when a document contains too many nested
// objects.

func (rb *MappingLimitSettingsNestedObjectsBuilder) Limit(limit int) *MappingLimitSettingsNestedObjectsBuilder {
	rb.v.Limit = &limit
	return rb
}
