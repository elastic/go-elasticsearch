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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _bengaliAnalyzer struct {
	v *types.BengaliAnalyzer
}

func NewBengaliAnalyzer() *_bengaliAnalyzer {

	return &_bengaliAnalyzer{v: types.NewBengaliAnalyzer()}

}

func (s *_bengaliAnalyzer) StemExclusion(stemexclusions ...string) *_bengaliAnalyzer {

	for _, v := range stemexclusions {

		s.v.StemExclusion = append(s.v.StemExclusion, v)

	}
	return s
}

func (s *_bengaliAnalyzer) Stopwords(stopwords ...string) *_bengaliAnalyzer {

	s.v.Stopwords = stopwords

	return s
}

func (s *_bengaliAnalyzer) StopwordsPath(stopwordspath string) *_bengaliAnalyzer {

	s.v.StopwordsPath = &stopwordspath

	return s
}

func (s *_bengaliAnalyzer) BengaliAnalyzerCaster() *types.BengaliAnalyzer {
	return s.v
}
