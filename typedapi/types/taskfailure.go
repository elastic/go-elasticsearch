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

// TaskFailure type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Errors.ts#L66-L71
type TaskFailure struct {
	NodeId NodeId     `json:"node_id"`
	Reason ErrorCause `json:"reason"`
	Status string     `json:"status"`
	TaskId int64      `json:"task_id"`
}

// TaskFailureBuilder holds TaskFailure struct and provides a builder API.
type TaskFailureBuilder struct {
	v *TaskFailure
}

// NewTaskFailure provides a builder for the TaskFailure struct.
func NewTaskFailureBuilder() *TaskFailureBuilder {
	r := TaskFailureBuilder{
		&TaskFailure{},
	}

	return &r
}

// Build finalize the chain and returns the TaskFailure struct
func (rb *TaskFailureBuilder) Build() TaskFailure {
	return *rb.v
}

func (rb *TaskFailureBuilder) NodeId(nodeid NodeId) *TaskFailureBuilder {
	rb.v.NodeId = nodeid
	return rb
}

func (rb *TaskFailureBuilder) Reason(reason *ErrorCauseBuilder) *TaskFailureBuilder {
	v := reason.Build()
	rb.v.Reason = v
	return rb
}

func (rb *TaskFailureBuilder) Status(status string) *TaskFailureBuilder {
	rb.v.Status = status
	return rb
}

func (rb *TaskFailureBuilder) TaskId(taskid int64) *TaskFailureBuilder {
	rb.v.TaskId = taskid
	return rb
}
