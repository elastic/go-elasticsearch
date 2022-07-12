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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// ShardFailure type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/Errors.ts#L50-L56
type ShardFailure struct {
	Index  *IndexName `json:"index,omitempty"`
	Node   *string    `json:"node,omitempty"`
	Reason ErrorCause `json:"reason"`
	Shard  int        `json:"shard"`
	Status *string    `json:"status,omitempty"`
}

// ShardFailureBuilder holds ShardFailure struct and provides a builder API.
type ShardFailureBuilder struct {
	v *ShardFailure
}

// NewShardFailure provides a builder for the ShardFailure struct.
func NewShardFailureBuilder() *ShardFailureBuilder {
	r := ShardFailureBuilder{
		&ShardFailure{},
	}

	return &r
}

// Build finalize the chain and returns the ShardFailure struct
func (rb *ShardFailureBuilder) Build() ShardFailure {
	return *rb.v
}

func (rb *ShardFailureBuilder) Index(index IndexName) *ShardFailureBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *ShardFailureBuilder) Node(node string) *ShardFailureBuilder {
	rb.v.Node = &node
	return rb
}

func (rb *ShardFailureBuilder) Reason(reason *ErrorCauseBuilder) *ShardFailureBuilder {
	v := reason.Build()
	rb.v.Reason = v
	return rb
}

func (rb *ShardFailureBuilder) Shard(shard int) *ShardFailureBuilder {
	rb.v.Shard = shard
	return rb
}

func (rb *ShardFailureBuilder) Status(status string) *ShardFailureBuilder {
	rb.v.Status = &status
	return rb
}
