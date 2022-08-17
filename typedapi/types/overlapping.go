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

// Overlapping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/simulate_template/IndicesSimulateTemplateResponse.ts#L39-L42
type Overlapping struct {
	IndexPatterns []string `json:"index_patterns"`
	Name          Name     `json:"name"`
}

// OverlappingBuilder holds Overlapping struct and provides a builder API.
type OverlappingBuilder struct {
	v *Overlapping
}

// NewOverlapping provides a builder for the Overlapping struct.
func NewOverlappingBuilder() *OverlappingBuilder {
	r := OverlappingBuilder{
		&Overlapping{},
	}

	return &r
}

// Build finalize the chain and returns the Overlapping struct
func (rb *OverlappingBuilder) Build() Overlapping {
	return *rb.v
}

func (rb *OverlappingBuilder) IndexPatterns(index_patterns ...string) *OverlappingBuilder {
	rb.v.IndexPatterns = index_patterns
	return rb
}

func (rb *OverlappingBuilder) Name(name Name) *OverlappingBuilder {
	rb.v.Name = name
	return rb
}
