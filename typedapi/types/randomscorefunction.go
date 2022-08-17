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

// RandomScoreFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L65-L68
type RandomScoreFunction struct {
	Field *Field `json:"field,omitempty"`
	Seed  string `json:"seed,omitempty"`
}

// RandomScoreFunctionBuilder holds RandomScoreFunction struct and provides a builder API.
type RandomScoreFunctionBuilder struct {
	v *RandomScoreFunction
}

// NewRandomScoreFunction provides a builder for the RandomScoreFunction struct.
func NewRandomScoreFunctionBuilder() *RandomScoreFunctionBuilder {
	r := RandomScoreFunctionBuilder{
		&RandomScoreFunction{},
	}

	return &r
}

// Build finalize the chain and returns the RandomScoreFunction struct
func (rb *RandomScoreFunctionBuilder) Build() RandomScoreFunction {
	return *rb.v
}

func (rb *RandomScoreFunctionBuilder) Field(field Field) *RandomScoreFunctionBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *RandomScoreFunctionBuilder) Seed(arg string) *RandomScoreFunctionBuilder {
	rb.v.Seed = arg
	return rb
}
