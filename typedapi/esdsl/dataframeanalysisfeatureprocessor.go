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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _dataframeAnalysisFeatureProcessor struct {
	v *types.DataframeAnalysisFeatureProcessor
}

func NewDataframeAnalysisFeatureProcessor() *_dataframeAnalysisFeatureProcessor {
	return &_dataframeAnalysisFeatureProcessor{v: types.NewDataframeAnalysisFeatureProcessor()}
}

// AdditionalDataframeAnalysisFeatureProcessorProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_dataframeAnalysisFeatureProcessor) AdditionalDataframeAnalysisFeatureProcessorProperty(key string, value json.RawMessage) *_dataframeAnalysisFeatureProcessor {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalDataframeAnalysisFeatureProcessorProperty = tmp
	return s
}

func (s *_dataframeAnalysisFeatureProcessor) FrequencyEncoding(frequencyencoding types.DataframeAnalysisFeatureProcessorFrequencyEncodingVariant) *_dataframeAnalysisFeatureProcessor {

	s.v.FrequencyEncoding = frequencyencoding.DataframeAnalysisFeatureProcessorFrequencyEncodingCaster()

	return s
}

func (s *_dataframeAnalysisFeatureProcessor) MultiEncoding(multiencoding types.DataframeAnalysisFeatureProcessorMultiEncodingVariant) *_dataframeAnalysisFeatureProcessor {

	s.v.MultiEncoding = multiencoding.DataframeAnalysisFeatureProcessorMultiEncodingCaster()

	return s
}

func (s *_dataframeAnalysisFeatureProcessor) NGramEncoding(ngramencoding types.DataframeAnalysisFeatureProcessorNGramEncodingVariant) *_dataframeAnalysisFeatureProcessor {

	s.v.NGramEncoding = ngramencoding.DataframeAnalysisFeatureProcessorNGramEncodingCaster()

	return s
}

func (s *_dataframeAnalysisFeatureProcessor) OneHotEncoding(onehotencoding types.DataframeAnalysisFeatureProcessorOneHotEncodingVariant) *_dataframeAnalysisFeatureProcessor {

	s.v.OneHotEncoding = onehotencoding.DataframeAnalysisFeatureProcessorOneHotEncodingCaster()

	return s
}

func (s *_dataframeAnalysisFeatureProcessor) TargetMeanEncoding(targetmeanencoding types.DataframeAnalysisFeatureProcessorTargetMeanEncodingVariant) *_dataframeAnalysisFeatureProcessor {

	s.v.TargetMeanEncoding = targetmeanencoding.DataframeAnalysisFeatureProcessorTargetMeanEncodingCaster()

	return s
}

func (s *_dataframeAnalysisFeatureProcessor) DataframeAnalysisFeatureProcessorCaster() *types.DataframeAnalysisFeatureProcessor {
	return s.v
}
