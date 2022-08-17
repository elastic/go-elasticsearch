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

// EqlFeaturesKeys type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L112-L118
type EqlFeaturesKeys struct {
	JoinKeysFiveOrMore uint `json:"join_keys_five_or_more"`
	JoinKeysFour       uint `json:"join_keys_four"`
	JoinKeysOne        uint `json:"join_keys_one"`
	JoinKeysThree      uint `json:"join_keys_three"`
	JoinKeysTwo        uint `json:"join_keys_two"`
}

// EqlFeaturesKeysBuilder holds EqlFeaturesKeys struct and provides a builder API.
type EqlFeaturesKeysBuilder struct {
	v *EqlFeaturesKeys
}

// NewEqlFeaturesKeys provides a builder for the EqlFeaturesKeys struct.
func NewEqlFeaturesKeysBuilder() *EqlFeaturesKeysBuilder {
	r := EqlFeaturesKeysBuilder{
		&EqlFeaturesKeys{},
	}

	return &r
}

// Build finalize the chain and returns the EqlFeaturesKeys struct
func (rb *EqlFeaturesKeysBuilder) Build() EqlFeaturesKeys {
	return *rb.v
}

func (rb *EqlFeaturesKeysBuilder) JoinKeysFiveOrMore(joinkeysfiveormore uint) *EqlFeaturesKeysBuilder {
	rb.v.JoinKeysFiveOrMore = joinkeysfiveormore
	return rb
}

func (rb *EqlFeaturesKeysBuilder) JoinKeysFour(joinkeysfour uint) *EqlFeaturesKeysBuilder {
	rb.v.JoinKeysFour = joinkeysfour
	return rb
}

func (rb *EqlFeaturesKeysBuilder) JoinKeysOne(joinkeysone uint) *EqlFeaturesKeysBuilder {
	rb.v.JoinKeysOne = joinkeysone
	return rb
}

func (rb *EqlFeaturesKeysBuilder) JoinKeysThree(joinkeysthree uint) *EqlFeaturesKeysBuilder {
	rb.v.JoinKeysThree = joinkeysthree
	return rb
}

func (rb *EqlFeaturesKeysBuilder) JoinKeysTwo(joinkeystwo uint) *EqlFeaturesKeysBuilder {
	rb.v.JoinKeysTwo = joinkeystwo
	return rb
}
