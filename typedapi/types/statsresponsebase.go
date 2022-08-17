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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// StatsResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/ClusterStatsResponse.ts#L25-L53
type StatsResponseBase struct {
	// ClusterName Name of the cluster, based on the Cluster name setting setting.
	ClusterName Name `json:"cluster_name"`
	// ClusterUuid Unique identifier for the cluster.
	ClusterUuid Uuid `json:"cluster_uuid"`
	// Indices Contains statistics about indices with shards assigned to selected nodes.
	Indices ClusterIndices `json:"indices"`
	// NodeStats Contains statistics about the number of nodes selected by the request’s node
	// filters.
	NodeStats *NodeStatistics `json:"_nodes,omitempty"`
	// Nodes Contains statistics about nodes selected by the request’s node filters.
	Nodes ClusterNodes `json:"nodes"`
	// Status Health status of the cluster, based on the state of its primary and replica
	// shards.
	Status healthstatus.HealthStatus `json:"status"`
	// Timestamp Unix timestamp, in milliseconds, of the last time the cluster statistics were
	// refreshed.
	Timestamp int64 `json:"timestamp"`
}

// StatsResponseBaseBuilder holds StatsResponseBase struct and provides a builder API.
type StatsResponseBaseBuilder struct {
	v *StatsResponseBase
}

// NewStatsResponseBase provides a builder for the StatsResponseBase struct.
func NewStatsResponseBaseBuilder() *StatsResponseBaseBuilder {
	r := StatsResponseBaseBuilder{
		&StatsResponseBase{},
	}

	return &r
}

// Build finalize the chain and returns the StatsResponseBase struct
func (rb *StatsResponseBaseBuilder) Build() StatsResponseBase {
	return *rb.v
}

// ClusterName Name of the cluster, based on the Cluster name setting setting.

func (rb *StatsResponseBaseBuilder) ClusterName(clustername Name) *StatsResponseBaseBuilder {
	rb.v.ClusterName = clustername
	return rb
}

// ClusterUuid Unique identifier for the cluster.

func (rb *StatsResponseBaseBuilder) ClusterUuid(clusteruuid Uuid) *StatsResponseBaseBuilder {
	rb.v.ClusterUuid = clusteruuid
	return rb
}

// Indices Contains statistics about indices with shards assigned to selected nodes.

func (rb *StatsResponseBaseBuilder) Indices(indices *ClusterIndicesBuilder) *StatsResponseBaseBuilder {
	v := indices.Build()
	rb.v.Indices = v
	return rb
}

// NodeStats Contains statistics about the number of nodes selected by the request’s node
// filters.

func (rb *StatsResponseBaseBuilder) NodeStats(nodestats *NodeStatisticsBuilder) *StatsResponseBaseBuilder {
	v := nodestats.Build()
	rb.v.NodeStats = &v
	return rb
}

// Nodes Contains statistics about nodes selected by the request’s node filters.

func (rb *StatsResponseBaseBuilder) Nodes(nodes *ClusterNodesBuilder) *StatsResponseBaseBuilder {
	v := nodes.Build()
	rb.v.Nodes = v
	return rb
}

// Status Health status of the cluster, based on the state of its primary and replica
// shards.

func (rb *StatsResponseBaseBuilder) Status(status healthstatus.HealthStatus) *StatsResponseBaseBuilder {
	rb.v.Status = status
	return rb
}

// Timestamp Unix timestamp, in milliseconds, of the last time the cluster statistics were
// refreshed.

func (rb *StatsResponseBaseBuilder) Timestamp(timestamp int64) *StatsResponseBaseBuilder {
	rb.v.Timestamp = timestamp
	return rb
}
