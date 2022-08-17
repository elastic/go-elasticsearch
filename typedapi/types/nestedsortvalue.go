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

// NestedSortValue type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/sort.ts#L29-L34
type NestedSortValue struct {
	Filter      *QueryContainer  `json:"filter,omitempty"`
	MaxChildren *int             `json:"max_children,omitempty"`
	Nested      *NestedSortValue `json:"nested,omitempty"`
	Path        Field            `json:"path"`
}

// NestedSortValueBuilder holds NestedSortValue struct and provides a builder API.
type NestedSortValueBuilder struct {
	v *NestedSortValue
}

// NewNestedSortValue provides a builder for the NestedSortValue struct.
func NewNestedSortValueBuilder() *NestedSortValueBuilder {
	r := NestedSortValueBuilder{
		&NestedSortValue{},
	}

	return &r
}

// Build finalize the chain and returns the NestedSortValue struct
func (rb *NestedSortValueBuilder) Build() NestedSortValue {
	return *rb.v
}

func (rb *NestedSortValueBuilder) Filter(filter *QueryContainerBuilder) *NestedSortValueBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *NestedSortValueBuilder) MaxChildren(maxchildren int) *NestedSortValueBuilder {
	rb.v.MaxChildren = &maxchildren
	return rb
}

func (rb *NestedSortValueBuilder) Nested(nested *NestedSortValueBuilder) *NestedSortValueBuilder {
	v := nested.Build()
	rb.v.Nested = &v
	return rb
}

func (rb *NestedSortValueBuilder) Path(path Field) *NestedSortValueBuilder {
	rb.v.Path = path
	return rb
}
