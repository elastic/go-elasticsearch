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

// PredicateTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L295-L298
type PredicateTokenFilter struct {
	Script  Script         `json:"script"`
	Type    string         `json:"type,omitempty"`
	Version *VersionString `json:"version,omitempty"`
}

// PredicateTokenFilterBuilder holds PredicateTokenFilter struct and provides a builder API.
type PredicateTokenFilterBuilder struct {
	v *PredicateTokenFilter
}

// NewPredicateTokenFilter provides a builder for the PredicateTokenFilter struct.
func NewPredicateTokenFilterBuilder() *PredicateTokenFilterBuilder {
	r := PredicateTokenFilterBuilder{
		&PredicateTokenFilter{},
	}

	r.v.Type = "predicate_token_filter"

	return &r
}

// Build finalize the chain and returns the PredicateTokenFilter struct
func (rb *PredicateTokenFilterBuilder) Build() PredicateTokenFilter {
	return *rb.v
}

func (rb *PredicateTokenFilterBuilder) Script(script *ScriptBuilder) *PredicateTokenFilterBuilder {
	v := script.Build()
	rb.v.Script = v
	return rb
}

func (rb *PredicateTokenFilterBuilder) Version(version VersionString) *PredicateTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
