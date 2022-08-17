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

// FillMaskInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L235-L243
type FillMaskInferenceOptions struct {
	// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.
	NumTopClasses *int `json:"num_top_classes,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
}

// FillMaskInferenceOptionsBuilder holds FillMaskInferenceOptions struct and provides a builder API.
type FillMaskInferenceOptionsBuilder struct {
	v *FillMaskInferenceOptions
}

// NewFillMaskInferenceOptions provides a builder for the FillMaskInferenceOptions struct.
func NewFillMaskInferenceOptionsBuilder() *FillMaskInferenceOptionsBuilder {
	r := FillMaskInferenceOptionsBuilder{
		&FillMaskInferenceOptions{},
	}

	return &r
}

// Build finalize the chain and returns the FillMaskInferenceOptions struct
func (rb *FillMaskInferenceOptionsBuilder) Build() FillMaskInferenceOptions {
	return *rb.v
}

// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.

func (rb *FillMaskInferenceOptionsBuilder) NumTopClasses(numtopclasses int) *FillMaskInferenceOptionsBuilder {
	rb.v.NumTopClasses = &numtopclasses
	return rb
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *FillMaskInferenceOptionsBuilder) ResultsField(resultsfield string) *FillMaskInferenceOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// Tokenization The tokenization options to update when inferring

func (rb *FillMaskInferenceOptionsBuilder) Tokenization(tokenization *TokenizationConfigContainerBuilder) *FillMaskInferenceOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}
