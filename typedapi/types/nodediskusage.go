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

// NodeDiskUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/allocation_explain/types.ts#L56-L60
type NodeDiskUsage struct {
	LeastAvailable DiskUsage `json:"least_available"`
	MostAvailable  DiskUsage `json:"most_available"`
	NodeName       Name      `json:"node_name"`
}

// NodeDiskUsageBuilder holds NodeDiskUsage struct and provides a builder API.
type NodeDiskUsageBuilder struct {
	v *NodeDiskUsage
}

// NewNodeDiskUsage provides a builder for the NodeDiskUsage struct.
func NewNodeDiskUsageBuilder() *NodeDiskUsageBuilder {
	r := NodeDiskUsageBuilder{
		&NodeDiskUsage{},
	}

	return &r
}

// Build finalize the chain and returns the NodeDiskUsage struct
func (rb *NodeDiskUsageBuilder) Build() NodeDiskUsage {
	return *rb.v
}

func (rb *NodeDiskUsageBuilder) LeastAvailable(leastavailable *DiskUsageBuilder) *NodeDiskUsageBuilder {
	v := leastavailable.Build()
	rb.v.LeastAvailable = v
	return rb
}

func (rb *NodeDiskUsageBuilder) MostAvailable(mostavailable *DiskUsageBuilder) *NodeDiskUsageBuilder {
	v := mostavailable.Build()
	rb.v.MostAvailable = v
	return rb
}

func (rb *NodeDiskUsageBuilder) NodeName(nodename Name) *NodeDiskUsageBuilder {
	rb.v.NodeName = nodename
	return rb
}
