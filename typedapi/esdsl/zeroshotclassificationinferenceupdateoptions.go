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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _zeroShotClassificationInferenceUpdateOptions struct {
	v *types.ZeroShotClassificationInferenceUpdateOptions
}

// Zeroshot classification configuration for inference.
func NewZeroShotClassificationInferenceUpdateOptions() *_zeroShotClassificationInferenceUpdateOptions {

	return &_zeroShotClassificationInferenceUpdateOptions{v: types.NewZeroShotClassificationInferenceUpdateOptions()}

}

// The labels to predict.
func (s *_zeroShotClassificationInferenceUpdateOptions) Labels(labels ...string) *_zeroShotClassificationInferenceUpdateOptions {

	for _, v := range labels {

		s.v.Labels = append(s.v.Labels, v)

	}
	return s
}

// Update the configured multi label option. Indicates if more than one true
// label exists. Defaults to the configured value.
func (s *_zeroShotClassificationInferenceUpdateOptions) MultiLabel(multilabel bool) *_zeroShotClassificationInferenceUpdateOptions {

	s.v.MultiLabel = &multilabel

	return s
}

// The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.
func (s *_zeroShotClassificationInferenceUpdateOptions) ResultsField(resultsfield string) *_zeroShotClassificationInferenceUpdateOptions {

	s.v.ResultsField = &resultsfield

	return s
}

// The tokenization options to update when inferring
func (s *_zeroShotClassificationInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_zeroShotClassificationInferenceUpdateOptions {

	s.v.Tokenization = tokenization.NlpTokenizationUpdateOptionsCaster()

	return s
}

func (s *_zeroShotClassificationInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	container := types.NewInferenceConfigUpdateContainer()

	container.ZeroShotClassification = s.v

	return container
}

func (s *_zeroShotClassificationInferenceUpdateOptions) ZeroShotClassificationInferenceUpdateOptionsCaster() *types.ZeroShotClassificationInferenceUpdateOptions {
	return s.v
}
