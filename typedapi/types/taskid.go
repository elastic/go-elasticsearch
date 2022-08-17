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

// TaskId holds the union for the following types:
//
//	int
//	string
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/common.ts#L106-L106
type TaskId interface{}

// TaskIdBuilder holds TaskId struct and provides a builder API.
type TaskIdBuilder struct {
	v TaskId
}

// NewTaskId provides a builder for the TaskId struct.
func NewTaskIdBuilder() *TaskIdBuilder {
	return &TaskIdBuilder{}
}

// Build finalize the chain and returns the TaskId struct
func (u *TaskIdBuilder) Build() TaskId {
	return u.v
}

func (u *TaskIdBuilder) Int(int int) *TaskIdBuilder {
	u.v = &int
	return u
}

func (u *TaskIdBuilder) String(string string) *TaskIdBuilder {
	u.v = &string
	return u
}
