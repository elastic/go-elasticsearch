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

// ClusterIndicesShards type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/cluster/stats/types.ts#L49-L61
type ClusterIndicesShards struct {
	// Index Contains statistics about shards assigned to selected nodes.
	Index *ClusterIndicesShardsIndex `json:"index,omitempty"`
	// Primaries Number of primary shards assigned to selected nodes.
	Primaries *Float64 `json:"primaries,omitempty"`
	// Replication Ratio of replica shards to primary shards across all selected nodes.
	Replication *Float64 `json:"replication,omitempty"`
	// Total Total number of shards assigned to selected nodes.
	Total *Float64 `json:"total,omitempty"`
}

// NewClusterIndicesShards returns a ClusterIndicesShards.
func NewClusterIndicesShards() *ClusterIndicesShards {
	r := &ClusterIndicesShards{}

	return r
}
