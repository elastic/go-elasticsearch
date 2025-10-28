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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hyphenationDecompounderTokenFilter struct {
	v *types.HyphenationDecompounderTokenFilter
}

func NewHyphenationDecompounderTokenFilter(hyphenationpatternspath string) *_hyphenationDecompounderTokenFilter {

	tmp := &_hyphenationDecompounderTokenFilter{v: types.NewHyphenationDecompounderTokenFilter()}

	tmp.HyphenationPatternsPath(hyphenationpatternspath)

	return tmp

}

func (s *_hyphenationDecompounderTokenFilter) HyphenationPatternsPath(hyphenationpatternspath string) *_hyphenationDecompounderTokenFilter {

	s.v.HyphenationPatternsPath = hyphenationpatternspath

	return s
}

func (s *_hyphenationDecompounderTokenFilter) MaxSubwordSize(maxsubwordsize int) *_hyphenationDecompounderTokenFilter {

	s.v.MaxSubwordSize = &maxsubwordsize

	return s
}

func (s *_hyphenationDecompounderTokenFilter) MinSubwordSize(minsubwordsize int) *_hyphenationDecompounderTokenFilter {

	s.v.MinSubwordSize = &minsubwordsize

	return s
}

func (s *_hyphenationDecompounderTokenFilter) MinWordSize(minwordsize int) *_hyphenationDecompounderTokenFilter {

	s.v.MinWordSize = &minwordsize

	return s
}

func (s *_hyphenationDecompounderTokenFilter) NoOverlappingMatches(nooverlappingmatches bool) *_hyphenationDecompounderTokenFilter {

	s.v.NoOverlappingMatches = &nooverlappingmatches

	return s
}

func (s *_hyphenationDecompounderTokenFilter) NoSubMatches(nosubmatches bool) *_hyphenationDecompounderTokenFilter {

	s.v.NoSubMatches = &nosubmatches

	return s
}

func (s *_hyphenationDecompounderTokenFilter) OnlyLongestMatch(onlylongestmatch bool) *_hyphenationDecompounderTokenFilter {

	s.v.OnlyLongestMatch = &onlylongestmatch

	return s
}

func (s *_hyphenationDecompounderTokenFilter) Version(versionstring string) *_hyphenationDecompounderTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_hyphenationDecompounderTokenFilter) WordList(wordlists ...string) *_hyphenationDecompounderTokenFilter {

	for _, v := range wordlists {

		s.v.WordList = append(s.v.WordList, v)

	}
	return s
}

func (s *_hyphenationDecompounderTokenFilter) WordListPath(wordlistpath string) *_hyphenationDecompounderTokenFilter {

	s.v.WordListPath = &wordlistpath

	return s
}

func (s *_hyphenationDecompounderTokenFilter) HyphenationDecompounderTokenFilterCaster() *types.HyphenationDecompounderTokenFilter {
	return s.v
}
