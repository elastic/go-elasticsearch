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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scorenormalizer"
)

type _innerRetriever struct {
	v *types.InnerRetriever
}

func NewInnerRetriever(normalizer scorenormalizer.ScoreNormalizer, retriever types.RetrieverContainerVariant, weight float32) *_innerRetriever {

	tmp := &_innerRetriever{v: types.NewInnerRetriever()}

	tmp.Normalizer(normalizer)

	tmp.Retriever(retriever)

	tmp.Weight(weight)

	return tmp

}

func (s *_innerRetriever) Normalizer(normalizer scorenormalizer.ScoreNormalizer) *_innerRetriever {

	s.v.Normalizer = normalizer
	return s
}

func (s *_innerRetriever) Retriever(retriever types.RetrieverContainerVariant) *_innerRetriever {

	s.v.Retriever = *retriever.RetrieverContainerCaster()

	return s
}

func (s *_innerRetriever) Weight(weight float32) *_innerRetriever {

	s.v.Weight = weight

	return s
}

func (s *_innerRetriever) InnerRetrieverCaster() *types.InnerRetriever {
	return s.v
}
