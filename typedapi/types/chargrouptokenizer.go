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

// CharGroupTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/tokenizers.ts#L55-L59
type CharGroupTokenizer struct {
	MaxTokenLength  *int           `json:"max_token_length,omitempty"`
	TokenizeOnChars []string       `json:"tokenize_on_chars"`
	Type            string         `json:"type,omitempty"`
	Version         *VersionString `json:"version,omitempty"`
}

// CharGroupTokenizerBuilder holds CharGroupTokenizer struct and provides a builder API.
type CharGroupTokenizerBuilder struct {
	v *CharGroupTokenizer
}

// NewCharGroupTokenizer provides a builder for the CharGroupTokenizer struct.
func NewCharGroupTokenizerBuilder() *CharGroupTokenizerBuilder {
	r := CharGroupTokenizerBuilder{
		&CharGroupTokenizer{},
	}

	r.v.Type = "char_group"

	return &r
}

// Build finalize the chain and returns the CharGroupTokenizer struct
func (rb *CharGroupTokenizerBuilder) Build() CharGroupTokenizer {
	return *rb.v
}

func (rb *CharGroupTokenizerBuilder) MaxTokenLength(maxtokenlength int) *CharGroupTokenizerBuilder {
	rb.v.MaxTokenLength = &maxtokenlength
	return rb
}

func (rb *CharGroupTokenizerBuilder) TokenizeOnChars(tokenize_on_chars ...string) *CharGroupTokenizerBuilder {
	rb.v.TokenizeOnChars = tokenize_on_chars
	return rb
}

func (rb *CharGroupTokenizerBuilder) Version(version VersionString) *CharGroupTokenizerBuilder {
	rb.v.Version = &version
	return rb
}
