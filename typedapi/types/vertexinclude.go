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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

// VertexInclude type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/graph/_types/Vertex.ts#L39-L42
type VertexInclude struct {
	Boost float64 `json:"boost"`
	Term  string  `json:"term"`
}

// VertexIncludeBuilder holds VertexInclude struct and provides a builder API.
type VertexIncludeBuilder struct {
	v *VertexInclude
}

// NewVertexInclude provides a builder for the VertexInclude struct.
func NewVertexIncludeBuilder() *VertexIncludeBuilder {
	r := VertexIncludeBuilder{
		&VertexInclude{},
	}

	return &r
}

// Build finalize the chain and returns the VertexInclude struct
func (rb *VertexIncludeBuilder) Build() VertexInclude {
	return *rb.v
}

func (rb *VertexIncludeBuilder) Boost(boost float64) *VertexIncludeBuilder {
	rb.v.Boost = boost
	return rb
}

func (rb *VertexIncludeBuilder) Term(term string) *VertexIncludeBuilder {
	rb.v.Term = term
	return rb
}
