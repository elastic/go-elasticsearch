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
// https://github.com/elastic/elasticsearch-specification/tree/f1932ce6b46a53a8342db522b1a7883bcc9e0996

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// IntervalsMatch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/f1932ce6b46a53a8342db522b1a7883bcc9e0996/specification/_types/query_dsl/fulltext.ts#L186-L216
type IntervalsMatch struct {
	// Analyzer Analyzer used to analyze terms in the query.
	Analyzer *string `json:"analyzer,omitempty"`
	// Filter An optional interval filter.
	Filter *IntervalsFilter `json:"filter,omitempty"`
	// MaxGaps Maximum number of positions between the matching terms.
	// Terms further apart than this are not considered matches.
	MaxGaps *int `json:"max_gaps,omitempty"`
	// Ordered If `true`, matching terms must appear in their specified order.
	Ordered *bool `json:"ordered,omitempty"`
	// Query Text you wish to find in the provided field.
	Query string `json:"query"`
	// UseField If specified, match intervals from this field rather than the top-level
	// field.
	// The `term` is normalized using the search analyzer from this field, unless
	// `analyzer` is specified separately.
	UseField *string `json:"use_field,omitempty"`
}

func (s *IntervalsMatch) UnmarshalJSON(data []byte) error {

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

		case "analyzer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Analyzer", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Analyzer = &o

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return fmt.Errorf("%s | %w", "Filter", err)
			}

		case "max_gaps":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxGaps", err)
				}
				s.MaxGaps = &value
			case float64:
				f := int(v)
				s.MaxGaps = &f
			}

		case "ordered":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Ordered", err)
				}
				s.Ordered = &value
			case bool:
				s.Ordered = &v
			}

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Query = o

		case "use_field":
			if err := dec.Decode(&s.UseField); err != nil {
				return fmt.Errorf("%s | %w", "UseField", err)
			}

		}
	}
	return nil
}

// NewIntervalsMatch returns a IntervalsMatch.
func NewIntervalsMatch() *IntervalsMatch {
	r := &IntervalsMatch{}

	return r
}

// true

type IntervalsMatchVariant interface {
	IntervalsMatchCaster() *IntervalsMatch
}

func (s *IntervalsMatch) IntervalsMatchCaster() *IntervalsMatch {
	return s
}
