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

// ClusterRemoteSniffInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/remote_info/ClusterRemoteInfoResponse.ts#L31-L39
type ClusterRemoteSniffInfo struct {
	Connected                bool     `json:"connected"`
	InitialConnectTimeout    Duration `json:"initial_connect_timeout"`
	MaxConnectionsPerCluster int      `json:"max_connections_per_cluster"`
	Mode                     string   `json:"mode,omitempty"`
	NumNodesConnected        int64    `json:"num_nodes_connected"`
	Seeds                    []string `json:"seeds"`
	SkipUnavailable          bool     `json:"skip_unavailable"`
}

// ClusterRemoteSniffInfoBuilder holds ClusterRemoteSniffInfo struct and provides a builder API.
type ClusterRemoteSniffInfoBuilder struct {
	v *ClusterRemoteSniffInfo
}

// NewClusterRemoteSniffInfo provides a builder for the ClusterRemoteSniffInfo struct.
func NewClusterRemoteSniffInfoBuilder() *ClusterRemoteSniffInfoBuilder {
	r := ClusterRemoteSniffInfoBuilder{
		&ClusterRemoteSniffInfo{},
	}

	r.v.Mode = "sniff"

	return &r
}

// Build finalize the chain and returns the ClusterRemoteSniffInfo struct
func (rb *ClusterRemoteSniffInfoBuilder) Build() ClusterRemoteSniffInfo {
	return *rb.v
}

func (rb *ClusterRemoteSniffInfoBuilder) Connected(connected bool) *ClusterRemoteSniffInfoBuilder {
	rb.v.Connected = connected
	return rb
}

func (rb *ClusterRemoteSniffInfoBuilder) InitialConnectTimeout(initialconnecttimeout *DurationBuilder) *ClusterRemoteSniffInfoBuilder {
	v := initialconnecttimeout.Build()
	rb.v.InitialConnectTimeout = v
	return rb
}

func (rb *ClusterRemoteSniffInfoBuilder) MaxConnectionsPerCluster(maxconnectionspercluster int) *ClusterRemoteSniffInfoBuilder {
	rb.v.MaxConnectionsPerCluster = maxconnectionspercluster
	return rb
}

func (rb *ClusterRemoteSniffInfoBuilder) NumNodesConnected(numnodesconnected int64) *ClusterRemoteSniffInfoBuilder {
	rb.v.NumNodesConnected = numnodesconnected
	return rb
}

func (rb *ClusterRemoteSniffInfoBuilder) Seeds(seeds ...string) *ClusterRemoteSniffInfoBuilder {
	rb.v.Seeds = seeds
	return rb
}

func (rb *ClusterRemoteSniffInfoBuilder) SkipUnavailable(skipunavailable bool) *ClusterRemoteSniffInfoBuilder {
	rb.v.SkipUnavailable = skipunavailable
	return rb
}
