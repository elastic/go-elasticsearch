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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scoremode"
)

type _rescoreQuery struct {
	v *types.RescoreQuery
}

func NewRescoreQuery(query types.QueryVariant) *_rescoreQuery {

	tmp := &_rescoreQuery{v: types.NewRescoreQuery()}

	tmp.Query(query)

	return tmp

}

// The query to use for rescoring.
// This query is only run on the Top-K results returned by the `query` and
// `post_filter` phases.
func (s *_rescoreQuery) Query(query types.QueryVariant) *_rescoreQuery {

	s.v.Query = *query.QueryCaster()

	return s
}

// Relative importance of the original query versus the rescore query.
func (s *_rescoreQuery) QueryWeight(queryweight types.Float64) *_rescoreQuery {

	s.v.QueryWeight = &queryweight

	return s
}

// Relative importance of the rescore query versus the original query.
func (s *_rescoreQuery) RescoreQueryWeight(rescorequeryweight types.Float64) *_rescoreQuery {

	s.v.RescoreQueryWeight = &rescorequeryweight

	return s
}

// Determines how scores are combined.
func (s *_rescoreQuery) ScoreMode(scoremode scoremode.ScoreMode) *_rescoreQuery {

	s.v.ScoreMode = &scoremode
	return s
}

func (s *_rescoreQuery) RescoreCaster() *types.Rescore {
	container := types.NewRescore()

	container.Query = s.v

	return container
}

func (s *_rescoreQuery) RescoreQueryCaster() *types.RescoreQuery {
	return s.v
}
