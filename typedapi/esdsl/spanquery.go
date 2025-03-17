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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _spanQuery struct {
	v *types.SpanQuery
}

func NewSpanQuery() *_spanQuery {
	return &_spanQuery{v: types.NewSpanQuery()}
}

// AdditionalSpanQueryProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_spanQuery) AdditionalSpanQueryProperty(key string, value json.RawMessage) *_spanQuery {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalSpanQueryProperty = tmp
	return s
}

// Accepts a list of span queries, but only returns those spans which also match
// a second span query.
func (s *_spanQuery) SpanContaining(spancontaining types.SpanContainingQueryVariant) *_spanQuery {

	s.v.SpanContaining = spancontaining.SpanContainingQueryCaster()

	return s
}

// Allows queries like `span_near` or `span_or` across different fields.
func (s *_spanQuery) SpanFieldMasking(spanfieldmasking types.SpanFieldMaskingQueryVariant) *_spanQuery {

	s.v.SpanFieldMasking = spanfieldmasking.SpanFieldMaskingQueryCaster()

	return s
}

// Accepts another span query whose matches must appear within the first N
// positions of the field.
func (s *_spanQuery) SpanFirst(spanfirst types.SpanFirstQueryVariant) *_spanQuery {

	s.v.SpanFirst = spanfirst.SpanFirstQueryCaster()

	return s
}

func (s *_spanQuery) SpanGap(spangapquery types.SpanGapQueryVariant) *_spanQuery {

	s.v.SpanGap = *spangapquery.SpanGapQueryCaster()

	return s
}

// Wraps a `term`, `range`, `prefix`, `wildcard`, `regexp`, or `fuzzy` query.
func (s *_spanQuery) SpanMulti(spanmulti types.SpanMultiTermQueryVariant) *_spanQuery {

	s.v.SpanMulti = spanmulti.SpanMultiTermQueryCaster()

	return s
}

// Accepts multiple span queries whose matches must be within the specified
// distance of each other, and possibly in the same order.
func (s *_spanQuery) SpanNear(spannear types.SpanNearQueryVariant) *_spanQuery {

	s.v.SpanNear = spannear.SpanNearQueryCaster()

	return s
}

// Wraps another span query, and excludes any documents which match that query.
func (s *_spanQuery) SpanNot(spannot types.SpanNotQueryVariant) *_spanQuery {

	s.v.SpanNot = spannot.SpanNotQueryCaster()

	return s
}

// Combines multiple span queries and returns documents which match any of the
// specified queries.
func (s *_spanQuery) SpanOr(spanor types.SpanOrQueryVariant) *_spanQuery {

	s.v.SpanOr = spanor.SpanOrQueryCaster()

	return s
}

// The equivalent of the `term` query but for use with other span queries.
// SpanTerm is a single key dictionnary.
// It will replace the current value on each call.
func (s *_spanQuery) SpanTerm(key string, value types.SpanTermQueryVariant) *_spanQuery {

	tmp := make(map[string]types.SpanTermQuery)

	tmp[key] = *value.SpanTermQueryCaster()

	s.v.SpanTerm = tmp
	return s
}

// The result from a single span query is returned as long is its span falls
// within the spans returned by a list of other span queries.
func (s *_spanQuery) SpanWithin(spanwithin types.SpanWithinQueryVariant) *_spanQuery {

	s.v.SpanWithin = spanwithin.SpanWithinQueryCaster()

	return s
}

func (s *_spanQuery) SpanQueryCaster() *types.SpanQuery {
	return s.v
}
