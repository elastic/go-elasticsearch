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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// SpanQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/query_dsl/span.ts#L158-L200
type SpanQuery struct {
	AdditionalSpanQueryProperty map[string]json.RawMessage `json:"-"`
	// SpanContaining Accepts a list of span queries, but only returns those spans which also match
	// a second span query.
	SpanContaining *SpanContainingQuery `json:"span_containing,omitempty"`
	// SpanFieldMasking Allows queries like `span_near` or `span_or` across different fields.
	SpanFieldMasking *SpanFieldMaskingQuery `json:"span_field_masking,omitempty"`
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

		case "span_containing":
			if err := dec.Decode(&s.SpanContaining); err != nil {
				return fmt.Errorf("%s | %w", "SpanContaining", err)
			}

		case "span_field_masking":
			if err := dec.Decode(&s.SpanFieldMasking); err != nil {
				return fmt.Errorf("%s | %w", "SpanFieldMasking", err)
			}

		case "span_first":
			if err := dec.Decode(&s.SpanFirst); err != nil {
				return fmt.Errorf("%s | %w", "SpanFirst", err)
			}

		case "span_gap":
			if err := dec.Decode(&s.SpanGap); err != nil {
				return fmt.Errorf("%s | %w", "SpanGap", err)
			}

		case "span_multi":
			if err := dec.Decode(&s.SpanMulti); err != nil {
				return fmt.Errorf("%s | %w", "SpanMulti", err)
			}

		case "span_near":
			if err := dec.Decode(&s.SpanNear); err != nil {
				return fmt.Errorf("%s | %w", "SpanNear", err)
			}

		case "span_not":
			if err := dec.Decode(&s.SpanNot); err != nil {
				return fmt.Errorf("%s | %w", "SpanNot", err)
			}

		case "span_or":
			if err := dec.Decode(&s.SpanOr); err != nil {
				return fmt.Errorf("%s | %w", "SpanOr", err)
			}

		case "span_term":
			if s.SpanTerm == nil {
				s.SpanTerm = make(map[string]SpanTermQuery, 0)
			}
			if err := dec.Decode(&s.SpanTerm); err != nil {
				return fmt.Errorf("%s | %w", "SpanTerm", err)
			}

		case "span_within":
			if err := dec.Decode(&s.SpanWithin); err != nil {
				return fmt.Errorf("%s | %w", "SpanWithin", err)
			}

		default:

			if key, ok := t.(string); ok {
				if s.AdditionalSpanQueryProperty == nil {
					s.AdditionalSpanQueryProperty = make(map[string]json.RawMessage, 0)
				}
				raw := new(json.RawMessage)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "AdditionalSpanQueryProperty", err)
				}
				s.AdditionalSpanQueryProperty[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s SpanQuery) MarshalJSON() ([]byte, error) {
	type opt SpanQuery
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalSpanQueryProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalSpanQueryProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewSpanQuery returns a SpanQuery.
func NewSpanQuery() *SpanQuery {
	r := &SpanQuery{
		AdditionalSpanQueryProperty: make(map[string]json.RawMessage),
		SpanTerm:                    make(map[string]SpanTermQuery),
	}

	return r
}
