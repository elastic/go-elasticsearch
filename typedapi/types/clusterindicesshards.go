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

// ClusterIndicesShards type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L49-L61
type ClusterIndicesShards struct {
	// Index Contains statistics about shards assigned to selected nodes.
	Index *ClusterIndicesShardsIndex `json:"index,omitempty"`
	// Primaries Number of primary shards assigned to selected nodes.
	Primaries *float64 `json:"primaries,omitempty"`
	// Replication Ratio of replica shards to primary shards across all selected nodes.
	Replication *float64 `json:"replication,omitempty"`
	// Total Total number of shards assigned to selected nodes.
	Total *float64 `json:"total,omitempty"`
}

// ClusterIndicesShardsBuilder holds ClusterIndicesShards struct and provides a builder API.
type ClusterIndicesShardsBuilder struct {
	v *ClusterIndicesShards
}

// NewClusterIndicesShards provides a builder for the ClusterIndicesShards struct.
func NewClusterIndicesShardsBuilder() *ClusterIndicesShardsBuilder {
	r := ClusterIndicesShardsBuilder{
		&ClusterIndicesShards{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterIndicesShards struct
func (rb *ClusterIndicesShardsBuilder) Build() ClusterIndicesShards {
	return *rb.v
}

// Index Contains statistics about shards assigned to selected nodes.

func (rb *ClusterIndicesShardsBuilder) Index(index *ClusterIndicesShardsIndexBuilder) *ClusterIndicesShardsBuilder {
	v := index.Build()
	rb.v.Index = &v
	return rb
}

// Primaries Number of primary shards assigned to selected nodes.

func (rb *ClusterIndicesShardsBuilder) Primaries(primaries float64) *ClusterIndicesShardsBuilder {
	rb.v.Primaries = &primaries
	return rb
}

// Replication Ratio of replica shards to primary shards across all selected nodes.

func (rb *ClusterIndicesShardsBuilder) Replication(replication float64) *ClusterIndicesShardsBuilder {
	rb.v.Replication = &replication
	return rb
}

// Total Total number of shards assigned to selected nodes.

func (rb *ClusterIndicesShardsBuilder) Total(total float64) *ClusterIndicesShardsBuilder {
	rb.v.Total = &total
	return rb
}
