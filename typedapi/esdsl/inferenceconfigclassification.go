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

type _inferenceConfigClassification struct {
	v *types.InferenceConfigClassification
}

// Classification configuration for inference.
func NewInferenceConfigClassification() *_inferenceConfigClassification {

	return &_inferenceConfigClassification{v: types.NewInferenceConfigClassification()}

}

// Specifies the number of top class predictions to return.
func (s *_inferenceConfigClassification) NumTopClasses(numtopclasses int) *_inferenceConfigClassification {

	s.v.NumTopClasses = &numtopclasses

	return s
}

// Specifies the maximum number of feature importance values per document.
func (s *_inferenceConfigClassification) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_inferenceConfigClassification {

	s.v.NumTopFeatureImportanceValues = &numtopfeatureimportancevalues

	return s
}

// Specifies the type of the predicted field to write.
// Valid values are: `string`, `number`, `boolean`.
func (s *_inferenceConfigClassification) PredictionFieldType(predictionfieldtype string) *_inferenceConfigClassification {

	s.v.PredictionFieldType = &predictionfieldtype

	return s
}

// The field that is added to incoming documents to contain the inference
// prediction.
func (s *_inferenceConfigClassification) ResultsField(field string) *_inferenceConfigClassification {

	s.v.ResultsField = &field

	return s
}

// Specifies the field to which the top classes are written.
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
