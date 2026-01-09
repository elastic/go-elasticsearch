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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _nerInferenceOptions struct {
	v *types.NerInferenceOptions
}

// Named entity recognition configuration for inference.
func NewNerInferenceOptions() *_nerInferenceOptions {

	return &_nerInferenceOptions{v: types.NewNerInferenceOptions()}

}

func (s *_nerInferenceOptions) ClassificationLabels(classificationlabels ...string) *_nerInferenceOptions {

	for _, v := range classificationlabels {

		s.v.ClassificationLabels = append(s.v.ClassificationLabels, v)

	}
	return s
}

func (s *_nerInferenceOptions) ResultsField(resultsfield string) *_nerInferenceOptions {

	s.v.ResultsField = &resultsfield

	return s
}

func (s *_nerInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_nerInferenceOptions {

	s.v.Tokenization = tokenization.TokenizationConfigContainerCaster()

	return s
}

func (s *_nerInferenceOptions) Vocabulary(vocabulary types.VocabularyVariant) *_nerInferenceOptions {

	s.v.Vocabulary = vocabulary.VocabularyCaster()

	return s
}

func (s *_nerInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	container := types.NewInferenceConfigCreateContainer()

	container.Ner = s.v

	return container
}

func (s *_nerInferenceOptions) NerInferenceOptionsCaster() *types.NerInferenceOptions {
	return s.v
}
