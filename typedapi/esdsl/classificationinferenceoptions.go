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

type _classificationInferenceOptions struct {
	v *types.ClassificationInferenceOptions
}

// Classification configuration for inference.
func NewClassificationInferenceOptions() *_classificationInferenceOptions {

	return &_classificationInferenceOptions{v: types.NewClassificationInferenceOptions()}

}

func (s *_classificationInferenceOptions) NumTopClasses(numtopclasses int) *_classificationInferenceOptions {

	s.v.NumTopClasses = &numtopclasses

	return s
}

func (s *_classificationInferenceOptions) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_classificationInferenceOptions {

	s.v.NumTopFeatureImportanceValues = &numtopfeatureimportancevalues

	return s
}

func (s *_classificationInferenceOptions) PredictionFieldType(predictionfieldtype string) *_classificationInferenceOptions {

	s.v.PredictionFieldType = &predictionfieldtype

	return s
}

func (s *_classificationInferenceOptions) ResultsField(resultsfield string) *_classificationInferenceOptions {

	s.v.ResultsField = &resultsfield

	return s
}

func (s *_classificationInferenceOptions) TopClassesResultsField(topclassesresultsfield string) *_classificationInferenceOptions {

	s.v.TopClassesResultsField = &topclassesresultsfield

	return s
}

func (s *_classificationInferenceOptions) InferenceConfigContainerCaster() *types.InferenceConfigContainer {
	container := types.NewInferenceConfigContainer()

	container.Classification = s.v

	return container
}

func (s *_classificationInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	container := types.NewInferenceConfigCreateContainer()

	container.Classification = s.v

	return container
}

func (s *_classificationInferenceOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	container := types.NewInferenceConfigUpdateContainer()

	container.Classification = s.v

	return container
}

func (s *_classificationInferenceOptions) ClassificationInferenceOptionsCaster() *types.ClassificationInferenceOptions {
	return s.v
}
