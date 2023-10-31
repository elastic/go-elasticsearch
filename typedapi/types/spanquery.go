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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// SpanQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/query_dsl/span.ts#L131-L170
type SpanQuery struct {
	// FieldMaskingSpan Allows queries like `span_near` or `span_or` across different fields.
	FieldMaskingSpan *SpanFieldMaskingQuery `json:"field_masking_span,omitempty"`
	// SpanContaining Accepts a list of span queries, but only returns those spans which also match
	// a second span query.
	SpanContaining *SpanContainingQuery `json:"span_containing,omitempty"`
	// SpanFirst Accepts another span query whose matches must appear within the first N
	// positions of the field.
	SpanFirst *SpanFirstQuery `json:"span_first,omitempty"`
	SpanGap   SpanGapQuery    `json:"span_gap,omitempty"`
	// SpanMulti Wraps a `term`, `range`, `prefix`, `wildcard`, `regexp`, or `fuzzy` query.
	SpanMulti *SpanMultiTermQuery `json:"span_multi,omitempty"`
	// SpanNear Accepts multiple span queries whose matches must be within the specified
	// distance of each other, and possibly in the same order.
	SpanNear *SpanNearQuery `json:"span_near,omitempty"`
	// SpanNot Wraps another span query, and excludes any documents which match that query.
	SpanNot *SpanNotQuery `json:"span_not,omitempty"`
	// SpanOr Combines multiple span queries and returns documents which match any of the
	// specified queries.
	SpanOr *SpanOrQuery `json:"span_or,omitempty"`
	// SpanTerm The equivalent of the `term` query but for use with other span queries.
	SpanTerm map[string]SpanTermQuery `json:"span_term,omitempty"`
	// SpanWithin The result from a single span query is returned as long is its span falls
	// within the spans returned by a list of other span queries.
	SpanWithin *SpanWithinQuery `json:"span_within,omitempty"`
}

func (s *SpanQuery) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "field_masking_span":
			if err := dec.Decode(&s.FieldMaskingSpan); err != nil {
				return err
			}

		case "span_containing":
			if err := dec.Decode(&s.SpanContaining); err != nil {
				return err
			}

		case "span_first":
			if err := dec.Decode(&s.SpanFirst); err != nil {
				return err
			}

		case "span_gap":
			if err := dec.Decode(&s.SpanGap); err != nil {
				return err
			}

		case "span_multi":
			if err := dec.Decode(&s.SpanMulti); err != nil {
				return err
			}

		case "span_near":
			if err := dec.Decode(&s.SpanNear); err != nil {
				return err
			}

		case "span_not":
			if err := dec.Decode(&s.SpanNot); err != nil {
				return err
			}

		case "span_or":
			if err := dec.Decode(&s.SpanOr); err != nil {
				return err
			}

		case "span_term":
			if s.SpanTerm == nil {
				s.SpanTerm = make(map[string]SpanTermQuery, 0)
			}
			if err := dec.Decode(&s.SpanTerm); err != nil {
				return err
			}

		case "span_within":
			if err := dec.Decode(&s.SpanWithin); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSpanQuery returns a SpanQuery.
func NewSpanQuery() *SpanQuery {
	r := &SpanQuery{
		SpanTerm: make(map[string]SpanTermQuery, 0),
	}

	return r
}
