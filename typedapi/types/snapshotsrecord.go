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

// SnapshotsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/snapshots/types.ts#L24-L90
type SnapshotsRecord struct {
	// Duration duration
	Duration *Duration `json:"duration,omitempty"`
	// EndEpoch end time in seconds since 1970-01-01 00:00:00
	EndEpoch *StringifiedEpochTimeUnitSeconds `json:"end_epoch,omitempty"`
	// EndTime end time in HH:MM:SS
	EndTime *string `json:"end_time,omitempty"`
	// FailedShards number of failed shards
	FailedShards *string `json:"failed_shards,omitempty"`
	// Id unique snapshot
	Id *string `json:"id,omitempty"`
	// Indices number of indices
	Indices *string `json:"indices,omitempty"`
	// Reason reason for failures
	Reason *string `json:"reason,omitempty"`
	// Repository repository name
	Repository *string `json:"repository,omitempty"`
	// StartEpoch start time in seconds since 1970-01-01 00:00:00
	StartEpoch *StringifiedEpochTimeUnitSeconds `json:"start_epoch,omitempty"`
	// StartTime start time in HH:MM:SS
	StartTime *ScheduleTimeOfDay `json:"start_time,omitempty"`
	// Status snapshot name
	Status *string `json:"status,omitempty"`
	// SuccessfulShards number of successful shards
	SuccessfulShards *string `json:"successful_shards,omitempty"`
	// TotalShards number of total shards
	TotalShards *string `json:"total_shards,omitempty"`
}

// NewSnapshotsRecord returns a SnapshotsRecord.
func NewSnapshotsRecord() *SnapshotsRecord {
	r := &SnapshotsRecord{}

	return r
}
