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

// CommonGramsTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L172-L178
type CommonGramsTokenFilter struct {
	CommonWords     []string       `json:"common_words,omitempty"`
	CommonWordsPath *string        `json:"common_words_path,omitempty"`
	IgnoreCase      *bool          `json:"ignore_case,omitempty"`
	QueryMode       *bool          `json:"query_mode,omitempty"`
	Type            string         `json:"type,omitempty"`
	Version         *VersionString `json:"version,omitempty"`
}

// CommonGramsTokenFilterBuilder holds CommonGramsTokenFilter struct and provides a builder API.
type CommonGramsTokenFilterBuilder struct {
	v *CommonGramsTokenFilter
}

// NewCommonGramsTokenFilter provides a builder for the CommonGramsTokenFilter struct.
func NewCommonGramsTokenFilterBuilder() *CommonGramsTokenFilterBuilder {
	r := CommonGramsTokenFilterBuilder{
		&CommonGramsTokenFilter{},
	}

	r.v.Type = "common_grams"

	return &r
}

// Build finalize the chain and returns the CommonGramsTokenFilter struct
func (rb *CommonGramsTokenFilterBuilder) Build() CommonGramsTokenFilter {
	return *rb.v
}

func (rb *CommonGramsTokenFilterBuilder) CommonWords(common_words ...string) *CommonGramsTokenFilterBuilder {
	rb.v.CommonWords = common_words
	return rb
}

func (rb *CommonGramsTokenFilterBuilder) CommonWordsPath(commonwordspath string) *CommonGramsTokenFilterBuilder {
	rb.v.CommonWordsPath = &commonwordspath
	return rb
}

func (rb *CommonGramsTokenFilterBuilder) IgnoreCase(ignorecase bool) *CommonGramsTokenFilterBuilder {
	rb.v.IgnoreCase = &ignorecase
	return rb
}

func (rb *CommonGramsTokenFilterBuilder) QueryMode(querymode bool) *CommonGramsTokenFilterBuilder {
	rb.v.QueryMode = &querymode
	return rb
}

func (rb *CommonGramsTokenFilterBuilder) Version(version VersionString) *CommonGramsTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
