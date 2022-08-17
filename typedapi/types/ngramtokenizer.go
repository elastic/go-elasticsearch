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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/tokenchar"
)

// NGramTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/tokenizers.ts#L38-L44
type NGramTokenizer struct {
	CustomTokenChars *string               `json:"custom_token_chars,omitempty"`
	MaxGram          int                   `json:"max_gram"`
	MinGram          int                   `json:"min_gram"`
	TokenChars       []tokenchar.TokenChar `json:"token_chars"`
	Type             string                `json:"type,omitempty"`
	Version          *VersionString        `json:"version,omitempty"`
}

// NGramTokenizerBuilder holds NGramTokenizer struct and provides a builder API.
type NGramTokenizerBuilder struct {
	v *NGramTokenizer
}

// NewNGramTokenizer provides a builder for the NGramTokenizer struct.
func NewNGramTokenizerBuilder() *NGramTokenizerBuilder {
	r := NGramTokenizerBuilder{
		&NGramTokenizer{},
	}

	r.v.Type = "ngram"

	return &r
}

// Build finalize the chain and returns the NGramTokenizer struct
func (rb *NGramTokenizerBuilder) Build() NGramTokenizer {
	return *rb.v
}

func (rb *NGramTokenizerBuilder) CustomTokenChars(customtokenchars string) *NGramTokenizerBuilder {
	rb.v.CustomTokenChars = &customtokenchars
	return rb
}

func (rb *NGramTokenizerBuilder) MaxGram(maxgram int) *NGramTokenizerBuilder {
	rb.v.MaxGram = maxgram
	return rb
}

func (rb *NGramTokenizerBuilder) MinGram(mingram int) *NGramTokenizerBuilder {
	rb.v.MinGram = mingram
	return rb
}

func (rb *NGramTokenizerBuilder) TokenChars(token_chars ...tokenchar.TokenChar) *NGramTokenizerBuilder {
	rb.v.TokenChars = token_chars
	return rb
}

func (rb *NGramTokenizerBuilder) Version(version VersionString) *NGramTokenizerBuilder {
	rb.v.Version = &version
	return rb
}
