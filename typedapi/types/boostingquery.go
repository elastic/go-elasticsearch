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

// BoostingQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L36-L40
type BoostingQuery struct {
	Boost         *float32        `json:"boost,omitempty"`
	Negative      *QueryContainer `json:"negative,omitempty"`
	NegativeBoost float64         `json:"negative_boost"`
	Positive      *QueryContainer `json:"positive,omitempty"`
	QueryName_    *string         `json:"_name,omitempty"`
}

// BoostingQueryBuilder holds BoostingQuery struct and provides a builder API.
type BoostingQueryBuilder struct {
	v *BoostingQuery
}

// NewBoostingQuery provides a builder for the BoostingQuery struct.
func NewBoostingQueryBuilder() *BoostingQueryBuilder {
	r := BoostingQueryBuilder{
		&BoostingQuery{},
	}

	return &r
}

// Build finalize the chain and returns the BoostingQuery struct
func (rb *BoostingQueryBuilder) Build() BoostingQuery {
	return *rb.v
}

func (rb *BoostingQueryBuilder) Boost(boost float32) *BoostingQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *BoostingQueryBuilder) Negative(negative *QueryContainerBuilder) *BoostingQueryBuilder {
	v := negative.Build()
	rb.v.Negative = &v
	return rb
}

func (rb *BoostingQueryBuilder) NegativeBoost(negativeboost float64) *BoostingQueryBuilder {
	rb.v.NegativeBoost = negativeboost
	return rb
}

func (rb *BoostingQueryBuilder) Positive(positive *QueryContainerBuilder) *BoostingQueryBuilder {
	v := positive.Build()
	rb.v.Positive = &v
	return rb
}

func (rb *BoostingQueryBuilder) QueryName_(queryname_ string) *BoostingQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
