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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

type _moreLikeThisQuery struct {
	v *types.MoreLikeThisQuery
}

// Returns documents that are "like" a given set of documents.
func NewMoreLikeThisQuery() *_moreLikeThisQuery {

	return &_moreLikeThisQuery{v: types.NewMoreLikeThisQuery()}

}

// The analyzer that is used to analyze the free form text.
// Defaults to the analyzer associated with the first field in fields.
func (s *_moreLikeThisQuery) Analyzer(analyzer string) *_moreLikeThisQuery {

	s.v.Analyzer = &analyzer

	return s
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_moreLikeThisQuery) Boost(boost float32) *_moreLikeThisQuery {

	s.v.Boost = &boost

	return s
}

// Each term in the formed query could be further boosted by their tf-idf score.
// This sets the boost factor to use when using this feature.
// Defaults to deactivated (0).
func (s *_moreLikeThisQuery) BoostTerms(boostterms types.Float64) *_moreLikeThisQuery {

	s.v.BoostTerms = &boostterms

	return s
}

// Controls whether the query should fail (throw an exception) if any of the
// specified fields are not of the supported types (`text` or `keyword`).
func (s *_moreLikeThisQuery) FailOnUnsupportedField(failonunsupportedfield bool) *_moreLikeThisQuery {

	s.v.FailOnUnsupportedField = &failonunsupportedfield

	return s
}

// A list of fields to fetch and analyze the text from.
// Defaults to the `index.query.default_field` index setting, which has a
// default value of `*`.
func (s *_moreLikeThisQuery) Fields(fields ...string) *_moreLikeThisQuery {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

// Specifies whether the input documents should also be included in the search
// results returned.
func (s *_moreLikeThisQuery) Include(include bool) *_moreLikeThisQuery {

	s.v.Include = &include

	return s
}

// Specifies free form text and/or a single or multiple documents for which you
// want to find similar documents.
func (s *_moreLikeThisQuery) Like(likes ...types.LikeVariant) *_moreLikeThisQuery {

	s.v.Like = make([]types.Like, len(likes))
	for i, v := range likes {
		s.v.Like[i] = *v.LikeCaster()
	}

	return s
}

// The maximum document frequency above which the terms are ignored from the
// input document.
func (s *_moreLikeThisQuery) MaxDocFreq(maxdocfreq int) *_moreLikeThisQuery {

	s.v.MaxDocFreq = &maxdocfreq

	return s
}

// The maximum number of query terms that can be selected.
func (s *_moreLikeThisQuery) MaxQueryTerms(maxqueryterms int) *_moreLikeThisQuery {

	s.v.MaxQueryTerms = &maxqueryterms

	return s
}

// The maximum word length above which the terms are ignored.
// Defaults to unbounded (`0`).
func (s *_moreLikeThisQuery) MaxWordLength(maxwordlength int) *_moreLikeThisQuery {

	s.v.MaxWordLength = &maxwordlength

	return s
}

// The minimum document frequency below which the terms are ignored from the
// input document.
func (s *_moreLikeThisQuery) MinDocFreq(mindocfreq int) *_moreLikeThisQuery {

	s.v.MinDocFreq = &mindocfreq

	return s
}

// The minimum term frequency below which the terms are ignored from the input
// document.
func (s *_moreLikeThisQuery) MinTermFreq(mintermfreq int) *_moreLikeThisQuery {

	s.v.MinTermFreq = &mintermfreq

	return s
}

// The minimum word length below which the terms are ignored.
func (s *_moreLikeThisQuery) MinWordLength(minwordlength int) *_moreLikeThisQuery {

	s.v.MinWordLength = &minwordlength

	return s
}

// After the disjunctive query has been formed, this parameter controls the
// number of terms that must match.
func (s *_moreLikeThisQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_moreLikeThisQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

func (s *_moreLikeThisQuery) QueryName_(queryname_ string) *_moreLikeThisQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_moreLikeThisQuery) Routing(routing string) *_moreLikeThisQuery {

	s.v.Routing = &routing

	return s
}

// An array of stop words.
// Any word in this set is ignored.
func (s *_moreLikeThisQuery) StopWords(stopwords ...string) *_moreLikeThisQuery {

	s.v.StopWords = stopwords

	return s
}

// Used in combination with `like` to exclude documents that match a set of
// terms.
func (s *_moreLikeThisQuery) Unlike(unlikes ...types.LikeVariant) *_moreLikeThisQuery {

	s.v.Unlike = make([]types.Like, len(unlikes))
	for i, v := range unlikes {
		s.v.Unlike[i] = *v.LikeCaster()
	}

	return s
}

func (s *_moreLikeThisQuery) Version(versionnumber int64) *_moreLikeThisQuery {

	s.v.Version = &versionnumber

	return s
}

func (s *_moreLikeThisQuery) VersionType(versiontype versiontype.VersionType) *_moreLikeThisQuery {

	s.v.VersionType = &versiontype
	return s
}

func (s *_moreLikeThisQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.MoreLikeThis = s.v

	return container
}

func (s *_moreLikeThisQuery) MoreLikeThisQueryCaster() *types.MoreLikeThisQuery {
	return s.v
}
