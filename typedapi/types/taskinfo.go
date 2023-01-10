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

// TaskInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/tasks/_types/TaskInfo.ts#L32-L46
type TaskInfo struct {
	Action             string            `json:"action"`
	Cancellable        bool              `json:"cancellable"`
	Cancelled          *bool             `json:"cancelled,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Headers            map[string]string `json:"headers"`
	Id                 int64             `json:"id"`
	Node               string            `json:"node"`
	ParentTaskId       *TaskId           `json:"parent_task_id,omitempty"`
	RunningTime        *Duration         `json:"running_time,omitempty"`
	RunningTimeInNanos int64             `json:"running_time_in_nanos"`
	StartTimeInMillis  int64             `json:"start_time_in_millis"`
	Status             *TaskStatus       `json:"status,omitempty"`
	Type               string            `json:"type"`
}

// NewTaskInfo returns a TaskInfo.
func NewTaskInfo() *TaskInfo {
	r := &TaskInfo{
		Headers: make(map[string]string, 0),
	}

	return r
}
