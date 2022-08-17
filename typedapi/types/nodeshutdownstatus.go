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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shutdownstatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shutdowntype"
)

// NodeShutdownStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/shutdown/get_node/ShutdownGetNodeResponse.ts#L29-L38
type NodeShutdownStatus struct {
	NodeId                NodeId                        `json:"node_id"`
	PersistentTasks       PersistentTaskStatus          `json:"persistent_tasks"`
	Plugins               PluginsStatus                 `json:"plugins"`
	Reason                string                        `json:"reason"`
	ShardMigration        ShardMigrationStatus          `json:"shard_migration"`
	ShutdownStartedmillis EpochTimeUnitMillis           `json:"shutdown_startedmillis"`
	Status                shutdownstatus.ShutdownStatus `json:"status"`
	Type                  shutdowntype.ShutdownType     `json:"type"`
}

// NodeShutdownStatusBuilder holds NodeShutdownStatus struct and provides a builder API.
type NodeShutdownStatusBuilder struct {
	v *NodeShutdownStatus
}

// NewNodeShutdownStatus provides a builder for the NodeShutdownStatus struct.
func NewNodeShutdownStatusBuilder() *NodeShutdownStatusBuilder {
	r := NodeShutdownStatusBuilder{
		&NodeShutdownStatus{},
	}

	return &r
}

// Build finalize the chain and returns the NodeShutdownStatus struct
func (rb *NodeShutdownStatusBuilder) Build() NodeShutdownStatus {
	return *rb.v
}

func (rb *NodeShutdownStatusBuilder) NodeId(nodeid NodeId) *NodeShutdownStatusBuilder {
	rb.v.NodeId = nodeid
	return rb
}

func (rb *NodeShutdownStatusBuilder) PersistentTasks(persistenttasks *PersistentTaskStatusBuilder) *NodeShutdownStatusBuilder {
	v := persistenttasks.Build()
	rb.v.PersistentTasks = v
	return rb
}

func (rb *NodeShutdownStatusBuilder) Plugins(plugins *PluginsStatusBuilder) *NodeShutdownStatusBuilder {
	v := plugins.Build()
	rb.v.Plugins = v
	return rb
}

func (rb *NodeShutdownStatusBuilder) Reason(reason string) *NodeShutdownStatusBuilder {
	rb.v.Reason = reason
	return rb
}

func (rb *NodeShutdownStatusBuilder) ShardMigration(shardmigration *ShardMigrationStatusBuilder) *NodeShutdownStatusBuilder {
	v := shardmigration.Build()
	rb.v.ShardMigration = v
	return rb
}

func (rb *NodeShutdownStatusBuilder) ShutdownStartedmillis(shutdownstartedmillis *EpochTimeUnitMillisBuilder) *NodeShutdownStatusBuilder {
	v := shutdownstartedmillis.Build()
	rb.v.ShutdownStartedmillis = v
	return rb
}

func (rb *NodeShutdownStatusBuilder) Status(status shutdownstatus.ShutdownStatus) *NodeShutdownStatusBuilder {
	rb.v.Status = status
	return rb
}

func (rb *NodeShutdownStatusBuilder) Type_(type_ shutdowntype.ShutdownType) *NodeShutdownStatusBuilder {
	rb.v.Type = type_
	return rb
}
