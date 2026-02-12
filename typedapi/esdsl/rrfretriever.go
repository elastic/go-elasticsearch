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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rRFRetriever struct {
	v *types.RRFRetriever
}

// A retriever that produces top documents from reciprocal rank fusion (RRF).
func NewRRFRetriever() *_rRFRetriever {

	return &_rRFRetriever{v: types.NewRRFRetriever()}

}

func (s *_rRFRetriever) Fields(fields ...string) *_rRFRetriever {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

func (s *_rRFRetriever) Query(query string) *_rRFRetriever {

	s.v.Query = &query

	return s
}

func (s *_rRFRetriever) RankConstant(rankconstant int) *_rRFRetriever {

	s.v.RankConstant = &rankconstant

	return s
}

func (s *_rRFRetriever) RankWindowSize(rankwindowsize int) *_rRFRetriever {

	s.v.RankWindowSize = &rankwindowsize

	return s
}

func (s *_rRFRetriever) Retrievers(retrievers ...types.RRFRetrieverEntryVariant) *_rRFRetriever {

	for _, v := range retrievers {

		s.v.Retrievers = append(s.v.Retrievers, *v.RRFRetrieverEntryCaster())

	}
	return s
}

func (s *_rRFRetriever) Filter(filters ...types.QueryVariant) *_rRFRetriever {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_rRFRetriever) MinScore(minscore float32) *_rRFRetriever {

	s.v.MinScore = &minscore

	return s
}

func (s *_rRFRetriever) Name_(name_ string) *_rRFRetriever {

	s.v.Name_ = &name_

	return s
}

func (s *_rRFRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	container := types.NewRetrieverContainer()

	container.Rrf = s.v

	return container
}

func (s *_rRFRetriever) RRFRetrieverCaster() *types.RRFRetriever {
	return s.v
}
