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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _questionAnsweringInferenceOptions struct {
	v *types.QuestionAnsweringInferenceOptions
}

// Question answering configuration for inference.
func NewQuestionAnsweringInferenceOptions() *_questionAnsweringInferenceOptions {

	return &_questionAnsweringInferenceOptions{v: types.NewQuestionAnsweringInferenceOptions()}

}

func (s *_questionAnsweringInferenceOptions) MaxAnswerLength(maxanswerlength int) *_questionAnsweringInferenceOptions {

	s.v.MaxAnswerLength = &maxanswerlength

	return s
}

func (s *_questionAnsweringInferenceOptions) NumTopClasses(numtopclasses int) *_questionAnsweringInferenceOptions {

	s.v.NumTopClasses = &numtopclasses

	return s
}

func (s *_questionAnsweringInferenceOptions) ResultsField(resultsfield string) *_questionAnsweringInferenceOptions {

	s.v.ResultsField = &resultsfield

	return s
}

func (s *_questionAnsweringInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_questionAnsweringInferenceOptions {

	s.v.Tokenization = tokenization.TokenizationConfigContainerCaster()

	return s
}

func (s *_questionAnsweringInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	container := types.NewInferenceConfigCreateContainer()

	container.QuestionAnswering = s.v

	return container
}

func (s *_questionAnsweringInferenceOptions) QuestionAnsweringInferenceOptionsCaster() *types.QuestionAnsweringInferenceOptions {
	return s.v
}
