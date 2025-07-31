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
	"strconv"
)

// SpanNotQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/query_dsl/span.ts#L95-L122
type SpanNotQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// Dist The number of tokens from within the include span that can’t have overlap
	// with the exclude span.
	// Equivalent to setting both `pre` and `post`.
	Dist *int `json:"dist,omitempty"`
	// Exclude Span query whose matches must not overlap those returned.
	Exclude SpanQuery `json:"exclude"`
	// Include Span query whose matches are filtered.
	Include SpanQuery `json:"include"`
	// Post The number of tokens after the include span that can’t have overlap with the
	// exclude span.
	Post *int `json:"post,omitempty"`
	// Pre The number of tokens before the include span that can’t have overlap with the
	// exclude span.
	Pre        *int    `json:"pre,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *SpanNotQuery) UnmarshalJSON(data []byte) error {

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

		case "boost":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Boost", err)
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "dist":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Dist", err)
				}
				s.Dist = &value
			case float64:
				f := int(v)
				s.Dist = &f
			}

		case "exclude":
			if err := dec.Decode(&s.Exclude); err != nil {
				return fmt.Errorf("%s | %w", "Exclude", err)
			}

		case "include":
			if err := dec.Decode(&s.Include); err != nil {
				return fmt.Errorf("%s | %w", "Include", err)
			}

		case "post":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Post", err)
				}
				s.Post = &value
			case float64:
				f := int(v)
				s.Post = &f
			}

		case "pre":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Pre", err)
				}
				s.Pre = &value
			case float64:
				f := int(v)
				s.Pre = &f
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "QueryName_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		}
	}
	return nil
}

// NewSpanNotQuery returns a SpanNotQuery.
func NewSpanNotQuery() *SpanNotQuery {
	r := &SpanNotQuery{}

	return r
}
