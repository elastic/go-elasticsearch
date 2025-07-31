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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ClusterRemoteProxyInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/remote_info/ClusterRemoteInfoResponse.ts#L58-L83
type ClusterRemoteProxyInfo struct {
	// ClusterCredentials This field is present and has a value of `::es_redacted::` only when the
	// remote cluster is configured with the API key based model. Otherwise, the
	// field is not present.
	ClusterCredentials *string `json:"cluster_credentials,omitempty"`
	// Connected If it is `true`, there is at least one open connection to the remote cluster.
	// If it is `false`, it means that the cluster no longer has an open connection
	// to the remote cluster.
	// It does not necessarily mean that the remote cluster is down or unavailable,
	// just that at some point a connection was lost.
	Connected bool `json:"connected"`
	// InitialConnectTimeout The initial connect timeout for remote cluster connections.
	InitialConnectTimeout Duration `json:"initial_connect_timeout"`
	// MaxProxySocketConnections The maximum number of socket connections to the remote cluster when proxy
	// mode is configured.
	MaxProxySocketConnections int `json:"max_proxy_socket_connections"`
	// Mode The connection mode for the remote cluster.
	Mode string `json:"mode,omitempty"`
	// NumProxySocketsConnected The number of open socket connections to the remote cluster when proxy mode
	// is configured.
	NumProxySocketsConnected int `json:"num_proxy_sockets_connected"`
	// ProxyAddress The address for remote connections when proxy mode is configured.
	ProxyAddress string `json:"proxy_address"`
	ServerName   string `json:"server_name"`
	// SkipUnavailable If `true`, cross-cluster search skips the remote cluster when its nodes are
	// unavailable during the search and ignores errors returned by the remote
	// cluster.
	SkipUnavailable bool `json:"skip_unavailable"`
}

func (s *ClusterRemoteProxyInfo) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "cluster_credentials":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ClusterCredentials", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ClusterCredentials = &o

		case "connected":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Connected", err)
				}
				s.Connected = value
			case bool:
				s.Connected = v
			}

		case "initial_connect_timeout":
			if err := dec.Decode(&s.InitialConnectTimeout); err != nil {
				return fmt.Errorf("%s | %w", "InitialConnectTimeout", err)
			}

		case "max_proxy_socket_connections":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxProxySocketConnections", err)
				}
				s.MaxProxySocketConnections = value
			case float64:
				f := int(v)
				s.MaxProxySocketConnections = f
			}

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return fmt.Errorf("%s | %w", "Mode", err)
			}

		case "num_proxy_sockets_connected":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumProxySocketsConnected", err)
				}
				s.NumProxySocketsConnected = value
			case float64:
				f := int(v)
				s.NumProxySocketsConnected = f
			}

		case "proxy_address":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ProxyAddress", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ProxyAddress = o

		case "server_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ServerName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ServerName = o

		case "skip_unavailable":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SkipUnavailable", err)
				}
				s.SkipUnavailable = value
			case bool:
				s.SkipUnavailable = v
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s ClusterRemoteProxyInfo) MarshalJSON() ([]byte, error) {
	type innerClusterRemoteProxyInfo ClusterRemoteProxyInfo
	tmp := innerClusterRemoteProxyInfo{
		ClusterCredentials:        s.ClusterCredentials,
		Connected:                 s.Connected,
		InitialConnectTimeout:     s.InitialConnectTimeout,
		MaxProxySocketConnections: s.MaxProxySocketConnections,
		Mode:                      s.Mode,
		NumProxySocketsConnected:  s.NumProxySocketsConnected,
		ProxyAddress:              s.ProxyAddress,
		ServerName:                s.ServerName,
		SkipUnavailable:           s.SkipUnavailable,
	}

	tmp.Mode = "proxy"

	return json.Marshal(tmp)
}

// NewClusterRemoteProxyInfo returns a ClusterRemoteProxyInfo.
func NewClusterRemoteProxyInfo() *ClusterRemoteProxyInfo {
	r := &ClusterRemoteProxyInfo{}

	return r
}
