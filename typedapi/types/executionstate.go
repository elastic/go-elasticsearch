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

// ExecutionState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Action.ts#L117-L121
type ExecutionState struct {
	Reason     *string  `json:"reason,omitempty"`
	Successful bool     `json:"successful"`
	Timestamp  DateTime `json:"timestamp"`
}

// ExecutionStateBuilder holds ExecutionState struct and provides a builder API.
type ExecutionStateBuilder struct {
	v *ExecutionState
}

// NewExecutionState provides a builder for the ExecutionState struct.
func NewExecutionStateBuilder() *ExecutionStateBuilder {
	r := ExecutionStateBuilder{
		&ExecutionState{},
	}

	return &r
}

// Build finalize the chain and returns the ExecutionState struct
func (rb *ExecutionStateBuilder) Build() ExecutionState {
	return *rb.v
}

func (rb *ExecutionStateBuilder) Reason(reason string) *ExecutionStateBuilder {
	rb.v.Reason = &reason
	return rb
}

func (rb *ExecutionStateBuilder) Successful(successful bool) *ExecutionStateBuilder {
	rb.v.Successful = successful
	return rb
}

func (rb *ExecutionStateBuilder) Timestamp(timestamp *DateTimeBuilder) *ExecutionStateBuilder {
	v := timestamp.Build()
	rb.v.Timestamp = v
	return rb
}
