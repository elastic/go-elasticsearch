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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _sparseVectorQuery struct {
	v *types.SparseVectorQuery
}

func NewSparseVectorQuery() *_sparseVectorQuery {
	return &_sparseVectorQuery{v: types.NewSparseVectorQuery()}
}

// AdditionalSparseVectorQueryProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_sparseVectorQuery) AdditionalSparseVectorQueryProperty(key string, value json.RawMessage) *_sparseVectorQuery {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalSparseVectorQueryProperty = tmp
	return s
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_sparseVectorQuery) Boost(boost float32) *_sparseVectorQuery {

	s.v.Boost = &boost

	return s
}

// The name of the field that contains the token-weight pairs to be searched
// against.
// This field must be a mapped sparse_vector field.
func (s *_sparseVectorQuery) Field(field string) *_sparseVectorQuery {

	s.v.Field = field

	return s
}

// The inference ID to use to convert the query text into token-weight pairs.
// It must be the same inference ID that was used to create the tokens from the
// input text.
// Only one of inference_id and query_vector is allowed.
// If inference_id is specified, query must also be specified.
// Only one of inference_id or query_vector may be supplied in a request.
func (s *_sparseVectorQuery) InferenceId(id string) *_sparseVectorQuery {

	s.v.InferenceId = &id

	return s
}

// Whether to perform pruning, omitting the non-significant tokens from the
// query to improve query performance.
// If prune is true but the pruning_config is not specified, pruning will occur
// but default values will be used.
// Default: false
func (s *_sparseVectorQuery) Prune(prune bool) *_sparseVectorQuery {

	s.v.Prune = &prune

	return s
}

// Optional pruning configuration.
// If enabled, this will omit non-significant tokens from the query in order to
// improve query performance.
// This is only used if prune is set to true.
// If prune is set to true but pruning_config is not specified, default values
// will be used.
func (s *_sparseVectorQuery) PruningConfig(pruningconfig types.TokenPruningConfigVariant) *_sparseVectorQuery {

	s.v.PruningConfig = pruningconfig.TokenPruningConfigCaster()

	return s
}

// The query text you want to use for search.
// If inference_id is specified, query must also be specified.
func (s *_sparseVectorQuery) Query(query string) *_sparseVectorQuery {

	s.v.Query = &query

	return s
}

func (s *_sparseVectorQuery) QueryName_(queryname_ string) *_sparseVectorQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Dictionary of precomputed sparse vectors and their associated weights.
// Only one of inference_id or query_vector may be supplied in a request.
func (s *_sparseVectorQuery) QueryVector(queryvector map[string]float32) *_sparseVectorQuery {

	s.v.QueryVector = queryvector
	return s
}

func (s *_sparseVectorQuery) AddQueryVector(key string, value float32) *_sparseVectorQuery {

	var tmp map[string]float32
	if s.v.QueryVector == nil {
		s.v.QueryVector = make(map[string]float32)
	} else {
		tmp = s.v.QueryVector
	}

	tmp[key] = value

	s.v.QueryVector = tmp
	return s
}

func (s *_sparseVectorQuery) SparseVectorQueryCaster() *types.SparseVectorQuery {
	return s.v
}
