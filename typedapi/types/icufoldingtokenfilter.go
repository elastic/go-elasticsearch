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

// IcuFoldingTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/icu-plugin.ts#L46-L49
type IcuFoldingTokenFilter struct {
	Type             string         `json:"type,omitempty"`
	UnicodeSetFilter string         `json:"unicode_set_filter"`
	Version          *VersionString `json:"version,omitempty"`
}

// IcuFoldingTokenFilterBuilder holds IcuFoldingTokenFilter struct and provides a builder API.
type IcuFoldingTokenFilterBuilder struct {
	v *IcuFoldingTokenFilter
}

// NewIcuFoldingTokenFilter provides a builder for the IcuFoldingTokenFilter struct.
func NewIcuFoldingTokenFilterBuilder() *IcuFoldingTokenFilterBuilder {
	r := IcuFoldingTokenFilterBuilder{
		&IcuFoldingTokenFilter{},
	}

	r.v.Type = "icu_folding"

	return &r
}

// Build finalize the chain and returns the IcuFoldingTokenFilter struct
func (rb *IcuFoldingTokenFilterBuilder) Build() IcuFoldingTokenFilter {
	return *rb.v
}

func (rb *IcuFoldingTokenFilterBuilder) UnicodeSetFilter(unicodesetfilter string) *IcuFoldingTokenFilterBuilder {
	rb.v.UnicodeSetFilter = unicodesetfilter
	return rb
}

func (rb *IcuFoldingTokenFilterBuilder) Version(version VersionString) *IcuFoldingTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
