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

type _spanNearQuery struct {
	v *types.SpanNearQuery
}

// Matches spans which are near one another.
// You can specify `slop`, the maximum number of intervening unmatched
// positions, as well as whether matches are required to be in-order.
func NewSpanNearQuery() *_spanNearQuery {

	return &_spanNearQuery{v: types.NewSpanNearQuery()}

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_spanNearQuery) Boost(boost float32) *_spanNearQuery {

	s.v.Boost = &boost

	return s
}

// Array of one or more other span type queries.
func (s *_spanNearQuery) Clauses(clauses ...types.SpanQueryVariant) *_spanNearQuery {

	for _, v := range clauses {

		s.v.Clauses = append(s.v.Clauses, *v.SpanQueryCaster())

	}
	return s
}

// Controls whether matches are required to be in-order.
func (s *_spanNearQuery) InOrder(inorder bool) *_spanNearQuery {

	s.v.InOrder = &inorder

	return s
}

func (s *_spanNearQuery) QueryName_(queryname_ string) *_spanNearQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Controls the maximum number of intervening unmatched positions permitted.
func (s *_spanNearQuery) Slop(slop int) *_spanNearQuery {

	s.v.Slop = &slop

	return s
}

func (s *_spanNearQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.SpanNear = s.v

	return container
}

func (s *_spanNearQuery) SpanQueryCaster() *types.SpanQuery {
	container := types.NewSpanQuery()

	container.SpanNear = s.v

	return container
}

func (s *_spanNearQuery) SpanNearQueryCaster() *types.SpanNearQuery {
	return s.v
}
