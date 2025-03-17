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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _constantScoreQuery struct {
	v *types.ConstantScoreQuery
}

// Wraps a filter query and returns every matching document with a relevance
// score equal to the `boost` parameter value.
func NewConstantScoreQuery(filter types.QueryVariant) *_constantScoreQuery {

	tmp := &_constantScoreQuery{v: types.NewConstantScoreQuery()}

	tmp.Filter(filter)

	return tmp

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_constantScoreQuery) Boost(boost float32) *_constantScoreQuery {

	s.v.Boost = &boost

	return s
}

// Filter query you wish to run. Any returned documents must match this query.
// Filter queries do not calculate relevance scores.
// To speed up performance, Elasticsearch automatically caches frequently used
// filter queries.
func (s *_constantScoreQuery) Filter(filter types.QueryVariant) *_constantScoreQuery {

	s.v.Filter = *filter.QueryCaster()

	return s
}

func (s *_constantScoreQuery) QueryName_(queryname_ string) *_constantScoreQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_constantScoreQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.ConstantScore = s.v

	return container
}

func (s *_constantScoreQuery) ConstantScoreQueryCaster() *types.ConstantScoreQuery {
	return s.v
}
