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

// RankFeatureFunctionSaturation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/specialized.ts#L145-L147
type RankFeatureFunctionSaturation struct {
	Pivot *float32 `json:"pivot,omitempty"`
}

// RankFeatureFunctionSaturationBuilder holds RankFeatureFunctionSaturation struct and provides a builder API.
type RankFeatureFunctionSaturationBuilder struct {
	v *RankFeatureFunctionSaturation
}

// NewRankFeatureFunctionSaturation provides a builder for the RankFeatureFunctionSaturation struct.
func NewRankFeatureFunctionSaturationBuilder() *RankFeatureFunctionSaturationBuilder {
	r := RankFeatureFunctionSaturationBuilder{
		&RankFeatureFunctionSaturation{},
	}

	return &r
}

// Build finalize the chain and returns the RankFeatureFunctionSaturation struct
func (rb *RankFeatureFunctionSaturationBuilder) Build() RankFeatureFunctionSaturation {
	return *rb.v
}

func (rb *RankFeatureFunctionSaturationBuilder) Pivot(pivot float32) *RankFeatureFunctionSaturationBuilder {
	rb.v.Pivot = &pivot
	return rb
}
