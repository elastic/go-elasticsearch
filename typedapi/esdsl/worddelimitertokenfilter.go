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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _wordDelimiterTokenFilter struct {
	v *types.WordDelimiterTokenFilter
}

func NewWordDelimiterTokenFilter() *_wordDelimiterTokenFilter {

	return &_wordDelimiterTokenFilter{v: types.NewWordDelimiterTokenFilter()}

}

func (s *_wordDelimiterTokenFilter) CatenateAll(catenateall bool) *_wordDelimiterTokenFilter {

	s.v.CatenateAll = &catenateall

	return s
}

func (s *_wordDelimiterTokenFilter) CatenateNumbers(catenatenumbers bool) *_wordDelimiterTokenFilter {

	s.v.CatenateNumbers = &catenatenumbers

	return s
}

func (s *_wordDelimiterTokenFilter) CatenateWords(catenatewords bool) *_wordDelimiterTokenFilter {

	s.v.CatenateWords = &catenatewords

	return s
}

func (s *_wordDelimiterTokenFilter) GenerateNumberParts(generatenumberparts bool) *_wordDelimiterTokenFilter {

	s.v.GenerateNumberParts = &generatenumberparts

	return s
}

func (s *_wordDelimiterTokenFilter) GenerateWordParts(generatewordparts bool) *_wordDelimiterTokenFilter {

	s.v.GenerateWordParts = &generatewordparts

	return s
}

func (s *_wordDelimiterTokenFilter) PreserveOriginal(stringifiedboolean types.StringifiedbooleanVariant) *_wordDelimiterTokenFilter {

	s.v.PreserveOriginal = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_wordDelimiterTokenFilter) ProtectedWords(protectedwords ...string) *_wordDelimiterTokenFilter {

	for _, v := range protectedwords {

		s.v.ProtectedWords = append(s.v.ProtectedWords, v)

	}
	return s
}

func (s *_wordDelimiterTokenFilter) ProtectedWordsPath(protectedwordspath string) *_wordDelimiterTokenFilter {

	s.v.ProtectedWordsPath = &protectedwordspath

	return s
}

func (s *_wordDelimiterTokenFilter) SplitOnCaseChange(splitoncasechange bool) *_wordDelimiterTokenFilter {

	s.v.SplitOnCaseChange = &splitoncasechange

	return s
}

func (s *_wordDelimiterTokenFilter) SplitOnNumerics(splitonnumerics bool) *_wordDelimiterTokenFilter {

	s.v.SplitOnNumerics = &splitonnumerics

	return s
}

func (s *_wordDelimiterTokenFilter) StemEnglishPossessive(stemenglishpossessive bool) *_wordDelimiterTokenFilter {

	s.v.StemEnglishPossessive = &stemenglishpossessive

	return s
}

func (s *_wordDelimiterTokenFilter) TypeTable(typetables ...string) *_wordDelimiterTokenFilter {

	for _, v := range typetables {

		s.v.TypeTable = append(s.v.TypeTable, v)

	}
	return s
}

func (s *_wordDelimiterTokenFilter) TypeTablePath(typetablepath string) *_wordDelimiterTokenFilter {

	s.v.TypeTablePath = &typetablepath

	return s
}

func (s *_wordDelimiterTokenFilter) Version(versionstring string) *_wordDelimiterTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_wordDelimiterTokenFilter) WordDelimiterTokenFilterCaster() *types.WordDelimiterTokenFilter {
	return s.v
}
