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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noridecompoundmode"
)

// NoriTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/tokenizers.ts#L80-L86
type NoriTokenizer struct {
	DecompoundMode      *noridecompoundmode.NoriDecompoundMode `json:"decompound_mode,omitempty"`
	DiscardPunctuation  *bool                                  `json:"discard_punctuation,omitempty"`
	Type                string                                 `json:"type,omitempty"`
	UserDictionary      *string                                `json:"user_dictionary,omitempty"`
	UserDictionaryRules []string                               `json:"user_dictionary_rules,omitempty"`
	Version             *VersionString                         `json:"version,omitempty"`
}

// NoriTokenizerBuilder holds NoriTokenizer struct and provides a builder API.
type NoriTokenizerBuilder struct {
	v *NoriTokenizer
}

// NewNoriTokenizer provides a builder for the NoriTokenizer struct.
func NewNoriTokenizerBuilder() *NoriTokenizerBuilder {
	r := NoriTokenizerBuilder{
		&NoriTokenizer{},
	}

	r.v.Type = "nori_tokenizer"

	return &r
}

// Build finalize the chain and returns the NoriTokenizer struct
func (rb *NoriTokenizerBuilder) Build() NoriTokenizer {
	return *rb.v
}

func (rb *NoriTokenizerBuilder) DecompoundMode(decompoundmode noridecompoundmode.NoriDecompoundMode) *NoriTokenizerBuilder {
	rb.v.DecompoundMode = &decompoundmode
	return rb
}

func (rb *NoriTokenizerBuilder) DiscardPunctuation(discardpunctuation bool) *NoriTokenizerBuilder {
	rb.v.DiscardPunctuation = &discardpunctuation
	return rb
}

func (rb *NoriTokenizerBuilder) UserDictionary(userdictionary string) *NoriTokenizerBuilder {
	rb.v.UserDictionary = &userdictionary
	return rb
}

func (rb *NoriTokenizerBuilder) UserDictionaryRules(user_dictionary_rules ...string) *NoriTokenizerBuilder {
	rb.v.UserDictionaryRules = user_dictionary_rules
	return rb
}

func (rb *NoriTokenizerBuilder) Version(version VersionString) *NoriTokenizerBuilder {
	rb.v.Version = &version
	return rb
}
