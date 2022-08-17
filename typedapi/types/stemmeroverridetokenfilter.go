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

// StemmerOverrideTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L313-L317
type StemmerOverrideTokenFilter struct {
	Rules     []string       `json:"rules,omitempty"`
	RulesPath *string        `json:"rules_path,omitempty"`
	Type      string         `json:"type,omitempty"`
	Version   *VersionString `json:"version,omitempty"`
}

// StemmerOverrideTokenFilterBuilder holds StemmerOverrideTokenFilter struct and provides a builder API.
type StemmerOverrideTokenFilterBuilder struct {
	v *StemmerOverrideTokenFilter
}

// NewStemmerOverrideTokenFilter provides a builder for the StemmerOverrideTokenFilter struct.
func NewStemmerOverrideTokenFilterBuilder() *StemmerOverrideTokenFilterBuilder {
	r := StemmerOverrideTokenFilterBuilder{
		&StemmerOverrideTokenFilter{},
	}

	r.v.Type = "stemmer_override"

	return &r
}

// Build finalize the chain and returns the StemmerOverrideTokenFilter struct
func (rb *StemmerOverrideTokenFilterBuilder) Build() StemmerOverrideTokenFilter {
	return *rb.v
}

func (rb *StemmerOverrideTokenFilterBuilder) Rules(rules ...string) *StemmerOverrideTokenFilterBuilder {
	rb.v.Rules = rules
	return rb
}

func (rb *StemmerOverrideTokenFilterBuilder) RulesPath(rulespath string) *StemmerOverrideTokenFilterBuilder {
	rb.v.RulesPath = &rulespath
	return rb
}

func (rb *StemmerOverrideTokenFilterBuilder) Version(version VersionString) *StemmerOverrideTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
