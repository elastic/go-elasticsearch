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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textClassificationInferenceUpdateOptions struct {
	v *types.TextClassificationInferenceUpdateOptions
}

// Text classification configuration for inference.
func NewTextClassificationInferenceUpdateOptions() *_textClassificationInferenceUpdateOptions {

	return &_textClassificationInferenceUpdateOptions{v: types.NewTextClassificationInferenceUpdateOptions()}

}

func (s *_textClassificationInferenceUpdateOptions) ClassificationLabels(classificationlabels ...string) *_textClassificationInferenceUpdateOptions {

	for _, v := range classificationlabels {

		s.v.ClassificationLabels = append(s.v.ClassificationLabels, v)

	}
	return s
}

func (s *_textClassificationInferenceUpdateOptions) NumTopClasses(numtopclasses int) *_textClassificationInferenceUpdateOptions {

	s.v.NumTopClasses = &numtopclasses

	return s
}

func (s *_textClassificationInferenceUpdateOptions) ResultsField(resultsfield string) *_textClassificationInferenceUpdateOptions {

	s.v.ResultsField = &resultsfield

	return s
}

func (s *_textClassificationInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_textClassificationInferenceUpdateOptions {

	s.v.Tokenization = tokenization.NlpTokenizationUpdateOptionsCaster()

	return s
}

func (s *_textClassificationInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	container := types.NewInferenceConfigUpdateContainer()

	container.TextClassification = s.v

	return container
}

func (s *_textClassificationInferenceUpdateOptions) TextClassificationInferenceUpdateOptionsCaster() *types.TextClassificationInferenceUpdateOptions {
	return s.v
}
