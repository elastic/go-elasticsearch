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

type _disMaxQuery struct {
	v *types.DisMaxQuery
}

// Returns documents matching one or more wrapped queries, called query clauses
// or clauses.
// If a returned document matches multiple query clauses, the `dis_max` query
// assigns the document the highest relevance score from any matching clause,
// plus a tie breaking increment for any additional matching subqueries.
func NewDisMaxQuery() *_disMaxQuery {

	return &_disMaxQuery{v: types.NewDisMaxQuery()}

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_disMaxQuery) Boost(boost float32) *_disMaxQuery {

	s.v.Boost = &boost

	return s
}

// One or more query clauses.
// Returned documents must match one or more of these queries.
// If a document matches multiple queries, Elasticsearch uses the highest
// relevance score.
func (s *_disMaxQuery) Queries(queries ...types.QueryVariant) *_disMaxQuery {

	for _, v := range queries {

		s.v.Queries = append(s.v.Queries, *v.QueryCaster())

	}
	return s
}

func (s *_disMaxQuery) QueryName_(queryname_ string) *_disMaxQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Floating point number between 0 and 1.0 used to increase the relevance scores
// of documents matching multiple query clauses.
func (s *_disMaxQuery) TieBreaker(tiebreaker types.Float64) *_disMaxQuery {

	s.v.TieBreaker = &tiebreaker

	return s
}

func (s *_disMaxQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.DisMax = s.v

	return container
}

func (s *_disMaxQuery) DisMaxQueryCaster() *types.DisMaxQuery {
	return s.v
}
