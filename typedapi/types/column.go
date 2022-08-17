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

// Column type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/sql/types.ts#L23-L26
type Column struct {
	Name Name   `json:"name"`
	Type string `json:"type"`
}

// ColumnBuilder holds Column struct and provides a builder API.
type ColumnBuilder struct {
	v *Column
}

// NewColumn provides a builder for the Column struct.
func NewColumnBuilder() *ColumnBuilder {
	r := ColumnBuilder{
		&Column{},
	}

	return &r
}

// Build finalize the chain and returns the Column struct
func (rb *ColumnBuilder) Build() Column {
	return *rb.v
}

func (rb *ColumnBuilder) Name(name Name) *ColumnBuilder {
	rb.v.Name = name
	return rb
}

func (rb *ColumnBuilder) Type_(type_ string) *ColumnBuilder {
	rb.v.Type = type_
	return rb
}
