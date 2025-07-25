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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _greekAnalyzer struct {
	v *types.GreekAnalyzer
}

func NewGreekAnalyzer() *_greekAnalyzer {

	return &_greekAnalyzer{v: types.NewGreekAnalyzer()}

}

func (s *_greekAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_greekAnalyzer {

	s.v.Stopwords = *stopwords.StopWordsCaster()

	return s
}

func (s *_greekAnalyzer) StopwordsPath(stopwordspath string) *_greekAnalyzer {

	s.v.StopwordsPath = &stopwordspath

	return s
}

func (s *_greekAnalyzer) GreekAnalyzerCaster() *types.GreekAnalyzer {
	return s.v
}
