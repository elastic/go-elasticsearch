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

type _knnRetriever struct {
	v *types.KnnRetriever
}

// A retriever that replaces the functionality  of a knn search.
func NewKnnRetriever(field string, k int, numcandidates int) *_knnRetriever {

	tmp := &_knnRetriever{v: types.NewKnnRetriever()}

	tmp.Field(field)

	tmp.K(k)

	tmp.NumCandidates(numcandidates)

	return tmp

}

func (s *_knnRetriever) Field(field string) *_knnRetriever {

	s.v.Field = field

	return s
}

func (s *_knnRetriever) K(k int) *_knnRetriever {

	s.v.K = k

	return s
}

func (s *_knnRetriever) NumCandidates(numcandidates int) *_knnRetriever {

	s.v.NumCandidates = numcandidates

	return s
}

func (s *_knnRetriever) QueryVector(queryvectors ...float32) *_knnRetriever {

	s.v.QueryVector = queryvectors

	return s
}

func (s *_knnRetriever) QueryVectorBuilder(queryvectorbuilder types.QueryVectorBuilderVariant) *_knnRetriever {

	s.v.QueryVectorBuilder = queryvectorbuilder.QueryVectorBuilderCaster()

	return s
}

func (s *_knnRetriever) RescoreVector(rescorevector types.RescoreVectorVariant) *_knnRetriever {

	s.v.RescoreVector = rescorevector.RescoreVectorCaster()

	return s
}

func (s *_knnRetriever) Similarity(similarity float32) *_knnRetriever {

	s.v.Similarity = &similarity

	return s
}

func (s *_knnRetriever) VisitPercentage(visitpercentage float32) *_knnRetriever {

	s.v.VisitPercentage = &visitpercentage

	return s
}

func (s *_knnRetriever) Filter(filters ...types.QueryVariant) *_knnRetriever {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_knnRetriever) MinScore(minscore float32) *_knnRetriever {

	s.v.MinScore = &minscore

	return s
}

func (s *_knnRetriever) Name_(name_ string) *_knnRetriever {

	s.v.Name_ = &name_

	return s
}

func (s *_knnRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	container := types.NewRetrieverContainer()

	container.Knn = s.v

	return container
}

func (s *_knnRetriever) KnnRetrieverCaster() *types.KnnRetriever {
	return s.v
}
