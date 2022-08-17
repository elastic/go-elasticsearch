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

// EqlFeaturesJoin type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L104-L110
type EqlFeaturesJoin struct {
	JoinQueriesFiveOrMore uint `json:"join_queries_five_or_more"`
	JoinQueriesFour       uint `json:"join_queries_four"`
	JoinQueriesThree      uint `json:"join_queries_three"`
	JoinQueriesTwo        uint `json:"join_queries_two"`
	JoinUntil             uint `json:"join_until"`
}

// EqlFeaturesJoinBuilder holds EqlFeaturesJoin struct and provides a builder API.
type EqlFeaturesJoinBuilder struct {
	v *EqlFeaturesJoin
}

// NewEqlFeaturesJoin provides a builder for the EqlFeaturesJoin struct.
func NewEqlFeaturesJoinBuilder() *EqlFeaturesJoinBuilder {
	r := EqlFeaturesJoinBuilder{
		&EqlFeaturesJoin{},
	}

	return &r
}

// Build finalize the chain and returns the EqlFeaturesJoin struct
func (rb *EqlFeaturesJoinBuilder) Build() EqlFeaturesJoin {
	return *rb.v
}

func (rb *EqlFeaturesJoinBuilder) JoinQueriesFiveOrMore(joinqueriesfiveormore uint) *EqlFeaturesJoinBuilder {
	rb.v.JoinQueriesFiveOrMore = joinqueriesfiveormore
	return rb
}

func (rb *EqlFeaturesJoinBuilder) JoinQueriesFour(joinqueriesfour uint) *EqlFeaturesJoinBuilder {
	rb.v.JoinQueriesFour = joinqueriesfour
	return rb
}

func (rb *EqlFeaturesJoinBuilder) JoinQueriesThree(joinqueriesthree uint) *EqlFeaturesJoinBuilder {
	rb.v.JoinQueriesThree = joinqueriesthree
	return rb
}

func (rb *EqlFeaturesJoinBuilder) JoinQueriesTwo(joinqueriestwo uint) *EqlFeaturesJoinBuilder {
	rb.v.JoinQueriesTwo = joinqueriestwo
	return rb
}

func (rb *EqlFeaturesJoinBuilder) JoinUntil(joinuntil uint) *EqlFeaturesJoinBuilder {
	rb.v.JoinUntil = joinuntil
	return rb
}
