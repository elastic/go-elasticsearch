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

// Flattened type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L336-L338
type Flattened struct {
	Available  bool `json:"available"`
	Enabled    bool `json:"enabled"`
	FieldCount int  `json:"field_count"`
}

// FlattenedBuilder holds Flattened struct and provides a builder API.
type FlattenedBuilder struct {
	v *Flattened
}

// NewFlattened provides a builder for the Flattened struct.
func NewFlattenedBuilder() *FlattenedBuilder {
	r := FlattenedBuilder{
		&Flattened{},
	}

	return &r
}

// Build finalize the chain and returns the Flattened struct
func (rb *FlattenedBuilder) Build() Flattened {
	return *rb.v
}

func (rb *FlattenedBuilder) Available(available bool) *FlattenedBuilder {
	rb.v.Available = available
	return rb
}

func (rb *FlattenedBuilder) Enabled(enabled bool) *FlattenedBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *FlattenedBuilder) FieldCount(fieldcount int) *FlattenedBuilder {
	rb.v.FieldCount = fieldcount
	return rb
}
