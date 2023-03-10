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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// SpanQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/query_dsl/span.ts#L79-L91
type SpanQuery struct {
	FieldMaskingSpan *SpanFieldMaskingQuery   `json:"field_masking_span,omitempty"`
	SpanContaining   *SpanContainingQuery     `json:"span_containing,omitempty"`
	SpanFirst        *SpanFirstQuery          `json:"span_first,omitempty"`
	SpanGap          map[string]int           `json:"span_gap,omitempty"`
	SpanMulti        *SpanMultiTermQuery      `json:"span_multi,omitempty"`
	SpanNear         *SpanNearQuery           `json:"span_near,omitempty"`
	SpanNot          *SpanNotQuery            `json:"span_not,omitempty"`
	SpanOr           *SpanOrQuery             `json:"span_or,omitempty"`
	SpanTerm         map[string]SpanTermQuery `json:"span_term,omitempty"`
	SpanWithin       *SpanWithinQuery         `json:"span_within,omitempty"`
}

// NewSpanQuery returns a SpanQuery.
func NewSpanQuery() *SpanQuery {
	r := &SpanQuery{
		SpanTerm: make(map[string]SpanTermQuery, 0),
	}

	return r
}
