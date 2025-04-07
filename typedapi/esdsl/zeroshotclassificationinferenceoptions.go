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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _zeroShotClassificationInferenceOptions struct {
	v *types.ZeroShotClassificationInferenceOptions
}

// Zeroshot classification configuration for inference.
func NewZeroShotClassificationInferenceOptions() *_zeroShotClassificationInferenceOptions {

	return &_zeroShotClassificationInferenceOptions{v: types.NewZeroShotClassificationInferenceOptions()}

}

func (s *_zeroShotClassificationInferenceOptions) ClassificationLabels(classificationlabels ...string) *_zeroShotClassificationInferenceOptions {

	for _, v := range classificationlabels {

		s.v.ClassificationLabels = append(s.v.ClassificationLabels, v)

	}
	return s
}

func (s *_zeroShotClassificationInferenceOptions) HypothesisTemplate(hypothesistemplate string) *_zeroShotClassificationInferenceOptions {

	s.v.HypothesisTemplate = &hypothesistemplate

	return s
}

func (s *_zeroShotClassificationInferenceOptions) Labels(labels ...string) *_zeroShotClassificationInferenceOptions {

	for _, v := range labels {

		s.v.Labels = append(s.v.Labels, v)

	}
	return s
}

func (s *_zeroShotClassificationInferenceOptions) MultiLabel(multilabel bool) *_zeroShotClassificationInferenceOptions {

	s.v.MultiLabel = &multilabel

	return s
}

func (s *_zeroShotClassificationInferenceOptions) ResultsField(resultsfield string) *_zeroShotClassificationInferenceOptions {

	s.v.ResultsField = &resultsfield

	return s
}

func (s *_zeroShotClassificationInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_zeroShotClassificationInferenceOptions {

	s.v.Tokenization = tokenization.TokenizationConfigContainerCaster()

	return s
}

func (s *_zeroShotClassificationInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	container := types.NewInferenceConfigCreateContainer()

	container.ZeroShotClassification = s.v

	return container
}

func (s *_zeroShotClassificationInferenceOptions) ZeroShotClassificationInferenceOptionsCaster() *types.ZeroShotClassificationInferenceOptions {
	return s.v
}
