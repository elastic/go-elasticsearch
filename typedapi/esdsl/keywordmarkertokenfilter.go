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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _keywordMarkerTokenFilter struct {
	v *types.KeywordMarkerTokenFilter
}

func NewKeywordMarkerTokenFilter() *_keywordMarkerTokenFilter {

	return &_keywordMarkerTokenFilter{v: types.NewKeywordMarkerTokenFilter()}

}

func (s *_keywordMarkerTokenFilter) IgnoreCase(ignorecase bool) *_keywordMarkerTokenFilter {

	s.v.IgnoreCase = &ignorecase

	return s
}

func (s *_keywordMarkerTokenFilter) Keywords(keywords ...string) *_keywordMarkerTokenFilter {

	s.v.Keywords = make([]string, len(keywords))
	s.v.Keywords = keywords

	return s
}

func (s *_keywordMarkerTokenFilter) KeywordsPath(keywordspath string) *_keywordMarkerTokenFilter {

	s.v.KeywordsPath = &keywordspath

	return s
}

func (s *_keywordMarkerTokenFilter) KeywordsPattern(keywordspattern string) *_keywordMarkerTokenFilter {

	s.v.KeywordsPattern = &keywordspattern

	return s
}

func (s *_keywordMarkerTokenFilter) Version(versionstring string) *_keywordMarkerTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_keywordMarkerTokenFilter) KeywordMarkerTokenFilterCaster() *types.KeywordMarkerTokenFilter {
	return s.v
}
