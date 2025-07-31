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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// RemoteClusterInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L786-L817
type RemoteClusterInfo struct {
	// ClusterUuid The UUID of the remote cluster.
	ClusterUuid string `json:"cluster_uuid"`
	// IndicesCount The total number of indices in the remote cluster.
	IndicesCount int `json:"indices_count"`
	// IndicesTotalSize Total data set size of all shards assigned to selected nodes, as a
	// human-readable string.
	IndicesTotalSize *string `json:"indices_total_size,omitempty"`
	// IndicesTotalSizeInBytes Total data set size, in bytes, of all shards assigned to selected nodes.
	IndicesTotalSizeInBytes int64 `json:"indices_total_size_in_bytes"`
	// MaxHeap Maximum amount of memory available for use by the heap across the nodes of
	// the remote cluster, as a human-readable string.
	MaxHeap *string `json:"max_heap,omitempty"`
	// MaxHeapInBytes Maximum amount of memory, in bytes, available for use by the heap across the
	// nodes of the remote cluster.
	MaxHeapInBytes int64 `json:"max_heap_in_bytes"`
	// MemTotal Total amount of physical memory across the nodes of the remote cluster, as a
	// human-readable string.
	MemTotal *string `json:"mem_total,omitempty"`
	// MemTotalInBytes Total amount, in bytes, of physical memory across the nodes of the remote
	// cluster.
	MemTotalInBytes int64 `json:"mem_total_in_bytes"`
	// Mode The connection mode used to communicate with the remote cluster.
	Mode string `json:"mode"`
	// NodesCount The total count of nodes in the remote cluster.
	NodesCount int `json:"nodes_count"`
	// ShardsCount The total number of shards in the remote cluster.
	ShardsCount int `json:"shards_count"`
	// SkipUnavailable The `skip_unavailable` setting used for this remote cluster.
	SkipUnavailable bool `json:"skip_unavailable"`
	// Status Health status of the cluster, based on the state of its primary and replica
	// shards.
	Status healthstatus.HealthStatus `json:"status"`
	// TransportCompress Transport compression setting used for this remote cluster.
	TransportCompress string `json:"transport.compress"`
	// Version The list of Elasticsearch versions used by the nodes on the remote cluster.
	Version []string `json:"version"`
}

func (s *RemoteClusterInfo) UnmarshalJSON(data []byte) error {

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

		case "cluster_uuid":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ClusterUuid", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ClusterUuid = o

		case "indices_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndicesCount", err)
				}
				s.IndicesCount = value
			case float64:
				f := int(v)
				s.IndicesCount = f
			}

		case "indices_total_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "IndicesTotalSize", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndicesTotalSize = &o

		case "indices_total_size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndicesTotalSizeInBytes", err)
				}
				s.IndicesTotalSizeInBytes = value
			case float64:
				f := int64(v)
				s.IndicesTotalSizeInBytes = f
			}

		case "max_heap":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MaxHeap", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxHeap = &o

		case "max_heap_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxHeapInBytes", err)
				}
				s.MaxHeapInBytes = value
			case float64:
				f := int64(v)
				s.MaxHeapInBytes = f
			}

		case "mem_total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MemTotal", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MemTotal = &o

		case "mem_total_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MemTotalInBytes", err)
				}
				s.MemTotalInBytes = value
			case float64:
				f := int64(v)
				s.MemTotalInBytes = f
			}

		case "mode":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Mode", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Mode = o

		case "nodes_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NodesCount", err)
				}
				s.NodesCount = value
			case float64:
				f := int(v)
				s.NodesCount = f
			}

		case "shards_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardsCount", err)
				}
				s.ShardsCount = value
			case float64:
				f := int(v)
				s.ShardsCount = f
			}

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

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "transport.compress":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TransportCompress", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TransportCompress = o

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewRemoteClusterInfo returns a RemoteClusterInfo.
func NewRemoteClusterInfo() *RemoteClusterInfo {
	r := &RemoteClusterInfo{}

	return r
}
