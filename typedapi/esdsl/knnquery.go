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

type _knnQuery struct {
	v *types.KnnQuery
}

// Finds the k nearest vectors to a query vector, as measured by a similarity
// metric. knn query finds nearest vectors through approximate search on indexed
// dense_vectors.
func NewKnnQuery() *_knnQuery {

	return &_knnQuery{v: types.NewKnnQuery()}

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_knnQuery) Boost(boost float32) *_knnQuery {

	s.v.Boost = &boost

	return s
}

// The name of the vector field to search against
func (s *_knnQuery) Field(field string) *_knnQuery {

	s.v.Field = field

	return s
}

// Filters for the kNN search query
func (s *_knnQuery) Filter(filters ...types.QueryVariant) *_knnQuery {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

// The final number of nearest neighbors to return as top hits
func (s *_knnQuery) K(k int) *_knnQuery {

	s.v.K = &k

	return s
}

// The number of nearest neighbor candidates to consider per shard
func (s *_knnQuery) NumCandidates(numcandidates int) *_knnQuery {

	s.v.NumCandidates = &numcandidates

	return s
}

func (s *_knnQuery) QueryName_(queryname_ string) *_knnQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// The query vector
func (s *_knnQuery) QueryVector(queryvectors ...float32) *_knnQuery {

	s.v.QueryVector = queryvectors

	return s
}

// The query vector builder. You must provide a query_vector_builder or
// query_vector, but not both.
func (s *_knnQuery) QueryVectorBuilder(queryvectorbuilder types.QueryVectorBuilderVariant) *_knnQuery {

	s.v.QueryVectorBuilder = queryvectorbuilder.QueryVectorBuilderCaster()

	return s
}

// Apply oversampling and rescoring to quantized vectors *
func (s *_knnQuery) RescoreVector(rescorevector types.RescoreVectorVariant) *_knnQuery {

	s.v.RescoreVector = rescorevector.RescoreVectorCaster()

	return s
}

// The minimum similarity for a vector to be considered a match
func (s *_knnQuery) Similarity(similarity float32) *_knnQuery {

	s.v.Similarity = &similarity

	return s
}

func (s *_knnQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Knn = s.v

	return container
}

func (s *_knnQuery) KnnQueryCaster() *types.KnnQuery {
	return s.v
}
