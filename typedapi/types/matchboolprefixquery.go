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
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
)

// MatchBoolPrefixQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/query_dsl/fulltext.ts#L160-L171
type MatchBoolPrefixQuery struct {
	Analyzer            *string            `json:"analyzer,omitempty"`
	Boost               *float32           `json:"boost,omitempty"`
	Fuzziness           Fuzziness          `json:"fuzziness,omitempty"`
	FuzzyRewrite        *string            `json:"fuzzy_rewrite,omitempty"`
	FuzzyTranspositions *bool              `json:"fuzzy_transpositions,omitempty"`
	MaxExpansions       *int               `json:"max_expansions,omitempty"`
	MinimumShouldMatch  MinimumShouldMatch `json:"minimum_should_match,omitempty"`
	Operator            *operator.Operator `json:"operator,omitempty"`
	PrefixLength        *int               `json:"prefix_length,omitempty"`
	Query               string             `json:"query"`
	QueryName_          *string            `json:"_name,omitempty"`
}

func (s *MatchBoolPrefixQuery) UnmarshalJSON(data []byte) error {

	if !bytes.HasPrefix(data, []byte(`{`)) {
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&s.Query)
		return err
	}

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

		case "fuzziness":
			if err := dec.Decode(&s.Fuzziness); err != nil {
				return err
			}

		case "fuzzy_rewrite":
			if err := dec.Decode(&s.FuzzyRewrite); err != nil {
				return err
			}

		case "fuzzy_transpositions":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.FuzzyTranspositions = &value
			case bool:
				s.FuzzyTranspositions = &v
			}

		case "max_expansions":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxExpansions = &value
			case float64:
				f := int(v)
				s.MaxExpansions = &f
			}

		case "minimum_should_match":
			if err := dec.Decode(&s.MinimumShouldMatch); err != nil {
				return err
			}

		case "operator":
			if err := dec.Decode(&s.Operator); err != nil {
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

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Query = o

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

		}
	}
	return nil
}

// NewMatchBoolPrefixQuery returns a MatchBoolPrefixQuery.
func NewMatchBoolPrefixQuery() *MatchBoolPrefixQuery {
	r := &MatchBoolPrefixQuery{}

	return r
}
