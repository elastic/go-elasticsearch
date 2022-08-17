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

// CoordinatorStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/enrich/stats/types.ts#L29-L35
type CoordinatorStats struct {
	ExecutedSearchesTotal int64 `json:"executed_searches_total"`
	NodeId                Id    `json:"node_id"`
	QueueSize             int   `json:"queue_size"`
	RemoteRequestsCurrent int   `json:"remote_requests_current"`
	RemoteRequestsTotal   int64 `json:"remote_requests_total"`
}

// CoordinatorStatsBuilder holds CoordinatorStats struct and provides a builder API.
type CoordinatorStatsBuilder struct {
	v *CoordinatorStats
}

// NewCoordinatorStats provides a builder for the CoordinatorStats struct.
func NewCoordinatorStatsBuilder() *CoordinatorStatsBuilder {
	r := CoordinatorStatsBuilder{
		&CoordinatorStats{},
	}

	return &r
}

// Build finalize the chain and returns the CoordinatorStats struct
func (rb *CoordinatorStatsBuilder) Build() CoordinatorStats {
	return *rb.v
}

func (rb *CoordinatorStatsBuilder) ExecutedSearchesTotal(executedsearchestotal int64) *CoordinatorStatsBuilder {
	rb.v.ExecutedSearchesTotal = executedsearchestotal
	return rb
}

func (rb *CoordinatorStatsBuilder) NodeId(nodeid Id) *CoordinatorStatsBuilder {
	rb.v.NodeId = nodeid
	return rb
}

func (rb *CoordinatorStatsBuilder) QueueSize(queuesize int) *CoordinatorStatsBuilder {
	rb.v.QueueSize = queuesize
	return rb
}

func (rb *CoordinatorStatsBuilder) RemoteRequestsCurrent(remoterequestscurrent int) *CoordinatorStatsBuilder {
	rb.v.RemoteRequestsCurrent = remoterequestscurrent
	return rb
}

func (rb *CoordinatorStatsBuilder) RemoteRequestsTotal(remoterequeststotal int64) *CoordinatorStatsBuilder {
	rb.v.RemoteRequestsTotal = remoterequeststotal
	return rb
}
