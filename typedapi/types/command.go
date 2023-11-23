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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

// Command type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/reroute/types.ts#L22-L43
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

// NewCommand returns a Command.
func NewCommand() *Command {
	r := &Command{}

	return r
}
