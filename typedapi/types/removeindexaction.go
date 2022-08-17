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

// RemoveIndexAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/update_aliases/types.ts#L55-L60
type RemoveIndexAction struct {
	Index     *IndexName `json:"index,omitempty"`
	Indices   *Indices   `json:"indices,omitempty"`
	MustExist *bool      `json:"must_exist,omitempty"`
}

// RemoveIndexActionBuilder holds RemoveIndexAction struct and provides a builder API.
type RemoveIndexActionBuilder struct {
	v *RemoveIndexAction
}

// NewRemoveIndexAction provides a builder for the RemoveIndexAction struct.
func NewRemoveIndexActionBuilder() *RemoveIndexActionBuilder {
	r := RemoveIndexActionBuilder{
		&RemoveIndexAction{},
	}

	return &r
}

// Build finalize the chain and returns the RemoveIndexAction struct
func (rb *RemoveIndexActionBuilder) Build() RemoveIndexAction {
	return *rb.v
}

func (rb *RemoveIndexActionBuilder) Index(index IndexName) *RemoveIndexActionBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *RemoveIndexActionBuilder) Indices(indices *IndicesBuilder) *RemoveIndexActionBuilder {
	v := indices.Build()
	rb.v.Indices = &v
	return rb
}

func (rb *RemoveIndexActionBuilder) MustExist(mustexist bool) *RemoveIndexActionBuilder {
	rb.v.MustExist = &mustexist
	return rb
}
