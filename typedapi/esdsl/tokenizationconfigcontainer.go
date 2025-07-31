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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _tokenizationConfigContainer struct {
	v *types.TokenizationConfigContainer
}

func NewTokenizationConfigContainer() *_tokenizationConfigContainer {
	return &_tokenizationConfigContainer{v: types.NewTokenizationConfigContainer()}
}

// AdditionalTokenizationConfigContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_tokenizationConfigContainer) AdditionalTokenizationConfigContainerProperty(key string, value json.RawMessage) *_tokenizationConfigContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalTokenizationConfigContainerProperty = tmp
	return s
}

func (s *_tokenizationConfigContainer) Bert(bert types.NlpBertTokenizationConfigVariant) *_tokenizationConfigContainer {

	s.v.Bert = bert.NlpBertTokenizationConfigCaster()

	return s
}

func (s *_tokenizationConfigContainer) BertJa(bertja types.NlpBertTokenizationConfigVariant) *_tokenizationConfigContainer {

	s.v.BertJa = bertja.NlpBertTokenizationConfigCaster()

	return s
}

func (s *_tokenizationConfigContainer) Mpnet(mpnet types.NlpBertTokenizationConfigVariant) *_tokenizationConfigContainer {

	s.v.Mpnet = mpnet.NlpBertTokenizationConfigCaster()

	return s
}

func (s *_tokenizationConfigContainer) Roberta(roberta types.NlpRobertaTokenizationConfigVariant) *_tokenizationConfigContainer {

	s.v.Roberta = roberta.NlpRobertaTokenizationConfigCaster()

	return s
}

func (s *_tokenizationConfigContainer) XlmRoberta(xlmroberta types.XlmRobertaTokenizationConfigVariant) *_tokenizationConfigContainer {

	s.v.XlmRoberta = xlmroberta.XlmRobertaTokenizationConfigCaster()

	return s
}

func (s *_tokenizationConfigContainer) TokenizationConfigContainerCaster() *types.TokenizationConfigContainer {
	return s.v
}
