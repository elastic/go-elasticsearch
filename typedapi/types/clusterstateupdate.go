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

// ClusterStateUpdate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/nodes/_types/Stats.ts#L120-L136
type ClusterStateUpdate struct {
	CommitTime                    *Duration `json:"commit_time,omitempty"`
	CommitTimeMillis              *int64    `json:"commit_time_millis,omitempty"`
	CompletionTime                *Duration `json:"completion_time,omitempty"`
	CompletionTimeMillis          *int64    `json:"completion_time_millis,omitempty"`
	ComputationTime               *Duration `json:"computation_time,omitempty"`
	ComputationTimeMillis         *int64    `json:"computation_time_millis,omitempty"`
	ContextConstructionTime       *Duration `json:"context_construction_time,omitempty"`
	ContextConstructionTimeMillis *int64    `json:"context_construction_time_millis,omitempty"`
	Count                         int64     `json:"count"`
	MasterApplyTime               *Duration `json:"master_apply_time,omitempty"`
	MasterApplyTimeMillis         *int64    `json:"master_apply_time_millis,omitempty"`
	NotificationTime              *Duration `json:"notification_time,omitempty"`
	NotificationTimeMillis        *int64    `json:"notification_time_millis,omitempty"`
	PublicationTime               *Duration `json:"publication_time,omitempty"`
	PublicationTimeMillis         *int64    `json:"publication_time_millis,omitempty"`
}

// NewClusterStateUpdate returns a ClusterStateUpdate.
func NewClusterStateUpdate() *ClusterStateUpdate {
	r := &ClusterStateUpdate{}

	return r
}
