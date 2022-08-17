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

// NGramTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L265-L270
type NGramTokenFilter struct {
	MaxGram          *int           `json:"max_gram,omitempty"`
	MinGram          *int           `json:"min_gram,omitempty"`
	PreserveOriginal *bool          `json:"preserve_original,omitempty"`
	Type             string         `json:"type,omitempty"`
	Version          *VersionString `json:"version,omitempty"`
}

// NGramTokenFilterBuilder holds NGramTokenFilter struct and provides a builder API.
type NGramTokenFilterBuilder struct {
	v *NGramTokenFilter
}

// NewNGramTokenFilter provides a builder for the NGramTokenFilter struct.
func NewNGramTokenFilterBuilder() *NGramTokenFilterBuilder {
	r := NGramTokenFilterBuilder{
		&NGramTokenFilter{},
	}

	r.v.Type = "ngram"

	return &r
}

// Build finalize the chain and returns the NGramTokenFilter struct
func (rb *NGramTokenFilterBuilder) Build() NGramTokenFilter {
	return *rb.v
}

func (rb *NGramTokenFilterBuilder) MaxGram(maxgram int) *NGramTokenFilterBuilder {
	rb.v.MaxGram = &maxgram
	return rb
}

func (rb *NGramTokenFilterBuilder) MinGram(mingram int) *NGramTokenFilterBuilder {
	rb.v.MinGram = &mingram
	return rb
}

func (rb *NGramTokenFilterBuilder) PreserveOriginal(preserveoriginal bool) *NGramTokenFilterBuilder {
	rb.v.PreserveOriginal = &preserveoriginal
	return rb
}

func (rb *NGramTokenFilterBuilder) Version(version VersionString) *NGramTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
