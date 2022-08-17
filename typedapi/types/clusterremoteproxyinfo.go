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

// ClusterRemoteProxyInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/remote_info/ClusterRemoteInfoResponse.ts#L41-L50
type ClusterRemoteProxyInfo struct {
	Connected                 bool     `json:"connected"`
	InitialConnectTimeout     Duration `json:"initial_connect_timeout"`
	MaxProxySocketConnections int      `json:"max_proxy_socket_connections"`
	Mode                      string   `json:"mode,omitempty"`
	NumProxySocketsConnected  int      `json:"num_proxy_sockets_connected"`
	ProxyAddress              string   `json:"proxy_address"`
	ServerName                string   `json:"server_name"`
	SkipUnavailable           bool     `json:"skip_unavailable"`
}

// ClusterRemoteProxyInfoBuilder holds ClusterRemoteProxyInfo struct and provides a builder API.
type ClusterRemoteProxyInfoBuilder struct {
	v *ClusterRemoteProxyInfo
}

// NewClusterRemoteProxyInfo provides a builder for the ClusterRemoteProxyInfo struct.
func NewClusterRemoteProxyInfoBuilder() *ClusterRemoteProxyInfoBuilder {
	r := ClusterRemoteProxyInfoBuilder{
		&ClusterRemoteProxyInfo{},
	}

	r.v.Mode = "proxy"

	return &r
}

// Build finalize the chain and returns the ClusterRemoteProxyInfo struct
func (rb *ClusterRemoteProxyInfoBuilder) Build() ClusterRemoteProxyInfo {
	return *rb.v
}

func (rb *ClusterRemoteProxyInfoBuilder) Connected(connected bool) *ClusterRemoteProxyInfoBuilder {
	rb.v.Connected = connected
	return rb
}

func (rb *ClusterRemoteProxyInfoBuilder) InitialConnectTimeout(initialconnecttimeout *DurationBuilder) *ClusterRemoteProxyInfoBuilder {
	v := initialconnecttimeout.Build()
	rb.v.InitialConnectTimeout = v
	return rb
}

func (rb *ClusterRemoteProxyInfoBuilder) MaxProxySocketConnections(maxproxysocketconnections int) *ClusterRemoteProxyInfoBuilder {
	rb.v.MaxProxySocketConnections = maxproxysocketconnections
	return rb
}

func (rb *ClusterRemoteProxyInfoBuilder) NumProxySocketsConnected(numproxysocketsconnected int) *ClusterRemoteProxyInfoBuilder {
	rb.v.NumProxySocketsConnected = numproxysocketsconnected
	return rb
}

func (rb *ClusterRemoteProxyInfoBuilder) ProxyAddress(proxyaddress string) *ClusterRemoteProxyInfoBuilder {
	rb.v.ProxyAddress = proxyaddress
	return rb
}

func (rb *ClusterRemoteProxyInfoBuilder) ServerName(servername string) *ClusterRemoteProxyInfoBuilder {
	rb.v.ServerName = servername
	return rb
}

func (rb *ClusterRemoteProxyInfoBuilder) SkipUnavailable(skipunavailable bool) *ClusterRemoteProxyInfoBuilder {
	rb.v.SkipUnavailable = skipunavailable
	return rb
}
