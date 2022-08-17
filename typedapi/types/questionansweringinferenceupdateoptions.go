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

// QuestionAnsweringInferenceUpdateOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L373-L384
type QuestionAnsweringInferenceUpdateOptions struct {
	// MaxAnswerLength The maximum answer length to consider for extraction
	MaxAnswerLength *int `json:"max_answer_length,omitempty"`
	// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.
	NumTopClasses *int `json:"num_top_classes,omitempty"`
	// Question The question to answer given the inference context
	Question string `json:"question"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

// QuestionAnsweringInferenceUpdateOptionsBuilder holds QuestionAnsweringInferenceUpdateOptions struct and provides a builder API.
type QuestionAnsweringInferenceUpdateOptionsBuilder struct {
	v *QuestionAnsweringInferenceUpdateOptions
}

// NewQuestionAnsweringInferenceUpdateOptions provides a builder for the QuestionAnsweringInferenceUpdateOptions struct.
func NewQuestionAnsweringInferenceUpdateOptionsBuilder() *QuestionAnsweringInferenceUpdateOptionsBuilder {
	r := QuestionAnsweringInferenceUpdateOptionsBuilder{
		&QuestionAnsweringInferenceUpdateOptions{},
	}

	return &r
}

// Build finalize the chain and returns the QuestionAnsweringInferenceUpdateOptions struct
func (rb *QuestionAnsweringInferenceUpdateOptionsBuilder) Build() QuestionAnsweringInferenceUpdateOptions {
	return *rb.v
}

// MaxAnswerLength The maximum answer length to consider for extraction

func (rb *QuestionAnsweringInferenceUpdateOptionsBuilder) MaxAnswerLength(maxanswerlength int) *QuestionAnsweringInferenceUpdateOptionsBuilder {
	rb.v.MaxAnswerLength = &maxanswerlength
	return rb
}

// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.

func (rb *QuestionAnsweringInferenceUpdateOptionsBuilder) NumTopClasses(numtopclasses int) *QuestionAnsweringInferenceUpdateOptionsBuilder {
	rb.v.NumTopClasses = &numtopclasses
	return rb
}

// Question The question to answer given the inference context

func (rb *QuestionAnsweringInferenceUpdateOptionsBuilder) Question(question string) *QuestionAnsweringInferenceUpdateOptionsBuilder {
	rb.v.Question = question
	return rb
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *QuestionAnsweringInferenceUpdateOptionsBuilder) ResultsField(resultsfield string) *QuestionAnsweringInferenceUpdateOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// Tokenization The tokenization options to update when inferring

func (rb *QuestionAnsweringInferenceUpdateOptionsBuilder) Tokenization(tokenization *NlpTokenizationUpdateOptionsBuilder) *QuestionAnsweringInferenceUpdateOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}
