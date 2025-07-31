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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

type _xlmRobertaTokenizationConfig struct {
	v *types.XlmRobertaTokenizationConfig
}

func NewXlmRobertaTokenizationConfig() *_xlmRobertaTokenizationConfig {

	return &_xlmRobertaTokenizationConfig{v: types.NewXlmRobertaTokenizationConfig()}

}

func (s *_xlmRobertaTokenizationConfig) DoLowerCase(dolowercase bool) *_xlmRobertaTokenizationConfig {

	s.v.DoLowerCase = &dolowercase

	return s
}

func (s *_xlmRobertaTokenizationConfig) MaxSequenceLength(maxsequencelength int) *_xlmRobertaTokenizationConfig {

	s.v.MaxSequenceLength = &maxsequencelength

	return s
}

func (s *_xlmRobertaTokenizationConfig) Span(span int) *_xlmRobertaTokenizationConfig {

	s.v.Span = &span

	return s
}

func (s *_xlmRobertaTokenizationConfig) Truncate(truncate tokenizationtruncate.TokenizationTruncate) *_xlmRobertaTokenizationConfig {

	s.v.Truncate = &truncate
	return s
}

func (s *_xlmRobertaTokenizationConfig) WithSpecialTokens(withspecialtokens bool) *_xlmRobertaTokenizationConfig {

	s.v.WithSpecialTokens = &withspecialtokens

	return s
}

func (s *_xlmRobertaTokenizationConfig) TokenizationConfigContainerCaster() *types.TokenizationConfigContainer {
	container := types.NewTokenizationConfigContainer()

	container.XlmRoberta = s.v

	return container
}

func (s *_xlmRobertaTokenizationConfig) XlmRobertaTokenizationConfigCaster() *types.XlmRobertaTokenizationConfig {
	return s.v
}
