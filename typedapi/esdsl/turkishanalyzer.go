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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _turkishAnalyzer struct {
	v *types.TurkishAnalyzer
}

func NewTurkishAnalyzer() *_turkishAnalyzer {

	return &_turkishAnalyzer{v: types.NewTurkishAnalyzer()}

}

func (s *_turkishAnalyzer) StemExclusion(stemexclusions ...string) *_turkishAnalyzer {

	for _, v := range stemexclusions {

		s.v.StemExclusion = append(s.v.StemExclusion, v)

	}
	return s
}

func (s *_turkishAnalyzer) Stopwords(stopwords ...string) *_turkishAnalyzer {

	s.v.Stopwords = stopwords

	return s
}

func (s *_turkishAnalyzer) StopwordsPath(stopwordspath string) *_turkishAnalyzer {

	s.v.StopwordsPath = &stopwordspath

	return s
}

func (s *_turkishAnalyzer) TurkishAnalyzerCaster() *types.TurkishAnalyzer {
	return s.v
}
