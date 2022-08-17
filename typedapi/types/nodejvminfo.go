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

// NodeJvmInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L348-L362
type NodeJvmInfo struct {
	BundledJdk                            bool                `json:"bundled_jdk"`
	GcCollectors                          []string            `json:"gc_collectors"`
	InputArguments                        []string            `json:"input_arguments"`
	Mem                                   NodeInfoJvmMemory   `json:"mem"`
	MemoryPools                           []string            `json:"memory_pools"`
	Pid                                   int                 `json:"pid"`
	StartTimeInMillis                     EpochTimeUnitMillis `json:"start_time_in_millis"`
	UsingBundledJdk                       bool                `json:"using_bundled_jdk"`
	UsingCompressedOrdinaryObjectPointers string              `json:"using_compressed_ordinary_object_pointers,omitempty"`
	Version                               VersionString       `json:"version"`
	VmName                                Name                `json:"vm_name"`
	VmVendor                              string              `json:"vm_vendor"`
	VmVersion                             VersionString       `json:"vm_version"`
}

// NodeJvmInfoBuilder holds NodeJvmInfo struct and provides a builder API.
type NodeJvmInfoBuilder struct {
	v *NodeJvmInfo
}

// NewNodeJvmInfo provides a builder for the NodeJvmInfo struct.
func NewNodeJvmInfoBuilder() *NodeJvmInfoBuilder {
	r := NodeJvmInfoBuilder{
		&NodeJvmInfo{},
	}

	return &r
}

// Build finalize the chain and returns the NodeJvmInfo struct
func (rb *NodeJvmInfoBuilder) Build() NodeJvmInfo {
	return *rb.v
}

func (rb *NodeJvmInfoBuilder) BundledJdk(bundledjdk bool) *NodeJvmInfoBuilder {
	rb.v.BundledJdk = bundledjdk
	return rb
}

func (rb *NodeJvmInfoBuilder) GcCollectors(gc_collectors ...string) *NodeJvmInfoBuilder {
	rb.v.GcCollectors = gc_collectors
	return rb
}

func (rb *NodeJvmInfoBuilder) InputArguments(input_arguments ...string) *NodeJvmInfoBuilder {
	rb.v.InputArguments = input_arguments
	return rb
}

func (rb *NodeJvmInfoBuilder) Mem(mem *NodeInfoJvmMemoryBuilder) *NodeJvmInfoBuilder {
	v := mem.Build()
	rb.v.Mem = v
	return rb
}

func (rb *NodeJvmInfoBuilder) MemoryPools(memory_pools ...string) *NodeJvmInfoBuilder {
	rb.v.MemoryPools = memory_pools
	return rb
}

func (rb *NodeJvmInfoBuilder) Pid(pid int) *NodeJvmInfoBuilder {
	rb.v.Pid = pid
	return rb
}

func (rb *NodeJvmInfoBuilder) StartTimeInMillis(starttimeinmillis *EpochTimeUnitMillisBuilder) *NodeJvmInfoBuilder {
	v := starttimeinmillis.Build()
	rb.v.StartTimeInMillis = v
	return rb
}

func (rb *NodeJvmInfoBuilder) UsingBundledJdk(usingbundledjdk bool) *NodeJvmInfoBuilder {
	rb.v.UsingBundledJdk = usingbundledjdk
	return rb
}

func (rb *NodeJvmInfoBuilder) UsingCompressedOrdinaryObjectPointers(arg string) *NodeJvmInfoBuilder {
	rb.v.UsingCompressedOrdinaryObjectPointers = arg
	return rb
}

func (rb *NodeJvmInfoBuilder) Version(version VersionString) *NodeJvmInfoBuilder {
	rb.v.Version = version
	return rb
}

func (rb *NodeJvmInfoBuilder) VmName(vmname Name) *NodeJvmInfoBuilder {
	rb.v.VmName = vmname
	return rb
}

func (rb *NodeJvmInfoBuilder) VmVendor(vmvendor string) *NodeJvmInfoBuilder {
	rb.v.VmVendor = vmvendor
	return rb
}

func (rb *NodeJvmInfoBuilder) VmVersion(vmversion VersionString) *NodeJvmInfoBuilder {
	rb.v.VmVersion = vmversion
	return rb
}
