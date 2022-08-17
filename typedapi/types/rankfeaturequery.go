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

// RankFeatureQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/specialized.ts#L154-L162
type RankFeatureQuery struct {
	Boost      *float32                       `json:"boost,omitempty"`
	Field      Field                          `json:"field"`
	Linear     *RankFeatureFunctionLinear     `json:"linear,omitempty"`
	Log        *RankFeatureFunctionLogarithm  `json:"log,omitempty"`
	QueryName_ *string                        `json:"_name,omitempty"`
	Saturation *RankFeatureFunctionSaturation `json:"saturation,omitempty"`
	Sigmoid    *RankFeatureFunctionSigmoid    `json:"sigmoid,omitempty"`
}

// RankFeatureQueryBuilder holds RankFeatureQuery struct and provides a builder API.
type RankFeatureQueryBuilder struct {
	v *RankFeatureQuery
}

// NewRankFeatureQuery provides a builder for the RankFeatureQuery struct.
func NewRankFeatureQueryBuilder() *RankFeatureQueryBuilder {
	r := RankFeatureQueryBuilder{
		&RankFeatureQuery{},
	}

	return &r
}

// Build finalize the chain and returns the RankFeatureQuery struct
func (rb *RankFeatureQueryBuilder) Build() RankFeatureQuery {
	return *rb.v
}

func (rb *RankFeatureQueryBuilder) Boost(boost float32) *RankFeatureQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *RankFeatureQueryBuilder) Field(field Field) *RankFeatureQueryBuilder {
	rb.v.Field = field
	return rb
}

func (rb *RankFeatureQueryBuilder) Linear(linear *RankFeatureFunctionLinearBuilder) *RankFeatureQueryBuilder {
	v := linear.Build()
	rb.v.Linear = &v
	return rb
}

func (rb *RankFeatureQueryBuilder) Log(log *RankFeatureFunctionLogarithmBuilder) *RankFeatureQueryBuilder {
	v := log.Build()
	rb.v.Log = &v
	return rb
}

func (rb *RankFeatureQueryBuilder) QueryName_(queryname_ string) *RankFeatureQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *RankFeatureQueryBuilder) Saturation(saturation *RankFeatureFunctionSaturationBuilder) *RankFeatureQueryBuilder {
	v := saturation.Build()
	rb.v.Saturation = &v
	return rb
}

func (rb *RankFeatureQueryBuilder) Sigmoid(sigmoid *RankFeatureFunctionSigmoidBuilder) *RankFeatureQueryBuilder {
	v := sigmoid.Build()
	rb.v.Sigmoid = &v
	return rb
}
