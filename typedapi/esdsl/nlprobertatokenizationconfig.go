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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/tokenizationtruncate"
)

type _nlpRobertaTokenizationConfig struct {
	v *types.NlpRobertaTokenizationConfig
}

// Indicates RoBERTa tokenization and its options
func NewNlpRobertaTokenizationConfig() *_nlpRobertaTokenizationConfig {

	return &_nlpRobertaTokenizationConfig{v: types.NewNlpRobertaTokenizationConfig()}

}

// Should the tokenizer prefix input with a space character
func (s *_nlpRobertaTokenizationConfig) AddPrefixSpace(addprefixspace bool) *_nlpRobertaTokenizationConfig {

	s.v.AddPrefixSpace = &addprefixspace

	return s
}

// Should the tokenizer lower case the text
func (s *_nlpRobertaTokenizationConfig) DoLowerCase(dolowercase bool) *_nlpRobertaTokenizationConfig {

	s.v.DoLowerCase = &dolowercase

	return s
}

// Maximum input sequence length for the model
func (s *_nlpRobertaTokenizationConfig) MaxSequenceLength(maxsequencelength int) *_nlpRobertaTokenizationConfig {

	s.v.MaxSequenceLength = &maxsequencelength

	return s
}

// Tokenization spanning options. Special value of -1 indicates no spanning
// takes place
func (s *_nlpRobertaTokenizationConfig) Span(span int) *_nlpRobertaTokenizationConfig {

	s.v.Span = &span

	return s
}

// Should tokenization input be automatically truncated before sending to the
// model for inference
func (s *_nlpRobertaTokenizationConfig) Truncate(truncate tokenizationtruncate.TokenizationTruncate) *_nlpRobertaTokenizationConfig {

	s.v.Truncate = &truncate
	return s
}

// Is tokenization completed with special tokens
func (s *_nlpRobertaTokenizationConfig) WithSpecialTokens(withspecialtokens bool) *_nlpRobertaTokenizationConfig {

	s.v.WithSpecialTokens = &withspecialtokens

	return s
}

func (s *_nlpRobertaTokenizationConfig) TokenizationConfigContainerCaster() *types.TokenizationConfigContainer {
	container := types.NewTokenizationConfigContainer()

	container.Roberta = s.v

	return container
}

func (s *_nlpRobertaTokenizationConfig) NlpRobertaTokenizationConfigCaster() *types.NlpRobertaTokenizationConfig {
	return s.v
}
