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
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package types

// NodeBufferPool type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/nodes/_types/Stats.ts#L310-L316
type NodeBufferPool struct {
	Count                *int64  `json:"count,omitempty"`
	TotalCapacity        *string `json:"total_capacity,omitempty"`
	TotalCapacityInBytes *int64  `json:"total_capacity_in_bytes,omitempty"`
	Used                 *string `json:"used,omitempty"`
	UsedInBytes          *int64  `json:"used_in_bytes,omitempty"`
}

// NewNodeBufferPool returns a NodeBufferPool.
func NewNodeBufferPool() *NodeBufferPool {
	r := &NodeBufferPool{}

	return r
}
