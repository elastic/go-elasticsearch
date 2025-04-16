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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanFirstQuery struct {
	v *types.SpanFirstQuery
}

// Matches spans near the beginning of a field.
func NewSpanFirstQuery(end int, match types.SpanQueryVariant) *_spanFirstQuery {

	tmp := &_spanFirstQuery{v: types.NewSpanFirstQuery()}

	tmp.End(end)

	tmp.Match(match)

	return tmp

}

func (s *_spanFirstQuery) Boost(boost float32) *_spanFirstQuery {

	s.v.Boost = &boost

	return s
}

func (s *_spanFirstQuery) End(end int) *_spanFirstQuery {

	s.v.End = end

	return s
}

func (s *_spanFirstQuery) Match(match types.SpanQueryVariant) *_spanFirstQuery {

	s.v.Match = *match.SpanQueryCaster()

	return s
}

func (s *_spanFirstQuery) QueryName_(queryname_ string) *_spanFirstQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_spanFirstQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.SpanFirst = s.v

	return container
}

func (s *_spanFirstQuery) SpanQueryCaster() *types.SpanQuery {
	container := types.NewSpanQuery()

	container.SpanFirst = s.v

	return container
}

func (s *_spanFirstQuery) SpanFirstQueryCaster() *types.SpanFirstQuery {
	return s.v
}
