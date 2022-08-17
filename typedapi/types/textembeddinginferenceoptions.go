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

// TextEmbeddingInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L217-L223
type TextEmbeddingInferenceOptions struct {
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options
	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
}

// TextEmbeddingInferenceOptionsBuilder holds TextEmbeddingInferenceOptions struct and provides a builder API.
type TextEmbeddingInferenceOptionsBuilder struct {
	v *TextEmbeddingInferenceOptions
}

// NewTextEmbeddingInferenceOptions provides a builder for the TextEmbeddingInferenceOptions struct.
func NewTextEmbeddingInferenceOptionsBuilder() *TextEmbeddingInferenceOptionsBuilder {
	r := TextEmbeddingInferenceOptionsBuilder{
		&TextEmbeddingInferenceOptions{},
	}

	return &r
}

// Build finalize the chain and returns the TextEmbeddingInferenceOptions struct
func (rb *TextEmbeddingInferenceOptionsBuilder) Build() TextEmbeddingInferenceOptions {
	return *rb.v
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *TextEmbeddingInferenceOptionsBuilder) ResultsField(resultsfield string) *TextEmbeddingInferenceOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// Tokenization The tokenization options

func (rb *TextEmbeddingInferenceOptionsBuilder) Tokenization(tokenization *TokenizationConfigContainerBuilder) *TextEmbeddingInferenceOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}
