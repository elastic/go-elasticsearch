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

type _inferenceConfigRegression struct {
	v *types.InferenceConfigRegression
}

// Regression configuration for inference.
func NewInferenceConfigRegression() *_inferenceConfigRegression {

	return &_inferenceConfigRegression{v: types.NewInferenceConfigRegression()}

}

// Specifies the maximum number of feature importance values per document.
func (s *_inferenceConfigRegression) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_inferenceConfigRegression {

	s.v.NumTopFeatureImportanceValues = &numtopfeatureimportancevalues

	return s
}

// The field that is added to incoming documents to contain the inference
// prediction.
func (s *_inferenceConfigRegression) ResultsField(field string) *_inferenceConfigRegression {

	s.v.ResultsField = &field

	return s
}

func (s *_inferenceConfigRegression) InferenceConfigCaster() *types.InferenceConfig {
	container := types.NewInferenceConfig()

	container.Regression = s.v

	return container
}

func (s *_inferenceConfigRegression) InferenceConfigRegressionCaster() *types.InferenceConfigRegression {
	return s.v
}
