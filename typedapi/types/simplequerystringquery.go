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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
)

// SimpleQueryStringQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_types/query_dsl/fulltext.ts#L294-L312
type SimpleQueryStringQuery struct {
	AnalyzeWildcard                 *bool                   `json:"analyze_wildcard,omitempty"`
	Analyzer                        *string                 `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery *bool                   `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float32                `json:"boost,omitempty"`
	DefaultOperator                 *operator.Operator      `json:"default_operator,omitempty"`
	Fields                          []string                `json:"fields,omitempty"`
	Flags                           *SimpleQueryStringFlags `json:"flags,omitempty"`
	FuzzyMaxExpansions              *int                    `json:"fuzzy_max_expansions,omitempty"`
	FuzzyPrefixLength               *int                    `json:"fuzzy_prefix_length,omitempty"`
	FuzzyTranspositions             *bool                   `json:"fuzzy_transpositions,omitempty"`
	Lenient                         *bool                   `json:"lenient,omitempty"`
	MinimumShouldMatch              *MinimumShouldMatch     `json:"minimum_should_match,omitempty"`
	Query                           string                  `json:"query"`
	QueryName_                      *string                 `json:"_name,omitempty"`
	QuoteFieldSuffix                *string                 `json:"quote_field_suffix,omitempty"`
}

// NewSimpleQueryStringQuery returns a SimpleQueryStringQuery.
func NewSimpleQueryStringQuery() *SimpleQueryStringQuery {
	r := &SimpleQueryStringQuery{}

	return r
}
