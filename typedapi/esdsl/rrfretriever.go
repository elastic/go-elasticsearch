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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _rRFRetriever struct {
	v *types.RRFRetriever
}

// A retriever that produces top documents from reciprocal rank fusion (RRF).
func NewRRFRetriever() *_rRFRetriever {

	return &_rRFRetriever{v: types.NewRRFRetriever()}

}

// Query to filter the documents that can match.
func (s *_rRFRetriever) Filter(filters ...types.QueryVariant) *_rRFRetriever {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

// Minimum _score for matching documents. Documents with a lower _score are not
// included in the top documents.
func (s *_rRFRetriever) MinScore(minscore float32) *_rRFRetriever {

	s.v.MinScore = &minscore

	return s
}

// This value determines how much influence documents in individual result sets
// per query have over the final ranked result set.
func (s *_rRFRetriever) RankConstant(rankconstant int) *_rRFRetriever {

	s.v.RankConstant = &rankconstant

	return s
}

// This value determines the size of the individual result sets per query.
func (s *_rRFRetriever) RankWindowSize(rankwindowsize int) *_rRFRetriever {

	s.v.RankWindowSize = &rankwindowsize

	return s
}

// A list of child retrievers to specify which sets of returned top documents
// will have the RRF formula applied to them.
func (s *_rRFRetriever) Retrievers(retrievers ...types.RetrieverContainerVariant) *_rRFRetriever {

	for _, v := range retrievers {

		s.v.Retrievers = append(s.v.Retrievers, *v.RetrieverContainerCaster())

	}
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
