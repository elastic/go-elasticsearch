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

// QuestionAnsweringInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L245-L255
type QuestionAnsweringInferenceOptions struct {
	// MaxAnswerLength The maximum answer length to consider
	MaxAnswerLength *int `json:"max_answer_length,omitempty"`
	// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.
	NumTopClasses *int `json:"num_top_classes,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
}

// QuestionAnsweringInferenceOptionsBuilder holds QuestionAnsweringInferenceOptions struct and provides a builder API.
type QuestionAnsweringInferenceOptionsBuilder struct {
	v *QuestionAnsweringInferenceOptions
}

// NewQuestionAnsweringInferenceOptions provides a builder for the QuestionAnsweringInferenceOptions struct.
func NewQuestionAnsweringInferenceOptionsBuilder() *QuestionAnsweringInferenceOptionsBuilder {
	r := QuestionAnsweringInferenceOptionsBuilder{
		&QuestionAnsweringInferenceOptions{},
	}

	return &r
}

// Build finalize the chain and returns the QuestionAnsweringInferenceOptions struct
func (rb *QuestionAnsweringInferenceOptionsBuilder) Build() QuestionAnsweringInferenceOptions {
	return *rb.v
}

// MaxAnswerLength The maximum answer length to consider

func (rb *QuestionAnsweringInferenceOptionsBuilder) MaxAnswerLength(maxanswerlength int) *QuestionAnsweringInferenceOptionsBuilder {
	rb.v.MaxAnswerLength = &maxanswerlength
	return rb
}

// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.

func (rb *QuestionAnsweringInferenceOptionsBuilder) NumTopClasses(numtopclasses int) *QuestionAnsweringInferenceOptionsBuilder {
	rb.v.NumTopClasses = &numtopclasses
	return rb
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *QuestionAnsweringInferenceOptionsBuilder) ResultsField(resultsfield string) *QuestionAnsweringInferenceOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// Tokenization The tokenization options to update when inferring

func (rb *QuestionAnsweringInferenceOptionsBuilder) Tokenization(tokenization *TokenizationConfigContainerBuilder) *QuestionAnsweringInferenceOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}
