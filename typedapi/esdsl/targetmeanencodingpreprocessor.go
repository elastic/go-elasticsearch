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

type _targetMeanEncodingPreprocessor struct {
	v *types.TargetMeanEncodingPreprocessor
}

func NewTargetMeanEncodingPreprocessor(defaultvalue types.Float64, featurename string, field string) *_targetMeanEncodingPreprocessor {

	tmp := &_targetMeanEncodingPreprocessor{v: types.NewTargetMeanEncodingPreprocessor()}

	tmp.DefaultValue(defaultvalue)

	tmp.FeatureName(featurename)

	tmp.Field(field)

	return tmp

}

func (s *_targetMeanEncodingPreprocessor) DefaultValue(defaultvalue types.Float64) *_targetMeanEncodingPreprocessor {

	s.v.DefaultValue = defaultvalue

	return s
}

func (s *_targetMeanEncodingPreprocessor) FeatureName(featurename string) *_targetMeanEncodingPreprocessor {

	s.v.FeatureName = featurename

	return s
}

func (s *_targetMeanEncodingPreprocessor) Field(field string) *_targetMeanEncodingPreprocessor {

	s.v.Field = field

	return s
}

func (s *_targetMeanEncodingPreprocessor) TargetMap(targetmap map[string]types.Float64) *_targetMeanEncodingPreprocessor {

	s.v.TargetMap = targetmap
	return s
}

func (s *_targetMeanEncodingPreprocessor) AddTargetMap(key string, value types.Float64) *_targetMeanEncodingPreprocessor {

	var tmp map[string]types.Float64
	if s.v.TargetMap == nil {
		s.v.TargetMap = make(map[string]types.Float64)
	} else {
		tmp = s.v.TargetMap
	}

	tmp[key] = value

	s.v.TargetMap = tmp
	return s
}

func (s *_targetMeanEncodingPreprocessor) PreprocessorCaster() *types.Preprocessor {
	container := types.NewPreprocessor()

	container.TargetMeanEncoding = s.v

	return container
}

func (s *_targetMeanEncodingPreprocessor) TargetMeanEncodingPreprocessorCaster() *types.TargetMeanEncodingPreprocessor {
	return s.v
}
