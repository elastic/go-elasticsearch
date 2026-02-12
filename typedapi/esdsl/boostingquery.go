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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _boostingQuery struct {
	v *types.BoostingQuery
}

// Returns documents matching a `positive` query while reducing the relevance
// score of documents that also match a `negative` query.
func NewBoostingQuery(negative types.QueryVariant, negativeboost types.Float64, positive types.QueryVariant) *_boostingQuery {

	tmp := &_boostingQuery{v: types.NewBoostingQuery()}

	tmp.Negative(negative)

	tmp.NegativeBoost(negativeboost)

	tmp.Positive(positive)

	return tmp

}

func (s *_boostingQuery) Negative(negative types.QueryVariant) *_boostingQuery {

	s.v.Negative = *negative.QueryCaster()

	return s
}

func (s *_boostingQuery) NegativeBoost(negativeboost types.Float64) *_boostingQuery {

	s.v.NegativeBoost = negativeboost

	return s
}

func (s *_boostingQuery) Positive(positive types.QueryVariant) *_boostingQuery {

	s.v.Positive = *positive.QueryCaster()

	return s
}

func (s *_boostingQuery) Boost(boost float32) *_boostingQuery {

	s.v.Boost = &boost

	return s
}

func (s *_boostingQuery) QueryName_(queryname_ string) *_boostingQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_boostingQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Boosting = s.v

	return container
}

func (s *_boostingQuery) BoostingQueryCaster() *types.BoostingQuery {
	return s.v
}
