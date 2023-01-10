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

// TasksRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/tasks/types.ts#L22-L101
type TasksRecord struct {
	// Action task action
	Action *string `json:"action,omitempty"`
	// Description task action
	Description *string `json:"description,omitempty"`
	// Id id of the task with the node
	Id *string `json:"id,omitempty"`
	// Ip ip address
	Ip *string `json:"ip,omitempty"`
	// Node node name
	Node *string `json:"node,omitempty"`
	// NodeId unique node id
	NodeId *string `json:"node_id,omitempty"`
	// ParentTaskId parent task id
	ParentTaskId *string `json:"parent_task_id,omitempty"`
	// Port bound transport port
	Port *string `json:"port,omitempty"`
	// RunningTime running time
	RunningTime *string `json:"running_time,omitempty"`
	// RunningTimeNs running time ns
	RunningTimeNs *string `json:"running_time_ns,omitempty"`
	// StartTime start time in ms
	StartTime *string `json:"start_time,omitempty"`
	// TaskId unique task id
	TaskId *string `json:"task_id,omitempty"`
	// Timestamp start time in HH:MM:SS
	Timestamp *string `json:"timestamp,omitempty"`
	// Type task type
	Type *string `json:"type,omitempty"`
	// Version es version
	Version *string `json:"version,omitempty"`
	// XOpaqueId X-Opaque-ID header
	XOpaqueId *string `json:"x_opaque_id,omitempty"`
}

// NewTasksRecord returns a TasksRecord.
func NewTasksRecord() *TasksRecord {
	r := &TasksRecord{}

	return r
}
