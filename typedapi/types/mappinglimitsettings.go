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

// MappingLimitSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L402-L414
type MappingLimitSettings struct {
	Depth           *MappingLimitSettingsDepth           `json:"depth,omitempty"`
	DimensionFields *MappingLimitSettingsDimensionFields `json:"dimension_fields,omitempty"`
	FieldNameLength *MappingLimitSettingsFieldNameLength `json:"field_name_length,omitempty"`
	IgnoreMalformed *bool                                `json:"ignore_malformed,omitempty"`
	NestedFields    *MappingLimitSettingsNestedFields    `json:"nested_fields,omitempty"`
	NestedObjects   *MappingLimitSettingsNestedObjects   `json:"nested_objects,omitempty"`
	TotalFields     *MappingLimitSettingsTotalFields     `json:"total_fields,omitempty"`
}

// MappingLimitSettingsBuilder holds MappingLimitSettings struct and provides a builder API.
type MappingLimitSettingsBuilder struct {
	v *MappingLimitSettings
}

// NewMappingLimitSettings provides a builder for the MappingLimitSettings struct.
func NewMappingLimitSettingsBuilder() *MappingLimitSettingsBuilder {
	r := MappingLimitSettingsBuilder{
		&MappingLimitSettings{},
	}

	return &r
}

// Build finalize the chain and returns the MappingLimitSettings struct
func (rb *MappingLimitSettingsBuilder) Build() MappingLimitSettings {
	return *rb.v
}

func (rb *MappingLimitSettingsBuilder) Depth(depth *MappingLimitSettingsDepthBuilder) *MappingLimitSettingsBuilder {
	v := depth.Build()
	rb.v.Depth = &v
	return rb
}

func (rb *MappingLimitSettingsBuilder) DimensionFields(dimensionfields *MappingLimitSettingsDimensionFieldsBuilder) *MappingLimitSettingsBuilder {
	v := dimensionfields.Build()
	rb.v.DimensionFields = &v
	return rb
}

func (rb *MappingLimitSettingsBuilder) FieldNameLength(fieldnamelength *MappingLimitSettingsFieldNameLengthBuilder) *MappingLimitSettingsBuilder {
	v := fieldnamelength.Build()
	rb.v.FieldNameLength = &v
	return rb
}

func (rb *MappingLimitSettingsBuilder) IgnoreMalformed(ignoremalformed bool) *MappingLimitSettingsBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *MappingLimitSettingsBuilder) NestedFields(nestedfields *MappingLimitSettingsNestedFieldsBuilder) *MappingLimitSettingsBuilder {
	v := nestedfields.Build()
	rb.v.NestedFields = &v
	return rb
}

func (rb *MappingLimitSettingsBuilder) NestedObjects(nestedobjects *MappingLimitSettingsNestedObjectsBuilder) *MappingLimitSettingsBuilder {
	v := nestedobjects.Build()
	rb.v.NestedObjects = &v
	return rb
}

func (rb *MappingLimitSettingsBuilder) TotalFields(totalfields *MappingLimitSettingsTotalFieldsBuilder) *MappingLimitSettingsBuilder {
	v := totalfields.Build()
	rb.v.TotalFields = &v
	return rb
}
