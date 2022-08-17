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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/kuromojitokenizationmode"
)

// KuromojiTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/kuromoji-plugin.ts#L58-L67
type KuromojiTokenizer struct {
	DiscardCompoundToken *bool                                             `json:"discard_compound_token,omitempty"`
	DiscardPunctuation   *bool                                             `json:"discard_punctuation,omitempty"`
	Mode                 kuromojitokenizationmode.KuromojiTokenizationMode `json:"mode"`
	NbestCost            *int                                              `json:"nbest_cost,omitempty"`
	NbestExamples        *string                                           `json:"nbest_examples,omitempty"`
	Type                 string                                            `json:"type,omitempty"`
	UserDictionary       *string                                           `json:"user_dictionary,omitempty"`
	UserDictionaryRules  []string                                          `json:"user_dictionary_rules,omitempty"`
	Version              *VersionString                                    `json:"version,omitempty"`
}

// KuromojiTokenizerBuilder holds KuromojiTokenizer struct and provides a builder API.
type KuromojiTokenizerBuilder struct {
	v *KuromojiTokenizer
}

// NewKuromojiTokenizer provides a builder for the KuromojiTokenizer struct.
func NewKuromojiTokenizerBuilder() *KuromojiTokenizerBuilder {
	r := KuromojiTokenizerBuilder{
		&KuromojiTokenizer{},
	}

	r.v.Type = "kuromoji_tokenizer"

	return &r
}

// Build finalize the chain and returns the KuromojiTokenizer struct
func (rb *KuromojiTokenizerBuilder) Build() KuromojiTokenizer {
	return *rb.v
}

func (rb *KuromojiTokenizerBuilder) DiscardCompoundToken(discardcompoundtoken bool) *KuromojiTokenizerBuilder {
	rb.v.DiscardCompoundToken = &discardcompoundtoken
	return rb
}

func (rb *KuromojiTokenizerBuilder) DiscardPunctuation(discardpunctuation bool) *KuromojiTokenizerBuilder {
	rb.v.DiscardPunctuation = &discardpunctuation
	return rb
}

func (rb *KuromojiTokenizerBuilder) Mode(mode kuromojitokenizationmode.KuromojiTokenizationMode) *KuromojiTokenizerBuilder {
	rb.v.Mode = mode
	return rb
}

func (rb *KuromojiTokenizerBuilder) NbestCost(nbestcost int) *KuromojiTokenizerBuilder {
	rb.v.NbestCost = &nbestcost
	return rb
}

func (rb *KuromojiTokenizerBuilder) NbestExamples(nbestexamples string) *KuromojiTokenizerBuilder {
	rb.v.NbestExamples = &nbestexamples
	return rb
}

func (rb *KuromojiTokenizerBuilder) UserDictionary(userdictionary string) *KuromojiTokenizerBuilder {
	rb.v.UserDictionary = &userdictionary
	return rb
}

func (rb *KuromojiTokenizerBuilder) UserDictionaryRules(user_dictionary_rules ...string) *KuromojiTokenizerBuilder {
	rb.v.UserDictionaryRules = user_dictionary_rules
	return rb
}

func (rb *KuromojiTokenizerBuilder) Version(version VersionString) *KuromojiTokenizerBuilder {
	rb.v.Version = &version
	return rb
}
