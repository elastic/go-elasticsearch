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

// ConfusionMatrixThreshold type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/evaluate_data_frame/types.ts#L96-L117
type ConfusionMatrixThreshold struct {
	// FalseNegative False Negative
	FalseNegative int `json:"fn"`
	// FalsePositive False Positive
	FalsePositive int `json:"fp"`
	// TrueNegative True Negative
	TrueNegative int `json:"tn"`
	// TruePositive True Positive
	TruePositive int `json:"tp"`
}

// ConfusionMatrixThresholdBuilder holds ConfusionMatrixThreshold struct and provides a builder API.
type ConfusionMatrixThresholdBuilder struct {
	v *ConfusionMatrixThreshold
}

// NewConfusionMatrixThreshold provides a builder for the ConfusionMatrixThreshold struct.
func NewConfusionMatrixThresholdBuilder() *ConfusionMatrixThresholdBuilder {
	r := ConfusionMatrixThresholdBuilder{
		&ConfusionMatrixThreshold{},
	}

	return &r
}

// Build finalize the chain and returns the ConfusionMatrixThreshold struct
func (rb *ConfusionMatrixThresholdBuilder) Build() ConfusionMatrixThreshold {
	return *rb.v
}

// FalseNegative False Negative

func (rb *ConfusionMatrixThresholdBuilder) FalseNegative(falsenegative int) *ConfusionMatrixThresholdBuilder {
	rb.v.FalseNegative = falsenegative
	return rb
}

// FalsePositive False Positive

func (rb *ConfusionMatrixThresholdBuilder) FalsePositive(falsepositive int) *ConfusionMatrixThresholdBuilder {
	rb.v.FalsePositive = falsepositive
	return rb
}

// TrueNegative True Negative

func (rb *ConfusionMatrixThresholdBuilder) TrueNegative(truenegative int) *ConfusionMatrixThresholdBuilder {
	rb.v.TrueNegative = truenegative
	return rb
}

// TruePositive True Positive

func (rb *ConfusionMatrixThresholdBuilder) TruePositive(truepositive int) *ConfusionMatrixThresholdBuilder {
	rb.v.TruePositive = truepositive
	return rb
}
