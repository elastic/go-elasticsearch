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

// AutoFollowStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ccr/stats/types.ts.ts#L33-L39
type AutoFollowStats struct {
	AutoFollowedClusters                     []AutoFollowedCluster `json:"auto_followed_clusters"`
	NumberOfFailedFollowIndices              int64                 `json:"number_of_failed_follow_indices"`
	NumberOfFailedRemoteClusterStateRequests int64                 `json:"number_of_failed_remote_cluster_state_requests"`
	NumberOfSuccessfulFollowIndices          int64                 `json:"number_of_successful_follow_indices"`
	RecentAutoFollowErrors                   []ErrorCause          `json:"recent_auto_follow_errors"`
}

// AutoFollowStatsBuilder holds AutoFollowStats struct and provides a builder API.
type AutoFollowStatsBuilder struct {
	v *AutoFollowStats
}

// NewAutoFollowStats provides a builder for the AutoFollowStats struct.
func NewAutoFollowStatsBuilder() *AutoFollowStatsBuilder {
	r := AutoFollowStatsBuilder{
		&AutoFollowStats{},
	}

	return &r
}

// Build finalize the chain and returns the AutoFollowStats struct
func (rb *AutoFollowStatsBuilder) Build() AutoFollowStats {
	return *rb.v
}

func (rb *AutoFollowStatsBuilder) AutoFollowedClusters(auto_followed_clusters []AutoFollowedClusterBuilder) *AutoFollowStatsBuilder {
	tmp := make([]AutoFollowedCluster, len(auto_followed_clusters))
	for _, value := range auto_followed_clusters {
		tmp = append(tmp, value.Build())
	}
	rb.v.AutoFollowedClusters = tmp
	return rb
}

func (rb *AutoFollowStatsBuilder) NumberOfFailedFollowIndices(numberoffailedfollowindices int64) *AutoFollowStatsBuilder {
	rb.v.NumberOfFailedFollowIndices = numberoffailedfollowindices
	return rb
}

func (rb *AutoFollowStatsBuilder) NumberOfFailedRemoteClusterStateRequests(numberoffailedremoteclusterstaterequests int64) *AutoFollowStatsBuilder {
	rb.v.NumberOfFailedRemoteClusterStateRequests = numberoffailedremoteclusterstaterequests
	return rb
}

func (rb *AutoFollowStatsBuilder) NumberOfSuccessfulFollowIndices(numberofsuccessfulfollowindices int64) *AutoFollowStatsBuilder {
	rb.v.NumberOfSuccessfulFollowIndices = numberofsuccessfulfollowindices
	return rb
}

func (rb *AutoFollowStatsBuilder) RecentAutoFollowErrors(recent_auto_follow_errors []ErrorCauseBuilder) *AutoFollowStatsBuilder {
	tmp := make([]ErrorCause, len(recent_auto_follow_errors))
	for _, value := range recent_auto_follow_errors {
		tmp = append(tmp, value.Build())
	}
	rb.v.RecentAutoFollowErrors = tmp
	return rb
}
