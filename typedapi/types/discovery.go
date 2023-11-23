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

package types

// Discovery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/nodes/_types/Stats.ts#L201-L219
type Discovery struct {
	ClusterApplierStats *ClusterAppliedStats `json:"cluster_applier_stats,omitempty"`
	// ClusterStateQueue Contains statistics for the cluster state queue of the node.
	ClusterStateQueue *ClusterStateQueue `json:"cluster_state_queue,omitempty"`
	// ClusterStateUpdate Contains low-level statistics about how long various activities took during
	// cluster state updates while the node was the elected master.
	// Omitted if the node is not master-eligible.
	// Every field whose name ends in `_time` within this object is also represented
	// as a raw number of milliseconds in a field whose name ends in `_time_millis`.
	// The human-readable fields with a `_time` suffix are only returned if
	// requested with the `?human=true` query parameter.
	ClusterStateUpdate map[string]ClusterStateUpdate `json:"cluster_state_update,omitempty"`
	// PublishedClusterStates Contains statistics for the published cluster states of the node.
	PublishedClusterStates  *PublishedClusterStates `json:"published_cluster_states,omitempty"`
	SerializedClusterStates *SerializedClusterState `json:"serialized_cluster_states,omitempty"`
}

// NewDiscovery returns a Discovery.
func NewDiscovery() *Discovery {
	r := &Discovery{
		ClusterStateUpdate: make(map[string]ClusterStateUpdate, 0),
	}

	return r
}
