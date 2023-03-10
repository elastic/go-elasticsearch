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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/simplequerystringflag"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// SimpleQueryStringQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/query_dsl/fulltext.ts#L294-L312
type SimpleQueryStringQuery struct {
	AnalyzeWildcard                 *bool                  `json:"analyze_wildcard,omitempty"`
	Analyzer                        *string                `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery *bool                  `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float32               `json:"boost,omitempty"`
	DefaultOperator                 *operator.Operator     `json:"default_operator,omitempty"`
	Fields                          []string               `json:"fields,omitempty"`
	Flags                           SimpleQueryStringFlags `json:"flags,omitempty"`
	FuzzyMaxExpansions              *int                   `json:"fuzzy_max_expansions,omitempty"`
	FuzzyPrefixLength               *int                   `json:"fuzzy_prefix_length,omitempty"`
	FuzzyTranspositions             *bool                  `json:"fuzzy_transpositions,omitempty"`
	Lenient                         *bool                  `json:"lenient,omitempty"`
	MinimumShouldMatch              MinimumShouldMatch     `json:"minimum_should_match,omitempty"`
	Query                           string                 `json:"query"`
	QueryName_                      *string                `json:"_name,omitempty"`
	QuoteFieldSuffix                *string                `json:"quote_field_suffix,omitempty"`
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
			if err := dec.Decode(&s.AnalyzeWildcard); err != nil {
				return err
			}

		case "analyzer":
			if err := dec.Decode(&s.Analyzer); err != nil {
				return err
			}

		case "auto_generate_synonyms_phrase_query":
			if err := dec.Decode(&s.AutoGenerateSynonymsPhraseQuery); err != nil {
				return err
			}

		case "boost":
			if err := dec.Decode(&s.Boost); err != nil {
				return err
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
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := &simplequerystringflag.SimpleQueryStringFlag{}
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Flags = *o

			default:
				if err := localDec.Decode(&s.Flags); err != nil {
					return err
				}
			}

		case "fuzzy_max_expansions":
			if err := dec.Decode(&s.FuzzyMaxExpansions); err != nil {
				return err
			}

		case "fuzzy_prefix_length":
			if err := dec.Decode(&s.FuzzyPrefixLength); err != nil {
				return err
			}

		case "fuzzy_transpositions":
			if err := dec.Decode(&s.FuzzyTranspositions); err != nil {
				return err
			}

		case "lenient":
			if err := dec.Decode(&s.Lenient); err != nil {
				return err
			}

		case "minimum_should_match":
			if err := dec.Decode(&s.MinimumShouldMatch); err != nil {
				return err
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return err
			}

		case "_name":
			if err := dec.Decode(&s.QueryName_); err != nil {
				return err
			}

		case "quote_field_suffix":
			if err := dec.Decode(&s.QuoteFieldSuffix); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSimpleQueryStringQuery returns a SimpleQueryStringQuery.
func NewSimpleQueryStringQuery() *SimpleQueryStringQuery {
	r := &SimpleQueryStringQuery{}

	return r
}
