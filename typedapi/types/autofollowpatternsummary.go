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

// AutoFollowPatternSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ccr/get_auto_follow_pattern/types.ts#L28-L51
type AutoFollowPatternSummary struct {
	Active bool `json:"active"`
	// FollowIndexPattern The name of follower index.
	FollowIndexPattern *IndexPattern `json:"follow_index_pattern,omitempty"`
	// LeaderIndexExclusionPatterns An array of simple index patterns that can be used to exclude indices from
	// being auto-followed.
	LeaderIndexExclusionPatterns IndexPatterns `json:"leader_index_exclusion_patterns"`
	// LeaderIndexPatterns An array of simple index patterns to match against indices in the remote
	// cluster specified by the remote_cluster field.
	LeaderIndexPatterns IndexPatterns `json:"leader_index_patterns"`
	// MaxOutstandingReadRequests The maximum number of outstanding reads requests from the remote cluster.
	MaxOutstandingReadRequests int `json:"max_outstanding_read_requests"`
	// RemoteCluster The remote cluster containing the leader indices to match against.
	RemoteCluster string `json:"remote_cluster"`
}

// AutoFollowPatternSummaryBuilder holds AutoFollowPatternSummary struct and provides a builder API.
type AutoFollowPatternSummaryBuilder struct {
	v *AutoFollowPatternSummary
}

// NewAutoFollowPatternSummary provides a builder for the AutoFollowPatternSummary struct.
func NewAutoFollowPatternSummaryBuilder() *AutoFollowPatternSummaryBuilder {
	r := AutoFollowPatternSummaryBuilder{
		&AutoFollowPatternSummary{},
	}

	return &r
}

// Build finalize the chain and returns the AutoFollowPatternSummary struct
func (rb *AutoFollowPatternSummaryBuilder) Build() AutoFollowPatternSummary {
	return *rb.v
}

func (rb *AutoFollowPatternSummaryBuilder) Active(active bool) *AutoFollowPatternSummaryBuilder {
	rb.v.Active = active
	return rb
}

// FollowIndexPattern The name of follower index.

func (rb *AutoFollowPatternSummaryBuilder) FollowIndexPattern(followindexpattern IndexPattern) *AutoFollowPatternSummaryBuilder {
	rb.v.FollowIndexPattern = &followindexpattern
	return rb
}

// LeaderIndexExclusionPatterns An array of simple index patterns that can be used to exclude indices from
// being auto-followed.

func (rb *AutoFollowPatternSummaryBuilder) LeaderIndexExclusionPatterns(leaderindexexclusionpatterns *IndexPatternsBuilder) *AutoFollowPatternSummaryBuilder {
	v := leaderindexexclusionpatterns.Build()
	rb.v.LeaderIndexExclusionPatterns = v
	return rb
}

// LeaderIndexPatterns An array of simple index patterns to match against indices in the remote
// cluster specified by the remote_cluster field.

func (rb *AutoFollowPatternSummaryBuilder) LeaderIndexPatterns(leaderindexpatterns *IndexPatternsBuilder) *AutoFollowPatternSummaryBuilder {
	v := leaderindexpatterns.Build()
	rb.v.LeaderIndexPatterns = v
	return rb
}

// MaxOutstandingReadRequests The maximum number of outstanding reads requests from the remote cluster.

func (rb *AutoFollowPatternSummaryBuilder) MaxOutstandingReadRequests(maxoutstandingreadrequests int) *AutoFollowPatternSummaryBuilder {
	rb.v.MaxOutstandingReadRequests = maxoutstandingreadrequests
	return rb
}

// RemoteCluster The remote cluster containing the leader indices to match against.

func (rb *AutoFollowPatternSummaryBuilder) RemoteCluster(remotecluster string) *AutoFollowPatternSummaryBuilder {
	rb.v.RemoteCluster = remotecluster
	return rb
}
