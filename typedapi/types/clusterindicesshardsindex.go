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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

// ClusterIndicesShardsIndex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cluster/stats/types.ts#L40-L47
type ClusterIndicesShardsIndex struct {
	// Primaries Contains statistics about the number of primary shards assigned to selected
	// nodes.
	Primaries ClusterShardMetrics `json:"primaries"`
	// Replication Contains statistics about the number of replication shards assigned to
	// selected nodes.
	Replication ClusterShardMetrics `json:"replication"`
	// Shards Contains statistics about the number of shards assigned to selected nodes.
	Shards ClusterShardMetrics `json:"shards"`
}

// NewClusterIndicesShardsIndex returns a ClusterIndicesShardsIndex.
func NewClusterIndicesShardsIndex() *ClusterIndicesShardsIndex {
	r := &ClusterIndicesShardsIndex{}

	return r
}
