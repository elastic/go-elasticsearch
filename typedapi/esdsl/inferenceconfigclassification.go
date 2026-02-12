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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _inferenceConfigClassification struct {
	v *types.InferenceConfigClassification
}

// Classification configuration for inference.
func NewInferenceConfigClassification() *_inferenceConfigClassification {

	return &_inferenceConfigClassification{v: types.NewInferenceConfigClassification()}

}

func (s *_inferenceConfigClassification) NumTopClasses(numtopclasses int) *_inferenceConfigClassification {

	s.v.NumTopClasses = &numtopclasses

	return s
}

func (s *_inferenceConfigClassification) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_inferenceConfigClassification {

	s.v.NumTopFeatureImportanceValues = &numtopfeatureimportancevalues

	return s
}

func (s *_inferenceConfigClassification) PredictionFieldType(predictionfieldtype string) *_inferenceConfigClassification {

	s.v.PredictionFieldType = &predictionfieldtype

	return s
}

func (s *_inferenceConfigClassification) ResultsField(field string) *_inferenceConfigClassification {

	s.v.ResultsField = &field

	return s
}

func (s *_inferenceConfigClassification) TopClassesResultsField(field string) *_inferenceConfigClassification {

	s.v.TopClassesResultsField = &field

	return s
}

func (s *_inferenceConfigClassification) InferenceConfigCaster() *types.InferenceConfig {
	container := types.NewInferenceConfig()

	container.Classification = s.v

	return container
}

func (s *_inferenceConfigClassification) InferenceConfigClassificationCaster() *types.InferenceConfigClassification {
	return s.v
}
