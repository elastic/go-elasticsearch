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

// ZeroShotClassificationInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L186-L207
type ZeroShotClassificationInferenceOptions struct {
	// ClassificationLabels The zero shot classification labels indicating entailment, neutral, and
	// contradiction
	// Must contain exactly and only entailment, neutral, and contradiction
	ClassificationLabels []string `json:"classification_labels"`
	// HypothesisTemplate Hypothesis template used when tokenizing labels for prediction
	HypothesisTemplate *string `json:"hypothesis_template,omitempty"`
	// Labels The labels to predict.
	Labels []string `json:"labels,omitempty"`
	// MultiLabel Indicates if more than one true label exists.
	MultiLabel *bool `json:"multi_label,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
}

// ZeroShotClassificationInferenceOptionsBuilder holds ZeroShotClassificationInferenceOptions struct and provides a builder API.
type ZeroShotClassificationInferenceOptionsBuilder struct {
	v *ZeroShotClassificationInferenceOptions
}

// NewZeroShotClassificationInferenceOptions provides a builder for the ZeroShotClassificationInferenceOptions struct.
func NewZeroShotClassificationInferenceOptionsBuilder() *ZeroShotClassificationInferenceOptionsBuilder {
	r := ZeroShotClassificationInferenceOptionsBuilder{
		&ZeroShotClassificationInferenceOptions{},
	}

	return &r
}

// Build finalize the chain and returns the ZeroShotClassificationInferenceOptions struct
func (rb *ZeroShotClassificationInferenceOptionsBuilder) Build() ZeroShotClassificationInferenceOptions {
	return *rb.v
}

// ClassificationLabels The zero shot classification labels indicating entailment, neutral, and
// contradiction
// Must contain exactly and only entailment, neutral, and contradiction

func (rb *ZeroShotClassificationInferenceOptionsBuilder) ClassificationLabels(classification_labels ...string) *ZeroShotClassificationInferenceOptionsBuilder {
	rb.v.ClassificationLabels = classification_labels
	return rb
}

// HypothesisTemplate Hypothesis template used when tokenizing labels for prediction

func (rb *ZeroShotClassificationInferenceOptionsBuilder) HypothesisTemplate(hypothesistemplate string) *ZeroShotClassificationInferenceOptionsBuilder {
	rb.v.HypothesisTemplate = &hypothesistemplate
	return rb
}

// Labels The labels to predict.

func (rb *ZeroShotClassificationInferenceOptionsBuilder) Labels(labels ...string) *ZeroShotClassificationInferenceOptionsBuilder {
	rb.v.Labels = labels
	return rb
}

// MultiLabel Indicates if more than one true label exists.

func (rb *ZeroShotClassificationInferenceOptionsBuilder) MultiLabel(multilabel bool) *ZeroShotClassificationInferenceOptionsBuilder {
	rb.v.MultiLabel = &multilabel
	return rb
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *ZeroShotClassificationInferenceOptionsBuilder) ResultsField(resultsfield string) *ZeroShotClassificationInferenceOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// Tokenization The tokenization options to update when inferring

func (rb *ZeroShotClassificationInferenceOptionsBuilder) Tokenization(tokenization *TokenizationConfigContainerBuilder) *ZeroShotClassificationInferenceOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}
