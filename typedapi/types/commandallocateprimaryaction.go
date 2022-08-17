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

// CommandAllocatePrimaryAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/reroute/types.ts#L78-L84
type CommandAllocatePrimaryAction struct {
	// AcceptDataLoss If a node which has a copy of the data rejoins the cluster later on, that
	// data will be deleted. To ensure that these implications are well-understood,
	// this command requires the flag accept_data_loss to be explicitly set to true
	AcceptDataLoss bool      `json:"accept_data_loss"`
	Index          IndexName `json:"index"`
	Node           string    `json:"node"`
	Shard          int       `json:"shard"`
}

// CommandAllocatePrimaryActionBuilder holds CommandAllocatePrimaryAction struct and provides a builder API.
type CommandAllocatePrimaryActionBuilder struct {
	v *CommandAllocatePrimaryAction
}

// NewCommandAllocatePrimaryAction provides a builder for the CommandAllocatePrimaryAction struct.
func NewCommandAllocatePrimaryActionBuilder() *CommandAllocatePrimaryActionBuilder {
	r := CommandAllocatePrimaryActionBuilder{
		&CommandAllocatePrimaryAction{},
	}

	return &r
}

// Build finalize the chain and returns the CommandAllocatePrimaryAction struct
func (rb *CommandAllocatePrimaryActionBuilder) Build() CommandAllocatePrimaryAction {
	return *rb.v
}

// AcceptDataLoss If a node which has a copy of the data rejoins the cluster later on, that
// data will be deleted. To ensure that these implications are well-understood,
// this command requires the flag accept_data_loss to be explicitly set to true

func (rb *CommandAllocatePrimaryActionBuilder) AcceptDataLoss(acceptdataloss bool) *CommandAllocatePrimaryActionBuilder {
	rb.v.AcceptDataLoss = acceptdataloss
	return rb
}

func (rb *CommandAllocatePrimaryActionBuilder) Index(index IndexName) *CommandAllocatePrimaryActionBuilder {
	rb.v.Index = index
	return rb
}

func (rb *CommandAllocatePrimaryActionBuilder) Node(node string) *CommandAllocatePrimaryActionBuilder {
	rb.v.Node = node
	return rb
}

func (rb *CommandAllocatePrimaryActionBuilder) Shard(shard int) *CommandAllocatePrimaryActionBuilder {
	rb.v.Shard = shard
	return rb
}
