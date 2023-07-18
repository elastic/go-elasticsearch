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
// https://github.com/elastic/elasticsearch-specification/tree/33e8a1c9cad22a5946ac735c4fba31af2da2cec2

package types

// SerializedClusterState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/33e8a1c9cad22a5946ac735c4fba31af2da2cec2/specification/nodes/_types/Stats.ts#L101-L104
type SerializedClusterState struct {
	Diffs      *SerializedClusterStateDetail `json:"diffs,omitempty"`
	FullStates *SerializedClusterStateDetail `json:"full_states,omitempty"`
}

// NewSerializedClusterState returns a SerializedClusterState.
func NewSerializedClusterState() *SerializedClusterState {
	r := &SerializedClusterState{}

	return r
}
