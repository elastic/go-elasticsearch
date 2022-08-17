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

// FieldTypesMappings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L96-L99
type FieldTypesMappings struct {
	FieldTypes        []FieldTypes        `json:"field_types"`
	RuntimeFieldTypes []RuntimeFieldTypes `json:"runtime_field_types,omitempty"`
}

// FieldTypesMappingsBuilder holds FieldTypesMappings struct and provides a builder API.
type FieldTypesMappingsBuilder struct {
	v *FieldTypesMappings
}

// NewFieldTypesMappings provides a builder for the FieldTypesMappings struct.
func NewFieldTypesMappingsBuilder() *FieldTypesMappingsBuilder {
	r := FieldTypesMappingsBuilder{
		&FieldTypesMappings{},
	}

	return &r
}

// Build finalize the chain and returns the FieldTypesMappings struct
func (rb *FieldTypesMappingsBuilder) Build() FieldTypesMappings {
	return *rb.v
}

func (rb *FieldTypesMappingsBuilder) FieldTypes(field_types []FieldTypesBuilder) *FieldTypesMappingsBuilder {
	tmp := make([]FieldTypes, len(field_types))
	for _, value := range field_types {
		tmp = append(tmp, value.Build())
	}
	rb.v.FieldTypes = tmp
	return rb
}

func (rb *FieldTypesMappingsBuilder) RuntimeFieldTypes(runtime_field_types []RuntimeFieldTypesBuilder) *FieldTypesMappingsBuilder {
	tmp := make([]RuntimeFieldTypes, len(runtime_field_types))
	for _, value := range runtime_field_types {
		tmp = append(tmp, value.Build())
	}
	rb.v.RuntimeFieldTypes = tmp
	return rb
}
