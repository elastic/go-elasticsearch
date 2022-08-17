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

// HotThread type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/hot_threads/types.ts#L23-L28
type HotThread struct {
	Hosts    []Host   `json:"hosts"`
	NodeId   Id       `json:"node_id"`
	NodeName Name     `json:"node_name"`
	Threads  []string `json:"threads"`
}

// HotThreadBuilder holds HotThread struct and provides a builder API.
type HotThreadBuilder struct {
	v *HotThread
}

// NewHotThread provides a builder for the HotThread struct.
func NewHotThreadBuilder() *HotThreadBuilder {
	r := HotThreadBuilder{
		&HotThread{},
	}

	return &r
}

// Build finalize the chain and returns the HotThread struct
func (rb *HotThreadBuilder) Build() HotThread {
	return *rb.v
}

func (rb *HotThreadBuilder) Hosts(hosts ...Host) *HotThreadBuilder {
	rb.v.Hosts = hosts
	return rb
}

func (rb *HotThreadBuilder) NodeId(nodeid Id) *HotThreadBuilder {
	rb.v.NodeId = nodeid
	return rb
}

func (rb *HotThreadBuilder) NodeName(nodename Name) *HotThreadBuilder {
	rb.v.NodeName = nodename
	return rb
}

func (rb *HotThreadBuilder) Threads(threads ...string) *HotThreadBuilder {
	rb.v.Threads = threads
	return rb
}
