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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// IntervalsQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_types/query_dsl/fulltext.ts#L235-L263
type IntervalsQuery struct {
	// AllOf Returns matches that span a combination of other rules.
	AllOf *IntervalsAllOf `json:"all_of,omitempty"`
	// AnyOf Returns intervals produced by any of its sub-rules.
	AnyOf *IntervalsAnyOf `json:"any_of,omitempty"`
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// Fuzzy Matches terms that are similar to the provided term, within an edit distance
	// defined by `fuzziness`.
	Fuzzy *IntervalsFuzzy `json:"fuzzy,omitempty"`
	// Match Matches analyzed text.
	Match *IntervalsMatch `json:"match,omitempty"`
	// Prefix Matches terms that start with a specified set of characters.
	Prefix     *IntervalsPrefix `json:"prefix,omitempty"`
	QueryName_ *string          `json:"_name,omitempty"`
	// Wildcard Matches terms using a wildcard pattern.
	Wildcard *IntervalsWildcard `json:"wildcard,omitempty"`
}

func (s *IntervalsQuery) UnmarshalJSON(data []byte) error {

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

		case "all_of":
			if err := dec.Decode(&s.AllOf); err != nil {
				return fmt.Errorf("%s | %w", "AllOf", err)
			}

		case "any_of":
			if err := dec.Decode(&s.AnyOf); err != nil {
				return fmt.Errorf("%s | %w", "AnyOf", err)
			}

		case "boost":
			var tmp interface{}
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

		case "fuzzy":
			if err := dec.Decode(&s.Fuzzy); err != nil {
				return fmt.Errorf("%s | %w", "Fuzzy", err)
			}

		case "match":
			if err := dec.Decode(&s.Match); err != nil {
				return fmt.Errorf("%s | %w", "Match", err)
			}

		case "prefix":
			if err := dec.Decode(&s.Prefix); err != nil {
				return fmt.Errorf("%s | %w", "Prefix", err)
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

		case "wildcard":
			if err := dec.Decode(&s.Wildcard); err != nil {
				return fmt.Errorf("%s | %w", "Wildcard", err)
			}

		}
	}
	return nil
}

// NewIntervalsQuery returns a IntervalsQuery.
func NewIntervalsQuery() *IntervalsQuery {
	r := &IntervalsQuery{}

	return r
}
