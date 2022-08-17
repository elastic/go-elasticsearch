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

// TaskInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/tasks/_types/TaskInfo.ts#L33-L46
type TaskInfo struct {
	Action             string                 `json:"action"`
	Cancellable        bool                   `json:"cancellable"`
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

// TaskInfoBuilder holds TaskInfo struct and provides a builder API.
type TaskInfoBuilder struct {
	v *TaskInfo
}

// NewTaskInfo provides a builder for the TaskInfo struct.
func NewTaskInfoBuilder() *TaskInfoBuilder {
	r := TaskInfoBuilder{
		&TaskInfo{
			Headers: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TaskInfo struct
func (rb *TaskInfoBuilder) Build() TaskInfo {
	return *rb.v
}

func (rb *TaskInfoBuilder) Action(action string) *TaskInfoBuilder {
	rb.v.Action = action
	return rb
}

func (rb *TaskInfoBuilder) Cancellable(cancellable bool) *TaskInfoBuilder {
	rb.v.Cancellable = cancellable
	return rb
}

func (rb *TaskInfoBuilder) Description(description string) *TaskInfoBuilder {
	rb.v.Description = &description
	return rb
}

func (rb *TaskInfoBuilder) Headers(value map[string]string) *TaskInfoBuilder {
	rb.v.Headers = value
	return rb
}

func (rb *TaskInfoBuilder) Id(id int64) *TaskInfoBuilder {
	rb.v.Id = id
	return rb
}

func (rb *TaskInfoBuilder) Node(node NodeId) *TaskInfoBuilder {
	rb.v.Node = node
	return rb
}

func (rb *TaskInfoBuilder) ParentTaskId(parenttaskid *TaskIdBuilder) *TaskInfoBuilder {
	v := parenttaskid.Build()
	rb.v.ParentTaskId = &v
	return rb
}

func (rb *TaskInfoBuilder) RunningTime(runningtime *DurationBuilder) *TaskInfoBuilder {
	v := runningtime.Build()
	rb.v.RunningTime = &v
	return rb
}

func (rb *TaskInfoBuilder) RunningTimeInNanos(runningtimeinnanos *DurationValueUnitNanosBuilder) *TaskInfoBuilder {
	v := runningtimeinnanos.Build()
	rb.v.RunningTimeInNanos = v
	return rb
}

func (rb *TaskInfoBuilder) StartTimeInMillis(starttimeinmillis *EpochTimeUnitMillisBuilder) *TaskInfoBuilder {
	v := starttimeinmillis.Build()
	rb.v.StartTimeInMillis = v
	return rb
}

func (rb *TaskInfoBuilder) Status(status *TaskStatusBuilder) *TaskInfoBuilder {
	v := status.Build()
	rb.v.Status = &v
	return rb
}

func (rb *TaskInfoBuilder) Type_(type_ string) *TaskInfoBuilder {
	rb.v.Type = type_
	return rb
}
