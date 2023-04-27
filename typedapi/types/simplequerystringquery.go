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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/simplequerystringflag"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// SimpleQueryStringQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/query_dsl/fulltext.ts#L294-L312
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
			o := string(tmp)
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
			o := string(tmp)
			s.Query = o

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.QueryName_ = &o

		case "quote_field_suffix":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
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
