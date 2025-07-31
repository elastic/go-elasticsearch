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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeAnalysisFeatureProcessorFrequencyEncoding struct {
	v *types.DataframeAnalysisFeatureProcessorFrequencyEncoding
}

// The configuration information necessary to perform frequency encoding.
func NewDataframeAnalysisFeatureProcessorFrequencyEncoding() *_dataframeAnalysisFeatureProcessorFrequencyEncoding {

	return &_dataframeAnalysisFeatureProcessorFrequencyEncoding{v: types.NewDataframeAnalysisFeatureProcessorFrequencyEncoding()}

}

func (s *_dataframeAnalysisFeatureProcessorFrequencyEncoding) FeatureName(name string) *_dataframeAnalysisFeatureProcessorFrequencyEncoding {

	s.v.FeatureName = name

	return s
}

func (s *_dataframeAnalysisFeatureProcessorFrequencyEncoding) Field(field string) *_dataframeAnalysisFeatureProcessorFrequencyEncoding {

	s.v.Field = field

	return s
}

func (s *_dataframeAnalysisFeatureProcessorFrequencyEncoding) FrequencyMap(frequencymap map[string]types.Float64) *_dataframeAnalysisFeatureProcessorFrequencyEncoding {

	s.v.FrequencyMap = frequencymap
	return s
}

func (s *_dataframeAnalysisFeatureProcessorFrequencyEncoding) AddFrequencyMap(key string, value types.Float64) *_dataframeAnalysisFeatureProcessorFrequencyEncoding {

	var tmp map[string]types.Float64
	if s.v.FrequencyMap == nil {
		s.v.FrequencyMap = make(map[string]types.Float64)
	} else {
		tmp = s.v.FrequencyMap
	}

	tmp[key] = value

	s.v.FrequencyMap = tmp
	return s
}

func (s *_dataframeAnalysisFeatureProcessorFrequencyEncoding) DataframeAnalysisFeatureProcessorCaster() *types.DataframeAnalysisFeatureProcessor {
	container := types.NewDataframeAnalysisFeatureProcessor()

	container.FrequencyEncoding = s.v

	return container
}

func (s *_dataframeAnalysisFeatureProcessorFrequencyEncoding) DataframeAnalysisFeatureProcessorFrequencyEncodingCaster() *types.DataframeAnalysisFeatureProcessorFrequencyEncoding {
	return s.v
}
