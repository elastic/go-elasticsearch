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

type _wordDelimiterGraphTokenFilter struct {
	v *types.WordDelimiterGraphTokenFilter
}

func NewWordDelimiterGraphTokenFilter() *_wordDelimiterGraphTokenFilter {

	return &_wordDelimiterGraphTokenFilter{v: types.NewWordDelimiterGraphTokenFilter()}

}

func (s *_wordDelimiterGraphTokenFilter) AdjustOffsets(adjustoffsets bool) *_wordDelimiterGraphTokenFilter {

	s.v.AdjustOffsets = &adjustoffsets

	return s
}

func (s *_wordDelimiterGraphTokenFilter) CatenateAll(catenateall bool) *_wordDelimiterGraphTokenFilter {

	s.v.CatenateAll = &catenateall

	return s
}

func (s *_wordDelimiterGraphTokenFilter) CatenateNumbers(catenatenumbers bool) *_wordDelimiterGraphTokenFilter {

	s.v.CatenateNumbers = &catenatenumbers

	return s
}

func (s *_wordDelimiterGraphTokenFilter) CatenateWords(catenatewords bool) *_wordDelimiterGraphTokenFilter {

	s.v.CatenateWords = &catenatewords

	return s
}

func (s *_wordDelimiterGraphTokenFilter) GenerateNumberParts(generatenumberparts bool) *_wordDelimiterGraphTokenFilter {

	s.v.GenerateNumberParts = &generatenumberparts

	return s
}

func (s *_wordDelimiterGraphTokenFilter) GenerateWordParts(generatewordparts bool) *_wordDelimiterGraphTokenFilter {

	s.v.GenerateWordParts = &generatewordparts

	return s
}

func (s *_wordDelimiterGraphTokenFilter) IgnoreKeywords(ignorekeywords bool) *_wordDelimiterGraphTokenFilter {

	s.v.IgnoreKeywords = &ignorekeywords

	return s
}

func (s *_wordDelimiterGraphTokenFilter) PreserveOriginal(stringifiedboolean types.StringifiedbooleanVariant) *_wordDelimiterGraphTokenFilter {

	s.v.PreserveOriginal = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_wordDelimiterGraphTokenFilter) ProtectedWords(protectedwords ...string) *_wordDelimiterGraphTokenFilter {

	for _, v := range protectedwords {

		s.v.ProtectedWords = append(s.v.ProtectedWords, v)

	}
	return s
}

func (s *_wordDelimiterGraphTokenFilter) ProtectedWordsPath(protectedwordspath string) *_wordDelimiterGraphTokenFilter {

	s.v.ProtectedWordsPath = &protectedwordspath

	return s
}

func (s *_wordDelimiterGraphTokenFilter) SplitOnCaseChange(splitoncasechange bool) *_wordDelimiterGraphTokenFilter {

	s.v.SplitOnCaseChange = &splitoncasechange

	return s
}

func (s *_wordDelimiterGraphTokenFilter) SplitOnNumerics(splitonnumerics bool) *_wordDelimiterGraphTokenFilter {

	s.v.SplitOnNumerics = &splitonnumerics

	return s
}

func (s *_wordDelimiterGraphTokenFilter) StemEnglishPossessive(stemenglishpossessive bool) *_wordDelimiterGraphTokenFilter {

	s.v.StemEnglishPossessive = &stemenglishpossessive

	return s
}

func (s *_wordDelimiterGraphTokenFilter) TypeTable(typetables ...string) *_wordDelimiterGraphTokenFilter {

	for _, v := range typetables {

		s.v.TypeTable = append(s.v.TypeTable, v)

	}
	return s
}

func (s *_wordDelimiterGraphTokenFilter) TypeTablePath(typetablepath string) *_wordDelimiterGraphTokenFilter {

	s.v.TypeTablePath = &typetablepath

	return s
}

func (s *_wordDelimiterGraphTokenFilter) Version(versionstring string) *_wordDelimiterGraphTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_wordDelimiterGraphTokenFilter) WordDelimiterGraphTokenFilterCaster() *types.WordDelimiterGraphTokenFilter {
	return s.v
}
