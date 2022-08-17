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

// InferenceConfigCreateContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L23-L67
type InferenceConfigCreateContainer struct {
	// Classification Classification configuration for inference.
	Classification *ClassificationInferenceOptions `json:"classification,omitempty"`
	// FillMask Fill mask configuration for inference.
	FillMask *FillMaskInferenceOptions `json:"fill_mask,omitempty"`
	// Ner Named entity recognition configuration for inference.
	Ner *NerInferenceOptions `json:"ner,omitempty"`
	// PassThrough Pass through configuration for inference.
	PassThrough *PassThroughInferenceOptions `json:"pass_through,omitempty"`
	// QuestionAnswering Question answering configuration for inference.
	QuestionAnswering *QuestionAnsweringInferenceOptions `json:"question_answering,omitempty"`
	// Regression Regression configuration for inference.
	Regression *RegressionInferenceOptions `json:"regression,omitempty"`
	// TextClassification Text classification configuration for inference.
	TextClassification *TextClassificationInferenceOptions `json:"text_classification,omitempty"`
	// TextEmbedding Text embedding configuration for inference.
	TextEmbedding *TextEmbeddingInferenceOptions `json:"text_embedding,omitempty"`
	// ZeroShotClassification Zeroshot classification configuration for inference.
	ZeroShotClassification *ZeroShotClassificationInferenceOptions `json:"zero_shot_classification,omitempty"`
}

// InferenceConfigCreateContainerBuilder holds InferenceConfigCreateContainer struct and provides a builder API.
type InferenceConfigCreateContainerBuilder struct {
	v *InferenceConfigCreateContainer
}

// NewInferenceConfigCreateContainer provides a builder for the InferenceConfigCreateContainer struct.
func NewInferenceConfigCreateContainerBuilder() *InferenceConfigCreateContainerBuilder {
	r := InferenceConfigCreateContainerBuilder{
		&InferenceConfigCreateContainer{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceConfigCreateContainer struct
func (rb *InferenceConfigCreateContainerBuilder) Build() InferenceConfigCreateContainer {
	return *rb.v
}

// Classification Classification configuration for inference.

func (rb *InferenceConfigCreateContainerBuilder) Classification(classification *ClassificationInferenceOptionsBuilder) *InferenceConfigCreateContainerBuilder {
	v := classification.Build()
	rb.v.Classification = &v
	return rb
}

// FillMask Fill mask configuration for inference.

func (rb *InferenceConfigCreateContainerBuilder) FillMask(fillmask *FillMaskInferenceOptionsBuilder) *InferenceConfigCreateContainerBuilder {
	v := fillmask.Build()
	rb.v.FillMask = &v
	return rb
}

// Ner Named entity recognition configuration for inference.

func (rb *InferenceConfigCreateContainerBuilder) Ner(ner *NerInferenceOptionsBuilder) *InferenceConfigCreateContainerBuilder {
	v := ner.Build()
	rb.v.Ner = &v
	return rb
}

// PassThrough Pass through configuration for inference.

func (rb *InferenceConfigCreateContainerBuilder) PassThrough(passthrough *PassThroughInferenceOptionsBuilder) *InferenceConfigCreateContainerBuilder {
	v := passthrough.Build()
	rb.v.PassThrough = &v
	return rb
}

// QuestionAnswering Question answering configuration for inference.

func (rb *InferenceConfigCreateContainerBuilder) QuestionAnswering(questionanswering *QuestionAnsweringInferenceOptionsBuilder) *InferenceConfigCreateContainerBuilder {
	v := questionanswering.Build()
	rb.v.QuestionAnswering = &v
	return rb
}

// Regression Regression configuration for inference.

func (rb *InferenceConfigCreateContainerBuilder) Regression(regression *RegressionInferenceOptionsBuilder) *InferenceConfigCreateContainerBuilder {
	v := regression.Build()
	rb.v.Regression = &v
	return rb
}

// TextClassification Text classification configuration for inference.

func (rb *InferenceConfigCreateContainerBuilder) TextClassification(textclassification *TextClassificationInferenceOptionsBuilder) *InferenceConfigCreateContainerBuilder {
	v := textclassification.Build()
	rb.v.TextClassification = &v
	return rb
}

// TextEmbedding Text embedding configuration for inference.

func (rb *InferenceConfigCreateContainerBuilder) TextEmbedding(textembedding *TextEmbeddingInferenceOptionsBuilder) *InferenceConfigCreateContainerBuilder {
	v := textembedding.Build()
	rb.v.TextEmbedding = &v
	return rb
}

// ZeroShotClassification Zeroshot classification configuration for inference.

func (rb *InferenceConfigCreateContainerBuilder) ZeroShotClassification(zeroshotclassification *ZeroShotClassificationInferenceOptionsBuilder) *InferenceConfigCreateContainerBuilder {
	v := zeroshotclassification.Build()
	rb.v.ZeroShotClassification = &v
	return rb
}
