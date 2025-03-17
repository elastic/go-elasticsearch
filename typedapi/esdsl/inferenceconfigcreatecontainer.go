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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _inferenceConfigCreateContainer struct {
	v *types.InferenceConfigCreateContainer
}

func NewInferenceConfigCreateContainer() *_inferenceConfigCreateContainer {
	return &_inferenceConfigCreateContainer{v: types.NewInferenceConfigCreateContainer()}
}

// AdditionalInferenceConfigCreateContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_inferenceConfigCreateContainer) AdditionalInferenceConfigCreateContainerProperty(key string, value json.RawMessage) *_inferenceConfigCreateContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalInferenceConfigCreateContainerProperty = tmp
	return s
}

// Classification configuration for inference.
func (s *_inferenceConfigCreateContainer) Classification(classification types.ClassificationInferenceOptionsVariant) *_inferenceConfigCreateContainer {

	s.v.Classification = classification.ClassificationInferenceOptionsCaster()

	return s
}

// Fill mask configuration for inference.
func (s *_inferenceConfigCreateContainer) FillMask(fillmask types.FillMaskInferenceOptionsVariant) *_inferenceConfigCreateContainer {

	s.v.FillMask = fillmask.FillMaskInferenceOptionsCaster()

	return s
}

// Named entity recognition configuration for inference.
func (s *_inferenceConfigCreateContainer) Ner(ner types.NerInferenceOptionsVariant) *_inferenceConfigCreateContainer {

	s.v.Ner = ner.NerInferenceOptionsCaster()

	return s
}

// Pass through configuration for inference.
func (s *_inferenceConfigCreateContainer) PassThrough(passthrough types.PassThroughInferenceOptionsVariant) *_inferenceConfigCreateContainer {

	s.v.PassThrough = passthrough.PassThroughInferenceOptionsCaster()

	return s
}

// Question answering configuration for inference.
func (s *_inferenceConfigCreateContainer) QuestionAnswering(questionanswering types.QuestionAnsweringInferenceOptionsVariant) *_inferenceConfigCreateContainer {

	s.v.QuestionAnswering = questionanswering.QuestionAnsweringInferenceOptionsCaster()

	return s
}

// Regression configuration for inference.
func (s *_inferenceConfigCreateContainer) Regression(regression types.RegressionInferenceOptionsVariant) *_inferenceConfigCreateContainer {

	s.v.Regression = regression.RegressionInferenceOptionsCaster()

	return s
}

// Text classification configuration for inference.
func (s *_inferenceConfigCreateContainer) TextClassification(textclassification types.TextClassificationInferenceOptionsVariant) *_inferenceConfigCreateContainer {

	s.v.TextClassification = textclassification.TextClassificationInferenceOptionsCaster()

	return s
}

// Text embedding configuration for inference.
func (s *_inferenceConfigCreateContainer) TextEmbedding(textembedding types.TextEmbeddingInferenceOptionsVariant) *_inferenceConfigCreateContainer {

	s.v.TextEmbedding = textembedding.TextEmbeddingInferenceOptionsCaster()

	return s
}

// Text expansion configuration for inference.
func (s *_inferenceConfigCreateContainer) TextExpansion(textexpansion types.TextExpansionInferenceOptionsVariant) *_inferenceConfigCreateContainer {

	s.v.TextExpansion = textexpansion.TextExpansionInferenceOptionsCaster()

	return s
}

// Zeroshot classification configuration for inference.
func (s *_inferenceConfigCreateContainer) ZeroShotClassification(zeroshotclassification types.ZeroShotClassificationInferenceOptionsVariant) *_inferenceConfigCreateContainer {

	s.v.ZeroShotClassification = zeroshotclassification.ZeroShotClassificationInferenceOptionsCaster()

	return s
}

func (s *_inferenceConfigCreateContainer) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	return s.v
}
