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

// TermsSetQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/term.ts#L249-L262
type TermsSetQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// MinimumShouldMatchField Numeric field containing the number of matching terms required to return a
	// document.
	MinimumShouldMatchField *string `json:"minimum_should_match_field,omitempty"`
	// MinimumShouldMatchScript Custom script containing the number of matching terms required to return a
	// document.
	MinimumShouldMatchScript Script  `json:"minimum_should_match_script,omitempty"`
	QueryName_               *string `json:"_name,omitempty"`
	// Terms Array of terms you wish to find in the provided field.
	Terms []string `json:"terms"`
}

func (s *TermsSetQuery) UnmarshalJSON(data []byte) error {

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
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "minimum_should_match_field":
			if err := dec.Decode(&s.MinimumShouldMatchField); err != nil {
				return err
			}

		case "minimum_should_match_script":
			if err := dec.Decode(&s.MinimumShouldMatchScript); err != nil {
				return err
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		case "terms":
			if err := dec.Decode(&s.Terms); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTermsSetQuery returns a TermsSetQuery.
func NewTermsSetQuery() *TermsSetQuery {
	r := &TermsSetQuery{}

	return r
}
