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

type _patternAnalyzer struct {
	v *types.PatternAnalyzer
}

func NewPatternAnalyzer() *_patternAnalyzer {

	return &_patternAnalyzer{v: types.NewPatternAnalyzer()}

}

func (s *_patternAnalyzer) Flags(flags string) *_patternAnalyzer {

	s.v.Flags = &flags

	return s
}

func (s *_patternAnalyzer) Lowercase(lowercase bool) *_patternAnalyzer {

	s.v.Lowercase = &lowercase

	return s
}

func (s *_patternAnalyzer) Pattern(pattern string) *_patternAnalyzer {

	s.v.Pattern = &pattern

	return s
}

func (s *_patternAnalyzer) Stopwords(stopwords ...string) *_patternAnalyzer {

	s.v.Stopwords = stopwords

	return s
}

func (s *_patternAnalyzer) StopwordsPath(stopwordspath string) *_patternAnalyzer {

	s.v.StopwordsPath = &stopwordspath

	return s
}

func (s *_patternAnalyzer) Version(versionstring string) *_patternAnalyzer {

	s.v.Version = &versionstring

	return s
}

func (s *_patternAnalyzer) PatternAnalyzerCaster() *types.PatternAnalyzer {
	return s.v
}
