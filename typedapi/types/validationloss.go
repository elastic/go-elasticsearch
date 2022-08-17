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

// ValidationLoss type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L428-L433
type ValidationLoss struct {
	// FoldValues Validation loss values for every added decision tree during the forest
	// growing procedure.
	FoldValues []string `json:"fold_values"`
	// LossType The type of the loss metric. For example, binomial_logistic.
	LossType string `json:"loss_type"`
}

// ValidationLossBuilder holds ValidationLoss struct and provides a builder API.
type ValidationLossBuilder struct {
	v *ValidationLoss
}

// NewValidationLoss provides a builder for the ValidationLoss struct.
func NewValidationLossBuilder() *ValidationLossBuilder {
	r := ValidationLossBuilder{
		&ValidationLoss{},
	}

	return &r
}

// Build finalize the chain and returns the ValidationLoss struct
func (rb *ValidationLossBuilder) Build() ValidationLoss {
	return *rb.v
}

// FoldValues Validation loss values for every added decision tree during the forest
// growing procedure.

func (rb *ValidationLossBuilder) FoldValues(fold_values ...string) *ValidationLossBuilder {
	rb.v.FoldValues = fold_values
	return rb
}

// LossType The type of the loss metric. For example, binomial_logistic.

func (rb *ValidationLossBuilder) LossType(losstype string) *ValidationLossBuilder {
	rb.v.LossType = losstype
	return rb
}
