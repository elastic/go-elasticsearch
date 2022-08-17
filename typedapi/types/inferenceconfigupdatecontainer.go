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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// InferenceConfigUpdateContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L259-L279
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

// InferenceConfigUpdateContainerBuilder holds InferenceConfigUpdateContainer struct and provides a builder API.
type InferenceConfigUpdateContainerBuilder struct {
	v *InferenceConfigUpdateContainer
}

// NewInferenceConfigUpdateContainer provides a builder for the InferenceConfigUpdateContainer struct.
func NewInferenceConfigUpdateContainerBuilder() *InferenceConfigUpdateContainerBuilder {
	r := InferenceConfigUpdateContainerBuilder{
		&InferenceConfigUpdateContainer{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceConfigUpdateContainer struct
func (rb *InferenceConfigUpdateContainerBuilder) Build() InferenceConfigUpdateContainer {
	return *rb.v
}

// Classification Classification configuration for inference.

func (rb *InferenceConfigUpdateContainerBuilder) Classification(classification *ClassificationInferenceOptionsBuilder) *InferenceConfigUpdateContainerBuilder {
	v := classification.Build()
	rb.v.Classification = &v
	return rb
}

// FillMask Fill mask configuration for inference.

func (rb *InferenceConfigUpdateContainerBuilder) FillMask(fillmask *FillMaskInferenceUpdateOptionsBuilder) *InferenceConfigUpdateContainerBuilder {
	v := fillmask.Build()
	rb.v.FillMask = &v
	return rb
}

// Ner Named entity recognition configuration for inference.

func (rb *InferenceConfigUpdateContainerBuilder) Ner(ner *NerInferenceUpdateOptionsBuilder) *InferenceConfigUpdateContainerBuilder {
	v := ner.Build()
	rb.v.Ner = &v
	return rb
}

// PassThrough Pass through configuration for inference.

func (rb *InferenceConfigUpdateContainerBuilder) PassThrough(passthrough *PassThroughInferenceUpdateOptionsBuilder) *InferenceConfigUpdateContainerBuilder {
	v := passthrough.Build()
	rb.v.PassThrough = &v
	return rb
}

// QuestionAnswering Question answering configuration for inference

func (rb *InferenceConfigUpdateContainerBuilder) QuestionAnswering(questionanswering *QuestionAnsweringInferenceUpdateOptionsBuilder) *InferenceConfigUpdateContainerBuilder {
	v := questionanswering.Build()
	rb.v.QuestionAnswering = &v
	return rb
}

// Regression Regression configuration for inference.

func (rb *InferenceConfigUpdateContainerBuilder) Regression(regression *RegressionInferenceOptionsBuilder) *InferenceConfigUpdateContainerBuilder {
	v := regression.Build()
	rb.v.Regression = &v
	return rb
}

// TextClassification Text classification configuration for inference.

func (rb *InferenceConfigUpdateContainerBuilder) TextClassification(textclassification *TextClassificationInferenceUpdateOptionsBuilder) *InferenceConfigUpdateContainerBuilder {
	v := textclassification.Build()
	rb.v.TextClassification = &v
	return rb
}

// TextEmbedding Text embedding configuration for inference.

func (rb *InferenceConfigUpdateContainerBuilder) TextEmbedding(textembedding *TextEmbeddingInferenceUpdateOptionsBuilder) *InferenceConfigUpdateContainerBuilder {
	v := textembedding.Build()
	rb.v.TextEmbedding = &v
	return rb
}

// ZeroShotClassification Zeroshot classification configuration for inference.

func (rb *InferenceConfigUpdateContainerBuilder) ZeroShotClassification(zeroshotclassification *ZeroShotClassificationInferenceUpdateOptionsBuilder) *InferenceConfigUpdateContainerBuilder {
	v := zeroshotclassification.Build()
	rb.v.ZeroShotClassification = &v
	return rb
}
