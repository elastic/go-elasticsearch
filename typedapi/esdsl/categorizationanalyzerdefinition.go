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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _categorizationAnalyzerDefinition struct {
	v *types.CategorizationAnalyzerDefinition
}

func NewCategorizationAnalyzerDefinition() *_categorizationAnalyzerDefinition {

	return &_categorizationAnalyzerDefinition{v: types.NewCategorizationAnalyzerDefinition()}

}

// One or more character filters. In addition to the built-in character filters,
// other plugins can provide more character filters. If this property is not
// specified, no character filters are applied prior to categorization. If you
// are customizing some other aspect of the analyzer and you need to achieve the
// equivalent of `categorization_filters` (which are not permitted when some
// other aspect of the analyzer is customized), add them here as pattern replace
// character filters.
func (s *_categorizationAnalyzerDefinition) CharFilter(charfilters ...types.CharFilterVariant) *_categorizationAnalyzerDefinition {

	for _, v := range charfilters {

		s.v.CharFilter = append(s.v.CharFilter, *v.CharFilterCaster())

	}
	return s
}

// One or more token filters. In addition to the built-in token filters, other
// plugins can provide more token filters. If this property is not specified, no
// token filters are applied prior to categorization.
func (s *_categorizationAnalyzerDefinition) Filter(filters ...types.TokenFilterVariant) *_categorizationAnalyzerDefinition {

	for _, v := range filters {

		s.v.Filter = append(s.v.Filter, *v.TokenFilterCaster())

	}
	return s
}

// The name or definition of the tokenizer to use after character filters are
// applied. This property is compulsory if `categorization_analyzer` is
// specified as an object. Machine learning provides a tokenizer called
// `ml_standard` that tokenizes in a way that has been determined to produce
// good categorization results on a variety of log file formats for logs in
// English. If you want to use that tokenizer but change the character or token
// filters, specify "tokenizer": "ml_standard" in your
// `categorization_analyzer`. Additionally, the `ml_classic` tokenizer is
// available, which tokenizes in the same way as the non-customizable tokenizer
// in old versions of the product (before 6.2). `ml_classic` was the default
// categorization tokenizer in versions 6.2 to 7.13, so if you need
// categorization identical to the default for jobs created in these versions,
// specify "tokenizer": "ml_classic" in your `categorization_analyzer`.
func (s *_categorizationAnalyzerDefinition) Tokenizer(tokenizer types.TokenizerVariant) *_categorizationAnalyzerDefinition {

	s.v.Tokenizer = *tokenizer.TokenizerCaster()

	return s
}

func (s *_categorizationAnalyzerDefinition) CategorizationAnalyzerDefinitionCaster() *types.CategorizationAnalyzerDefinition {
	return s.v
}
