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

// KeepWordsTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L223-L228
type KeepWordsTokenFilter struct {
	KeepWords     []string       `json:"keep_words,omitempty"`
	KeepWordsCase *bool          `json:"keep_words_case,omitempty"`
	KeepWordsPath *string        `json:"keep_words_path,omitempty"`
	Type          string         `json:"type,omitempty"`
	Version       *VersionString `json:"version,omitempty"`
}

// KeepWordsTokenFilterBuilder holds KeepWordsTokenFilter struct and provides a builder API.
type KeepWordsTokenFilterBuilder struct {
	v *KeepWordsTokenFilter
}

// NewKeepWordsTokenFilter provides a builder for the KeepWordsTokenFilter struct.
func NewKeepWordsTokenFilterBuilder() *KeepWordsTokenFilterBuilder {
	r := KeepWordsTokenFilterBuilder{
		&KeepWordsTokenFilter{},
	}

	r.v.Type = "keep"

	return &r
}

// Build finalize the chain and returns the KeepWordsTokenFilter struct
func (rb *KeepWordsTokenFilterBuilder) Build() KeepWordsTokenFilter {
	return *rb.v
}

func (rb *KeepWordsTokenFilterBuilder) KeepWords(keep_words ...string) *KeepWordsTokenFilterBuilder {
	rb.v.KeepWords = keep_words
	return rb
}

func (rb *KeepWordsTokenFilterBuilder) KeepWordsCase(keepwordscase bool) *KeepWordsTokenFilterBuilder {
	rb.v.KeepWordsCase = &keepwordscase
	return rb
}

func (rb *KeepWordsTokenFilterBuilder) KeepWordsPath(keepwordspath string) *KeepWordsTokenFilterBuilder {
	rb.v.KeepWordsPath = &keepwordspath
	return rb
}

func (rb *KeepWordsTokenFilterBuilder) Version(version VersionString) *KeepWordsTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
