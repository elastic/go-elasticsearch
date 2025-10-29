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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
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

func (s *_spanQuery) SpanContaining(spancontaining types.SpanContainingQueryVariant) *_spanQuery {

	s.v.SpanContaining = spancontaining.SpanContainingQueryCaster()

	return s
}

func (s *_spanQuery) SpanFieldMasking(spanfieldmasking types.SpanFieldMaskingQueryVariant) *_spanQuery {

	s.v.SpanFieldMasking = spanfieldmasking.SpanFieldMaskingQueryCaster()

	return s
}

func (s *_spanQuery) SpanFirst(spanfirst types.SpanFirstQueryVariant) *_spanQuery {

	s.v.SpanFirst = spanfirst.SpanFirstQueryCaster()

	return s
}

func (s *_spanQuery) SpanGap(spangapquery types.SpanGapQueryVariant) *_spanQuery {

	s.v.SpanGap = *spangapquery.SpanGapQueryCaster()

	return s
}

func (s *_spanQuery) SpanMulti(spanmulti types.SpanMultiTermQueryVariant) *_spanQuery {

	s.v.SpanMulti = spanmulti.SpanMultiTermQueryCaster()

	return s
}

func (s *_spanQuery) SpanNear(spannear types.SpanNearQueryVariant) *_spanQuery {

	s.v.SpanNear = spannear.SpanNearQueryCaster()

	return s
}

func (s *_spanQuery) SpanNot(spannot types.SpanNotQueryVariant) *_spanQuery {

	s.v.SpanNot = spannot.SpanNotQueryCaster()

	return s
}

func (s *_spanQuery) SpanOr(spanor types.SpanOrQueryVariant) *_spanQuery {

	s.v.SpanOr = spanor.SpanOrQueryCaster()

	return s
}

// SpanTerm is a single key dictionnary.
// It will replace the current value on each call.
func (s *_spanQuery) SpanTerm(key string, value types.SpanTermQueryVariant) *_spanQuery {

	tmp := make(map[string]types.SpanTermQuery)

	tmp[key] = *value.SpanTermQueryCaster()

	s.v.SpanTerm = tmp
	return s
}

func (s *_spanQuery) SpanWithin(spanwithin types.SpanWithinQueryVariant) *_spanQuery {

	s.v.SpanWithin = spanwithin.SpanWithinQueryCaster()

	return s
}

func (s *_spanQuery) SpanQueryCaster() *types.SpanQuery {
	return s.v
}
