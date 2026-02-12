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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _knnQuery struct {
	v *types.KnnQuery
}

// Finds the k nearest vectors to a query vector, as measured by a similarity
// metric. knn query finds nearest vectors through approximate search on indexed
// dense_vectors.
func NewKnnQuery() *_knnQuery {

	return &_knnQuery{v: types.NewKnnQuery()}

}

func (s *_knnQuery) Field(field string) *_knnQuery {

	s.v.Field = field

	return s
}

func (s *_knnQuery) Filter(filters ...types.QueryVariant) *_knnQuery {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_knnQuery) K(k int) *_knnQuery {

	s.v.K = &k

	return s
}

func (s *_knnQuery) NumCandidates(numcandidates int) *_knnQuery {

	s.v.NumCandidates = &numcandidates

	return s
}

func (s *_knnQuery) QueryVector(queryvectors ...float32) *_knnQuery {

	s.v.QueryVector = queryvectors

	return s
}

func (s *_knnQuery) QueryVectorBuilder(queryvectorbuilder types.QueryVectorBuilderVariant) *_knnQuery {

	s.v.QueryVectorBuilder = queryvectorbuilder.QueryVectorBuilderCaster()

	return s
}

func (s *_knnQuery) RescoreVector(rescorevector types.RescoreVectorVariant) *_knnQuery {

	s.v.RescoreVector = rescorevector.RescoreVectorCaster()

	return s
}

func (s *_knnQuery) Similarity(similarity float32) *_knnQuery {

	s.v.Similarity = &similarity

	return s
}

func (s *_knnQuery) VisitPercentage(visitpercentage float32) *_knnQuery {

	s.v.VisitPercentage = &visitpercentage

	return s
}

func (s *_knnQuery) Boost(boost float32) *_knnQuery {

	s.v.Boost = &boost

	return s
}

func (s *_knnQuery) QueryName_(queryname_ string) *_knnQuery {

	s.v.QueryName_ = &queryname_

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
