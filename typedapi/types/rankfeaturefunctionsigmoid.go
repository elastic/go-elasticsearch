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

// RankFeatureFunctionSigmoid type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/specialized.ts#L149-L152
type RankFeatureFunctionSigmoid struct {
	Exponent float32 `json:"exponent"`
	Pivot    float32 `json:"pivot"`
}

// RankFeatureFunctionSigmoidBuilder holds RankFeatureFunctionSigmoid struct and provides a builder API.
type RankFeatureFunctionSigmoidBuilder struct {
	v *RankFeatureFunctionSigmoid
}

// NewRankFeatureFunctionSigmoid provides a builder for the RankFeatureFunctionSigmoid struct.
func NewRankFeatureFunctionSigmoidBuilder() *RankFeatureFunctionSigmoidBuilder {
	r := RankFeatureFunctionSigmoidBuilder{
		&RankFeatureFunctionSigmoid{},
	}

	return &r
}

// Build finalize the chain and returns the RankFeatureFunctionSigmoid struct
func (rb *RankFeatureFunctionSigmoidBuilder) Build() RankFeatureFunctionSigmoid {
	return *rb.v
}

func (rb *RankFeatureFunctionSigmoidBuilder) Exponent(exponent float32) *RankFeatureFunctionSigmoidBuilder {
	rb.v.Exponent = exponent
	return rb
}

func (rb *RankFeatureFunctionSigmoidBuilder) Pivot(pivot float32) *RankFeatureFunctionSigmoidBuilder {
	rb.v.Pivot = pivot
	return rb
}
