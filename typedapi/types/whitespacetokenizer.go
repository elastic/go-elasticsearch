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

// WhitespaceTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/tokenizers.ts#L114-L117
type WhitespaceTokenizer struct {
	MaxTokenLength *int           `json:"max_token_length,omitempty"`
	Type           string         `json:"type,omitempty"`
	Version        *VersionString `json:"version,omitempty"`
}

// WhitespaceTokenizerBuilder holds WhitespaceTokenizer struct and provides a builder API.
type WhitespaceTokenizerBuilder struct {
	v *WhitespaceTokenizer
}

// NewWhitespaceTokenizer provides a builder for the WhitespaceTokenizer struct.
func NewWhitespaceTokenizerBuilder() *WhitespaceTokenizerBuilder {
	r := WhitespaceTokenizerBuilder{
		&WhitespaceTokenizer{},
	}

	r.v.Type = "whitespace"

	return &r
}

// Build finalize the chain and returns the WhitespaceTokenizer struct
func (rb *WhitespaceTokenizerBuilder) Build() WhitespaceTokenizer {
	return *rb.v
}

func (rb *WhitespaceTokenizerBuilder) MaxTokenLength(maxtokenlength int) *WhitespaceTokenizerBuilder {
	rb.v.MaxTokenLength = &maxtokenlength
	return rb
}

func (rb *WhitespaceTokenizerBuilder) Version(version VersionString) *WhitespaceTokenizerBuilder {
	rb.v.Version = &version
	return rb
}
