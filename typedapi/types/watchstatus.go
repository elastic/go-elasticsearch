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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

// WatchStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/watcher/_types/Watch.ts#L49-L56
type WatchStatus struct {
	Actions          WatcherStatusActions `json:"actions"`
	ExecutionState   *string              `json:"execution_state,omitempty"`
	LastChecked      *DateTime            `json:"last_checked,omitempty"`
	LastMetCondition *DateTime            `json:"last_met_condition,omitempty"`
	State            ActivationState      `json:"state"`
	Version          int64                `json:"version"`
}

// NewWatchStatus returns a WatchStatus.
func NewWatchStatus() *WatchStatus {
	r := &WatchStatus{}

	return r
}
