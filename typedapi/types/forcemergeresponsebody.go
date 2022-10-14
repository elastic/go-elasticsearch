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
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package types

// ForceMergeResponseBody type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/forcemerge/_types/response.ts#L22-L28
type ForceMergeResponseBody struct {
	Shards_ ShardStatistics `json:"_shards"`
	// Task task contains a task id returned when wait_for_completion=false,
	// you can use the task_id to get the status of the task at _tasks/<task_id>
	Task *string `json:"task,omitempty"`
}

// ForceMergeResponseBodyBuilder holds ForceMergeResponseBody struct and provides a builder API.
type ForceMergeResponseBodyBuilder struct {
	v *ForceMergeResponseBody
}

// NewForceMergeResponseBody provides a builder for the ForceMergeResponseBody struct.
func NewForceMergeResponseBodyBuilder() *ForceMergeResponseBodyBuilder {
	r := ForceMergeResponseBodyBuilder{
		&ForceMergeResponseBody{},
	}

	return &r
}

// Build finalize the chain and returns the ForceMergeResponseBody struct
func (rb *ForceMergeResponseBodyBuilder) Build() ForceMergeResponseBody {
	return *rb.v
}

func (rb *ForceMergeResponseBodyBuilder) Shards_(shards_ *ShardStatisticsBuilder) *ForceMergeResponseBodyBuilder {
	v := shards_.Build()
	rb.v.Shards_ = v
	return rb
}

// Task task contains a task id returned when wait_for_completion=false,
// you can use the task_id to get the status of the task at _tasks/<task_id>

func (rb *ForceMergeResponseBodyBuilder) Task(task string) *ForceMergeResponseBodyBuilder {
	rb.v.Task = &task
	return rb
}
