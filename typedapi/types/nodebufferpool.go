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

// NodeBufferPool type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L310-L316
type NodeBufferPool struct {
	Count                *int64  `json:"count,omitempty"`
	TotalCapacity        *string `json:"total_capacity,omitempty"`
	TotalCapacityInBytes *int64  `json:"total_capacity_in_bytes,omitempty"`
	Used                 *string `json:"used,omitempty"`
	UsedInBytes          *int64  `json:"used_in_bytes,omitempty"`
}

// NodeBufferPoolBuilder holds NodeBufferPool struct and provides a builder API.
type NodeBufferPoolBuilder struct {
	v *NodeBufferPool
}

// NewNodeBufferPool provides a builder for the NodeBufferPool struct.
func NewNodeBufferPoolBuilder() *NodeBufferPoolBuilder {
	r := NodeBufferPoolBuilder{
		&NodeBufferPool{},
	}

	return &r
}

// Build finalize the chain and returns the NodeBufferPool struct
func (rb *NodeBufferPoolBuilder) Build() NodeBufferPool {
	return *rb.v
}

func (rb *NodeBufferPoolBuilder) Count(count int64) *NodeBufferPoolBuilder {
	rb.v.Count = &count
	return rb
}

func (rb *NodeBufferPoolBuilder) TotalCapacity(totalcapacity string) *NodeBufferPoolBuilder {
	rb.v.TotalCapacity = &totalcapacity
	return rb
}

func (rb *NodeBufferPoolBuilder) TotalCapacityInBytes(totalcapacityinbytes int64) *NodeBufferPoolBuilder {
	rb.v.TotalCapacityInBytes = &totalcapacityinbytes
	return rb
}

func (rb *NodeBufferPoolBuilder) Used(used string) *NodeBufferPoolBuilder {
	rb.v.Used = &used
	return rb
}

func (rb *NodeBufferPoolBuilder) UsedInBytes(usedinbytes int64) *NodeBufferPoolBuilder {
	rb.v.UsedInBytes = &usedinbytes
	return rb
}
