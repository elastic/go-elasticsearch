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

// NodeOperatingSystemInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L364-L381
type NodeOperatingSystemInfo struct {
	// AllocatedProcessors The number of processors actually used to calculate thread pool size. This
	// number can be set with the node.processors setting of a node and defaults to
	// the number of processors reported by the OS.
	AllocatedProcessors *int `json:"allocated_processors,omitempty"`
	// Arch Name of the JVM architecture (ex: amd64, x86)
	Arch string `json:"arch"`
	// AvailableProcessors Number of processors available to the Java virtual machine
	AvailableProcessors int             `json:"available_processors"`
	Cpu                 *NodeInfoOSCPU  `json:"cpu,omitempty"`
	Mem                 *NodeInfoMemory `json:"mem,omitempty"`
	// Name Name of the operating system (ex: Linux, Windows, Mac OS X)
	Name       Name `json:"name"`
	PrettyName Name `json:"pretty_name"`
	// RefreshIntervalInMillis Refresh interval for the OS statistics
	RefreshIntervalInMillis DurationValueUnitMillis `json:"refresh_interval_in_millis"`
	Swap                    *NodeInfoMemory         `json:"swap,omitempty"`
	// Version Version of the operating system
	Version VersionString `json:"version"`
}

// NodeOperatingSystemInfoBuilder holds NodeOperatingSystemInfo struct and provides a builder API.
type NodeOperatingSystemInfoBuilder struct {
	v *NodeOperatingSystemInfo
}

// NewNodeOperatingSystemInfo provides a builder for the NodeOperatingSystemInfo struct.
func NewNodeOperatingSystemInfoBuilder() *NodeOperatingSystemInfoBuilder {
	r := NodeOperatingSystemInfoBuilder{
		&NodeOperatingSystemInfo{},
	}

	return &r
}

// Build finalize the chain and returns the NodeOperatingSystemInfo struct
func (rb *NodeOperatingSystemInfoBuilder) Build() NodeOperatingSystemInfo {
	return *rb.v
}

// AllocatedProcessors The number of processors actually used to calculate thread pool size. This
// number can be set with the node.processors setting of a node and defaults to
// the number of processors reported by the OS.

func (rb *NodeOperatingSystemInfoBuilder) AllocatedProcessors(allocatedprocessors int) *NodeOperatingSystemInfoBuilder {
	rb.v.AllocatedProcessors = &allocatedprocessors
	return rb
}

// Arch Name of the JVM architecture (ex: amd64, x86)

func (rb *NodeOperatingSystemInfoBuilder) Arch(arch string) *NodeOperatingSystemInfoBuilder {
	rb.v.Arch = arch
	return rb
}

// AvailableProcessors Number of processors available to the Java virtual machine

func (rb *NodeOperatingSystemInfoBuilder) AvailableProcessors(availableprocessors int) *NodeOperatingSystemInfoBuilder {
	rb.v.AvailableProcessors = availableprocessors
	return rb
}

func (rb *NodeOperatingSystemInfoBuilder) Cpu(cpu *NodeInfoOSCPUBuilder) *NodeOperatingSystemInfoBuilder {
	v := cpu.Build()
	rb.v.Cpu = &v
	return rb
}

func (rb *NodeOperatingSystemInfoBuilder) Mem(mem *NodeInfoMemoryBuilder) *NodeOperatingSystemInfoBuilder {
	v := mem.Build()
	rb.v.Mem = &v
	return rb
}

// Name Name of the operating system (ex: Linux, Windows, Mac OS X)

func (rb *NodeOperatingSystemInfoBuilder) Name(name Name) *NodeOperatingSystemInfoBuilder {
	rb.v.Name = name
	return rb
}

func (rb *NodeOperatingSystemInfoBuilder) PrettyName(prettyname Name) *NodeOperatingSystemInfoBuilder {
	rb.v.PrettyName = prettyname
	return rb
}

// RefreshIntervalInMillis Refresh interval for the OS statistics

func (rb *NodeOperatingSystemInfoBuilder) RefreshIntervalInMillis(refreshintervalinmillis *DurationValueUnitMillisBuilder) *NodeOperatingSystemInfoBuilder {
	v := refreshintervalinmillis.Build()
	rb.v.RefreshIntervalInMillis = v
	return rb
}

func (rb *NodeOperatingSystemInfoBuilder) Swap(swap *NodeInfoMemoryBuilder) *NodeOperatingSystemInfoBuilder {
	v := swap.Build()
	rb.v.Swap = &v
	return rb
}

// Version Version of the operating system

func (rb *NodeOperatingSystemInfoBuilder) Version(version VersionString) *NodeOperatingSystemInfoBuilder {
	rb.v.Version = version
	return rb
}
