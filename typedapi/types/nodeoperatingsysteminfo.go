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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// NodeOperatingSystemInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/nodes/info/types.ts#L364-L381
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
	Name       string `json:"name"`
	PrettyName string `json:"pretty_name"`
	// RefreshIntervalInMillis Refresh interval for the OS statistics
	RefreshIntervalInMillis int64           `json:"refresh_interval_in_millis"`
	Swap                    *NodeInfoMemory `json:"swap,omitempty"`
	// Version Version of the operating system
	Version string `json:"version"`
}

// NewNodeOperatingSystemInfo returns a NodeOperatingSystemInfo.
func NewNodeOperatingSystemInfo() *NodeOperatingSystemInfo {
	r := &NodeOperatingSystemInfo{}

	return r
}
