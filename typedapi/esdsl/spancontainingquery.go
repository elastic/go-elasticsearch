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

type _spanContainingQuery struct {
	v *types.SpanContainingQuery
}

// Accepts a list of span queries, but only returns those spans which also match
// a second span query.
func NewSpanContainingQuery(big types.SpanQueryVariant, little types.SpanQueryVariant) *_spanContainingQuery {

	tmp := &_spanContainingQuery{v: types.NewSpanContainingQuery()}

	tmp.Big(big)

	tmp.Little(little)

	return tmp

}

// Can be any span query.
// Matching spans from `big` that contain matches from `little` are returned.
func (s *_spanContainingQuery) Big(big types.SpanQueryVariant) *_spanContainingQuery {

	s.v.Big = *big.SpanQueryCaster()

	return s
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_spanContainingQuery) Boost(boost float32) *_spanContainingQuery {

	s.v.Boost = &boost

	return s
}

// Can be any span query.
// Matching spans from `big` that contain matches from `little` are returned.
func (s *_spanContainingQuery) Little(little types.SpanQueryVariant) *_spanContainingQuery {

	s.v.Little = *little.SpanQueryCaster()

	return s
}

func (s *_spanContainingQuery) QueryName_(queryname_ string) *_spanContainingQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_spanContainingQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.SpanContaining = s.v

	return container
}

func (s *_spanContainingQuery) SpanQueryCaster() *types.SpanQuery {
	container := types.NewSpanQuery()

	container.SpanContaining = s.v

	return container
}

func (s *_spanContainingQuery) SpanContainingQueryCaster() *types.SpanContainingQuery {
	return s.v
}
