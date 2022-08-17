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

// Discovery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L76-L82
type Discovery struct {
	ClusterApplierStats     *ClusterAppliedStats          `json:"cluster_applier_stats,omitempty"`
	ClusterStateQueue       *ClusterStateQueue            `json:"cluster_state_queue,omitempty"`
	ClusterStateUpdate      map[string]ClusterStateUpdate `json:"cluster_state_update,omitempty"`
	PublishedClusterStates  *PublishedClusterStates       `json:"published_cluster_states,omitempty"`
	SerializedClusterStates *SerializedClusterState       `json:"serialized_cluster_states,omitempty"`
}

// DiscoveryBuilder holds Discovery struct and provides a builder API.
type DiscoveryBuilder struct {
	v *Discovery
}

// NewDiscovery provides a builder for the Discovery struct.
func NewDiscoveryBuilder() *DiscoveryBuilder {
	r := DiscoveryBuilder{
		&Discovery{
			ClusterStateUpdate: make(map[string]ClusterStateUpdate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Discovery struct
func (rb *DiscoveryBuilder) Build() Discovery {
	return *rb.v
}

func (rb *DiscoveryBuilder) ClusterApplierStats(clusterapplierstats *ClusterAppliedStatsBuilder) *DiscoveryBuilder {
	v := clusterapplierstats.Build()
	rb.v.ClusterApplierStats = &v
	return rb
}

func (rb *DiscoveryBuilder) ClusterStateQueue(clusterstatequeue *ClusterStateQueueBuilder) *DiscoveryBuilder {
	v := clusterstatequeue.Build()
	rb.v.ClusterStateQueue = &v
	return rb
}

func (rb *DiscoveryBuilder) ClusterStateUpdate(values map[string]*ClusterStateUpdateBuilder) *DiscoveryBuilder {
	tmp := make(map[string]ClusterStateUpdate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.ClusterStateUpdate = tmp
	return rb
}

func (rb *DiscoveryBuilder) PublishedClusterStates(publishedclusterstates *PublishedClusterStatesBuilder) *DiscoveryBuilder {
	v := publishedclusterstates.Build()
	rb.v.PublishedClusterStates = &v
	return rb
}

func (rb *DiscoveryBuilder) SerializedClusterStates(serializedclusterstates *SerializedClusterStateBuilder) *DiscoveryBuilder {
	v := serializedclusterstates.Build()
	rb.v.SerializedClusterStates = &v
	return rb
}
