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

// ParentTaskInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/tasks/_types/TaskListResponseBase.ts#L45-L47
type ParentTaskInfo struct {
	Action             string                 `json:"action"`
	Cancellable        bool                   `json:"cancellable"`
	Children           []TaskInfo             `json:"children,omitempty"`
	Description        *string                `json:"description,omitempty"`
	Headers            map[string]string      `json:"headers"`
	Id                 int64                  `json:"id"`
	Node               NodeId                 `json:"node"`
	ParentTaskId       *TaskId                `json:"parent_task_id,omitempty"`
	RunningTime        *Duration              `json:"running_time,omitempty"`
	RunningTimeInNanos DurationValueUnitNanos `json:"running_time_in_nanos"`
	StartTimeInMillis  EpochTimeUnitMillis    `json:"start_time_in_millis"`
	Status             *TaskStatus            `json:"status,omitempty"`
	Type               string                 `json:"type"`
}

// ParentTaskInfoBuilder holds ParentTaskInfo struct and provides a builder API.
type ParentTaskInfoBuilder struct {
	v *ParentTaskInfo
}

// NewParentTaskInfo provides a builder for the ParentTaskInfo struct.
func NewParentTaskInfoBuilder() *ParentTaskInfoBuilder {
	r := ParentTaskInfoBuilder{
		&ParentTaskInfo{
			Headers: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ParentTaskInfo struct
func (rb *ParentTaskInfoBuilder) Build() ParentTaskInfo {
	return *rb.v
}

func (rb *ParentTaskInfoBuilder) Action(action string) *ParentTaskInfoBuilder {
	rb.v.Action = action
	return rb
}

func (rb *ParentTaskInfoBuilder) Cancellable(cancellable bool) *ParentTaskInfoBuilder {
	rb.v.Cancellable = cancellable
	return rb
}

func (rb *ParentTaskInfoBuilder) Children(children []TaskInfoBuilder) *ParentTaskInfoBuilder {
	tmp := make([]TaskInfo, len(children))
	for _, value := range children {
		tmp = append(tmp, value.Build())
	}
	rb.v.Children = tmp
	return rb
}

func (rb *ParentTaskInfoBuilder) Description(description string) *ParentTaskInfoBuilder {
	rb.v.Description = &description
	return rb
}

func (rb *ParentTaskInfoBuilder) Headers(value map[string]string) *ParentTaskInfoBuilder {
	rb.v.Headers = value
	return rb
}

func (rb *ParentTaskInfoBuilder) Id(id int64) *ParentTaskInfoBuilder {
	rb.v.Id = id
	return rb
}

func (rb *ParentTaskInfoBuilder) Node(node NodeId) *ParentTaskInfoBuilder {
	rb.v.Node = node
	return rb
}

func (rb *ParentTaskInfoBuilder) ParentTaskId(parenttaskid *TaskIdBuilder) *ParentTaskInfoBuilder {
	v := parenttaskid.Build()
	rb.v.ParentTaskId = &v
	return rb
}

func (rb *ParentTaskInfoBuilder) RunningTime(runningtime *DurationBuilder) *ParentTaskInfoBuilder {
	v := runningtime.Build()
	rb.v.RunningTime = &v
	return rb
}

func (rb *ParentTaskInfoBuilder) RunningTimeInNanos(runningtimeinnanos *DurationValueUnitNanosBuilder) *ParentTaskInfoBuilder {
	v := runningtimeinnanos.Build()
	rb.v.RunningTimeInNanos = v
	return rb
}

func (rb *ParentTaskInfoBuilder) StartTimeInMillis(starttimeinmillis *EpochTimeUnitMillisBuilder) *ParentTaskInfoBuilder {
	v := starttimeinmillis.Build()
	rb.v.StartTimeInMillis = v
	return rb
}

func (rb *ParentTaskInfoBuilder) Status(status *TaskStatusBuilder) *ParentTaskInfoBuilder {
	v := status.Build()
	rb.v.Status = &v
	return rb
}

func (rb *ParentTaskInfoBuilder) Type_(type_ string) *ParentTaskInfoBuilder {
	rb.v.Type = type_
	return rb
}
