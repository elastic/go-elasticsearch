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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _categorizationAnalyzerDefinition struct {
	v *types.CategorizationAnalyzerDefinition
}

func NewCategorizationAnalyzerDefinition() *_categorizationAnalyzerDefinition {

	return &_categorizationAnalyzerDefinition{v: types.NewCategorizationAnalyzerDefinition()}

}

func (s *_categorizationAnalyzerDefinition) CharFilter(charfilters ...types.CharFilterVariant) *_categorizationAnalyzerDefinition {

	for _, v := range charfilters {

		s.v.CharFilter = append(s.v.CharFilter, *v.CharFilterCaster())

	}
	return s
}

func (s *_categorizationAnalyzerDefinition) Filter(filters ...types.TokenFilterVariant) *_categorizationAnalyzerDefinition {

	for _, v := range filters {

		s.v.Filter = append(s.v.Filter, *v.TokenFilterCaster())

	}
	return s
}

func (s *_categorizationAnalyzerDefinition) Tokenizer(tokenizer types.TokenizerVariant) *_categorizationAnalyzerDefinition {

	s.v.Tokenizer = *tokenizer.TokenizerCaster()

	return s
}

func (s *_categorizationAnalyzerDefinition) CategorizationAnalyzerDefinitionCaster() *types.CategorizationAnalyzerDefinition {
	return s.v
}
