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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ClusterRemoteSniffInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/remote_info/ClusterRemoteInfoResponse.ts#L31-L39
type ClusterRemoteSniffInfo struct {
	Connected                bool     `json:"connected"`
	InitialConnectTimeout    Duration `json:"initial_connect_timeout"`
	MaxConnectionsPerCluster int      `json:"max_connections_per_cluster"`
	Mode                     string   `json:"mode,omitempty"`
	NumNodesConnected        int64    `json:"num_nodes_connected"`
	Seeds                    []string `json:"seeds"`
	SkipUnavailable          bool     `json:"skip_unavailable"`
}

func (s *ClusterRemoteSniffInfo) UnmarshalJSON(data []byte) error {

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

		case "connected":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Connected = value
			case bool:
				s.Connected = v
			}

		case "initial_connect_timeout":
			if err := dec.Decode(&s.InitialConnectTimeout); err != nil {
				return err
			}

		case "max_connections_per_cluster":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxConnectionsPerCluster = value
			case float64:
				f := int(v)
				s.MaxConnectionsPerCluster = f
			}

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return err
			}

		case "num_nodes_connected":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NumNodesConnected = value
			case float64:
				f := int64(v)
				s.NumNodesConnected = f
			}

		case "seeds":
			if err := dec.Decode(&s.Seeds); err != nil {
				return err
			}

		case "skip_unavailable":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
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
func (s ClusterRemoteSniffInfo) MarshalJSON() ([]byte, error) {
	type innerClusterRemoteSniffInfo ClusterRemoteSniffInfo
	tmp := innerClusterRemoteSniffInfo{
		Connected:                s.Connected,
		InitialConnectTimeout:    s.InitialConnectTimeout,
		MaxConnectionsPerCluster: s.MaxConnectionsPerCluster,
		Mode:                     s.Mode,
		NumNodesConnected:        s.NumNodesConnected,
		Seeds:                    s.Seeds,
		SkipUnavailable:          s.SkipUnavailable,
	}

	tmp.Mode = "sniff"

	return json.Marshal(tmp)
}

// NewClusterRemoteSniffInfo returns a ClusterRemoteSniffInfo.
func NewClusterRemoteSniffInfo() *ClusterRemoteSniffInfo {
	r := &ClusterRemoteSniffInfo{}

	return r
}
