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

// MappingCharFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/char_filters.ts#L47-L51
type MappingCharFilter struct {
	Mappings     []string       `json:"mappings,omitempty"`
	MappingsPath *string        `json:"mappings_path,omitempty"`
	Type         string         `json:"type,omitempty"`
	Version      *VersionString `json:"version,omitempty"`
}

// MappingCharFilterBuilder holds MappingCharFilter struct and provides a builder API.
type MappingCharFilterBuilder struct {
	v *MappingCharFilter
}

// NewMappingCharFilter provides a builder for the MappingCharFilter struct.
func NewMappingCharFilterBuilder() *MappingCharFilterBuilder {
	r := MappingCharFilterBuilder{
		&MappingCharFilter{},
	}

	r.v.Type = "mapping"

	return &r
}

// Build finalize the chain and returns the MappingCharFilter struct
func (rb *MappingCharFilterBuilder) Build() MappingCharFilter {
	return *rb.v
}

func (rb *MappingCharFilterBuilder) Mappings(mappings ...string) *MappingCharFilterBuilder {
	rb.v.Mappings = mappings
	return rb
}

func (rb *MappingCharFilterBuilder) MappingsPath(mappingspath string) *MappingCharFilterBuilder {
	rb.v.MappingsPath = &mappingspath
	return rb
}

func (rb *MappingCharFilterBuilder) Version(version VersionString) *MappingCharFilterBuilder {
	rb.v.Version = &version
	return rb
}
