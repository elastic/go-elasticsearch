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

// ShardsTotalStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/stats/types.ts#L170-L172
type ShardsTotalStats struct {
	TotalCount int64 `json:"total_count"`
}

// ShardsTotalStatsBuilder holds ShardsTotalStats struct and provides a builder API.
type ShardsTotalStatsBuilder struct {
	v *ShardsTotalStats
}

// NewShardsTotalStats provides a builder for the ShardsTotalStats struct.
func NewShardsTotalStatsBuilder() *ShardsTotalStatsBuilder {
	r := ShardsTotalStatsBuilder{
		&ShardsTotalStats{},
	}

	return &r
}

// Build finalize the chain and returns the ShardsTotalStats struct
func (rb *ShardsTotalStatsBuilder) Build() ShardsTotalStats {
	return *rb.v
}

func (rb *ShardsTotalStatsBuilder) TotalCount(totalcount int64) *ShardsTotalStatsBuilder {
	rb.v.TotalCount = totalcount
	return rb
}
