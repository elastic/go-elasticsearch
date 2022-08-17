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

// HealthRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/health/types.ts#L23-L94
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
	Timestamp *TimeOfDay `json:"timestamp,omitempty"`
	// Unassign number of unassigned shards
	Unassign *string `json:"unassign,omitempty"`
}

// HealthRecordBuilder holds HealthRecord struct and provides a builder API.
type HealthRecordBuilder struct {
	v *HealthRecord
}

// NewHealthRecord provides a builder for the HealthRecord struct.
func NewHealthRecordBuilder() *HealthRecordBuilder {
	r := HealthRecordBuilder{
		&HealthRecord{},
	}

	return &r
}

// Build finalize the chain and returns the HealthRecord struct
func (rb *HealthRecordBuilder) Build() HealthRecord {
	return *rb.v
}

// ActiveShardsPercent active number of shards in percent

func (rb *HealthRecordBuilder) ActiveShardsPercent(activeshardspercent string) *HealthRecordBuilder {
	rb.v.ActiveShardsPercent = &activeshardspercent
	return rb
}

// Cluster cluster name

func (rb *HealthRecordBuilder) Cluster(cluster string) *HealthRecordBuilder {
	rb.v.Cluster = &cluster
	return rb
}

// Epoch seconds since 1970-01-01 00:00:00

func (rb *HealthRecordBuilder) Epoch(epoch *StringifiedEpochTimeUnitSecondsBuilder) *HealthRecordBuilder {
	v := epoch.Build()
	rb.v.Epoch = &v
	return rb
}

// Init number of initializing nodes

func (rb *HealthRecordBuilder) Init(init string) *HealthRecordBuilder {
	rb.v.Init = &init
	return rb
}

// MaxTaskWaitTime wait time of longest task pending

func (rb *HealthRecordBuilder) MaxTaskWaitTime(maxtaskwaittime string) *HealthRecordBuilder {
	rb.v.MaxTaskWaitTime = &maxtaskwaittime
	return rb
}

// NodeData number of nodes that can store data

func (rb *HealthRecordBuilder) NodeData(nodedata string) *HealthRecordBuilder {
	rb.v.NodeData = &nodedata
	return rb
}

// NodeTotal total number of nodes

func (rb *HealthRecordBuilder) NodeTotal(nodetotal string) *HealthRecordBuilder {
	rb.v.NodeTotal = &nodetotal
	return rb
}

// PendingTasks number of pending tasks

func (rb *HealthRecordBuilder) PendingTasks(pendingtasks string) *HealthRecordBuilder {
	rb.v.PendingTasks = &pendingtasks
	return rb
}

// Pri number of primary shards

func (rb *HealthRecordBuilder) Pri(pri string) *HealthRecordBuilder {
	rb.v.Pri = &pri
	return rb
}

// Relo number of relocating nodes

func (rb *HealthRecordBuilder) Relo(relo string) *HealthRecordBuilder {
	rb.v.Relo = &relo
	return rb
}

// Shards total number of shards

func (rb *HealthRecordBuilder) Shards(shards string) *HealthRecordBuilder {
	rb.v.Shards = &shards
	return rb
}

// Status health status

func (rb *HealthRecordBuilder) Status(status string) *HealthRecordBuilder {
	rb.v.Status = &status
	return rb
}

// Timestamp time in HH:MM:SS

func (rb *HealthRecordBuilder) Timestamp(timestamp TimeOfDay) *HealthRecordBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}

// Unassign number of unassigned shards

func (rb *HealthRecordBuilder) Unassign(unassign string) *HealthRecordBuilder {
	rb.v.Unassign = &unassign
	return rb
}
