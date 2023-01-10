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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// NodeJvmInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/nodes/info/types.ts#L348-L362
type NodeJvmInfo struct {
	BundledJdk                            bool              `json:"bundled_jdk"`
	GcCollectors                          []string          `json:"gc_collectors"`
	InputArguments                        []string          `json:"input_arguments"`
	Mem                                   NodeInfoJvmMemory `json:"mem"`
	MemoryPools                           []string          `json:"memory_pools"`
	Pid                                   int               `json:"pid"`
	StartTimeInMillis                     int64             `json:"start_time_in_millis"`
	UsingBundledJdk                       bool              `json:"using_bundled_jdk"`
	UsingCompressedOrdinaryObjectPointers string            `json:"using_compressed_ordinary_object_pointers,omitempty"`
	Version                               string            `json:"version"`
	VmName                                string            `json:"vm_name"`
	VmVendor                              string            `json:"vm_vendor"`
	VmVersion                             string            `json:"vm_version"`
}

// NewNodeJvmInfo returns a NodeJvmInfo.
func NewNodeJvmInfo() *NodeJvmInfo {
	r := &NodeJvmInfo{}

	return r
}
