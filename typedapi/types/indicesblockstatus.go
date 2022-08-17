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

// IndicesBlockStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/add_block/IndicesAddBlockResponse.ts#L30-L33
type IndicesBlockStatus struct {
	Blocked bool      `json:"blocked"`
	Name    IndexName `json:"name"`
}

// IndicesBlockStatusBuilder holds IndicesBlockStatus struct and provides a builder API.
type IndicesBlockStatusBuilder struct {
	v *IndicesBlockStatus
}

// NewIndicesBlockStatus provides a builder for the IndicesBlockStatus struct.
func NewIndicesBlockStatusBuilder() *IndicesBlockStatusBuilder {
	r := IndicesBlockStatusBuilder{
		&IndicesBlockStatus{},
	}

	return &r
}

// Build finalize the chain and returns the IndicesBlockStatus struct
func (rb *IndicesBlockStatusBuilder) Build() IndicesBlockStatus {
	return *rb.v
}

func (rb *IndicesBlockStatusBuilder) Blocked(blocked bool) *IndicesBlockStatusBuilder {
	rb.v.Blocked = blocked
	return rb
}

func (rb *IndicesBlockStatusBuilder) Name(name IndexName) *IndicesBlockStatusBuilder {
	rb.v.Name = name
	return rb
}
