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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/combinedfieldsoperator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/combinedfieldszeroterms"
)

// CombinedFieldsQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/abstractions.ts#L181-L195
type CombinedFieldsQuery struct {
	AutoGenerateSynonymsPhraseQuery *bool                                            `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float32                                         `json:"boost,omitempty"`
	Fields                          []Field                                          `json:"fields"`
	MinimumShouldMatch              *MinimumShouldMatch                              `json:"minimum_should_match,omitempty"`
	Operator                        *combinedfieldsoperator.CombinedFieldsOperator   `json:"operator,omitempty"`
	Query                           string                                           `json:"query"`
	QueryName_                      *string                                          `json:"_name,omitempty"`
	ZeroTermsQuery                  *combinedfieldszeroterms.CombinedFieldsZeroTerms `json:"zero_terms_query,omitempty"`
}

// CombinedFieldsQueryBuilder holds CombinedFieldsQuery struct and provides a builder API.
type CombinedFieldsQueryBuilder struct {
	v *CombinedFieldsQuery
}

// NewCombinedFieldsQuery provides a builder for the CombinedFieldsQuery struct.
func NewCombinedFieldsQueryBuilder() *CombinedFieldsQueryBuilder {
	r := CombinedFieldsQueryBuilder{
		&CombinedFieldsQuery{},
	}

	return &r
}

// Build finalize the chain and returns the CombinedFieldsQuery struct
func (rb *CombinedFieldsQueryBuilder) Build() CombinedFieldsQuery {
	return *rb.v
}

func (rb *CombinedFieldsQueryBuilder) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *CombinedFieldsQueryBuilder {
	rb.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery
	return rb
}

func (rb *CombinedFieldsQueryBuilder) Boost(boost float32) *CombinedFieldsQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *CombinedFieldsQueryBuilder) Fields(fields ...Field) *CombinedFieldsQueryBuilder {
	rb.v.Fields = fields
	return rb
}

func (rb *CombinedFieldsQueryBuilder) MinimumShouldMatch(minimumshouldmatch *MinimumShouldMatchBuilder) *CombinedFieldsQueryBuilder {
	v := minimumshouldmatch.Build()
	rb.v.MinimumShouldMatch = &v
	return rb
}

func (rb *CombinedFieldsQueryBuilder) Operator(operator combinedfieldsoperator.CombinedFieldsOperator) *CombinedFieldsQueryBuilder {
	rb.v.Operator = &operator
	return rb
}

func (rb *CombinedFieldsQueryBuilder) Query(query string) *CombinedFieldsQueryBuilder {
	rb.v.Query = query
	return rb
}

func (rb *CombinedFieldsQueryBuilder) QueryName_(queryname_ string) *CombinedFieldsQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *CombinedFieldsQueryBuilder) ZeroTermsQuery(zerotermsquery combinedfieldszeroterms.CombinedFieldsZeroTerms) *CombinedFieldsQueryBuilder {
	rb.v.ZeroTermsQuery = &zerotermsquery
	return rb
}
