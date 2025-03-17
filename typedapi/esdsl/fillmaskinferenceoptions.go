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

type _fillMaskInferenceOptions struct {
	v *types.FillMaskInferenceOptions
}

// Fill mask configuration for inference.
func NewFillMaskInferenceOptions(vocabulary types.VocabularyVariant) *_fillMaskInferenceOptions {

	tmp := &_fillMaskInferenceOptions{v: types.NewFillMaskInferenceOptions()}

	tmp.Vocabulary(vocabulary)

	return tmp

}

// The string/token which will be removed from incoming documents and replaced
// with the inference prediction(s).
// In a response, this field contains the mask token for the specified
// model/tokenizer. Each model and tokenizer
// has a predefined mask token which cannot be changed. Thus, it is recommended
// not to set this value in requests.
// However, if this field is present in a request, its value must match the
// predefined value for that model/tokenizer,
// otherwise the request will fail.
func (s *_fillMaskInferenceOptions) MaskToken(masktoken string) *_fillMaskInferenceOptions {

	s.v.MaskToken = &masktoken

	return s
}

// Specifies the number of top class predictions to return. Defaults to 0.
func (s *_fillMaskInferenceOptions) NumTopClasses(numtopclasses int) *_fillMaskInferenceOptions {

	s.v.NumTopClasses = &numtopclasses

	return s
}

// The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.
func (s *_fillMaskInferenceOptions) ResultsField(resultsfield string) *_fillMaskInferenceOptions {

	s.v.ResultsField = &resultsfield

	return s
}

// The tokenization options to update when inferring
func (s *_fillMaskInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_fillMaskInferenceOptions {

	s.v.Tokenization = tokenization.TokenizationConfigContainerCaster()

	return s
}

func (s *_fillMaskInferenceOptions) Vocabulary(vocabulary types.VocabularyVariant) *_fillMaskInferenceOptions {

	s.v.Vocabulary = *vocabulary.VocabularyCaster()

	return s
}

func (s *_fillMaskInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	container := types.NewInferenceConfigCreateContainer()

	container.FillMask = s.v

	return container
}

func (s *_fillMaskInferenceOptions) FillMaskInferenceOptionsCaster() *types.FillMaskInferenceOptions {
	return s.v
}
