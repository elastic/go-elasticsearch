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

// NerInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L225-L233
type NerInferenceOptions struct {
	// ClassificationLabels The token classification labels. Must be IOB formatted tags
	ClassificationLabels []string `json:"classification_labels,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options
	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
}

// NerInferenceOptionsBuilder holds NerInferenceOptions struct and provides a builder API.
type NerInferenceOptionsBuilder struct {
	v *NerInferenceOptions
}

// NewNerInferenceOptions provides a builder for the NerInferenceOptions struct.
func NewNerInferenceOptionsBuilder() *NerInferenceOptionsBuilder {
	r := NerInferenceOptionsBuilder{
		&NerInferenceOptions{},
	}

	return &r
}

// Build finalize the chain and returns the NerInferenceOptions struct
func (rb *NerInferenceOptionsBuilder) Build() NerInferenceOptions {
	return *rb.v
}

// ClassificationLabels The token classification labels. Must be IOB formatted tags

func (rb *NerInferenceOptionsBuilder) ClassificationLabels(classification_labels ...string) *NerInferenceOptionsBuilder {
	rb.v.ClassificationLabels = classification_labels
	return rb
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *NerInferenceOptionsBuilder) ResultsField(resultsfield string) *NerInferenceOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// Tokenization The tokenization options

func (rb *NerInferenceOptionsBuilder) Tokenization(tokenization *TokenizationConfigContainerBuilder) *NerInferenceOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}
