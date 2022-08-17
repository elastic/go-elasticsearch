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

// FollowStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ccr/stats/types.ts.ts#L41-L43
type FollowStats struct {
	Indices []FollowIndexStats `json:"indices"`
}

// FollowStatsBuilder holds FollowStats struct and provides a builder API.
type FollowStatsBuilder struct {
	v *FollowStats
}

// NewFollowStats provides a builder for the FollowStats struct.
func NewFollowStatsBuilder() *FollowStatsBuilder {
	r := FollowStatsBuilder{
		&FollowStats{},
	}

	return &r
}

// Build finalize the chain and returns the FollowStats struct
func (rb *FollowStatsBuilder) Build() FollowStats {
	return *rb.v
}

func (rb *FollowStatsBuilder) Indices(indices []FollowIndexStatsBuilder) *FollowStatsBuilder {
	tmp := make([]FollowIndexStats, len(indices))
	for _, value := range indices {
		tmp = append(tmp, value.Build())
	}
	rb.v.Indices = tmp
	return rb
}
