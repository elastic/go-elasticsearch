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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
)

// QueryStringQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/fulltext.ts#L580-L700
type QueryStringQuery struct {
	// AllowLeadingWildcard If `true`, the wildcard characters `*` and `?` are allowed as the first
	// character of the query string.
	AllowLeadingWildcard *bool `json:"allow_leading_wildcard,omitempty"`
	// AnalyzeWildcard If `true`, the query attempts to analyze wildcard terms in the query string.
	AnalyzeWildcard *bool `json:"analyze_wildcard,omitempty"`
	// Analyzer Analyzer used to convert text in the query string into tokens.
	Analyzer *string `json:"analyzer,omitempty"`
	// AutoGenerateSynonymsPhraseQuery If `true`, match phrase queries are automatically created for multi-term
	// synonyms.
	AutoGenerateSynonymsPhraseQuery *bool `json:"auto_generate_synonyms_phrase_query,omitempty"`
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// DefaultField Default field to search if no field is provided in the query string.
	// Supports wildcards (`*`).
	// Defaults to the `index.query.default_field` index setting, which has a
	// default value of `*`.
	DefaultField *string `json:"default_field,omitempty"`
	// DefaultOperator Default boolean logic used to interpret text in the query string if no
	// operators are specified.
	DefaultOperator *operator.Operator `json:"default_operator,omitempty"`
	// EnablePositionIncrements If `true`, enable position increments in queries constructed from a
	// `query_string` search.
	EnablePositionIncrements *bool `json:"enable_position_increments,omitempty"`
	Escape                   *bool `json:"escape,omitempty"`
	// Fields Array of fields to search. Supports wildcards (`*`).
	Fields []string `json:"fields,omitempty"`
	// Fuzziness Maximum edit distance allowed for fuzzy matching.
	Fuzziness Fuzziness `json:"fuzziness,omitempty"`
	// FuzzyMaxExpansions Maximum number of terms to which the query expands for fuzzy matching.
	FuzzyMaxExpansions *int `json:"fuzzy_max_expansions,omitempty"`
	// FuzzyPrefixLength Number of beginning characters left unchanged for fuzzy matching.
	FuzzyPrefixLength *int `json:"fuzzy_prefix_length,omitempty"`
	// FuzzyRewrite Method used to rewrite the query.
	FuzzyRewrite *string `json:"fuzzy_rewrite,omitempty"`
	// FuzzyTranspositions If `true`, edits for fuzzy matching include transpositions of two adjacent
	// characters (for example, `ab` to `ba`).
	FuzzyTranspositions *bool `json:"fuzzy_transpositions,omitempty"`
	// Lenient If `true`, format-based errors, such as providing a text value for a numeric
	// field, are ignored.
	Lenient *bool `json:"lenient,omitempty"`
	// MaxDeterminizedStates Maximum number of automaton states required for the query.
	MaxDeterminizedStates *int `json:"max_determinized_states,omitempty"`
	// MinimumShouldMatch Minimum number of clauses that must match for a document to be returned.
	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`
	// PhraseSlop Maximum number of positions allowed between matching tokens for phrases.
	PhraseSlop *Float64 `json:"phrase_slop,omitempty"`
	// Query Query string you wish to parse and use for search.
	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`
	// QuoteAnalyzer Analyzer used to convert quoted text in the query string into tokens.
	// For quoted text, this parameter overrides the analyzer specified in the
	// `analyzer` parameter.
	QuoteAnalyzer *string `json:"quote_analyzer,omitempty"`
	// QuoteFieldSuffix Suffix appended to quoted text in the query string.
	// You can use this suffix to use a different analysis method for exact matches.
	QuoteFieldSuffix *string `json:"quote_field_suffix,omitempty"`
	// Rewrite Method used to rewrite the query.
	Rewrite *string `json:"rewrite,omitempty"`
	// TieBreaker How to combine the queries generated from the individual search terms in the
	// resulting `dis_max` query.
	TieBreaker *Float64 `json:"tie_breaker,omitempty"`
	// TimeZone Coordinated Universal Time (UTC) offset or IANA time zone used to convert
	// date values in the query string to UTC.
	TimeZone *string `json:"time_zone,omitempty"`
	// Type Determines how the query matches and scores documents.
	Type *textquerytype.TextQueryType `json:"type,omitempty"`
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
