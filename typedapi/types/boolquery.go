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

// BoolQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L28-L34
type BoolQuery struct {
	Boost              *float32            `json:"boost,omitempty"`
	Filter             []QueryContainer    `json:"filter,omitempty"`
	MinimumShouldMatch *MinimumShouldMatch `json:"minimum_should_match,omitempty"`
	Must               []QueryContainer    `json:"must,omitempty"`
	MustNot            []QueryContainer    `json:"must_not,omitempty"`
	QueryName_         *string             `json:"_name,omitempty"`
	Should             []QueryContainer    `json:"should,omitempty"`
}

// BoolQueryBuilder holds BoolQuery struct and provides a builder API.
type BoolQueryBuilder struct {
	v *BoolQuery
}

// NewBoolQuery provides a builder for the BoolQuery struct.
func NewBoolQueryBuilder() *BoolQueryBuilder {
	r := BoolQueryBuilder{
		&BoolQuery{},
	}

	return &r
}

// Build finalize the chain and returns the BoolQuery struct
func (rb *BoolQueryBuilder) Build() BoolQuery {
	return *rb.v
}

func (rb *BoolQueryBuilder) Boost(boost float32) *BoolQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *BoolQueryBuilder) Filter(arg []QueryContainer) *BoolQueryBuilder {
	rb.v.Filter = arg
	return rb
}

func (rb *BoolQueryBuilder) MinimumShouldMatch(minimumshouldmatch *MinimumShouldMatchBuilder) *BoolQueryBuilder {
	v := minimumshouldmatch.Build()
	rb.v.MinimumShouldMatch = &v
	return rb
}

func (rb *BoolQueryBuilder) Must(arg []QueryContainer) *BoolQueryBuilder {
	rb.v.Must = arg
	return rb
}

func (rb *BoolQueryBuilder) MustNot(arg []QueryContainer) *BoolQueryBuilder {
	rb.v.MustNot = arg
	return rb
}

func (rb *BoolQueryBuilder) QueryName_(queryname_ string) *BoolQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *BoolQueryBuilder) Should(arg []QueryContainer) *BoolQueryBuilder {
	rb.v.Should = arg
	return rb
}
