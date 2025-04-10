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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _learningToRankConfig struct {
	v *types.LearningToRankConfig
}

func NewLearningToRankConfig(numtopfeatureimportancevalues int) *_learningToRankConfig {

	tmp := &_learningToRankConfig{v: types.NewLearningToRankConfig()}

	tmp.NumTopFeatureImportanceValues(numtopfeatureimportancevalues)

	return tmp

}

func (s *_learningToRankConfig) DefaultParams(defaultparams map[string]json.RawMessage) *_learningToRankConfig {

	s.v.DefaultParams = defaultparams
	return s
}

func (s *_learningToRankConfig) AddDefaultParam(key string, value json.RawMessage) *_learningToRankConfig {

	var tmp map[string]json.RawMessage
	if s.v.DefaultParams == nil {
		s.v.DefaultParams = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.DefaultParams
	}

	tmp[key] = value

	s.v.DefaultParams = tmp
	return s
}

func (s *_learningToRankConfig) FeatureExtractors(featureextractors []map[string]types.QueryFeatureExtractor) *_learningToRankConfig {

	s.v.FeatureExtractors = featureextractors

	return s
}

func (s *_learningToRankConfig) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_learningToRankConfig {

	s.v.NumTopFeatureImportanceValues = numtopfeatureimportancevalues

	return s
}

func (s *_learningToRankConfig) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	container := types.NewInferenceConfigCreateContainer()

	container.LearningToRank = s.v

	return container
}

func (s *_learningToRankConfig) LearningToRankConfigCaster() *types.LearningToRankConfig {
	return s.v
}
