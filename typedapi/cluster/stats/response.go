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

package stats

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// Response holds the response body struct for the package stats
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/stats/ClusterStatsResponse.ts#L53-L55

type Response struct {

	// ClusterName Name of the cluster, based on the cluster name setting.
	ClusterName string `json:"cluster_name"`
	// ClusterUuid Unique identifier for the cluster.
	ClusterUuid string `json:"cluster_uuid"`
	// Indices Contains statistics about indices with shards assigned to selected nodes.
	Indices types.ClusterIndices `json:"indices"`
	// NodeStats Contains statistics about the number of nodes selected by the request’s node
	// filters.
	NodeStats *types.NodeStatistics `json:"_nodes,omitempty"`
	// Nodes Contains statistics about nodes selected by the request’s node filters.
	Nodes types.ClusterNodes `json:"nodes"`
	// Status Health status of the cluster, based on the state of its primary and replica
	// shards.
	Status healthstatus.HealthStatus `json:"status"`
	// Timestamp Unix timestamp, in milliseconds, for the last time the cluster statistics
	// were refreshed.
	Timestamp int64 `json:"timestamp"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
