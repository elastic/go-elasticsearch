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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// MoreLikeThisQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/specialized.ts#L62-L89
type MoreLikeThisQuery struct {
	Analyzer               *string                  `json:"analyzer,omitempty"`
	Boost                  *float32                 `json:"boost,omitempty"`
	BoostTerms             *float64                 `json:"boost_terms,omitempty"`
	FailOnUnsupportedField *bool                    `json:"fail_on_unsupported_field,omitempty"`
	Fields                 []Field                  `json:"fields,omitempty"`
	Include                *bool                    `json:"include,omitempty"`
	Like                   []Like                   `json:"like"`
	MaxDocFreq             *int                     `json:"max_doc_freq,omitempty"`
	MaxQueryTerms          *int                     `json:"max_query_terms,omitempty"`
	MaxWordLength          *int                     `json:"max_word_length,omitempty"`
	MinDocFreq             *int                     `json:"min_doc_freq,omitempty"`
	MinTermFreq            *int                     `json:"min_term_freq,omitempty"`
	MinWordLength          *int                     `json:"min_word_length,omitempty"`
	MinimumShouldMatch     *MinimumShouldMatch      `json:"minimum_should_match,omitempty"`
	PerFieldAnalyzer       map[Field]string         `json:"per_field_analyzer,omitempty"`
	QueryName_             *string                  `json:"_name,omitempty"`
	Routing                *Routing                 `json:"routing,omitempty"`
	StopWords              *StopWords               `json:"stop_words,omitempty"`
	Unlike                 []Like                   `json:"unlike,omitempty"`
	Version                *VersionNumber           `json:"version,omitempty"`
	VersionType            *versiontype.VersionType `json:"version_type,omitempty"`
}

// MoreLikeThisQueryBuilder holds MoreLikeThisQuery struct and provides a builder API.
type MoreLikeThisQueryBuilder struct {
	v *MoreLikeThisQuery
}

// NewMoreLikeThisQuery provides a builder for the MoreLikeThisQuery struct.
func NewMoreLikeThisQueryBuilder() *MoreLikeThisQueryBuilder {
	r := MoreLikeThisQueryBuilder{
		&MoreLikeThisQuery{
			PerFieldAnalyzer: make(map[Field]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MoreLikeThisQuery struct
func (rb *MoreLikeThisQueryBuilder) Build() MoreLikeThisQuery {
	return *rb.v
}

func (rb *MoreLikeThisQueryBuilder) Analyzer(analyzer string) *MoreLikeThisQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *MoreLikeThisQueryBuilder) Boost(boost float32) *MoreLikeThisQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *MoreLikeThisQueryBuilder) BoostTerms(boostterms float64) *MoreLikeThisQueryBuilder {
	rb.v.BoostTerms = &boostterms
	return rb
}

func (rb *MoreLikeThisQueryBuilder) FailOnUnsupportedField(failonunsupportedfield bool) *MoreLikeThisQueryBuilder {
	rb.v.FailOnUnsupportedField = &failonunsupportedfield
	return rb
}

func (rb *MoreLikeThisQueryBuilder) Fields(fields ...Field) *MoreLikeThisQueryBuilder {
	rb.v.Fields = fields
	return rb
}

func (rb *MoreLikeThisQueryBuilder) Include(include bool) *MoreLikeThisQueryBuilder {
	rb.v.Include = &include
	return rb
}

func (rb *MoreLikeThisQueryBuilder) Like(arg []Like) *MoreLikeThisQueryBuilder {
	rb.v.Like = arg
	return rb
}

func (rb *MoreLikeThisQueryBuilder) MaxDocFreq(maxdocfreq int) *MoreLikeThisQueryBuilder {
	rb.v.MaxDocFreq = &maxdocfreq
	return rb
}

func (rb *MoreLikeThisQueryBuilder) MaxQueryTerms(maxqueryterms int) *MoreLikeThisQueryBuilder {
	rb.v.MaxQueryTerms = &maxqueryterms
	return rb
}

func (rb *MoreLikeThisQueryBuilder) MaxWordLength(maxwordlength int) *MoreLikeThisQueryBuilder {
	rb.v.MaxWordLength = &maxwordlength
	return rb
}

func (rb *MoreLikeThisQueryBuilder) MinDocFreq(mindocfreq int) *MoreLikeThisQueryBuilder {
	rb.v.MinDocFreq = &mindocfreq
	return rb
}

func (rb *MoreLikeThisQueryBuilder) MinTermFreq(mintermfreq int) *MoreLikeThisQueryBuilder {
	rb.v.MinTermFreq = &mintermfreq
	return rb
}

func (rb *MoreLikeThisQueryBuilder) MinWordLength(minwordlength int) *MoreLikeThisQueryBuilder {
	rb.v.MinWordLength = &minwordlength
	return rb
}

func (rb *MoreLikeThisQueryBuilder) MinimumShouldMatch(minimumshouldmatch *MinimumShouldMatchBuilder) *MoreLikeThisQueryBuilder {
	v := minimumshouldmatch.Build()
	rb.v.MinimumShouldMatch = &v
	return rb
}

func (rb *MoreLikeThisQueryBuilder) PerFieldAnalyzer(value map[Field]string) *MoreLikeThisQueryBuilder {
	rb.v.PerFieldAnalyzer = value
	return rb
}

func (rb *MoreLikeThisQueryBuilder) QueryName_(queryname_ string) *MoreLikeThisQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *MoreLikeThisQueryBuilder) Routing(routing Routing) *MoreLikeThisQueryBuilder {
	rb.v.Routing = &routing
	return rb
}

func (rb *MoreLikeThisQueryBuilder) StopWords(stopwords *StopWordsBuilder) *MoreLikeThisQueryBuilder {
	v := stopwords.Build()
	rb.v.StopWords = &v
	return rb
}

func (rb *MoreLikeThisQueryBuilder) Unlike(arg []Like) *MoreLikeThisQueryBuilder {
	rb.v.Unlike = arg
	return rb
}

func (rb *MoreLikeThisQueryBuilder) Version(version VersionNumber) *MoreLikeThisQueryBuilder {
	rb.v.Version = &version
	return rb
}

func (rb *MoreLikeThisQueryBuilder) VersionType(versiontype versiontype.VersionType) *MoreLikeThisQueryBuilder {
	rb.v.VersionType = &versiontype
	return rb
}
