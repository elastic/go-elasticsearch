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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _moreLikeThisQuery struct {
	v *types.MoreLikeThisQuery
}

// Returns documents that are "like" a given set of documents.
func NewMoreLikeThisQuery() *_moreLikeThisQuery {

	return &_moreLikeThisQuery{v: types.NewMoreLikeThisQuery()}

}

func (s *_moreLikeThisQuery) Analyzer(analyzer string) *_moreLikeThisQuery {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_moreLikeThisQuery) Boost(boost float32) *_moreLikeThisQuery {

	s.v.Boost = &boost

	return s
}

func (s *_moreLikeThisQuery) BoostTerms(boostterms types.Float64) *_moreLikeThisQuery {

	s.v.BoostTerms = &boostterms

	return s
}

func (s *_moreLikeThisQuery) FailOnUnsupportedField(failonunsupportedfield bool) *_moreLikeThisQuery {

	s.v.FailOnUnsupportedField = &failonunsupportedfield

	return s
}

func (s *_moreLikeThisQuery) Fields(fields ...string) *_moreLikeThisQuery {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

func (s *_moreLikeThisQuery) Include(include bool) *_moreLikeThisQuery {

	s.v.Include = &include

	return s
}

func (s *_moreLikeThisQuery) Like(likes ...types.LikeVariant) *_moreLikeThisQuery {

	s.v.Like = make([]types.Like, len(likes))
	for i, v := range likes {
		s.v.Like[i] = *v.LikeCaster()
	}

	return s
}

func (s *_moreLikeThisQuery) MaxDocFreq(maxdocfreq int) *_moreLikeThisQuery {

	s.v.MaxDocFreq = &maxdocfreq

	return s
}

func (s *_moreLikeThisQuery) MaxQueryTerms(maxqueryterms int) *_moreLikeThisQuery {

	s.v.MaxQueryTerms = &maxqueryterms

	return s
}

func (s *_moreLikeThisQuery) MaxWordLength(maxwordlength int) *_moreLikeThisQuery {

	s.v.MaxWordLength = &maxwordlength

	return s
}

func (s *_moreLikeThisQuery) MinDocFreq(mindocfreq int) *_moreLikeThisQuery {

	s.v.MinDocFreq = &mindocfreq

	return s
}

func (s *_moreLikeThisQuery) MinTermFreq(mintermfreq int) *_moreLikeThisQuery {

	s.v.MinTermFreq = &mintermfreq

	return s
}

func (s *_moreLikeThisQuery) MinWordLength(minwordlength int) *_moreLikeThisQuery {

	s.v.MinWordLength = &minwordlength

	return s
}

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

func (s *_moreLikeThisQuery) StopWords(stopwords types.StopWordsVariant) *_moreLikeThisQuery {

	s.v.StopWords = *stopwords.StopWordsCaster()

	return s
}

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
