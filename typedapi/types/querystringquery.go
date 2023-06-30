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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
)

// QueryStringQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/query_dsl/fulltext.ts#L233-L269
type QueryStringQuery struct {
	AllowLeadingWildcard            *bool                        `json:"allow_leading_wildcard,omitempty"`
	AnalyzeWildcard                 *bool                        `json:"analyze_wildcard,omitempty"`
	Analyzer                        *string                      `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery *bool                        `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float32                     `json:"boost,omitempty"`
	DefaultField                    *string                      `json:"default_field,omitempty"`
	DefaultOperator                 *operator.Operator           `json:"default_operator,omitempty"`
	EnablePositionIncrements        *bool                        `json:"enable_position_increments,omitempty"`
	Escape                          *bool                        `json:"escape,omitempty"`
	Fields                          []string                     `json:"fields,omitempty"`
	Fuzziness                       Fuzziness                    `json:"fuzziness,omitempty"`
	FuzzyMaxExpansions              *int                         `json:"fuzzy_max_expansions,omitempty"`
	FuzzyPrefixLength               *int                         `json:"fuzzy_prefix_length,omitempty"`
	FuzzyRewrite                    *string                      `json:"fuzzy_rewrite,omitempty"`
	FuzzyTranspositions             *bool                        `json:"fuzzy_transpositions,omitempty"`
	Lenient                         *bool                        `json:"lenient,omitempty"`
	MaxDeterminizedStates           *int                         `json:"max_determinized_states,omitempty"`
	MinimumShouldMatch              MinimumShouldMatch           `json:"minimum_should_match,omitempty"`
	PhraseSlop                      *Float64                     `json:"phrase_slop,omitempty"`
	Query                           string                       `json:"query"`
	QueryName_                      *string                      `json:"_name,omitempty"`
	QuoteAnalyzer                   *string                      `json:"quote_analyzer,omitempty"`
	QuoteFieldSuffix                *string                      `json:"quote_field_suffix,omitempty"`
	Rewrite                         *string                      `json:"rewrite,omitempty"`
	TieBreaker                      *Float64                     `json:"tie_breaker,omitempty"`
	TimeZone                        *string                      `json:"time_zone,omitempty"`
	Type                            *textquerytype.TextQueryType `json:"type,omitempty"`
}

func (s *QueryStringQuery) UnmarshalJSON(data []byte) error {

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

		case "allow_leading_wildcard":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AllowLeadingWildcard = &value
			case bool:
				s.AllowLeadingWildcard = &v
			}

		case "analyze_wildcard":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AnalyzeWildcard = &value
			case bool:
				s.AnalyzeWildcard = &v
			}

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

		case "auto_generate_synonyms_phrase_query":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AutoGenerateSynonymsPhraseQuery = &value
			case bool:
				s.AutoGenerateSynonymsPhraseQuery = &v
			}

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

		case "default_field":
			if err := dec.Decode(&s.DefaultField); err != nil {
				return err
			}

		case "default_operator":
			if err := dec.Decode(&s.DefaultOperator); err != nil {
				return err
			}

		case "enable_position_increments":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.EnablePositionIncrements = &value
			case bool:
				s.EnablePositionIncrements = &v
			}

		case "escape":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Escape = &value
			case bool:
				s.Escape = &v
			}

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "fuzziness":
			if err := dec.Decode(&s.Fuzziness); err != nil {
				return err
			}

		case "fuzzy_max_expansions":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FuzzyMaxExpansions = &value
			case float64:
				f := int(v)
				s.FuzzyMaxExpansions = &f
			}

		case "fuzzy_prefix_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FuzzyPrefixLength = &value
			case float64:
				f := int(v)
				s.FuzzyPrefixLength = &f
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

		case "lenient":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Lenient = &value
			case bool:
				s.Lenient = &v
			}

		case "max_determinized_states":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxDeterminizedStates = &value
			case float64:
				f := int(v)
				s.MaxDeterminizedStates = &f
			}

		case "minimum_should_match":
			if err := dec.Decode(&s.MinimumShouldMatch); err != nil {
				return err
			}

		case "phrase_slop":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.PhraseSlop = &f
			case float64:
				f := Float64(v)
				s.PhraseSlop = &f
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

		case "quote_analyzer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QuoteAnalyzer = &o

		case "quote_field_suffix":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QuoteFieldSuffix = &o

		case "rewrite":
			if err := dec.Decode(&s.Rewrite); err != nil {
				return err
			}

		case "tie_breaker":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.TieBreaker = &f
			case float64:
				f := Float64(v)
				s.TieBreaker = &f
			}

		case "time_zone":
			if err := dec.Decode(&s.TimeZone); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewQueryStringQuery returns a QueryStringQuery.
func NewQueryStringQuery() *QueryStringQuery {
	r := &QueryStringQuery{}

	return r
}
