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

// EqlFeaturesPipes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L120-L123
type EqlFeaturesPipes struct {
	PipeHead uint `json:"pipe_head"`
	PipeTail uint `json:"pipe_tail"`
}

// EqlFeaturesPipesBuilder holds EqlFeaturesPipes struct and provides a builder API.
type EqlFeaturesPipesBuilder struct {
	v *EqlFeaturesPipes
}

// NewEqlFeaturesPipes provides a builder for the EqlFeaturesPipes struct.
func NewEqlFeaturesPipesBuilder() *EqlFeaturesPipesBuilder {
	r := EqlFeaturesPipesBuilder{
		&EqlFeaturesPipes{},
	}

	return &r
}

// Build finalize the chain and returns the EqlFeaturesPipes struct
func (rb *EqlFeaturesPipesBuilder) Build() EqlFeaturesPipes {
	return *rb.v
}

func (rb *EqlFeaturesPipesBuilder) PipeHead(pipehead uint) *EqlFeaturesPipesBuilder {
	rb.v.PipeHead = pipehead
	return rb
}

func (rb *EqlFeaturesPipesBuilder) PipeTail(pipetail uint) *EqlFeaturesPipesBuilder {
	rb.v.PipeTail = pipetail
	return rb
}
