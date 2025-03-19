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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _dataframeAnalysisFeatureProcessorNGramEncoding struct {
	v *types.DataframeAnalysisFeatureProcessorNGramEncoding
}

// The configuration information necessary to perform n-gram encoding. Features
// created by this encoder have the following name format:
// <feature_prefix>.<ngram><string position>. For example, if the feature_prefix
// is f, the feature name for the second unigram in a string is f.11.
func NewDataframeAnalysisFeatureProcessorNGramEncoding() *_dataframeAnalysisFeatureProcessorNGramEncoding {

	return &_dataframeAnalysisFeatureProcessorNGramEncoding{v: types.NewDataframeAnalysisFeatureProcessorNGramEncoding()}

}

func (s *_dataframeAnalysisFeatureProcessorNGramEncoding) Custom(custom bool) *_dataframeAnalysisFeatureProcessorNGramEncoding {

	s.v.Custom = &custom

	return s
}

// The feature name prefix. Defaults to ngram_<start>_<length>.
func (s *_dataframeAnalysisFeatureProcessorNGramEncoding) FeaturePrefix(featureprefix string) *_dataframeAnalysisFeatureProcessorNGramEncoding {

	s.v.FeaturePrefix = &featureprefix

	return s
}

// The name of the text field to encode.
func (s *_dataframeAnalysisFeatureProcessorNGramEncoding) Field(field string) *_dataframeAnalysisFeatureProcessorNGramEncoding {

	s.v.Field = field

	return s
}

// Specifies the length of the n-gram substring. Defaults to 50. Must be greater
// than 0.
func (s *_dataframeAnalysisFeatureProcessorNGramEncoding) Length(length int) *_dataframeAnalysisFeatureProcessorNGramEncoding {

	s.v.Length = &length

	return s
}

// Specifies which n-grams to gather. It’s an array of integer values where the
// minimum value is 1, and a maximum value is 5.
func (s *_dataframeAnalysisFeatureProcessorNGramEncoding) NGrams(ngrams ...int) *_dataframeAnalysisFeatureProcessorNGramEncoding {

	for _, v := range ngrams {

		s.v.NGrams = append(s.v.NGrams, v)

	}
	return s
}

// Specifies the zero-indexed start of the n-gram substring. Negative values are
// allowed for encoding n-grams of string suffixes. Defaults to 0.
func (s *_dataframeAnalysisFeatureProcessorNGramEncoding) Start(start int) *_dataframeAnalysisFeatureProcessorNGramEncoding {

	s.v.Start = &start

	return s
}

func (s *_dataframeAnalysisFeatureProcessorNGramEncoding) DataframeAnalysisFeatureProcessorCaster() *types.DataframeAnalysisFeatureProcessor {
	container := types.NewDataframeAnalysisFeatureProcessor()

	container.NGramEncoding = s.v

	return container
}

func (s *_dataframeAnalysisFeatureProcessorNGramEncoding) DataframeAnalysisFeatureProcessorNGramEncodingCaster() *types.DataframeAnalysisFeatureProcessorNGramEncoding {
	return s.v
}
