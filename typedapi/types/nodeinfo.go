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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noderole"
)

// NodeInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/nodes/info/types.ts#L30-L66
type NodeInfo struct {
	Aggregations map[string]NodeInfoAggregation `json:"aggregations,omitempty"`
	Attributes   map[string]string              `json:"attributes"`
	BuildFlavor  string                         `json:"build_flavor"`
	// BuildHash Short hash of the last git commit in this release.
	BuildHash string `json:"build_hash"`
	BuildType string `json:"build_type"`
	// Host The node’s host name.
	Host   string          `json:"host"`
	Http   *NodeInfoHttp   `json:"http,omitempty"`
	Ingest *NodeInfoIngest `json:"ingest,omitempty"`
	// Ip The node’s IP address.
	Ip      string        `json:"ip"`
	Jvm     *NodeJvmInfo  `json:"jvm,omitempty"`
	Modules []PluginStats `json:"modules,omitempty"`
	// Name The node's name
	Name       string                        `json:"name"`
	Network    *NodeInfoNetwork              `json:"network,omitempty"`
	Os         *NodeOperatingSystemInfo      `json:"os,omitempty"`
	Plugins    []PluginStats                 `json:"plugins,omitempty"`
	Process    *NodeProcessInfo              `json:"process,omitempty"`
	Roles      []noderole.NodeRole           `json:"roles"`
	Settings   *NodeInfoSettings             `json:"settings,omitempty"`
	ThreadPool map[string]NodeThreadPoolInfo `json:"thread_pool,omitempty"`
	// TotalIndexingBuffer Total heap allowed to be used to hold recently indexed documents before they
	// must be written to disk. This size is a shared pool across all shards on this
	// node, and is controlled by Indexing Buffer settings.
	TotalIndexingBuffer *int64 `json:"total_indexing_buffer,omitempty"`
	// TotalIndexingBufferInBytes Same as total_indexing_buffer, but expressed in bytes.
	TotalIndexingBufferInBytes ByteSize           `json:"total_indexing_buffer_in_bytes,omitempty"`
	Transport                  *NodeInfoTransport `json:"transport,omitempty"`
	// TransportAddress Host and port where transport HTTP connections are accepted.
	TransportAddress string `json:"transport_address"`
	// Version Elasticsearch version running on this node.
	Version string `json:"version"`
}

// NewNodeInfo returns a NodeInfo.
func NewNodeInfo() *NodeInfo {
	r := &NodeInfo{
		Aggregations: make(map[string]NodeInfoAggregation, 0),
		Attributes:   make(map[string]string, 0),
		ThreadPool:   make(map[string]NodeThreadPoolInfo, 0),
	}

	return r
}
