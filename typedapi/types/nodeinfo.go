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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noderole"
)

// NodeInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/nodes/info/types.ts#L32-L72
type NodeInfo struct {
	Aggregations map[string]NodeInfoAggregation `json:"aggregations,omitempty"`
	Attributes   map[string]string              `json:"attributes"`
	BuildFlavor  string                         `json:"build_flavor"`
	// BuildHash Short hash of the last git commit in this release.
	BuildHash         string         `json:"build_hash"`
	BuildType         string         `json:"build_type"`
	ComponentVersions map[string]int `json:"component_versions"`
	// Host The node’s host name.
	Host         string          `json:"host"`
	Http         *NodeInfoHttp   `json:"http,omitempty"`
	IndexVersion int64           `json:"index_version"`
	Ingest       *NodeInfoIngest `json:"ingest,omitempty"`
	// Ip The node’s IP address.
	Ip      string        `json:"ip"`
	Jvm     *NodeJvmInfo  `json:"jvm,omitempty"`
	Modules []PluginStats `json:"modules,omitempty"`
	// Name The node's name
	Name                string                        `json:"name"`
	Os                  *NodeOperatingSystemInfo      `json:"os,omitempty"`
	Plugins             []PluginStats                 `json:"plugins,omitempty"`
	Process             *NodeProcessInfo              `json:"process,omitempty"`
	RemoteClusterServer *RemoveClusterServer          `json:"remote_cluster_server,omitempty"`
	Roles               []noderole.NodeRole           `json:"roles"`
	Settings            *NodeInfoSettings             `json:"settings,omitempty"`
	ThreadPool          map[string]NodeThreadPoolInfo `json:"thread_pool,omitempty"`
	// TotalIndexingBuffer Total heap allowed to be used to hold recently indexed documents before they
	// must be written to disk. This size is a shared pool across all shards on this
	// node, and is controlled by Indexing Buffer settings.
	TotalIndexingBuffer *int64 `json:"total_indexing_buffer,omitempty"`
	// TotalIndexingBufferInBytes Same as total_indexing_buffer, but expressed in bytes.
	TotalIndexingBufferInBytes ByteSize           `json:"total_indexing_buffer_in_bytes,omitempty"`
	Transport                  *NodeInfoTransport `json:"transport,omitempty"`
	// TransportAddress Host and port where transport HTTP connections are accepted.
	TransportAddress string `json:"transport_address"`
	TransportVersion int64  `json:"transport_version"`
	// Version Elasticsearch version running on this node.
	Version string `json:"version"`
}

func (s *NodeInfo) UnmarshalJSON(data []byte) error {

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

		case "aggregations":
			if s.Aggregations == nil {
				s.Aggregations = make(map[string]NodeInfoAggregation, 0)
			}
			if err := dec.Decode(&s.Aggregations); err != nil {
				return fmt.Errorf("%s | %w", "Aggregations", err)
			}

		case "attributes":
			if s.Attributes == nil {
				s.Attributes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Attributes); err != nil {
				return fmt.Errorf("%s | %w", "Attributes", err)
			}

		case "build_flavor":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "BuildFlavor", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BuildFlavor = o

		case "build_hash":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "BuildHash", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BuildHash = o

		case "build_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "BuildType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BuildType = o

		case "component_versions":
			if s.ComponentVersions == nil {
				s.ComponentVersions = make(map[string]int, 0)
			}
			if err := dec.Decode(&s.ComponentVersions); err != nil {
				return fmt.Errorf("%s | %w", "ComponentVersions", err)
			}

		case "host":
			if err := dec.Decode(&s.Host); err != nil {
				return fmt.Errorf("%s | %w", "Host", err)
			}

		case "http":
			if err := dec.Decode(&s.Http); err != nil {
				return fmt.Errorf("%s | %w", "Http", err)
			}

		case "index_version":
			if err := dec.Decode(&s.IndexVersion); err != nil {
				return fmt.Errorf("%s | %w", "IndexVersion", err)
			}

		case "ingest":
			if err := dec.Decode(&s.Ingest); err != nil {
				return fmt.Errorf("%s | %w", "Ingest", err)
			}

		case "ip":
			if err := dec.Decode(&s.Ip); err != nil {
				return fmt.Errorf("%s | %w", "Ip", err)
			}

		case "jvm":
			if err := dec.Decode(&s.Jvm); err != nil {
				return fmt.Errorf("%s | %w", "Jvm", err)
			}

		case "modules":
			if err := dec.Decode(&s.Modules); err != nil {
				return fmt.Errorf("%s | %w", "Modules", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "os":
			if err := dec.Decode(&s.Os); err != nil {
				return fmt.Errorf("%s | %w", "Os", err)
			}

		case "plugins":
			if err := dec.Decode(&s.Plugins); err != nil {
				return fmt.Errorf("%s | %w", "Plugins", err)
			}

		case "process":
			if err := dec.Decode(&s.Process); err != nil {
				return fmt.Errorf("%s | %w", "Process", err)
			}

		case "remote_cluster_server":
			if err := dec.Decode(&s.RemoteClusterServer); err != nil {
				return fmt.Errorf("%s | %w", "RemoteClusterServer", err)
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return fmt.Errorf("%s | %w", "Roles", err)
			}

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		case "thread_pool":
			if s.ThreadPool == nil {
				s.ThreadPool = make(map[string]NodeThreadPoolInfo, 0)
			}
			if err := dec.Decode(&s.ThreadPool); err != nil {
				return fmt.Errorf("%s | %w", "ThreadPool", err)
			}

		case "total_indexing_buffer":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalIndexingBuffer", err)
				}
				s.TotalIndexingBuffer = &value
			case float64:
				f := int64(v)
				s.TotalIndexingBuffer = &f
			}

		case "total_indexing_buffer_in_bytes":
			if err := dec.Decode(&s.TotalIndexingBufferInBytes); err != nil {
				return fmt.Errorf("%s | %w", "TotalIndexingBufferInBytes", err)
			}

		case "transport":
			if err := dec.Decode(&s.Transport); err != nil {
				return fmt.Errorf("%s | %w", "Transport", err)
			}

		case "transport_address":
			if err := dec.Decode(&s.TransportAddress); err != nil {
				return fmt.Errorf("%s | %w", "TransportAddress", err)
			}

		case "transport_version":
			if err := dec.Decode(&s.TransportVersion); err != nil {
				return fmt.Errorf("%s | %w", "TransportVersion", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewNodeInfo returns a NodeInfo.
func NewNodeInfo() *NodeInfo {
	r := &NodeInfo{
		Aggregations:      make(map[string]NodeInfoAggregation),
		Attributes:        make(map[string]string),
		ComponentVersions: make(map[string]int),
		ThreadPool:        make(map[string]NodeThreadPoolInfo),
	}

	return r
}
