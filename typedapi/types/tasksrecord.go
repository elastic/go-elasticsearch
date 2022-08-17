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

// TasksRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/tasks/types.ts#L22-L101
type TasksRecord struct {
	// Action task action
	Action *string `json:"action,omitempty"`
	// Description task action
	Description *string `json:"description,omitempty"`
	// Id id of the task with the node
	Id *Id `json:"id,omitempty"`
	// Ip ip address
	Ip *string `json:"ip,omitempty"`
	// Node node name
	Node *string `json:"node,omitempty"`
	// NodeId unique node id
	NodeId *NodeId `json:"node_id,omitempty"`
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
	TaskId *Id `json:"task_id,omitempty"`
	// Timestamp start time in HH:MM:SS
	Timestamp *string `json:"timestamp,omitempty"`
	// Type task type
	Type *string `json:"type,omitempty"`
	// Version es version
	Version *VersionString `json:"version,omitempty"`
	// XOpaqueId X-Opaque-ID header
	XOpaqueId *string `json:"x_opaque_id,omitempty"`
}

// TasksRecordBuilder holds TasksRecord struct and provides a builder API.
type TasksRecordBuilder struct {
	v *TasksRecord
}

// NewTasksRecord provides a builder for the TasksRecord struct.
func NewTasksRecordBuilder() *TasksRecordBuilder {
	r := TasksRecordBuilder{
		&TasksRecord{},
	}

	return &r
}

// Build finalize the chain and returns the TasksRecord struct
func (rb *TasksRecordBuilder) Build() TasksRecord {
	return *rb.v
}

// Action task action

func (rb *TasksRecordBuilder) Action(action string) *TasksRecordBuilder {
	rb.v.Action = &action
	return rb
}

// Description task action

func (rb *TasksRecordBuilder) Description(description string) *TasksRecordBuilder {
	rb.v.Description = &description
	return rb
}

// Id id of the task with the node

func (rb *TasksRecordBuilder) Id(id Id) *TasksRecordBuilder {
	rb.v.Id = &id
	return rb
}

// Ip ip address

func (rb *TasksRecordBuilder) Ip(ip string) *TasksRecordBuilder {
	rb.v.Ip = &ip
	return rb
}

// Node node name

func (rb *TasksRecordBuilder) Node(node string) *TasksRecordBuilder {
	rb.v.Node = &node
	return rb
}

// NodeId unique node id

func (rb *TasksRecordBuilder) NodeId(nodeid NodeId) *TasksRecordBuilder {
	rb.v.NodeId = &nodeid
	return rb
}

// ParentTaskId parent task id

func (rb *TasksRecordBuilder) ParentTaskId(parenttaskid string) *TasksRecordBuilder {
	rb.v.ParentTaskId = &parenttaskid
	return rb
}

// Port bound transport port

func (rb *TasksRecordBuilder) Port(port string) *TasksRecordBuilder {
	rb.v.Port = &port
	return rb
}

// RunningTime running time

func (rb *TasksRecordBuilder) RunningTime(runningtime string) *TasksRecordBuilder {
	rb.v.RunningTime = &runningtime
	return rb
}

// RunningTimeNs running time ns

func (rb *TasksRecordBuilder) RunningTimeNs(runningtimens string) *TasksRecordBuilder {
	rb.v.RunningTimeNs = &runningtimens
	return rb
}

// StartTime start time in ms

func (rb *TasksRecordBuilder) StartTime(starttime string) *TasksRecordBuilder {
	rb.v.StartTime = &starttime
	return rb
}

// TaskId unique task id

func (rb *TasksRecordBuilder) TaskId(taskid Id) *TasksRecordBuilder {
	rb.v.TaskId = &taskid
	return rb
}

// Timestamp start time in HH:MM:SS

func (rb *TasksRecordBuilder) Timestamp(timestamp string) *TasksRecordBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}

// Type task type

func (rb *TasksRecordBuilder) Type_(type_ string) *TasksRecordBuilder {
	rb.v.Type = &type_
	return rb
}

// Version es version

func (rb *TasksRecordBuilder) Version(version VersionString) *TasksRecordBuilder {
	rb.v.Version = &version
	return rb
}

// XOpaqueId X-Opaque-ID header

func (rb *TasksRecordBuilder) XOpaqueId(xopaqueid string) *TasksRecordBuilder {
	rb.v.XOpaqueId = &xopaqueid
	return rb
}
