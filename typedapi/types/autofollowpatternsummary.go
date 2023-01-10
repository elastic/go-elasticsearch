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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// AutoFollowPatternSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ccr/get_auto_follow_pattern/types.ts#L28-L51
type AutoFollowPatternSummary struct {
	Active bool `json:"active"`
	// FollowIndexPattern The name of follower index.
	FollowIndexPattern *string `json:"follow_index_pattern,omitempty"`
	// LeaderIndexExclusionPatterns An array of simple index patterns that can be used to exclude indices from
	// being auto-followed.
	LeaderIndexExclusionPatterns []string `json:"leader_index_exclusion_patterns"`
	// LeaderIndexPatterns An array of simple index patterns to match against indices in the remote
	// cluster specified by the remote_cluster field.
	LeaderIndexPatterns []string `json:"leader_index_patterns"`
	// MaxOutstandingReadRequests The maximum number of outstanding reads requests from the remote cluster.
	MaxOutstandingReadRequests int `json:"max_outstanding_read_requests"`
	// RemoteCluster The remote cluster containing the leader indices to match against.
	RemoteCluster string `json:"remote_cluster"`
}

// NewAutoFollowPatternSummary returns a AutoFollowPatternSummary.
func NewAutoFollowPatternSummary() *AutoFollowPatternSummary {
	r := &AutoFollowPatternSummary{}

	return r
}
