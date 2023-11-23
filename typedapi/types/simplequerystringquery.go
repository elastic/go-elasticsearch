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
)

// SimpleQueryStringQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/fulltext.ts#L765-L830
type SimpleQueryStringQuery struct {
	// AnalyzeWildcard If `true`, the query attempts to analyze wildcard terms in the query string.
	AnalyzeWildcard *bool `json:"analyze_wildcard,omitempty"`
	// Analyzer Analyzer used to convert text in the query string into tokens.
	Analyzer *string `json:"analyzer,omitempty"`
	// AutoGenerateSynonymsPhraseQuery If `true`, the parser creates a match_phrase query for each multi-position
	// token.
	AutoGenerateSynonymsPhraseQuery *bool `json:"auto_generate_synonyms_phrase_query,omitempty"`
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// DefaultOperator Default boolean logic used to interpret text in the query string if no
	// operators are specified.
	DefaultOperator *operator.Operator `json:"default_operator,omitempty"`
	// Fields Array of fields you wish to search.
	// Accepts wildcard expressions.
	// You also can boost relevance scores for matches to particular fields using a
	// caret (`^`) notation.
	// Defaults to the `index.query.default_field index` setting, which has a
	// default value of `*`.
	Fields []string `json:"fields,omitempty"`
	// Flags List of enabled operators for the simple query string syntax.
	Flags PipeSeparatedFlagsSimpleQueryStringFlag `json:"flags,omitempty"`
	// FuzzyMaxExpansions Maximum number of terms to which the query expands for fuzzy matching.
	FuzzyMaxExpansions *int `json:"fuzzy_max_expansions,omitempty"`
	// FuzzyPrefixLength Number of beginning characters left unchanged for fuzzy matching.
	FuzzyPrefixLength *int `json:"fuzzy_prefix_length,omitempty"`
	// FuzzyTranspositions If `true`, edits for fuzzy matching include transpositions of two adjacent
	// characters (for example, `ab` to `ba`).
	FuzzyTranspositions *bool `json:"fuzzy_transpositions,omitempty"`
	// Lenient If `true`, format-based errors, such as providing a text value for a numeric
	// field, are ignored.
	Lenient *bool `json:"lenient,omitempty"`
	// MinimumShouldMatch Minimum number of clauses that must match for a document to be returned.
	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`
	// Query Query string in the simple query string syntax you wish to parse and use for
	// search.
	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`
	// QuoteFieldSuffix Suffix appended to quoted text in the query string.
	QuoteFieldSuffix *string `json:"quote_field_suffix,omitempty"`
}

func (s *SimpleQueryStringQuery) UnmarshalJSON(data []byte) error {

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

		case "default_operator":
			if err := dec.Decode(&s.DefaultOperator); err != nil {
				return err
			}

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "flags":
			if err := dec.Decode(&s.Flags); err != nil {
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

		case "minimum_should_match":
			if err := dec.Decode(&s.MinimumShouldMatch); err != nil {
				return err
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

		}
	}
	return nil
}

// NewSimpleQueryStringQuery returns a SimpleQueryStringQuery.
func NewSimpleQueryStringQuery() *SimpleQueryStringQuery {
	r := &SimpleQueryStringQuery{}

	return r
}
