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

// ConstantScoreQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L42-L44
type ConstantScoreQuery struct {
	Boost      *float32        `json:"boost,omitempty"`
	Filter     *QueryContainer `json:"filter,omitempty"`
	QueryName_ *string         `json:"_name,omitempty"`
}

// ConstantScoreQueryBuilder holds ConstantScoreQuery struct and provides a builder API.
type ConstantScoreQueryBuilder struct {
	v *ConstantScoreQuery
}

// NewConstantScoreQuery provides a builder for the ConstantScoreQuery struct.
func NewConstantScoreQueryBuilder() *ConstantScoreQueryBuilder {
	r := ConstantScoreQueryBuilder{
		&ConstantScoreQuery{},
	}

	return &r
}

// Build finalize the chain and returns the ConstantScoreQuery struct
func (rb *ConstantScoreQueryBuilder) Build() ConstantScoreQuery {
	return *rb.v
}

func (rb *ConstantScoreQueryBuilder) Boost(boost float32) *ConstantScoreQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *ConstantScoreQueryBuilder) Filter(filter *QueryContainerBuilder) *ConstantScoreQueryBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *ConstantScoreQueryBuilder) QueryName_(queryname_ string) *ConstantScoreQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
