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

// FollowIndexStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ccr/_types/FollowIndexStats.ts#L30-L33
type FollowIndexStats struct {
	Index  IndexName    `json:"index"`
	Shards []ShardStats `json:"shards"`
}

// FollowIndexStatsBuilder holds FollowIndexStats struct and provides a builder API.
type FollowIndexStatsBuilder struct {
	v *FollowIndexStats
}

// NewFollowIndexStats provides a builder for the FollowIndexStats struct.
func NewFollowIndexStatsBuilder() *FollowIndexStatsBuilder {
	r := FollowIndexStatsBuilder{
		&FollowIndexStats{},
	}

	return &r
}

// Build finalize the chain and returns the FollowIndexStats struct
func (rb *FollowIndexStatsBuilder) Build() FollowIndexStats {
	return *rb.v
}

func (rb *FollowIndexStatsBuilder) Index(index IndexName) *FollowIndexStatsBuilder {
	rb.v.Index = index
	return rb
}

func (rb *FollowIndexStatsBuilder) Shards(shards []ShardStatsBuilder) *FollowIndexStatsBuilder {
	tmp := make([]ShardStats, len(shards))
	for _, value := range shards {
		tmp = append(tmp, value.Build())
	}
	rb.v.Shards = tmp
	return rb
}
