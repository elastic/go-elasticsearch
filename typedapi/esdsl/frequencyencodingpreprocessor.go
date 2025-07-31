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

type _frequencyEncodingPreprocessor struct {
	v *types.FrequencyEncodingPreprocessor
}

func NewFrequencyEncodingPreprocessor(featurename string, field string) *_frequencyEncodingPreprocessor {

	tmp := &_frequencyEncodingPreprocessor{v: types.NewFrequencyEncodingPreprocessor()}

	tmp.FeatureName(featurename)

	tmp.Field(field)

	return tmp

}

func (s *_frequencyEncodingPreprocessor) FeatureName(featurename string) *_frequencyEncodingPreprocessor {

	s.v.FeatureName = featurename

	return s
}

func (s *_frequencyEncodingPreprocessor) Field(field string) *_frequencyEncodingPreprocessor {

	s.v.Field = field

	return s
}

func (s *_frequencyEncodingPreprocessor) FrequencyMap(frequencymap map[string]types.Float64) *_frequencyEncodingPreprocessor {

	s.v.FrequencyMap = frequencymap
	return s
}

func (s *_frequencyEncodingPreprocessor) AddFrequencyMap(key string, value types.Float64) *_frequencyEncodingPreprocessor {

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

func (s *_frequencyEncodingPreprocessor) PreprocessorCaster() *types.Preprocessor {
	container := types.NewPreprocessor()

	container.FrequencyEncoding = s.v

	return container
}

func (s *_frequencyEncodingPreprocessor) FrequencyEncodingPreprocessorCaster() *types.FrequencyEncodingPreprocessor {
	return s.v
}
