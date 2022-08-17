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

// ClusterInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/allocation_explain/types.ts#L48-L54
type ClusterInfo struct {
	Nodes             map[string]NodeDiskUsage `json:"nodes"`
	ReservedSizes     []ReservedSize           `json:"reserved_sizes"`
	ShardDataSetSizes map[string]string        `json:"shard_data_set_sizes,omitempty"`
	ShardPaths        map[string]string        `json:"shard_paths"`
	ShardSizes        map[string]int64         `json:"shard_sizes"`
}

// ClusterInfoBuilder holds ClusterInfo struct and provides a builder API.
type ClusterInfoBuilder struct {
	v *ClusterInfo
}

// NewClusterInfo provides a builder for the ClusterInfo struct.
func NewClusterInfoBuilder() *ClusterInfoBuilder {
	r := ClusterInfoBuilder{
		&ClusterInfo{
			Nodes:             make(map[string]NodeDiskUsage, 0),
			ShardDataSetSizes: make(map[string]string, 0),
			ShardPaths:        make(map[string]string, 0),
			ShardSizes:        make(map[string]int64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ClusterInfo struct
func (rb *ClusterInfoBuilder) Build() ClusterInfo {
	return *rb.v
}

func (rb *ClusterInfoBuilder) Nodes(values map[string]*NodeDiskUsageBuilder) *ClusterInfoBuilder {
	tmp := make(map[string]NodeDiskUsage, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Nodes = tmp
	return rb
}

func (rb *ClusterInfoBuilder) ReservedSizes(reserved_sizes []ReservedSizeBuilder) *ClusterInfoBuilder {
	tmp := make([]ReservedSize, len(reserved_sizes))
	for _, value := range reserved_sizes {
		tmp = append(tmp, value.Build())
	}
	rb.v.ReservedSizes = tmp
	return rb
}

func (rb *ClusterInfoBuilder) ShardDataSetSizes(value map[string]string) *ClusterInfoBuilder {
	rb.v.ShardDataSetSizes = value
	return rb
}

func (rb *ClusterInfoBuilder) ShardPaths(value map[string]string) *ClusterInfoBuilder {
	rb.v.ShardPaths = value
	return rb
}

func (rb *ClusterInfoBuilder) ShardSizes(value map[string]int64) *ClusterInfoBuilder {
	rb.v.ShardSizes = value
	return rb
}
