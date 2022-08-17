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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

// MatchPhraseQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L173-L180
type MatchPhraseQuery struct {
	Analyzer       *string                        `json:"analyzer,omitempty"`
	Boost          *float32                       `json:"boost,omitempty"`
	Query          string                         `json:"query"`
	QueryName_     *string                        `json:"_name,omitempty"`
	Slop           *int                           `json:"slop,omitempty"`
	ZeroTermsQuery *zerotermsquery.ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

// MatchPhraseQueryBuilder holds MatchPhraseQuery struct and provides a builder API.
type MatchPhraseQueryBuilder struct {
	v *MatchPhraseQuery
}

// NewMatchPhraseQuery provides a builder for the MatchPhraseQuery struct.
func NewMatchPhraseQueryBuilder() *MatchPhraseQueryBuilder {
	r := MatchPhraseQueryBuilder{
		&MatchPhraseQuery{},
	}

	return &r
}

// Build finalize the chain and returns the MatchPhraseQuery struct
func (rb *MatchPhraseQueryBuilder) Build() MatchPhraseQuery {
	return *rb.v
}

func (rb *MatchPhraseQueryBuilder) Analyzer(analyzer string) *MatchPhraseQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *MatchPhraseQueryBuilder) Boost(boost float32) *MatchPhraseQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *MatchPhraseQueryBuilder) Query(query string) *MatchPhraseQueryBuilder {
	rb.v.Query = query
	return rb
}

func (rb *MatchPhraseQueryBuilder) QueryName_(queryname_ string) *MatchPhraseQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *MatchPhraseQueryBuilder) Slop(slop int) *MatchPhraseQueryBuilder {
	rb.v.Slop = &slop
	return rb
}

func (rb *MatchPhraseQueryBuilder) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *MatchPhraseQueryBuilder {
	rb.v.ZeroTermsQuery = &zerotermsquery
	return rb
}
