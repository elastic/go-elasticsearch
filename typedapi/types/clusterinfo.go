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

// ClusterInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/allocation_explain/types.ts#L48-L54
type ClusterInfo struct {
	Nodes             map[string]NodeDiskUsage `json:"nodes"`
	ReservedSizes     []ReservedSize           `json:"reserved_sizes"`
	ShardDataSetSizes map[string]string        `json:"shard_data_set_sizes,omitempty"`
	ShardPaths        map[string]string        `json:"shard_paths"`
	ShardSizes        map[string]int64         `json:"shard_sizes"`
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
