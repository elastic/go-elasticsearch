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

// ContextMethod type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/get_script_context/types.ts#L27-L31
type ContextMethod struct {
	Name       Name                 `json:"name"`
	Params     []ContextMethodParam `json:"params"`
	ReturnType string               `json:"return_type"`
}

// ContextMethodBuilder holds ContextMethod struct and provides a builder API.
type ContextMethodBuilder struct {
	v *ContextMethod
}

// NewContextMethod provides a builder for the ContextMethod struct.
func NewContextMethodBuilder() *ContextMethodBuilder {
	r := ContextMethodBuilder{
		&ContextMethod{},
	}

	return &r
}

// Build finalize the chain and returns the ContextMethod struct
func (rb *ContextMethodBuilder) Build() ContextMethod {
	return *rb.v
}

func (rb *ContextMethodBuilder) Name(name Name) *ContextMethodBuilder {
	rb.v.Name = name
	return rb
}

func (rb *ContextMethodBuilder) Params(params []ContextMethodParamBuilder) *ContextMethodBuilder {
	tmp := make([]ContextMethodParam, len(params))
	for _, value := range params {
		tmp = append(tmp, value.Build())
	}
	rb.v.Params = tmp
	return rb
}

func (rb *ContextMethodBuilder) ReturnType(returntype string) *ContextMethodBuilder {
	rb.v.ReturnType = returntype
	return rb
}
