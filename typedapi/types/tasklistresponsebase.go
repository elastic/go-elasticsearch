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

// TaskListResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/tasks/_types/TaskListResponseBase.ts#L26-L38
type TaskListResponseBase struct {
	NodeFailures []ErrorCause `json:"node_failures,omitempty"`
	// Nodes Task information grouped by node, if `group_by` was set to `node` (the
	// default).
	Nodes        map[string]NodeTasks `json:"nodes,omitempty"`
	TaskFailures []TaskFailure        `json:"task_failures,omitempty"`
	// Tasks Either a flat list of tasks if `group_by` was set to `none`, or grouped by
	// parents if
	// `group_by` was set to `parents`.
	Tasks *TaskInfos `json:"tasks,omitempty"`
}

// TaskListResponseBaseBuilder holds TaskListResponseBase struct and provides a builder API.
type TaskListResponseBaseBuilder struct {
	v *TaskListResponseBase
}

// NewTaskListResponseBase provides a builder for the TaskListResponseBase struct.
func NewTaskListResponseBaseBuilder() *TaskListResponseBaseBuilder {
	r := TaskListResponseBaseBuilder{
		&TaskListResponseBase{
			Nodes: make(map[string]NodeTasks, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TaskListResponseBase struct
func (rb *TaskListResponseBaseBuilder) Build() TaskListResponseBase {
	return *rb.v
}

func (rb *TaskListResponseBaseBuilder) NodeFailures(node_failures []ErrorCauseBuilder) *TaskListResponseBaseBuilder {
	tmp := make([]ErrorCause, len(node_failures))
	for _, value := range node_failures {
		tmp = append(tmp, value.Build())
	}
	rb.v.NodeFailures = tmp
	return rb
}

// Nodes Task information grouped by node, if `group_by` was set to `node` (the
// default).

func (rb *TaskListResponseBaseBuilder) Nodes(values map[string]*NodeTasksBuilder) *TaskListResponseBaseBuilder {
	tmp := make(map[string]NodeTasks, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Nodes = tmp
	return rb
}

func (rb *TaskListResponseBaseBuilder) TaskFailures(task_failures []TaskFailureBuilder) *TaskListResponseBaseBuilder {
	tmp := make([]TaskFailure, len(task_failures))
	for _, value := range task_failures {
		tmp = append(tmp, value.Build())
	}
	rb.v.TaskFailures = tmp
	return rb
}

// Tasks Either a flat list of tasks if `group_by` was set to `none`, or grouped by
// parents if
// `group_by` was set to `parents`.

func (rb *TaskListResponseBaseBuilder) Tasks(tasks *TaskInfosBuilder) *TaskListResponseBaseBuilder {
	v := tasks.Build()
	rb.v.Tasks = &v
	return rb
}
