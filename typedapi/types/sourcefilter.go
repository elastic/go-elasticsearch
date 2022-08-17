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

// SourceFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/SourceFilter.ts#L23-L31
type SourceFilter struct {
	Excludes *Fields `json:"excludes,omitempty"`
	Includes *Fields `json:"includes,omitempty"`
}

// SourceFilterBuilder holds SourceFilter struct and provides a builder API.
type SourceFilterBuilder struct {
	v *SourceFilter
}

// NewSourceFilter provides a builder for the SourceFilter struct.
func NewSourceFilterBuilder() *SourceFilterBuilder {
	r := SourceFilterBuilder{
		&SourceFilter{},
	}

	return &r
}

// Build finalize the chain and returns the SourceFilter struct
func (rb *SourceFilterBuilder) Build() SourceFilter {
	return *rb.v
}

func (rb *SourceFilterBuilder) Excludes(excludes *FieldsBuilder) *SourceFilterBuilder {
	v := excludes.Build()
	rb.v.Excludes = &v
	return rb
}

func (rb *SourceFilterBuilder) Includes(includes *FieldsBuilder) *SourceFilterBuilder {
	v := includes.Build()
	rb.v.Includes = &v
	return rb
}
