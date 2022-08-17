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

// WatchStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Watch.ts#L49-L56
type WatchStatus struct {
	Actions          Actions         `json:"actions"`
	ExecutionState   *string         `json:"execution_state,omitempty"`
	LastChecked      *DateTime       `json:"last_checked,omitempty"`
	LastMetCondition *DateTime       `json:"last_met_condition,omitempty"`
	State            ActivationState `json:"state"`
	Version          VersionNumber   `json:"version"`
}

// WatchStatusBuilder holds WatchStatus struct and provides a builder API.
type WatchStatusBuilder struct {
	v *WatchStatus
}

// NewWatchStatus provides a builder for the WatchStatus struct.
func NewWatchStatusBuilder() *WatchStatusBuilder {
	r := WatchStatusBuilder{
		&WatchStatus{},
	}

	return &r
}

// Build finalize the chain and returns the WatchStatus struct
func (rb *WatchStatusBuilder) Build() WatchStatus {
	return *rb.v
}

func (rb *WatchStatusBuilder) Actions(actions *ActionsBuilder) *WatchStatusBuilder {
	v := actions.Build()
	rb.v.Actions = v
	return rb
}

func (rb *WatchStatusBuilder) ExecutionState(executionstate string) *WatchStatusBuilder {
	rb.v.ExecutionState = &executionstate
	return rb
}

func (rb *WatchStatusBuilder) LastChecked(lastchecked *DateTimeBuilder) *WatchStatusBuilder {
	v := lastchecked.Build()
	rb.v.LastChecked = &v
	return rb
}

func (rb *WatchStatusBuilder) LastMetCondition(lastmetcondition *DateTimeBuilder) *WatchStatusBuilder {
	v := lastmetcondition.Build()
	rb.v.LastMetCondition = &v
	return rb
}

func (rb *WatchStatusBuilder) State(state *ActivationStateBuilder) *WatchStatusBuilder {
	v := state.Build()
	rb.v.State = v
	return rb
}

func (rb *WatchStatusBuilder) Version(version VersionNumber) *WatchStatusBuilder {
	rb.v.Version = version
	return rb
}
