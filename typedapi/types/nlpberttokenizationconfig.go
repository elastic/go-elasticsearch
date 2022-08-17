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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/tokenizationtruncate"
)

// NlpBertTokenizationConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L116-L143
type NlpBertTokenizationConfig struct {
	// DoLowerCase Should the tokenizer lower case the text
	DoLowerCase *bool `json:"do_lower_case,omitempty"`
	// MaxSequenceLength Maximum input sequence length for the model
	MaxSequenceLength *int `json:"max_sequence_length,omitempty"`
	// Span Tokenization spanning options. Special value of -1 indicates no spanning
	// takes place
	Span *int `json:"span,omitempty"`
	// Truncate Should tokenization input be automatically truncated before sending to the
	// model for inference
	Truncate *tokenizationtruncate.TokenizationTruncate `json:"truncate,omitempty"`
	// WithSpecialTokens Is tokenization completed with special tokens
	WithSpecialTokens *bool `json:"with_special_tokens,omitempty"`
}

// NlpBertTokenizationConfigBuilder holds NlpBertTokenizationConfig struct and provides a builder API.
type NlpBertTokenizationConfigBuilder struct {
	v *NlpBertTokenizationConfig
}

// NewNlpBertTokenizationConfig provides a builder for the NlpBertTokenizationConfig struct.
func NewNlpBertTokenizationConfigBuilder() *NlpBertTokenizationConfigBuilder {
	r := NlpBertTokenizationConfigBuilder{
		&NlpBertTokenizationConfig{},
	}

	return &r
}

// Build finalize the chain and returns the NlpBertTokenizationConfig struct
func (rb *NlpBertTokenizationConfigBuilder) Build() NlpBertTokenizationConfig {
	return *rb.v
}

// DoLowerCase Should the tokenizer lower case the text

func (rb *NlpBertTokenizationConfigBuilder) DoLowerCase(dolowercase bool) *NlpBertTokenizationConfigBuilder {
	rb.v.DoLowerCase = &dolowercase
	return rb
}

// MaxSequenceLength Maximum input sequence length for the model

func (rb *NlpBertTokenizationConfigBuilder) MaxSequenceLength(maxsequencelength int) *NlpBertTokenizationConfigBuilder {
	rb.v.MaxSequenceLength = &maxsequencelength
	return rb
}

// Span Tokenization spanning options. Special value of -1 indicates no spanning
// takes place

func (rb *NlpBertTokenizationConfigBuilder) Span(span int) *NlpBertTokenizationConfigBuilder {
	rb.v.Span = &span
	return rb
}

// Truncate Should tokenization input be automatically truncated before sending to the
// model for inference

func (rb *NlpBertTokenizationConfigBuilder) Truncate(truncate tokenizationtruncate.TokenizationTruncate) *NlpBertTokenizationConfigBuilder {
	rb.v.Truncate = &truncate
	return rb
}

// WithSpecialTokens Is tokenization completed with special tokens

func (rb *NlpBertTokenizationConfigBuilder) WithSpecialTokens(withspecialtokens bool) *NlpBertTokenizationConfigBuilder {
	rb.v.WithSpecialTokens = &withspecialtokens
	return rb
}
