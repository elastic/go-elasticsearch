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

// NodeInfoOSCPU type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L331-L340
type NodeInfoOSCPU struct {
	CacheSize        string `json:"cache_size"`
	CacheSizeInBytes int    `json:"cache_size_in_bytes"`
	CoresPerSocket   int    `json:"cores_per_socket"`
	Mhz              int    `json:"mhz"`
	Model            string `json:"model"`
	TotalCores       int    `json:"total_cores"`
	TotalSockets     int    `json:"total_sockets"`
	Vendor           string `json:"vendor"`
}

// NodeInfoOSCPUBuilder holds NodeInfoOSCPU struct and provides a builder API.
type NodeInfoOSCPUBuilder struct {
	v *NodeInfoOSCPU
}

// NewNodeInfoOSCPU provides a builder for the NodeInfoOSCPU struct.
func NewNodeInfoOSCPUBuilder() *NodeInfoOSCPUBuilder {
	r := NodeInfoOSCPUBuilder{
		&NodeInfoOSCPU{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoOSCPU struct
func (rb *NodeInfoOSCPUBuilder) Build() NodeInfoOSCPU {
	return *rb.v
}

func (rb *NodeInfoOSCPUBuilder) CacheSize(cachesize string) *NodeInfoOSCPUBuilder {
	rb.v.CacheSize = cachesize
	return rb
}

func (rb *NodeInfoOSCPUBuilder) CacheSizeInBytes(cachesizeinbytes int) *NodeInfoOSCPUBuilder {
	rb.v.CacheSizeInBytes = cachesizeinbytes
	return rb
}

func (rb *NodeInfoOSCPUBuilder) CoresPerSocket(corespersocket int) *NodeInfoOSCPUBuilder {
	rb.v.CoresPerSocket = corespersocket
	return rb
}

func (rb *NodeInfoOSCPUBuilder) Mhz(mhz int) *NodeInfoOSCPUBuilder {
	rb.v.Mhz = mhz
	return rb
}

func (rb *NodeInfoOSCPUBuilder) Model(model string) *NodeInfoOSCPUBuilder {
	rb.v.Model = model
	return rb
}

func (rb *NodeInfoOSCPUBuilder) TotalCores(totalcores int) *NodeInfoOSCPUBuilder {
	rb.v.TotalCores = totalcores
	return rb
}

func (rb *NodeInfoOSCPUBuilder) TotalSockets(totalsockets int) *NodeInfoOSCPUBuilder {
	rb.v.TotalSockets = totalsockets
	return rb
}

func (rb *NodeInfoOSCPUBuilder) Vendor(vendor string) *NodeInfoOSCPUBuilder {
	rb.v.Vendor = vendor
	return rb
}
