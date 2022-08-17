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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/followerindexstatus"
)

// FollowerIndex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ccr/follow_info/types.ts#L22-L28
type FollowerIndex struct {
	FollowerIndex IndexName                               `json:"follower_index"`
	LeaderIndex   IndexName                               `json:"leader_index"`
	Parameters    *FollowerIndexParameters                `json:"parameters,omitempty"`
	RemoteCluster Name                                    `json:"remote_cluster"`
	Status        followerindexstatus.FollowerIndexStatus `json:"status"`
}

// FollowerIndexBuilder holds FollowerIndex struct and provides a builder API.
type FollowerIndexBuilder struct {
	v *FollowerIndex
}

// NewFollowerIndex provides a builder for the FollowerIndex struct.
func NewFollowerIndexBuilder() *FollowerIndexBuilder {
	r := FollowerIndexBuilder{
		&FollowerIndex{},
	}

	return &r
}

// Build finalize the chain and returns the FollowerIndex struct
func (rb *FollowerIndexBuilder) Build() FollowerIndex {
	return *rb.v
}

func (rb *FollowerIndexBuilder) FollowerIndex(followerindex IndexName) *FollowerIndexBuilder {
	rb.v.FollowerIndex = followerindex
	return rb
}

func (rb *FollowerIndexBuilder) LeaderIndex(leaderindex IndexName) *FollowerIndexBuilder {
	rb.v.LeaderIndex = leaderindex
	return rb
}

func (rb *FollowerIndexBuilder) Parameters(parameters *FollowerIndexParametersBuilder) *FollowerIndexBuilder {
	v := parameters.Build()
	rb.v.Parameters = &v
	return rb
}

func (rb *FollowerIndexBuilder) RemoteCluster(remotecluster Name) *FollowerIndexBuilder {
	rb.v.RemoteCluster = remotecluster
	return rb
}

func (rb *FollowerIndexBuilder) Status(status followerindexstatus.FollowerIndexStatus) *FollowerIndexBuilder {
	rb.v.Status = status
	return rb
}
