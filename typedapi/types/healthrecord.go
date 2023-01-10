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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// HealthRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/health/types.ts#L23-L94
type HealthRecord struct {
	// ActiveShardsPercent active number of shards in percent
	ActiveShardsPercent *string `json:"active_shards_percent,omitempty"`
	// Cluster cluster name
	Cluster *string `json:"cluster,omitempty"`
	// Epoch seconds since 1970-01-01 00:00:00
	Epoch *StringifiedEpochTimeUnitSeconds `json:"epoch,omitempty"`
	// Init number of initializing nodes
	Init *string `json:"init,omitempty"`
	// MaxTaskWaitTime wait time of longest task pending
	MaxTaskWaitTime *string `json:"max_task_wait_time,omitempty"`
	// NodeData number of nodes that can store data
	NodeData *string `json:"node.data,omitempty"`
	// NodeTotal total number of nodes
	NodeTotal *string `json:"node.total,omitempty"`
	// PendingTasks number of pending tasks
	PendingTasks *string `json:"pending_tasks,omitempty"`
	// Pri number of primary shards
	Pri *string `json:"pri,omitempty"`
	// Relo number of relocating nodes
	Relo *string `json:"relo,omitempty"`
	// Shards total number of shards
	Shards *string `json:"shards,omitempty"`
	// Status health status
	Status *string `json:"status,omitempty"`
	// Timestamp time in HH:MM:SS
	Timestamp *string `json:"timestamp,omitempty"`
	// Unassign number of unassigned shards
	Unassign *string `json:"unassign,omitempty"`
}

// NewHealthRecord returns a HealthRecord.
func NewHealthRecord() *HealthRecord {
	r := &HealthRecord{}

	return r
}
