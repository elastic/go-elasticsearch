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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// IntervalsFuzzy type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/fulltext.ts#L154-L184
type IntervalsFuzzy struct {
	// Analyzer Analyzer used to normalize the term.
	Analyzer *string `json:"analyzer,omitempty"`
	// Fuzziness Maximum edit distance allowed for matching.
	Fuzziness Fuzziness `json:"fuzziness,omitempty"`
	// PrefixLength Number of beginning characters left unchanged when creating expansions.
	PrefixLength *int `json:"prefix_length,omitempty"`
	// Term The term to match.
	Term string `json:"term"`
	// Transpositions Indicates whether edits include transpositions of two adjacent characters
	// (for example, `ab` to `ba`).
	Transpositions *bool `json:"transpositions,omitempty"`
	// UseField If specified, match intervals from this field rather than the top-level
	// field.
	// The `term` is normalized using the search analyzer from this field, unless
	// `analyzer` is specified separately.
	UseField *string `json:"use_field,omitempty"`
}

func (s *IntervalsFuzzy) UnmarshalJSON(data []byte) error {

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
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Analyzer = &o

		case "fuzziness":
			if err := dec.Decode(&s.Fuzziness); err != nil {
				return err
			}

		case "prefix_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PrefixLength = &value
			case float64:
				f := int(v)
				s.PrefixLength = &f
			}

		case "term":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Term = o

		case "transpositions":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Transpositions = &value
			case bool:
				s.Transpositions = &v
			}

		case "use_field":
			if err := dec.Decode(&s.UseField); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIntervalsFuzzy returns a IntervalsFuzzy.
func NewIntervalsFuzzy() *IntervalsFuzzy {
	r := &IntervalsFuzzy{}

	return r
}
