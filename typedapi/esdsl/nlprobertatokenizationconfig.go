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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

type _nlpRobertaTokenizationConfig struct {
	v *types.NlpRobertaTokenizationConfig
}

// Indicates RoBERTa tokenization and its options
func NewNlpRobertaTokenizationConfig() *_nlpRobertaTokenizationConfig {

	return &_nlpRobertaTokenizationConfig{v: types.NewNlpRobertaTokenizationConfig()}

}

func (s *_nlpRobertaTokenizationConfig) AddPrefixSpace(addprefixspace bool) *_nlpRobertaTokenizationConfig {

	s.v.AddPrefixSpace = &addprefixspace

	return s
}

func (s *_nlpRobertaTokenizationConfig) DoLowerCase(dolowercase bool) *_nlpRobertaTokenizationConfig {

	s.v.DoLowerCase = &dolowercase

	return s
}

func (s *_nlpRobertaTokenizationConfig) MaxSequenceLength(maxsequencelength int) *_nlpRobertaTokenizationConfig {

	s.v.MaxSequenceLength = &maxsequencelength

	return s
}

func (s *_nlpRobertaTokenizationConfig) Span(span int) *_nlpRobertaTokenizationConfig {

	s.v.Span = &span

	return s
}

func (s *_nlpRobertaTokenizationConfig) Truncate(truncate tokenizationtruncate.TokenizationTruncate) *_nlpRobertaTokenizationConfig {

	s.v.Truncate = &truncate
	return s
}

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
