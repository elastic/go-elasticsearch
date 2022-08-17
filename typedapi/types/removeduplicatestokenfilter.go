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

// RemoveDuplicatesTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L300-L302
type RemoveDuplicatesTokenFilter struct {
	Type    string         `json:"type,omitempty"`
	Version *VersionString `json:"version,omitempty"`
}

// RemoveDuplicatesTokenFilterBuilder holds RemoveDuplicatesTokenFilter struct and provides a builder API.
type RemoveDuplicatesTokenFilterBuilder struct {
	v *RemoveDuplicatesTokenFilter
}

// NewRemoveDuplicatesTokenFilter provides a builder for the RemoveDuplicatesTokenFilter struct.
func NewRemoveDuplicatesTokenFilterBuilder() *RemoveDuplicatesTokenFilterBuilder {
	r := RemoveDuplicatesTokenFilterBuilder{
		&RemoveDuplicatesTokenFilter{},
	}

	r.v.Type = "remove_duplicates"

	return &r
}

// Build finalize the chain and returns the RemoveDuplicatesTokenFilter struct
func (rb *RemoveDuplicatesTokenFilterBuilder) Build() RemoveDuplicatesTokenFilter {
	return *rb.v
}

func (rb *RemoveDuplicatesTokenFilterBuilder) Version(version VersionString) *RemoveDuplicatesTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
