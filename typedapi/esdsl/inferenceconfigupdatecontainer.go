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

type _inferenceConfigUpdateContainer struct {
	v *types.InferenceConfigUpdateContainer
}

func NewInferenceConfigUpdateContainer() *_inferenceConfigUpdateContainer {
	return &_inferenceConfigUpdateContainer{v: types.NewInferenceConfigUpdateContainer()}
}

func (s *_inferenceConfigUpdateContainer) Classification(classification types.ClassificationInferenceOptionsVariant) *_inferenceConfigUpdateContainer {

	s.v.Classification = classification.ClassificationInferenceOptionsCaster()

	return s
}

func (s *_inferenceConfigUpdateContainer) FillMask(fillmask types.FillMaskInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {

	s.v.FillMask = fillmask.FillMaskInferenceUpdateOptionsCaster()

	return s
}

func (s *_inferenceConfigUpdateContainer) Ner(ner types.NerInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {

	s.v.Ner = ner.NerInferenceUpdateOptionsCaster()

	return s
}

func (s *_inferenceConfigUpdateContainer) PassThrough(passthrough types.PassThroughInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {

	s.v.PassThrough = passthrough.PassThroughInferenceUpdateOptionsCaster()

	return s
}

func (s *_inferenceConfigUpdateContainer) QuestionAnswering(questionanswering types.QuestionAnsweringInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {

	s.v.QuestionAnswering = questionanswering.QuestionAnsweringInferenceUpdateOptionsCaster()

	return s
}

func (s *_inferenceConfigUpdateContainer) Regression(regression types.RegressionInferenceOptionsVariant) *_inferenceConfigUpdateContainer {

	s.v.Regression = regression.RegressionInferenceOptionsCaster()

	return s
}

func (s *_inferenceConfigUpdateContainer) TextClassification(textclassification types.TextClassificationInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {

	s.v.TextClassification = textclassification.TextClassificationInferenceUpdateOptionsCaster()

	return s
}

func (s *_inferenceConfigUpdateContainer) TextEmbedding(textembedding types.TextEmbeddingInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {

	s.v.TextEmbedding = textembedding.TextEmbeddingInferenceUpdateOptionsCaster()

	return s
}

func (s *_inferenceConfigUpdateContainer) TextExpansion(textexpansion types.TextExpansionInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {

	s.v.TextExpansion = textexpansion.TextExpansionInferenceUpdateOptionsCaster()

	return s
}

func (s *_inferenceConfigUpdateContainer) ZeroShotClassification(zeroshotclassification types.ZeroShotClassificationInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {

	s.v.ZeroShotClassification = zeroshotclassification.ZeroShotClassificationInferenceUpdateOptionsCaster()

	return s
}

func (s *_inferenceConfigUpdateContainer) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	return s.v
}
