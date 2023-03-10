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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/combinedfieldsoperator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/combinedfieldszeroterms"
)

// CombinedFieldsQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/query_dsl/abstractions.ts#L181-L195
type CombinedFieldsQuery struct {
	AutoGenerateSynonymsPhraseQuery *bool                                            `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float32                                         `json:"boost,omitempty"`
	Fields                          []string                                         `json:"fields"`
	MinimumShouldMatch              MinimumShouldMatch                               `json:"minimum_should_match,omitempty"`
	Operator                        *combinedfieldsoperator.CombinedFieldsOperator   `json:"operator,omitempty"`
	Query                           string                                           `json:"query"`
	QueryName_                      *string                                          `json:"_name,omitempty"`
	ZeroTermsQuery                  *combinedfieldszeroterms.CombinedFieldsZeroTerms `json:"zero_terms_query,omitempty"`
}

// NewCombinedFieldsQuery returns a CombinedFieldsQuery.
func NewCombinedFieldsQuery() *CombinedFieldsQuery {
	r := &CombinedFieldsQuery{}

	return r
}
