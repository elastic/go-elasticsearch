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

type _knnSearch struct {
	v *types.KnnSearch
}

func NewKnnSearch() *_knnSearch {

	return &_knnSearch{v: types.NewKnnSearch()}

}

// Boost value to apply to kNN scores
func (s *_knnSearch) Boost(boost float32) *_knnSearch {

	s.v.Boost = &boost

	return s
}

// The name of the vector field to search against
func (s *_knnSearch) Field(field string) *_knnSearch {

	s.v.Field = field

	return s
}

// Filters for the kNN search query
func (s *_knnSearch) Filter(filters ...types.QueryVariant) *_knnSearch {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

// If defined, each search hit will contain inner hits.
func (s *_knnSearch) InnerHits(innerhits types.InnerHitsVariant) *_knnSearch {

	s.v.InnerHits = innerhits.InnerHitsCaster()

	return s
}

// The final number of nearest neighbors to return as top hits
func (s *_knnSearch) K(k int) *_knnSearch {

	s.v.K = &k

	return s
}

// The number of nearest neighbor candidates to consider per shard
func (s *_knnSearch) NumCandidates(numcandidates int) *_knnSearch {

	s.v.NumCandidates = &numcandidates

	return s
}

// The query vector
func (s *_knnSearch) QueryVector(queryvectors ...float32) *_knnSearch {

	s.v.QueryVector = queryvectors

	return s
}

// The query vector builder. You must provide a query_vector_builder or
// query_vector, but not both.
func (s *_knnSearch) QueryVectorBuilder(queryvectorbuilder types.QueryVectorBuilderVariant) *_knnSearch {

	s.v.QueryVectorBuilder = queryvectorbuilder.QueryVectorBuilderCaster()

	return s
}

// Apply oversampling and rescoring to quantized vectors *
func (s *_knnSearch) RescoreVector(rescorevector types.RescoreVectorVariant) *_knnSearch {

	s.v.RescoreVector = rescorevector.RescoreVectorCaster()

	return s
}

// The minimum similarity for a vector to be considered a match
func (s *_knnSearch) Similarity(similarity float32) *_knnSearch {

	s.v.Similarity = &similarity

	return s
}

func (s *_knnSearch) KnnSearchCaster() *types.KnnSearch {
	return s.v
}
