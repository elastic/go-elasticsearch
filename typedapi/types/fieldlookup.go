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

// FieldLookup type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/abstractions.ts#L164-L169
type FieldLookup struct {
	Id      Id         `json:"id"`
	Index   *IndexName `json:"index,omitempty"`
	Path    *Field     `json:"path,omitempty"`
	Routing *Routing   `json:"routing,omitempty"`
}

// FieldLookupBuilder holds FieldLookup struct and provides a builder API.
type FieldLookupBuilder struct {
	v *FieldLookup
}

// NewFieldLookup provides a builder for the FieldLookup struct.
func NewFieldLookupBuilder() *FieldLookupBuilder {
	r := FieldLookupBuilder{
		&FieldLookup{},
	}

	return &r
}

// Build finalize the chain and returns the FieldLookup struct
func (rb *FieldLookupBuilder) Build() FieldLookup {
	return *rb.v
}

func (rb *FieldLookupBuilder) Id(id Id) *FieldLookupBuilder {
	rb.v.Id = id
	return rb
}

func (rb *FieldLookupBuilder) Index(index IndexName) *FieldLookupBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *FieldLookupBuilder) Path(path Field) *FieldLookupBuilder {
	rb.v.Path = &path
	return rb
}

func (rb *FieldLookupBuilder) Routing(routing Routing) *FieldLookupBuilder {
	rb.v.Routing = &routing
	return rb
}
