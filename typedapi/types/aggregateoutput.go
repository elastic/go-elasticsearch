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

// AggregateOutput type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/put_trained_model/types.ts#L101-L106
type AggregateOutput struct {
	Exponent           *Weights `json:"exponent,omitempty"`
	LogisticRegression *Weights `json:"logistic_regression,omitempty"`
	WeightedMode       *Weights `json:"weighted_mode,omitempty"`
	WeightedSum        *Weights `json:"weighted_sum,omitempty"`
}

// AggregateOutputBuilder holds AggregateOutput struct and provides a builder API.
type AggregateOutputBuilder struct {
	v *AggregateOutput
}

// NewAggregateOutput provides a builder for the AggregateOutput struct.
func NewAggregateOutputBuilder() *AggregateOutputBuilder {
	r := AggregateOutputBuilder{
		&AggregateOutput{},
	}

	return &r
}

// Build finalize the chain and returns the AggregateOutput struct
func (rb *AggregateOutputBuilder) Build() AggregateOutput {
	return *rb.v
}

func (rb *AggregateOutputBuilder) Exponent(exponent *WeightsBuilder) *AggregateOutputBuilder {
	v := exponent.Build()
	rb.v.Exponent = &v
	return rb
}

func (rb *AggregateOutputBuilder) LogisticRegression(logisticregression *WeightsBuilder) *AggregateOutputBuilder {
	v := logisticregression.Build()
	rb.v.LogisticRegression = &v
	return rb
}

func (rb *AggregateOutputBuilder) WeightedMode(weightedmode *WeightsBuilder) *AggregateOutputBuilder {
	v := weightedmode.Build()
	rb.v.WeightedMode = &v
	return rb
}

func (rb *AggregateOutputBuilder) WeightedSum(weightedsum *WeightsBuilder) *AggregateOutputBuilder {
	v := weightedsum.Build()
	rb.v.WeightedSum = &v
	return rb
}
