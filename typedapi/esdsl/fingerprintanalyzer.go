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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fingerprintAnalyzer struct {
	v *types.FingerprintAnalyzer
}

func NewFingerprintAnalyzer() *_fingerprintAnalyzer {

	return &_fingerprintAnalyzer{v: types.NewFingerprintAnalyzer()}

}

func (s *_fingerprintAnalyzer) MaxOutputSize(maxoutputsize int) *_fingerprintAnalyzer {

	s.v.MaxOutputSize = &maxoutputsize

	return s
}

func (s *_fingerprintAnalyzer) Separator(separator string) *_fingerprintAnalyzer {

	s.v.Separator = &separator

	return s
}

func (s *_fingerprintAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_fingerprintAnalyzer {

	s.v.Stopwords = *stopwords.StopWordsCaster()

	return s
}

func (s *_fingerprintAnalyzer) StopwordsPath(stopwordspath string) *_fingerprintAnalyzer {

	s.v.StopwordsPath = &stopwordspath

	return s
}

func (s *_fingerprintAnalyzer) Version(versionstring string) *_fingerprintAnalyzer {

	s.v.Version = &versionstring

	return s
}

func (s *_fingerprintAnalyzer) FingerprintAnalyzerCaster() *types.FingerprintAnalyzer {
	return s.v
}
