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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanWithinQuery struct {
	v *types.SpanWithinQuery
}

// Returns matches which are enclosed inside another span query.
func NewSpanWithinQuery(big types.SpanQueryVariant, little types.SpanQueryVariant) *_spanWithinQuery {

	tmp := &_spanWithinQuery{v: types.NewSpanWithinQuery()}

	tmp.Big(big)

	tmp.Little(little)

	return tmp

}

func (s *_spanWithinQuery) Big(big types.SpanQueryVariant) *_spanWithinQuery {

	s.v.Big = *big.SpanQueryCaster()

	return s
}

func (s *_spanWithinQuery) Boost(boost float32) *_spanWithinQuery {

	s.v.Boost = &boost

	return s
}

func (s *_spanWithinQuery) Little(little types.SpanQueryVariant) *_spanWithinQuery {

	s.v.Little = *little.SpanQueryCaster()

	return s
}

func (s *_spanWithinQuery) QueryName_(queryname_ string) *_spanWithinQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_spanWithinQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.SpanWithin = s.v

	return container
}

func (s *_spanWithinQuery) SpanQueryCaster() *types.SpanQuery {
	container := types.NewSpanQuery()

	container.SpanWithin = s.v

	return container
}

func (s *_spanWithinQuery) SpanWithinQueryCaster() *types.SpanWithinQuery {
	return s.v
}
