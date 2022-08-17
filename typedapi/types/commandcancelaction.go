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

// CommandCancelAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/reroute/types.ts#L45-L50
type CommandCancelAction struct {
	AllowPrimary *bool     `json:"allow_primary,omitempty"`
	Index        IndexName `json:"index"`
	Node         string    `json:"node"`
	Shard        int       `json:"shard"`
}

// CommandCancelActionBuilder holds CommandCancelAction struct and provides a builder API.
type CommandCancelActionBuilder struct {
	v *CommandCancelAction
}

// NewCommandCancelAction provides a builder for the CommandCancelAction struct.
func NewCommandCancelActionBuilder() *CommandCancelActionBuilder {
	r := CommandCancelActionBuilder{
		&CommandCancelAction{},
	}

	return &r
}

// Build finalize the chain and returns the CommandCancelAction struct
func (rb *CommandCancelActionBuilder) Build() CommandCancelAction {
	return *rb.v
}

func (rb *CommandCancelActionBuilder) AllowPrimary(allowprimary bool) *CommandCancelActionBuilder {
	rb.v.AllowPrimary = &allowprimary
	return rb
}

func (rb *CommandCancelActionBuilder) Index(index IndexName) *CommandCancelActionBuilder {
	rb.v.Index = index
	return rb
}

func (rb *CommandCancelActionBuilder) Node(node string) *CommandCancelActionBuilder {
	rb.v.Node = node
	return rb
}

func (rb *CommandCancelActionBuilder) Shard(shard int) *CommandCancelActionBuilder {
	rb.v.Shard = shard
	return rb
}
