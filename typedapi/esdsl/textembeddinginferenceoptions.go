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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textEmbeddingInferenceOptions struct {
	v *types.TextEmbeddingInferenceOptions
}

// Text embedding configuration for inference.
func NewTextEmbeddingInferenceOptions() *_textEmbeddingInferenceOptions {

	return &_textEmbeddingInferenceOptions{v: types.NewTextEmbeddingInferenceOptions()}

}

func (s *_textEmbeddingInferenceOptions) EmbeddingSize(embeddingsize int) *_textEmbeddingInferenceOptions {

	s.v.EmbeddingSize = &embeddingsize

	return s
}

func (s *_textEmbeddingInferenceOptions) ResultsField(resultsfield string) *_textEmbeddingInferenceOptions {

	s.v.ResultsField = &resultsfield

	return s
}

func (s *_textEmbeddingInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_textEmbeddingInferenceOptions {

	s.v.Tokenization = tokenization.TokenizationConfigContainerCaster()

	return s
}

func (s *_textEmbeddingInferenceOptions) Vocabulary(vocabulary types.VocabularyVariant) *_textEmbeddingInferenceOptions {

	s.v.Vocabulary = vocabulary.VocabularyCaster()

	return s
}

func (s *_textEmbeddingInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	container := types.NewInferenceConfigCreateContainer()

	container.TextEmbedding = s.v

	return container
}

func (s *_textEmbeddingInferenceOptions) TextEmbeddingInferenceOptionsCaster() *types.TextEmbeddingInferenceOptions {
	return s.v
}
