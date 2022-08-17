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

// MappingLimitSettingsFieldNameLength type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L454-L461
type MappingLimitSettingsFieldNameLength struct {
	// Limit Setting for the maximum length of a field name. This setting isn’t really
	// something that addresses mappings explosion but
	// might still be useful if you want to limit the field length. It usually
	// shouldn’t be necessary to set this setting. The
	// default is okay unless a user starts to add a huge number of fields with
	// really long names. Default is `Long.MAX_VALUE` (no limit).
	Limit *int64 `json:"limit,omitempty"`
}

// MappingLimitSettingsFieldNameLengthBuilder holds MappingLimitSettingsFieldNameLength struct and provides a builder API.
type MappingLimitSettingsFieldNameLengthBuilder struct {
	v *MappingLimitSettingsFieldNameLength
}

// NewMappingLimitSettingsFieldNameLength provides a builder for the MappingLimitSettingsFieldNameLength struct.
func NewMappingLimitSettingsFieldNameLengthBuilder() *MappingLimitSettingsFieldNameLengthBuilder {
	r := MappingLimitSettingsFieldNameLengthBuilder{
		&MappingLimitSettingsFieldNameLength{},
	}

	return &r
}

// Build finalize the chain and returns the MappingLimitSettingsFieldNameLength struct
func (rb *MappingLimitSettingsFieldNameLengthBuilder) Build() MappingLimitSettingsFieldNameLength {
	return *rb.v
}

// Limit Setting for the maximum length of a field name. This setting isn’t really
// something that addresses mappings explosion but
// might still be useful if you want to limit the field length. It usually
// shouldn’t be necessary to set this setting. The
// default is okay unless a user starts to add a huge number of fields with
// really long names. Default is `Long.MAX_VALUE` (no limit).

func (rb *MappingLimitSettingsFieldNameLengthBuilder) Limit(limit int64) *MappingLimitSettingsFieldNameLengthBuilder {
	rb.v.Limit = &limit
	return rb
}
