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

// TextClassificationInferenceUpdateOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L322-L331
type TextClassificationInferenceUpdateOptions struct {
	// ClassificationLabels Classification labels to apply other than the stored labels. Must have the
	// same deminsions as the default configured labels
	ClassificationLabels []string `json:"classification_labels,omitempty"`
	// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.
	NumTopClasses *int `json:"num_top_classes,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

// TextClassificationInferenceUpdateOptionsBuilder holds TextClassificationInferenceUpdateOptions struct and provides a builder API.
type TextClassificationInferenceUpdateOptionsBuilder struct {
	v *TextClassificationInferenceUpdateOptions
}

// NewTextClassificationInferenceUpdateOptions provides a builder for the TextClassificationInferenceUpdateOptions struct.
func NewTextClassificationInferenceUpdateOptionsBuilder() *TextClassificationInferenceUpdateOptionsBuilder {
	r := TextClassificationInferenceUpdateOptionsBuilder{
		&TextClassificationInferenceUpdateOptions{},
	}

	return &r
}

// Build finalize the chain and returns the TextClassificationInferenceUpdateOptions struct
func (rb *TextClassificationInferenceUpdateOptionsBuilder) Build() TextClassificationInferenceUpdateOptions {
	return *rb.v
}

// ClassificationLabels Classification labels to apply other than the stored labels. Must have the
// same deminsions as the default configured labels

func (rb *TextClassificationInferenceUpdateOptionsBuilder) ClassificationLabels(classification_labels ...string) *TextClassificationInferenceUpdateOptionsBuilder {
	rb.v.ClassificationLabels = classification_labels
	return rb
}

// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.

func (rb *TextClassificationInferenceUpdateOptionsBuilder) NumTopClasses(numtopclasses int) *TextClassificationInferenceUpdateOptionsBuilder {
	rb.v.NumTopClasses = &numtopclasses
	return rb
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *TextClassificationInferenceUpdateOptionsBuilder) ResultsField(resultsfield string) *TextClassificationInferenceUpdateOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// Tokenization The tokenization options to update when inferring

func (rb *TextClassificationInferenceUpdateOptionsBuilder) Tokenization(tokenization *NlpTokenizationUpdateOptionsBuilder) *TextClassificationInferenceUpdateOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}
