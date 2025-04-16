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

type _indexSettingsAnalysis struct {
	v *types.IndexSettingsAnalysis
}

func NewIndexSettingsAnalysis() *_indexSettingsAnalysis {

	return &_indexSettingsAnalysis{v: types.NewIndexSettingsAnalysis()}

}

func (s *_indexSettingsAnalysis) Analyzer(analyzer map[string]types.Analyzer) *_indexSettingsAnalysis {

	s.v.Analyzer = analyzer
	return s
}

func (s *_indexSettingsAnalysis) AddAnalyzer(key string, value types.AnalyzerVariant) *_indexSettingsAnalysis {

	var tmp map[string]types.Analyzer
	if s.v.Analyzer == nil {
		s.v.Analyzer = make(map[string]types.Analyzer)
	} else {
		tmp = s.v.Analyzer
	}

	tmp[key] = *value.AnalyzerCaster()

	s.v.Analyzer = tmp
	return s
}

func (s *_indexSettingsAnalysis) CharFilter(charfilter map[string]types.CharFilter) *_indexSettingsAnalysis {

	s.v.CharFilter = charfilter
	return s
}

func (s *_indexSettingsAnalysis) AddCharFilter(key string, value types.CharFilterVariant) *_indexSettingsAnalysis {

	var tmp map[string]types.CharFilter
	if s.v.CharFilter == nil {
		s.v.CharFilter = make(map[string]types.CharFilter)
	} else {
		tmp = s.v.CharFilter
	}

	tmp[key] = *value.CharFilterCaster()

	s.v.CharFilter = tmp
	return s
}

func (s *_indexSettingsAnalysis) Filter(filter map[string]types.TokenFilter) *_indexSettingsAnalysis {

	s.v.Filter = filter
	return s
}

func (s *_indexSettingsAnalysis) AddFilter(key string, value types.TokenFilterVariant) *_indexSettingsAnalysis {

	var tmp map[string]types.TokenFilter
	if s.v.Filter == nil {
		s.v.Filter = make(map[string]types.TokenFilter)
	} else {
		tmp = s.v.Filter
	}

	tmp[key] = *value.TokenFilterCaster()

	s.v.Filter = tmp
	return s
}

func (s *_indexSettingsAnalysis) Normalizer(normalizer map[string]types.Normalizer) *_indexSettingsAnalysis {

	s.v.Normalizer = normalizer
	return s
}

func (s *_indexSettingsAnalysis) AddNormalizer(key string, value types.NormalizerVariant) *_indexSettingsAnalysis {

	var tmp map[string]types.Normalizer
	if s.v.Normalizer == nil {
		s.v.Normalizer = make(map[string]types.Normalizer)
	} else {
		tmp = s.v.Normalizer
	}

	tmp[key] = *value.NormalizerCaster()

	s.v.Normalizer = tmp
	return s
}

func (s *_indexSettingsAnalysis) Tokenizer(tokenizer map[string]types.Tokenizer) *_indexSettingsAnalysis {

	s.v.Tokenizer = tokenizer
	return s
}

func (s *_indexSettingsAnalysis) AddTokenizer(key string, value types.TokenizerVariant) *_indexSettingsAnalysis {

	var tmp map[string]types.Tokenizer
	if s.v.Tokenizer == nil {
		s.v.Tokenizer = make(map[string]types.Tokenizer)
	} else {
		tmp = s.v.Tokenizer
	}

	tmp[key] = *value.TokenizerCaster()

	s.v.Tokenizer = tmp
	return s
}

func (s *_indexSettingsAnalysis) IndexSettingsAnalysisCaster() *types.IndexSettingsAnalysis {
	return s.v
}
