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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dictionaryDecompounderTokenFilter struct {
	v *types.DictionaryDecompounderTokenFilter
}

func NewDictionaryDecompounderTokenFilter() *_dictionaryDecompounderTokenFilter {

	return &_dictionaryDecompounderTokenFilter{v: types.NewDictionaryDecompounderTokenFilter()}

}

func (s *_dictionaryDecompounderTokenFilter) HyphenationPatternsPath(hyphenationpatternspath string) *_dictionaryDecompounderTokenFilter {

	s.v.HyphenationPatternsPath = &hyphenationpatternspath

	return s
}

func (s *_dictionaryDecompounderTokenFilter) MaxSubwordSize(maxsubwordsize int) *_dictionaryDecompounderTokenFilter {

	s.v.MaxSubwordSize = &maxsubwordsize

	return s
}

func (s *_dictionaryDecompounderTokenFilter) MinSubwordSize(minsubwordsize int) *_dictionaryDecompounderTokenFilter {

	s.v.MinSubwordSize = &minsubwordsize

	return s
}

func (s *_dictionaryDecompounderTokenFilter) MinWordSize(minwordsize int) *_dictionaryDecompounderTokenFilter {

	s.v.MinWordSize = &minwordsize

	return s
}

func (s *_dictionaryDecompounderTokenFilter) OnlyLongestMatch(onlylongestmatch bool) *_dictionaryDecompounderTokenFilter {

	s.v.OnlyLongestMatch = &onlylongestmatch

	return s
}

func (s *_dictionaryDecompounderTokenFilter) Version(versionstring string) *_dictionaryDecompounderTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_dictionaryDecompounderTokenFilter) WordList(wordlists ...string) *_dictionaryDecompounderTokenFilter {

	for _, v := range wordlists {

		s.v.WordList = append(s.v.WordList, v)

	}
	return s
}

func (s *_dictionaryDecompounderTokenFilter) WordListPath(wordlistpath string) *_dictionaryDecompounderTokenFilter {

	s.v.WordListPath = &wordlistpath

	return s
}

func (s *_dictionaryDecompounderTokenFilter) DictionaryDecompounderTokenFilterCaster() *types.DictionaryDecompounderTokenFilter {
	return s.v
}
