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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorsimilarity"
)

type _denseVectorQuery struct {
	v *types.DenseVectorQuery
}

// Scores every document that has a value for a `dense_vector` field, with the
// exact similarity between its vector and a query vector. Unlike the `knn`
// query, it doesn't use an index structure, or a cutoff based on `k` or
// `num_candidates`.
func NewDenseVectorQuery() *_denseVectorQuery {

	return &_denseVectorQuery{v: types.NewDenseVectorQuery()}

}

func (s *_denseVectorQuery) Field(field string) *_denseVectorQuery {

	s.v.Field = field

	return s
}

func (s *_denseVectorQuery) Quantized(quantized bool) *_denseVectorQuery {

	s.v.Quantized = &quantized

	return s
}

func (s *_denseVectorQuery) QueryVector(queryvector types.QueryVectorVariant) *_denseVectorQuery {

	s.v.QueryVector = *queryvector.QueryVectorCaster()

	return s
}

func (s *_denseVectorQuery) QueryVectorBuilder(queryvectorbuilder types.QueryVectorBuilderVariant) *_denseVectorQuery {

	s.v.QueryVectorBuilder = queryvectorbuilder.QueryVectorBuilderCaster()

	return s
}

func (s *_denseVectorQuery) SimilarityFunction(similarityfunction densevectorsimilarity.DenseVectorSimilarity) *_denseVectorQuery {

	s.v.SimilarityFunction = &similarityfunction
	return s
}

func (s *_denseVectorQuery) Boost(boost float32) *_denseVectorQuery {

	s.v.Boost = &boost

	return s
}

func (s *_denseVectorQuery) QueryName_(queryname_ string) *_denseVectorQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_denseVectorQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.DenseVector = s.v

	return container
}

func (s *_denseVectorQuery) DenseVectorQueryCaster() *types.DenseVectorQuery {
	return s.v
}
