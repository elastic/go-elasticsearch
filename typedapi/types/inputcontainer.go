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

// InputContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Input.ts#L92-L100
type InputContainer struct {
	Chain  *ChainInput            `json:"chain,omitempty"`
	Http   *HttpInput             `json:"http,omitempty"`
	Search *SearchInput           `json:"search,omitempty"`
	Simple map[string]interface{} `json:"simple,omitempty"`
}

// InputContainerBuilder holds InputContainer struct and provides a builder API.
type InputContainerBuilder struct {
	v *InputContainer
}

// NewInputContainer provides a builder for the InputContainer struct.
func NewInputContainerBuilder() *InputContainerBuilder {
	r := InputContainerBuilder{
		&InputContainer{
			Simple: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the InputContainer struct
func (rb *InputContainerBuilder) Build() InputContainer {
	return *rb.v
}

func (rb *InputContainerBuilder) Chain(chain *ChainInputBuilder) *InputContainerBuilder {
	v := chain.Build()
	rb.v.Chain = &v
	return rb
}

func (rb *InputContainerBuilder) Http(http *HttpInputBuilder) *InputContainerBuilder {
	v := http.Build()
	rb.v.Http = &v
	return rb
}

func (rb *InputContainerBuilder) Search(search *SearchInputBuilder) *InputContainerBuilder {
	v := search.Build()
	rb.v.Search = &v
	return rb
}

func (rb *InputContainerBuilder) Simple(value map[string]interface{}) *InputContainerBuilder {
	rb.v.Simple = value
	return rb
}
