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

// ContextMethodParam type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/get_script_context/types.ts#L33-L36
type ContextMethodParam struct {
	Name Name   `json:"name"`
	Type string `json:"type"`
}

// ContextMethodParamBuilder holds ContextMethodParam struct and provides a builder API.
type ContextMethodParamBuilder struct {
	v *ContextMethodParam
}

// NewContextMethodParam provides a builder for the ContextMethodParam struct.
func NewContextMethodParamBuilder() *ContextMethodParamBuilder {
	r := ContextMethodParamBuilder{
		&ContextMethodParam{},
	}

	return &r
}

// Build finalize the chain and returns the ContextMethodParam struct
func (rb *ContextMethodParamBuilder) Build() ContextMethodParam {
	return *rb.v
}

func (rb *ContextMethodParamBuilder) Name(name Name) *ContextMethodParamBuilder {
	rb.v.Name = name
	return rb
}

func (rb *ContextMethodParamBuilder) Type_(type_ string) *ContextMethodParamBuilder {
	rb.v.Type = type_
	return rb
}
