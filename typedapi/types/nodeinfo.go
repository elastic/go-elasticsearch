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

// NodeInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L30-L66
type NodeInfo struct {
	Aggregations map[string]NodeInfoAggregation `json:"aggregations,omitempty"`
	Attributes   map[string]string              `json:"attributes"`
	BuildFlavor  string                         `json:"build_flavor"`
	// BuildHash Short hash of the last git commit in this release.
	BuildHash string `json:"build_hash"`
	BuildType string `json:"build_type"`
	// Host The node’s host name.
	Host   Host            `json:"host"`
	Http   *NodeInfoHttp   `json:"http,omitempty"`
	Ingest *NodeInfoIngest `json:"ingest,omitempty"`
	// Ip The node’s IP address.
	Ip      Ip            `json:"ip"`
	Jvm     *NodeJvmInfo  `json:"jvm,omitempty"`
	Modules []PluginStats `json:"modules,omitempty"`
	// Name The node's name
	Name       Name                          `json:"name"`
	Network    *NodeInfoNetwork              `json:"network,omitempty"`
	Os         *NodeOperatingSystemInfo      `json:"os,omitempty"`
	Plugins    []PluginStats                 `json:"plugins,omitempty"`
	Process    *NodeProcessInfo              `json:"process,omitempty"`
	Roles      NodeRoles                     `json:"roles"`
	Settings   *NodeInfoSettings             `json:"settings,omitempty"`
	ThreadPool map[string]NodeThreadPoolInfo `json:"thread_pool,omitempty"`
	// TotalIndexingBuffer Total heap allowed to be used to hold recently indexed documents before they
	// must be written to disk. This size is a shared pool across all shards on this
	// node, and is controlled by Indexing Buffer settings.
	TotalIndexingBuffer *int64 `json:"total_indexing_buffer,omitempty"`
	// TotalIndexingBufferInBytes Same as total_indexing_buffer, but expressed in bytes.
	TotalIndexingBufferInBytes *ByteSize          `json:"total_indexing_buffer_in_bytes,omitempty"`
	Transport                  *NodeInfoTransport `json:"transport,omitempty"`
	// TransportAddress Host and port where transport HTTP connections are accepted.
	TransportAddress TransportAddress `json:"transport_address"`
	// Version Elasticsearch version running on this node.
	Version VersionString `json:"version"`
}

// NodeInfoBuilder holds NodeInfo struct and provides a builder API.
type NodeInfoBuilder struct {
	v *NodeInfo
}

// NewNodeInfo provides a builder for the NodeInfo struct.
func NewNodeInfoBuilder() *NodeInfoBuilder {
	r := NodeInfoBuilder{
		&NodeInfo{
			Aggregations: make(map[string]NodeInfoAggregation, 0),
			Attributes:   make(map[string]string, 0),
			ThreadPool:   make(map[string]NodeThreadPoolInfo, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfo struct
func (rb *NodeInfoBuilder) Build() NodeInfo {
	return *rb.v
}

func (rb *NodeInfoBuilder) Aggregations(values map[string]*NodeInfoAggregationBuilder) *NodeInfoBuilder {
	tmp := make(map[string]NodeInfoAggregation, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *NodeInfoBuilder) Attributes(value map[string]string) *NodeInfoBuilder {
	rb.v.Attributes = value
	return rb
}

func (rb *NodeInfoBuilder) BuildFlavor(buildflavor string) *NodeInfoBuilder {
	rb.v.BuildFlavor = buildflavor
	return rb
}

// BuildHash Short hash of the last git commit in this release.

func (rb *NodeInfoBuilder) BuildHash(buildhash string) *NodeInfoBuilder {
	rb.v.BuildHash = buildhash
	return rb
}

func (rb *NodeInfoBuilder) BuildType(buildtype string) *NodeInfoBuilder {
	rb.v.BuildType = buildtype
	return rb
}

// Host The node’s host name.

func (rb *NodeInfoBuilder) Host(host Host) *NodeInfoBuilder {
	rb.v.Host = host
	return rb
}

func (rb *NodeInfoBuilder) Http(http *NodeInfoHttpBuilder) *NodeInfoBuilder {
	v := http.Build()
	rb.v.Http = &v
	return rb
}

func (rb *NodeInfoBuilder) Ingest(ingest *NodeInfoIngestBuilder) *NodeInfoBuilder {
	v := ingest.Build()
	rb.v.Ingest = &v
	return rb
}

// Ip The node’s IP address.

func (rb *NodeInfoBuilder) Ip(ip Ip) *NodeInfoBuilder {
	rb.v.Ip = ip
	return rb
}

func (rb *NodeInfoBuilder) Jvm(jvm *NodeJvmInfoBuilder) *NodeInfoBuilder {
	v := jvm.Build()
	rb.v.Jvm = &v
	return rb
}

func (rb *NodeInfoBuilder) Modules(modules []PluginStatsBuilder) *NodeInfoBuilder {
	tmp := make([]PluginStats, len(modules))
	for _, value := range modules {
		tmp = append(tmp, value.Build())
	}
	rb.v.Modules = tmp
	return rb
}

// Name The node's name

func (rb *NodeInfoBuilder) Name(name Name) *NodeInfoBuilder {
	rb.v.Name = name
	return rb
}

func (rb *NodeInfoBuilder) Network(network *NodeInfoNetworkBuilder) *NodeInfoBuilder {
	v := network.Build()
	rb.v.Network = &v
	return rb
}

func (rb *NodeInfoBuilder) Os(os *NodeOperatingSystemInfoBuilder) *NodeInfoBuilder {
	v := os.Build()
	rb.v.Os = &v
	return rb
}

func (rb *NodeInfoBuilder) Plugins(plugins []PluginStatsBuilder) *NodeInfoBuilder {
	tmp := make([]PluginStats, len(plugins))
	for _, value := range plugins {
		tmp = append(tmp, value.Build())
	}
	rb.v.Plugins = tmp
	return rb
}

func (rb *NodeInfoBuilder) Process(process *NodeProcessInfoBuilder) *NodeInfoBuilder {
	v := process.Build()
	rb.v.Process = &v
	return rb
}

func (rb *NodeInfoBuilder) Roles(roles *NodeRolesBuilder) *NodeInfoBuilder {
	v := roles.Build()
	rb.v.Roles = v
	return rb
}

func (rb *NodeInfoBuilder) Settings(settings *NodeInfoSettingsBuilder) *NodeInfoBuilder {
	v := settings.Build()
	rb.v.Settings = &v
	return rb
}

func (rb *NodeInfoBuilder) ThreadPool(values map[string]*NodeThreadPoolInfoBuilder) *NodeInfoBuilder {
	tmp := make(map[string]NodeThreadPoolInfo, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.ThreadPool = tmp
	return rb
}

// TotalIndexingBuffer Total heap allowed to be used to hold recently indexed documents before they
// must be written to disk. This size is a shared pool across all shards on this
// node, and is controlled by Indexing Buffer settings.

func (rb *NodeInfoBuilder) TotalIndexingBuffer(totalindexingbuffer int64) *NodeInfoBuilder {
	rb.v.TotalIndexingBuffer = &totalindexingbuffer
	return rb
}

// TotalIndexingBufferInBytes Same as total_indexing_buffer, but expressed in bytes.

func (rb *NodeInfoBuilder) TotalIndexingBufferInBytes(totalindexingbufferinbytes *ByteSizeBuilder) *NodeInfoBuilder {
	v := totalindexingbufferinbytes.Build()
	rb.v.TotalIndexingBufferInBytes = &v
	return rb
}

func (rb *NodeInfoBuilder) Transport(transport *NodeInfoTransportBuilder) *NodeInfoBuilder {
	v := transport.Build()
	rb.v.Transport = &v
	return rb
}

// TransportAddress Host and port where transport HTTP connections are accepted.

func (rb *NodeInfoBuilder) TransportAddress(transportaddress TransportAddress) *NodeInfoBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}

// Version Elasticsearch version running on this node.

func (rb *NodeInfoBuilder) Version(version VersionString) *NodeInfoBuilder {
	rb.v.Version = version
	return rb
}
