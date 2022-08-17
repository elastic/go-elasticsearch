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

// FieldTypes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L101-L107
type FieldTypes struct {
	Count       int  `json:"count"`
	IndexCount  int  `json:"index_count"`
	Name        Name `json:"name"`
	ScriptCount *int `json:"script_count,omitempty"`
}

// FieldTypesBuilder holds FieldTypes struct and provides a builder API.
type FieldTypesBuilder struct {
	v *FieldTypes
}

// NewFieldTypes provides a builder for the FieldTypes struct.
func NewFieldTypesBuilder() *FieldTypesBuilder {
	r := FieldTypesBuilder{
		&FieldTypes{},
	}

	return &r
}

// Build finalize the chain and returns the FieldTypes struct
func (rb *FieldTypesBuilder) Build() FieldTypes {
	return *rb.v
}

func (rb *FieldTypesBuilder) Count(count int) *FieldTypesBuilder {
	rb.v.Count = count
	return rb
}

func (rb *FieldTypesBuilder) IndexCount(indexcount int) *FieldTypesBuilder {
	rb.v.IndexCount = indexcount
	return rb
}

func (rb *FieldTypesBuilder) Name(name Name) *FieldTypesBuilder {
	rb.v.Name = name
	return rb
}

func (rb *FieldTypesBuilder) ScriptCount(scriptcount int) *FieldTypesBuilder {
	rb.v.ScriptCount = &scriptcount
	return rb
}
