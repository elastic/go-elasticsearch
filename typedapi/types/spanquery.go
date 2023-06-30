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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// SpanQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/query_dsl/span.ts#L79-L91
type SpanQuery struct {
	FieldMaskingSpan *SpanFieldMaskingQuery   `json:"field_masking_span,omitempty"`
	SpanContaining   *SpanContainingQuery     `json:"span_containing,omitempty"`
	SpanFirst        *SpanFirstQuery          `json:"span_first,omitempty"`
	SpanGap          SpanGapQuery             `json:"span_gap,omitempty"`
	SpanMulti        *SpanMultiTermQuery      `json:"span_multi,omitempty"`
	SpanNear         *SpanNearQuery           `json:"span_near,omitempty"`
	SpanNot          *SpanNotQuery            `json:"span_not,omitempty"`
	SpanOr           *SpanOrQuery             `json:"span_or,omitempty"`
	SpanTerm         map[string]SpanTermQuery `json:"span_term,omitempty"`
	SpanWithin       *SpanWithinQuery         `json:"span_within,omitempty"`
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
