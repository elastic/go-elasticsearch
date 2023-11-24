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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package health

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// Response holds the response body struct for the package health
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/health/ClusterHealthResponse.ts#L26-L37
type Response struct {

	// ActivePrimaryShards The number of active primary shards.
	ActivePrimaryShards int `json:"active_primary_shards"`
	// ActiveShards The total number of active primary and replica shards.
	ActiveShards int `json:"active_shards"`
	// ActiveShardsPercentAsNumber The ratio of active shards in the cluster expressed as a percentage.
	ActiveShardsPercentAsNumber types.Percentage `json:"active_shards_percent_as_number"`
	// ClusterName The name of the cluster.
	ClusterName string `json:"cluster_name"`
	// DelayedUnassignedShards The number of shards whose allocation has been delayed by the timeout
	// settings.
	DelayedUnassignedShards int                               `json:"delayed_unassigned_shards"`
	Indices                 map[string]types.IndexHealthStats `json:"indices,omitempty"`
	// InitializingShards The number of shards that are under initialization.
	InitializingShards int `json:"initializing_shards"`
	// NumberOfDataNodes The number of nodes that are dedicated data nodes.
	NumberOfDataNodes int `json:"number_of_data_nodes"`
	// NumberOfInFlightFetch The number of unfinished fetches.
	NumberOfInFlightFetch int `json:"number_of_in_flight_fetch"`
	// NumberOfNodes The number of nodes within the cluster.
	NumberOfNodes int `json:"number_of_nodes"`
	// NumberOfPendingTasks The number of cluster-level changes that have not yet been executed.
	NumberOfPendingTasks int `json:"number_of_pending_tasks"`
	// RelocatingShards The number of shards that are under relocation.
	RelocatingShards int                       `json:"relocating_shards"`
	Status           healthstatus.HealthStatus `json:"status"`
	// TaskMaxWaitingInQueue The time since the earliest initiated task is waiting for being performed.
	TaskMaxWaitingInQueue types.Duration `json:"task_max_waiting_in_queue,omitempty"`
	// TaskMaxWaitingInQueueMillis The time expressed in milliseconds since the earliest initiated task is
	// waiting for being performed.
	TaskMaxWaitingInQueueMillis int64 `json:"task_max_waiting_in_queue_millis"`
	// TimedOut If false the response returned within the period of time that is specified by
	// the timeout parameter (30s by default)
	TimedOut bool `json:"timed_out"`
	// UnassignedShards The number of shards that are not allocated.
	UnassignedShards int `json:"unassigned_shards"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Indices: make(map[string]types.IndexHealthStats, 0),
	}
	return r
}
