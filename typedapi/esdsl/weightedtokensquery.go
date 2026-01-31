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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _weightedTokensQuery struct {
	k string
	v *types.WeightedTokensQuery
}

// Supports returning text_expansion query results by sending in precomputed
// tokens with the query.
func NewWeightedTokensQuery(key string) *_weightedTokensQuery {
	return &_weightedTokensQuery{
		k: key,
		v: types.NewWeightedTokensQuery(),
	}
}

func (s *_weightedTokensQuery) PruningConfig(pruningconfig types.TokenPruningConfigVariant) *_weightedTokensQuery {

	s.v.PruningConfig = pruningconfig.TokenPruningConfigCaster()

	return s
}

func (s *_weightedTokensQuery) Tokens(tokens []map[string]float32) *_weightedTokensQuery {

	s.v.Tokens = tokens

	return s
}

func (s *_weightedTokensQuery) Boost(boost float32) *_weightedTokensQuery {

	s.v.Boost = &boost

	return s
}

func (s *_weightedTokensQuery) QueryName_(queryname_ string) *_weightedTokensQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_weightedTokensQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.WeightedTokens = map[string]types.WeightedTokensQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleWeightedTokensQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleWeightedTokensQuery() *_weightedTokensQuery {
	return &_weightedTokensQuery{
		k: "",
		v: types.NewWeightedTokensQuery(),
	}
}

func (s *_weightedTokensQuery) WeightedTokensQueryCaster() *types.WeightedTokensQuery {
	return s.v.WeightedTokensQueryCaster()
}
