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

// LengthTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L242-L246
type LengthTokenFilter struct {
	Max     *int           `json:"max,omitempty"`
	Min     *int           `json:"min,omitempty"`
	Type    string         `json:"type,omitempty"`
	Version *VersionString `json:"version,omitempty"`
}

// LengthTokenFilterBuilder holds LengthTokenFilter struct and provides a builder API.
type LengthTokenFilterBuilder struct {
	v *LengthTokenFilter
}

// NewLengthTokenFilter provides a builder for the LengthTokenFilter struct.
func NewLengthTokenFilterBuilder() *LengthTokenFilterBuilder {
	r := LengthTokenFilterBuilder{
		&LengthTokenFilter{},
	}

	r.v.Type = "length"

	return &r
}

// Build finalize the chain and returns the LengthTokenFilter struct
func (rb *LengthTokenFilterBuilder) Build() LengthTokenFilter {
	return *rb.v
}

func (rb *LengthTokenFilterBuilder) Max(max int) *LengthTokenFilterBuilder {
	rb.v.Max = &max
	return rb
}

func (rb *LengthTokenFilterBuilder) Min(min int) *LengthTokenFilterBuilder {
	rb.v.Min = &min
	return rb
}

func (rb *LengthTokenFilterBuilder) Version(version VersionString) *LengthTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
