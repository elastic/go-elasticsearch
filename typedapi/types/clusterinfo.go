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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// ClusterInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/cluster/allocation_explain/types.ts#L48-L54
type ClusterInfo struct {
	Nodes             map[string]NodeDiskUsage `json:"nodes"`
	ReservedSizes     []ReservedSize           `json:"reserved_sizes"`
	ShardDataSetSizes map[string]string        `json:"shard_data_set_sizes,omitempty"`
	ShardPaths        map[string]string        `json:"shard_paths"`
	ShardSizes        map[string]int64         `json:"shard_sizes"`
}

func (s *ClusterInfo) UnmarshalJSON(data []byte) error {

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

		case "nodes":
			if s.Nodes == nil {
				s.Nodes = make(map[string]NodeDiskUsage, 0)
			}
			if err := dec.Decode(&s.Nodes); err != nil {
				return err
			}

		case "reserved_sizes":
			if err := dec.Decode(&s.ReservedSizes); err != nil {
				return err
			}

		case "shard_data_set_sizes":
			if s.ShardDataSetSizes == nil {
				s.ShardDataSetSizes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.ShardDataSetSizes); err != nil {
				return err
			}

		case "shard_paths":
			if s.ShardPaths == nil {
				s.ShardPaths = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.ShardPaths); err != nil {
				return err
			}

		case "shard_sizes":
			if s.ShardSizes == nil {
				s.ShardSizes = make(map[string]int64, 0)
			}
			if err := dec.Decode(&s.ShardSizes); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewClusterInfo returns a ClusterInfo.
func NewClusterInfo() *ClusterInfo {
	r := &ClusterInfo{
		Nodes:             make(map[string]NodeDiskUsage, 0),
		ShardDataSetSizes: make(map[string]string, 0),
		ShardPaths:        make(map[string]string, 0),
		ShardSizes:        make(map[string]int64, 0),
	}

	return r
}
