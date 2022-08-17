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

// TaskInfos holds the union for the following types:
//
//	map[string]ParentTaskInfo
//	[]TaskInfo
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/tasks/_types/TaskListResponseBase.ts#L40-L43
type TaskInfos interface{}

// TaskInfosBuilder holds TaskInfos struct and provides a builder API.
type TaskInfosBuilder struct {
	v TaskInfos
}

// NewTaskInfos provides a builder for the TaskInfos struct.
func NewTaskInfosBuilder() *TaskInfosBuilder {
	return &TaskInfosBuilder{}
}

// Build finalize the chain and returns the TaskInfos struct
func (u *TaskInfosBuilder) Build() TaskInfos {
	return u.v
}

func (u *TaskInfosBuilder) Map(values map[string]*ParentTaskInfoBuilder) *TaskInfosBuilder {
	tmp := make(map[string]ParentTaskInfo, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	u.v = tmp
	return u
}

func (u *TaskInfosBuilder) TaskInfos(taskinfos []TaskInfoBuilder) *TaskInfosBuilder {
	tmp := make([]TaskInfo, len(taskinfos))
	for _, value := range taskinfos {
		tmp = append(tmp, value.Build())
	}
	u.v = tmp
	return u
}
