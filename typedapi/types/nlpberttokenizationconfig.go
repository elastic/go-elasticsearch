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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/tokenizationtruncate"
)

// NlpBertTokenizationConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/inference.ts#L116-L143
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

// NewNlpBertTokenizationConfig returns a NlpBertTokenizationConfig.
func NewNlpBertTokenizationConfig() *NlpBertTokenizationConfig {
	r := &NlpBertTokenizationConfig{}

	return r
}
