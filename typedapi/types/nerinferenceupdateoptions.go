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

// NerInferenceUpdateOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L357-L362
type NerInferenceUpdateOptions struct {
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

// NerInferenceUpdateOptionsBuilder holds NerInferenceUpdateOptions struct and provides a builder API.
type NerInferenceUpdateOptionsBuilder struct {
	v *NerInferenceUpdateOptions
}

// NewNerInferenceUpdateOptions provides a builder for the NerInferenceUpdateOptions struct.
func NewNerInferenceUpdateOptionsBuilder() *NerInferenceUpdateOptionsBuilder {
	r := NerInferenceUpdateOptionsBuilder{
		&NerInferenceUpdateOptions{},
	}

	return &r
}

// Build finalize the chain and returns the NerInferenceUpdateOptions struct
func (rb *NerInferenceUpdateOptionsBuilder) Build() NerInferenceUpdateOptions {
	return *rb.v
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *NerInferenceUpdateOptionsBuilder) ResultsField(resultsfield string) *NerInferenceUpdateOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// Tokenization The tokenization options to update when inferring

func (rb *NerInferenceUpdateOptionsBuilder) Tokenization(tokenization *NlpTokenizationUpdateOptionsBuilder) *NerInferenceUpdateOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}
