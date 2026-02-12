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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _romanianAnalyzer struct {
	v *types.RomanianAnalyzer
}

func NewRomanianAnalyzer() *_romanianAnalyzer {

	return &_romanianAnalyzer{v: types.NewRomanianAnalyzer()}

}

func (s *_romanianAnalyzer) StemExclusion(stemexclusions ...string) *_romanianAnalyzer {

	for _, v := range stemexclusions {

		s.v.StemExclusion = append(s.v.StemExclusion, v)

	}
	return s
}

func (s *_romanianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_romanianAnalyzer {

	s.v.Stopwords = *stopwords.StopWordsCaster()

	return s
}

func (s *_romanianAnalyzer) StopwordsPath(stopwordspath string) *_romanianAnalyzer {

	s.v.StopwordsPath = &stopwordspath

	return s
}

func (s *_romanianAnalyzer) RomanianAnalyzerCaster() *types.RomanianAnalyzer {
	return s.v
}
