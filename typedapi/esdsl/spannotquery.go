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

type _spanNotQuery struct {
	v *types.SpanNotQuery
}

// Removes matches which overlap with another span query or which are within x
// tokens before (controlled by the parameter `pre`) or y tokens after
// (controlled by the parameter `post`) another span query.
func NewSpanNotQuery(exclude types.SpanQueryVariant, include types.SpanQueryVariant) *_spanNotQuery {

	tmp := &_spanNotQuery{v: types.NewSpanNotQuery()}

	tmp.Exclude(exclude)

	tmp.Include(include)

	return tmp

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_spanNotQuery) Boost(boost float32) *_spanNotQuery {

	s.v.Boost = &boost

	return s
}

// The number of tokens from within the include span that can’t have overlap
// with the exclude span.
// Equivalent to setting both `pre` and `post`.
func (s *_spanNotQuery) Dist(dist int) *_spanNotQuery {

	s.v.Dist = &dist

	return s
}

// Span query whose matches must not overlap those returned.
func (s *_spanNotQuery) Exclude(exclude types.SpanQueryVariant) *_spanNotQuery {

	s.v.Exclude = *exclude.SpanQueryCaster()

	return s
}

// Span query whose matches are filtered.
func (s *_spanNotQuery) Include(include types.SpanQueryVariant) *_spanNotQuery {

	s.v.Include = *include.SpanQueryCaster()

	return s
}

// The number of tokens after the include span that can’t have overlap with the
// exclude span.
func (s *_spanNotQuery) Post(post int) *_spanNotQuery {

	s.v.Post = &post

	return s
}

// The number of tokens before the include span that can’t have overlap with the
// exclude span.
func (s *_spanNotQuery) Pre(pre int) *_spanNotQuery {

	s.v.Pre = &pre

	return s
}

func (s *_spanNotQuery) QueryName_(queryname_ string) *_spanNotQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_spanNotQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.SpanNot = s.v

	return container
}

func (s *_spanNotQuery) SpanQueryCaster() *types.SpanQuery {
	container := types.NewSpanQuery()

	container.SpanNot = s.v

	return container
}

func (s *_spanNotQuery) SpanNotQueryCaster() *types.SpanNotQuery {
	return s.v
}
