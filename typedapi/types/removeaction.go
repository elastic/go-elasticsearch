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

// RemoveAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/update_aliases/types.ts#L46-L53
type RemoveAction struct {
	Alias     *IndexAlias  `json:"alias,omitempty"`
	Aliases   []IndexAlias `json:"aliases,omitempty"`
	Index     *IndexName   `json:"index,omitempty"`
	Indices   *Indices     `json:"indices,omitempty"`
	MustExist *bool        `json:"must_exist,omitempty"`
}

// RemoveActionBuilder holds RemoveAction struct and provides a builder API.
type RemoveActionBuilder struct {
	v *RemoveAction
}

// NewRemoveAction provides a builder for the RemoveAction struct.
func NewRemoveActionBuilder() *RemoveActionBuilder {
	r := RemoveActionBuilder{
		&RemoveAction{},
	}

	return &r
}

// Build finalize the chain and returns the RemoveAction struct
func (rb *RemoveActionBuilder) Build() RemoveAction {
	return *rb.v
}

func (rb *RemoveActionBuilder) Alias(alias IndexAlias) *RemoveActionBuilder {
	rb.v.Alias = &alias
	return rb
}

func (rb *RemoveActionBuilder) Aliases(arg []IndexAlias) *RemoveActionBuilder {
	rb.v.Aliases = arg
	return rb
}

func (rb *RemoveActionBuilder) Index(index IndexName) *RemoveActionBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *RemoveActionBuilder) Indices(indices *IndicesBuilder) *RemoveActionBuilder {
	v := indices.Build()
	rb.v.Indices = &v
	return rb
}

func (rb *RemoveActionBuilder) MustExist(mustexist bool) *RemoveActionBuilder {
	rb.v.MustExist = &mustexist
	return rb
}
