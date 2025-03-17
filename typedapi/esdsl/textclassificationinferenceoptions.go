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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _textClassificationInferenceOptions struct {
	v *types.TextClassificationInferenceOptions
}

// Text classification configuration for inference.
func NewTextClassificationInferenceOptions() *_textClassificationInferenceOptions {

	return &_textClassificationInferenceOptions{v: types.NewTextClassificationInferenceOptions()}

}

// Classification labels to apply other than the stored labels. Must have the
// same deminsions as the default configured labels
func (s *_textClassificationInferenceOptions) ClassificationLabels(classificationlabels ...string) *_textClassificationInferenceOptions {

	for _, v := range classificationlabels {

		s.v.ClassificationLabels = append(s.v.ClassificationLabels, v)

	}
	return s
}

// Specifies the number of top class predictions to return. Defaults to 0.
func (s *_textClassificationInferenceOptions) NumTopClasses(numtopclasses int) *_textClassificationInferenceOptions {

	s.v.NumTopClasses = &numtopclasses

	return s
}

// The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.
func (s *_textClassificationInferenceOptions) ResultsField(resultsfield string) *_textClassificationInferenceOptions {

	s.v.ResultsField = &resultsfield

	return s
}

// The tokenization options
func (s *_textClassificationInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_textClassificationInferenceOptions {

	s.v.Tokenization = tokenization.TokenizationConfigContainerCaster()

	return s
}

func (s *_textClassificationInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	container := types.NewInferenceConfigCreateContainer()

	container.TextClassification = s.v

	return container
}

func (s *_textClassificationInferenceOptions) TextClassificationInferenceOptionsCaster() *types.TextClassificationInferenceOptions {
	return s.v
}
