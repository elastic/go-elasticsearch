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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textExpansionInferenceOptions struct {
	v *types.TextExpansionInferenceOptions
}

// Text expansion configuration for inference.
func NewTextExpansionInferenceOptions(vocabulary types.VocabularyVariant) *_textExpansionInferenceOptions {

	tmp := &_textExpansionInferenceOptions{v: types.NewTextExpansionInferenceOptions()}

	tmp.Vocabulary(vocabulary)

	return tmp

}

func (s *_textExpansionInferenceOptions) ResultsField(resultsfield string) *_textExpansionInferenceOptions {

	s.v.ResultsField = &resultsfield

	return s
}

func (s *_textExpansionInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_textExpansionInferenceOptions {

	s.v.Tokenization = tokenization.TokenizationConfigContainerCaster()

	return s
}

func (s *_textExpansionInferenceOptions) Vocabulary(vocabulary types.VocabularyVariant) *_textExpansionInferenceOptions {

	s.v.Vocabulary = *vocabulary.VocabularyCaster()

	return s
}

func (s *_textExpansionInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	container := types.NewInferenceConfigCreateContainer()

	container.TextExpansion = s.v

	return container
}

func (s *_textExpansionInferenceOptions) TextExpansionInferenceOptionsCaster() *types.TextExpansionInferenceOptions {
	return s.v
}
