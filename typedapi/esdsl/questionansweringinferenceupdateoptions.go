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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _questionAnsweringInferenceUpdateOptions struct {
	v *types.QuestionAnsweringInferenceUpdateOptions
}

// Question answering configuration for inference
func NewQuestionAnsweringInferenceUpdateOptions(question string) *_questionAnsweringInferenceUpdateOptions {

	tmp := &_questionAnsweringInferenceUpdateOptions{v: types.NewQuestionAnsweringInferenceUpdateOptions()}

	tmp.Question(question)

	return tmp

}

func (s *_questionAnsweringInferenceUpdateOptions) MaxAnswerLength(maxanswerlength int) *_questionAnsweringInferenceUpdateOptions {

	s.v.MaxAnswerLength = &maxanswerlength

	return s
}

func (s *_questionAnsweringInferenceUpdateOptions) NumTopClasses(numtopclasses int) *_questionAnsweringInferenceUpdateOptions {

	s.v.NumTopClasses = &numtopclasses

	return s
}

func (s *_questionAnsweringInferenceUpdateOptions) Question(question string) *_questionAnsweringInferenceUpdateOptions {

	s.v.Question = question

	return s
}

func (s *_questionAnsweringInferenceUpdateOptions) ResultsField(resultsfield string) *_questionAnsweringInferenceUpdateOptions {

	s.v.ResultsField = &resultsfield

	return s
}

func (s *_questionAnsweringInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_questionAnsweringInferenceUpdateOptions {

	s.v.Tokenization = tokenization.NlpTokenizationUpdateOptionsCaster()

	return s
}

func (s *_questionAnsweringInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	container := types.NewInferenceConfigUpdateContainer()

	container.QuestionAnswering = s.v

	return container
}

func (s *_questionAnsweringInferenceUpdateOptions) QuestionAnsweringInferenceUpdateOptionsCaster() *types.QuestionAnsweringInferenceUpdateOptions {
	return s.v
}
