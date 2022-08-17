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

// SnapshotIndexStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/_types/SnapshotIndexStats.ts#L25-L29
type SnapshotIndexStats struct {
	Shards      map[string]SnapshotShardsStatus `json:"shards"`
	ShardsStats ShardsStats                     `json:"shards_stats"`
	Stats       SnapshotStats                   `json:"stats"`
}

// SnapshotIndexStatsBuilder holds SnapshotIndexStats struct and provides a builder API.
type SnapshotIndexStatsBuilder struct {
	v *SnapshotIndexStats
}

// NewSnapshotIndexStats provides a builder for the SnapshotIndexStats struct.
func NewSnapshotIndexStatsBuilder() *SnapshotIndexStatsBuilder {
	r := SnapshotIndexStatsBuilder{
		&SnapshotIndexStats{
			Shards: make(map[string]SnapshotShardsStatus, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotIndexStats struct
func (rb *SnapshotIndexStatsBuilder) Build() SnapshotIndexStats {
	return *rb.v
}

func (rb *SnapshotIndexStatsBuilder) Shards(values map[string]*SnapshotShardsStatusBuilder) *SnapshotIndexStatsBuilder {
	tmp := make(map[string]SnapshotShardsStatus, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Shards = tmp
	return rb
}

func (rb *SnapshotIndexStatsBuilder) ShardsStats(shardsstats *ShardsStatsBuilder) *SnapshotIndexStatsBuilder {
	v := shardsstats.Build()
	rb.v.ShardsStats = v
	return rb
}

func (rb *SnapshotIndexStatsBuilder) Stats(stats *SnapshotStatsBuilder) *SnapshotIndexStatsBuilder {
	v := stats.Build()
	rb.v.Stats = v
	return rb
}
