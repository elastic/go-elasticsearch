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

// NodeThreadPoolInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L286-L293
type NodeThreadPoolInfo struct {
	Core      *int      `json:"core,omitempty"`
	KeepAlive *Duration `json:"keep_alive,omitempty"`
	Max       *int      `json:"max,omitempty"`
	QueueSize int       `json:"queue_size"`
	Size      *int      `json:"size,omitempty"`
	Type      string    `json:"type"`
}

// NodeThreadPoolInfoBuilder holds NodeThreadPoolInfo struct and provides a builder API.
type NodeThreadPoolInfoBuilder struct {
	v *NodeThreadPoolInfo
}

// NewNodeThreadPoolInfo provides a builder for the NodeThreadPoolInfo struct.
func NewNodeThreadPoolInfoBuilder() *NodeThreadPoolInfoBuilder {
	r := NodeThreadPoolInfoBuilder{
		&NodeThreadPoolInfo{},
	}

	return &r
}

// Build finalize the chain and returns the NodeThreadPoolInfo struct
func (rb *NodeThreadPoolInfoBuilder) Build() NodeThreadPoolInfo {
	return *rb.v
}

func (rb *NodeThreadPoolInfoBuilder) Core(core int) *NodeThreadPoolInfoBuilder {
	rb.v.Core = &core
	return rb
}

func (rb *NodeThreadPoolInfoBuilder) KeepAlive(keepalive *DurationBuilder) *NodeThreadPoolInfoBuilder {
	v := keepalive.Build()
	rb.v.KeepAlive = &v
	return rb
}

func (rb *NodeThreadPoolInfoBuilder) Max(max int) *NodeThreadPoolInfoBuilder {
	rb.v.Max = &max
	return rb
}

func (rb *NodeThreadPoolInfoBuilder) QueueSize(queuesize int) *NodeThreadPoolInfoBuilder {
	rb.v.QueueSize = queuesize
	return rb
}

func (rb *NodeThreadPoolInfoBuilder) Size(size int) *NodeThreadPoolInfoBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *NodeThreadPoolInfoBuilder) Type_(type_ string) *NodeThreadPoolInfoBuilder {
	rb.v.Type = type_
	return rb
}
