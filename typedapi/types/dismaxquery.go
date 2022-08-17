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

// DisMaxQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L46-L50
type DisMaxQuery struct {
	Boost      *float32         `json:"boost,omitempty"`
	Queries    []QueryContainer `json:"queries"`
	QueryName_ *string          `json:"_name,omitempty"`
	TieBreaker *float64         `json:"tie_breaker,omitempty"`
}

// DisMaxQueryBuilder holds DisMaxQuery struct and provides a builder API.
type DisMaxQueryBuilder struct {
	v *DisMaxQuery
}

// NewDisMaxQuery provides a builder for the DisMaxQuery struct.
func NewDisMaxQueryBuilder() *DisMaxQueryBuilder {
	r := DisMaxQueryBuilder{
		&DisMaxQuery{},
	}

	return &r
}

// Build finalize the chain and returns the DisMaxQuery struct
func (rb *DisMaxQueryBuilder) Build() DisMaxQuery {
	return *rb.v
}

func (rb *DisMaxQueryBuilder) Boost(boost float32) *DisMaxQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *DisMaxQueryBuilder) Queries(queries []QueryContainerBuilder) *DisMaxQueryBuilder {
	tmp := make([]QueryContainer, len(queries))
	for _, value := range queries {
		tmp = append(tmp, value.Build())
	}
	rb.v.Queries = tmp
	return rb
}

func (rb *DisMaxQueryBuilder) QueryName_(queryname_ string) *DisMaxQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *DisMaxQueryBuilder) TieBreaker(tiebreaker float64) *DisMaxQueryBuilder {
	rb.v.TieBreaker = &tiebreaker
	return rb
}
