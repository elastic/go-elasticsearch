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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
)

// QueryStringQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/query_dsl/fulltext.ts#L233-L269
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

// NewQueryStringQuery returns a QueryStringQuery.
func NewQueryStringQuery() *QueryStringQuery {
	r := &QueryStringQuery{}

	return r
}
