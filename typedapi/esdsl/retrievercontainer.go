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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _retrieverContainer struct {
	v *types.RetrieverContainer
}

func NewRetrieverContainer() *_retrieverContainer {
	return &_retrieverContainer{v: types.NewRetrieverContainer()}
}

// AdditionalRetrieverContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_retrieverContainer) AdditionalRetrieverContainerProperty(key string, value json.RawMessage) *_retrieverContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalRetrieverContainerProperty = tmp
	return s
}

func (s *_retrieverContainer) Knn(knn types.KnnRetrieverVariant) *_retrieverContainer {

	s.v.Knn = knn.KnnRetrieverCaster()

	return s
}

func (s *_retrieverContainer) Rrf(rrf types.RRFRetrieverVariant) *_retrieverContainer {

	s.v.Rrf = rrf.RRFRetrieverCaster()

	return s
}

func (s *_retrieverContainer) Rule(rule types.RuleRetrieverVariant) *_retrieverContainer {

	s.v.Rule = rule.RuleRetrieverCaster()

	return s
}

func (s *_retrieverContainer) Standard(standard types.StandardRetrieverVariant) *_retrieverContainer {

	s.v.Standard = standard.StandardRetrieverCaster()

	return s
}

func (s *_retrieverContainer) TextSimilarityReranker(textsimilarityreranker types.TextSimilarityRerankerVariant) *_retrieverContainer {

	s.v.TextSimilarityReranker = textsimilarityreranker.TextSimilarityRerankerCaster()

	return s
}

func (s *_retrieverContainer) RetrieverContainerCaster() *types.RetrieverContainer {
	return s.v
}
