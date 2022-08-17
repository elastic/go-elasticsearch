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

// Command type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/reroute/types.ts#L22-L43
type Command struct {
	// AllocateEmptyPrimary Allocate an empty primary shard to a node. Accepts the index and shard for
	// index name and shard number, and node to allocate the shard to. Using this
	// command leads to a complete loss of all data that was indexed into this
	// shard, if it was previously started. If a node which has a copy of the data
	// rejoins the cluster later on, that data will be deleted. To ensure that these
	// implications are well-understood, this command requires the flag
	// accept_data_loss to be explicitly set to true.
	AllocateEmptyPrimary *CommandAllocatePrimaryAction `json:"allocate_empty_primary,omitempty"`
	// AllocateReplica Allocate an unassigned replica shard to a node. Accepts index and shard for
	// index name and shard number, and node to allocate the shard to. Takes
	// allocation deciders into account.
	AllocateReplica *CommandAllocateReplicaAction `json:"allocate_replica,omitempty"`
	// AllocateStalePrimary Allocate a primary shard to a node that holds a stale copy. Accepts the index
	// and shard for index name and shard number, and node to allocate the shard to.
	// Using this command may lead to data loss for the provided shard id. If a node
	// which has the good copy of the data rejoins the cluster later on, that data
	// will be deleted or overwritten with the data of the stale copy that was
	// forcefully allocated with this command. To ensure that these implications are
	// well-understood, this command requires the flag accept_data_loss to be
	// explicitly set to true.
	AllocateStalePrimary *CommandAllocatePrimaryAction `json:"allocate_stale_primary,omitempty"`
	// Cancel Cancel allocation of a shard (or recovery). Accepts index and shard for index
	// name and shard number, and node for the node to cancel the shard allocation
	// on. This can be used to force resynchronization of existing replicas from the
	// primary shard by cancelling them and allowing them to be reinitialized
	// through the standard recovery process. By default only replica shard
	// allocations can be cancelled. If it is necessary to cancel the allocation of
	// a primary shard then the allow_primary flag must also be included in the
	// request.
	Cancel *CommandCancelAction `json:"cancel,omitempty"`
	// Move Move a started shard from one node to another node. Accepts index and shard
	// for index name and shard number, from_node for the node to move the shard
	// from, and to_node for the node to move the shard to.
	Move *CommandMoveAction `json:"move,omitempty"`
}

// CommandBuilder holds Command struct and provides a builder API.
type CommandBuilder struct {
	v *Command
}

// NewCommand provides a builder for the Command struct.
func NewCommandBuilder() *CommandBuilder {
	r := CommandBuilder{
		&Command{},
	}

	return &r
}

// Build finalize the chain and returns the Command struct
func (rb *CommandBuilder) Build() Command {
	return *rb.v
}

// AllocateEmptyPrimary Allocate an empty primary shard to a node. Accepts the index and shard for
// index name and shard number, and node to allocate the shard to. Using this
// command leads to a complete loss of all data that was indexed into this
// shard, if it was previously started. If a node which has a copy of the data
// rejoins the cluster later on, that data will be deleted. To ensure that these
// implications are well-understood, this command requires the flag
// accept_data_loss to be explicitly set to true.

func (rb *CommandBuilder) AllocateEmptyPrimary(allocateemptyprimary *CommandAllocatePrimaryActionBuilder) *CommandBuilder {
	v := allocateemptyprimary.Build()
	rb.v.AllocateEmptyPrimary = &v
	return rb
}

// AllocateReplica Allocate an unassigned replica shard to a node. Accepts index and shard for
// index name and shard number, and node to allocate the shard to. Takes
// allocation deciders into account.

func (rb *CommandBuilder) AllocateReplica(allocatereplica *CommandAllocateReplicaActionBuilder) *CommandBuilder {
	v := allocatereplica.Build()
	rb.v.AllocateReplica = &v
	return rb
}

// AllocateStalePrimary Allocate a primary shard to a node that holds a stale copy. Accepts the index
// and shard for index name and shard number, and node to allocate the shard to.
// Using this command may lead to data loss for the provided shard id. If a node
// which has the good copy of the data rejoins the cluster later on, that data
// will be deleted or overwritten with the data of the stale copy that was
// forcefully allocated with this command. To ensure that these implications are
// well-understood, this command requires the flag accept_data_loss to be
// explicitly set to true.

func (rb *CommandBuilder) AllocateStalePrimary(allocatestaleprimary *CommandAllocatePrimaryActionBuilder) *CommandBuilder {
	v := allocatestaleprimary.Build()
	rb.v.AllocateStalePrimary = &v
	return rb
}

// Cancel Cancel allocation of a shard (or recovery). Accepts index and shard for index
// name and shard number, and node for the node to cancel the shard allocation
// on. This can be used to force resynchronization of existing replicas from the
// primary shard by cancelling them and allowing them to be reinitialized
// through the standard recovery process. By default only replica shard
// allocations can be cancelled. If it is necessary to cancel the allocation of
// a primary shard then the allow_primary flag must also be included in the
// request.

func (rb *CommandBuilder) Cancel(cancel *CommandCancelActionBuilder) *CommandBuilder {
	v := cancel.Build()
	rb.v.Cancel = &v
	return rb
}

// Move Move a started shard from one node to another node. Accepts index and shard
// for index name and shard number, from_node for the node to move the shard
// from, and to_node for the node to move the shard to.

func (rb *CommandBuilder) Move(move *CommandMoveActionBuilder) *CommandBuilder {
	v := move.Build()
	rb.v.Move = &v
	return rb
}
