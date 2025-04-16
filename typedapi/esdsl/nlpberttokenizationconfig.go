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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

type _nlpBertTokenizationConfig struct {
	v *types.NlpBertTokenizationConfig
}

// Indicates MPNET tokenization and its options
func NewNlpBertTokenizationConfig() *_nlpBertTokenizationConfig {

	return &_nlpBertTokenizationConfig{v: types.NewNlpBertTokenizationConfig()}

}

func (s *_nlpBertTokenizationConfig) DoLowerCase(dolowercase bool) *_nlpBertTokenizationConfig {

	s.v.DoLowerCase = &dolowercase

	return s
}

func (s *_nlpBertTokenizationConfig) MaxSequenceLength(maxsequencelength int) *_nlpBertTokenizationConfig {

	s.v.MaxSequenceLength = &maxsequencelength

	return s
}

func (s *_nlpBertTokenizationConfig) Span(span int) *_nlpBertTokenizationConfig {

	s.v.Span = &span

	return s
}

func (s *_nlpBertTokenizationConfig) Truncate(truncate tokenizationtruncate.TokenizationTruncate) *_nlpBertTokenizationConfig {

	s.v.Truncate = &truncate
	return s
}

func (s *_nlpBertTokenizationConfig) WithSpecialTokens(withspecialtokens bool) *_nlpBertTokenizationConfig {

	s.v.WithSpecialTokens = &withspecialtokens

	return s
}

func (s *_nlpBertTokenizationConfig) TokenizationConfigContainerCaster() *types.TokenizationConfigContainer {
	container := types.NewTokenizationConfigContainer()

	container.Mpnet = s.v

	return container
}

func (s *_nlpBertTokenizationConfig) NlpBertTokenizationConfigCaster() *types.NlpBertTokenizationConfig {
	return s.v
}
