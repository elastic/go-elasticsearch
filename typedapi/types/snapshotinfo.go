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

// SnapshotInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/snapshot/_types/SnapshotInfo.ts#L41-L65
type SnapshotInfo struct {
	DataStreams        []string                `json:"data_streams"`
	Duration           *Duration               `json:"duration,omitempty"`
	DurationInMillis   *int64                  `json:"duration_in_millis,omitempty"`
	EndTime            *DateTime               `json:"end_time,omitempty"`
	EndTimeInMillis    *int64                  `json:"end_time_in_millis,omitempty"`
	Failures           []SnapshotShardFailure  `json:"failures,omitempty"`
	FeatureStates      []InfoFeatureState      `json:"feature_states,omitempty"`
	IncludeGlobalState *bool                   `json:"include_global_state,omitempty"`
	IndexDetails       map[string]IndexDetails `json:"index_details,omitempty"`
	Indices            []string                `json:"indices,omitempty"`
	Metadata           map[string]interface{}  `json:"metadata,omitempty"`
	Reason             *string                 `json:"reason,omitempty"`
	Repository         *string                 `json:"repository,omitempty"`
	Shards             *ShardStatistics        `json:"shards,omitempty"`
	Snapshot           string                  `json:"snapshot"`
	StartTime          *DateTime               `json:"start_time,omitempty"`
	StartTimeInMillis  *int64                  `json:"start_time_in_millis,omitempty"`
	State              *string                 `json:"state,omitempty"`
	Uuid               string                  `json:"uuid"`
	Version            *string                 `json:"version,omitempty"`
	VersionId          *int64                  `json:"version_id,omitempty"`
}

// NewSnapshotInfo returns a SnapshotInfo.
func NewSnapshotInfo() *SnapshotInfo {
	r := &SnapshotInfo{
		IndexDetails: make(map[string]IndexDetails, 0),
	}

	return r
}
