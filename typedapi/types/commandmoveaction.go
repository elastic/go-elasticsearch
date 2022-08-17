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

// CommandMoveAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/reroute/types.ts#L60-L67
type CommandMoveAction struct {
	// FromNode The node to move the shard from
	FromNode string    `json:"from_node"`
	Index    IndexName `json:"index"`
	Shard    int       `json:"shard"`
	// ToNode The node to move the shard to
	ToNode string `json:"to_node"`
}

// CommandMoveActionBuilder holds CommandMoveAction struct and provides a builder API.
type CommandMoveActionBuilder struct {
	v *CommandMoveAction
}

// NewCommandMoveAction provides a builder for the CommandMoveAction struct.
func NewCommandMoveActionBuilder() *CommandMoveActionBuilder {
	r := CommandMoveActionBuilder{
		&CommandMoveAction{},
	}

	return &r
}

// Build finalize the chain and returns the CommandMoveAction struct
func (rb *CommandMoveActionBuilder) Build() CommandMoveAction {
	return *rb.v
}

// FromNode The node to move the shard from

func (rb *CommandMoveActionBuilder) FromNode(fromnode string) *CommandMoveActionBuilder {
	rb.v.FromNode = fromnode
	return rb
}

func (rb *CommandMoveActionBuilder) Index(index IndexName) *CommandMoveActionBuilder {
	rb.v.Index = index
	return rb
}

func (rb *CommandMoveActionBuilder) Shard(shard int) *CommandMoveActionBuilder {
	rb.v.Shard = shard
	return rb
}

// ToNode The node to move the shard to

func (rb *CommandMoveActionBuilder) ToNode(tonode string) *CommandMoveActionBuilder {
	rb.v.ToNode = tonode
	return rb
}
