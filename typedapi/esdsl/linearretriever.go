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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scorenormalizer"
)

type _linearRetriever struct {
	v *types.LinearRetriever
}

// A retriever that supports the combination of different retrievers through a
// weighted linear combination.
func NewLinearRetriever() *_linearRetriever {

	return &_linearRetriever{v: types.NewLinearRetriever()}

}

func (s *_linearRetriever) Fields(fields ...string) *_linearRetriever {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

func (s *_linearRetriever) Normalizer(normalizer scorenormalizer.ScoreNormalizer) *_linearRetriever {

	s.v.Normalizer = &normalizer
	return s
}

func (s *_linearRetriever) Query(query string) *_linearRetriever {

	s.v.Query = &query

	return s
}

func (s *_linearRetriever) RankWindowSize(rankwindowsize int) *_linearRetriever {

	s.v.RankWindowSize = &rankwindowsize

	return s
}

func (s *_linearRetriever) Retrievers(retrievers ...types.InnerRetrieverVariant) *_linearRetriever {

	for _, v := range retrievers {

		s.v.Retrievers = append(s.v.Retrievers, *v.InnerRetrieverCaster())

	}
	return s
}

func (s *_linearRetriever) Filter(filters ...types.QueryVariant) *_linearRetriever {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_linearRetriever) MinScore(minscore float32) *_linearRetriever {

	s.v.MinScore = &minscore

	return s
}

func (s *_linearRetriever) Name_(name_ string) *_linearRetriever {

	s.v.Name_ = &name_

	return s
}

func (s *_linearRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	container := types.NewRetrieverContainer()

	container.Linear = s.v

	return container
}

func (s *_linearRetriever) LinearRetrieverCaster() *types.LinearRetriever {
	return s.v
}
