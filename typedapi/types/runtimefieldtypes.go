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

// RuntimeFieldTypes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L258-L260
type RuntimeFieldTypes struct {
	Available  bool                `json:"available"`
	Enabled    bool                `json:"enabled"`
	FieldTypes []RuntimeFieldsType `json:"field_types"`
}

// RuntimeFieldTypesBuilder holds RuntimeFieldTypes struct and provides a builder API.
type RuntimeFieldTypesBuilder struct {
	v *RuntimeFieldTypes
}

// NewRuntimeFieldTypes provides a builder for the RuntimeFieldTypes struct.
func NewRuntimeFieldTypesBuilder() *RuntimeFieldTypesBuilder {
	r := RuntimeFieldTypesBuilder{
		&RuntimeFieldTypes{},
	}

	return &r
}

// Build finalize the chain and returns the RuntimeFieldTypes struct
func (rb *RuntimeFieldTypesBuilder) Build() RuntimeFieldTypes {
	return *rb.v
}

func (rb *RuntimeFieldTypesBuilder) Available(available bool) *RuntimeFieldTypesBuilder {
	rb.v.Available = available
	return rb
}

func (rb *RuntimeFieldTypesBuilder) Enabled(enabled bool) *RuntimeFieldTypesBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *RuntimeFieldTypesBuilder) FieldTypes(field_types []RuntimeFieldsTypeBuilder) *RuntimeFieldTypesBuilder {
	tmp := make([]RuntimeFieldsType, len(field_types))
	for _, value := range field_types {
		tmp = append(tmp, value.Build())
	}
	rb.v.FieldTypes = tmp
	return rb
}
