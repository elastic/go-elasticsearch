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
// https://github.com/elastic/elasticsearch-specification/tree/1ad7fe36297b3a8e187b2259dedaf68a47bc236e

package types

// NodesIndexingPressureMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1ad7fe36297b3a8e187b2259dedaf68a47bc236e/specification/nodes/_types/Stats.ts#L59-L64
type NodesIndexingPressureMemory struct {
	Current      *PressureMemory `json:"current,omitempty"`
	Limit        ByteSize        `json:"limit,omitempty"`
	LimitInBytes *int64          `json:"limit_in_bytes,omitempty"`
	Total        *PressureMemory `json:"total,omitempty"`
}

// NewNodesIndexingPressureMemory returns a NodesIndexingPressureMemory.
func NewNodesIndexingPressureMemory() *NodesIndexingPressureMemory {
	r := &NodesIndexingPressureMemory{}

	return r
}
