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

// PatternReplaceCharFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/char_filters.ts#L53-L58
type PatternReplaceCharFilter struct {
	Flags       *string        `json:"flags,omitempty"`
	Pattern     string         `json:"pattern"`
	Replacement *string        `json:"replacement,omitempty"`
	Type        string         `json:"type,omitempty"`
	Version     *VersionString `json:"version,omitempty"`
}

// PatternReplaceCharFilterBuilder holds PatternReplaceCharFilter struct and provides a builder API.
type PatternReplaceCharFilterBuilder struct {
	v *PatternReplaceCharFilter
}

// NewPatternReplaceCharFilter provides a builder for the PatternReplaceCharFilter struct.
func NewPatternReplaceCharFilterBuilder() *PatternReplaceCharFilterBuilder {
	r := PatternReplaceCharFilterBuilder{
		&PatternReplaceCharFilter{},
	}

	r.v.Type = "pattern_replace"

	return &r
}

// Build finalize the chain and returns the PatternReplaceCharFilter struct
func (rb *PatternReplaceCharFilterBuilder) Build() PatternReplaceCharFilter {
	return *rb.v
}

func (rb *PatternReplaceCharFilterBuilder) Flags(flags string) *PatternReplaceCharFilterBuilder {
	rb.v.Flags = &flags
	return rb
}

func (rb *PatternReplaceCharFilterBuilder) Pattern(pattern string) *PatternReplaceCharFilterBuilder {
	rb.v.Pattern = pattern
	return rb
}

func (rb *PatternReplaceCharFilterBuilder) Replacement(replacement string) *PatternReplaceCharFilterBuilder {
	rb.v.Replacement = &replacement
	return rb
}

func (rb *PatternReplaceCharFilterBuilder) Version(version VersionString) *PatternReplaceCharFilterBuilder {
	rb.v.Version = &version
	return rb
}
