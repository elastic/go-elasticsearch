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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// InferenceConfigUpdateContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/inference.ts#L259-L279
type InferenceConfigUpdateContainer struct {
	// Classification Classification configuration for inference.
	Classification *ClassificationInferenceOptions `json:"classification,omitempty"`
	// FillMask Fill mask configuration for inference.
	FillMask *FillMaskInferenceUpdateOptions `json:"fill_mask,omitempty"`
	// Ner Named entity recognition configuration for inference.
	Ner *NerInferenceUpdateOptions `json:"ner,omitempty"`
	// PassThrough Pass through configuration for inference.
	PassThrough *PassThroughInferenceUpdateOptions `json:"pass_through,omitempty"`
	// QuestionAnswering Question answering configuration for inference
	QuestionAnswering *QuestionAnsweringInferenceUpdateOptions `json:"question_answering,omitempty"`
	// Regression Regression configuration for inference.
	Regression *RegressionInferenceOptions `json:"regression,omitempty"`
	// TextClassification Text classification configuration for inference.
	TextClassification *TextClassificationInferenceUpdateOptions `json:"text_classification,omitempty"`
	// TextEmbedding Text embedding configuration for inference.
	TextEmbedding *TextEmbeddingInferenceUpdateOptions `json:"text_embedding,omitempty"`
	// ZeroShotClassification Zeroshot classification configuration for inference.
	ZeroShotClassification *ZeroShotClassificationInferenceUpdateOptions `json:"zero_shot_classification,omitempty"`
}

// NewInferenceConfigUpdateContainer returns a InferenceConfigUpdateContainer.
func NewInferenceConfigUpdateContainer() *InferenceConfigUpdateContainer {
	r := &InferenceConfigUpdateContainer{}

	return r
}
