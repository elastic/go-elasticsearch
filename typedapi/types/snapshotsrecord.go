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

// SnapshotsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/snapshots/types.ts#L24-L90
type SnapshotsRecord struct {
	// Duration duration
	Duration *Duration `json:"duration,omitempty"`
	// EndEpoch end time in seconds since 1970-01-01 00:00:00
	EndEpoch *StringifiedEpochTimeUnitSeconds `json:"end_epoch,omitempty"`
	// EndTime end time in HH:MM:SS
	EndTime *TimeOfDay `json:"end_time,omitempty"`
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

// SnapshotsRecordBuilder holds SnapshotsRecord struct and provides a builder API.
type SnapshotsRecordBuilder struct {
	v *SnapshotsRecord
}

// NewSnapshotsRecord provides a builder for the SnapshotsRecord struct.
func NewSnapshotsRecordBuilder() *SnapshotsRecordBuilder {
	r := SnapshotsRecordBuilder{
		&SnapshotsRecord{},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotsRecord struct
func (rb *SnapshotsRecordBuilder) Build() SnapshotsRecord {
	return *rb.v
}

// Duration duration

func (rb *SnapshotsRecordBuilder) Duration(duration *DurationBuilder) *SnapshotsRecordBuilder {
	v := duration.Build()
	rb.v.Duration = &v
	return rb
}

// EndEpoch end time in seconds since 1970-01-01 00:00:00

func (rb *SnapshotsRecordBuilder) EndEpoch(endepoch *StringifiedEpochTimeUnitSecondsBuilder) *SnapshotsRecordBuilder {
	v := endepoch.Build()
	rb.v.EndEpoch = &v
	return rb
}

// EndTime end time in HH:MM:SS

func (rb *SnapshotsRecordBuilder) EndTime(endtime TimeOfDay) *SnapshotsRecordBuilder {
	rb.v.EndTime = &endtime
	return rb
}

// FailedShards number of failed shards

func (rb *SnapshotsRecordBuilder) FailedShards(failedshards string) *SnapshotsRecordBuilder {
	rb.v.FailedShards = &failedshards
	return rb
}

// Id unique snapshot

func (rb *SnapshotsRecordBuilder) Id(id string) *SnapshotsRecordBuilder {
	rb.v.Id = &id
	return rb
}

// Indices number of indices

func (rb *SnapshotsRecordBuilder) Indices(indices string) *SnapshotsRecordBuilder {
	rb.v.Indices = &indices
	return rb
}

// Reason reason for failures

func (rb *SnapshotsRecordBuilder) Reason(reason string) *SnapshotsRecordBuilder {
	rb.v.Reason = &reason
	return rb
}

// Repository repository name

func (rb *SnapshotsRecordBuilder) Repository(repository string) *SnapshotsRecordBuilder {
	rb.v.Repository = &repository
	return rb
}

// StartEpoch start time in seconds since 1970-01-01 00:00:00

func (rb *SnapshotsRecordBuilder) StartEpoch(startepoch *StringifiedEpochTimeUnitSecondsBuilder) *SnapshotsRecordBuilder {
	v := startepoch.Build()
	rb.v.StartEpoch = &v
	return rb
}

// StartTime start time in HH:MM:SS

func (rb *SnapshotsRecordBuilder) StartTime(starttime *ScheduleTimeOfDayBuilder) *SnapshotsRecordBuilder {
	v := starttime.Build()
	rb.v.StartTime = &v
	return rb
}

// Status snapshot name

func (rb *SnapshotsRecordBuilder) Status(status string) *SnapshotsRecordBuilder {
	rb.v.Status = &status
	return rb
}

// SuccessfulShards number of successful shards

func (rb *SnapshotsRecordBuilder) SuccessfulShards(successfulshards string) *SnapshotsRecordBuilder {
	rb.v.SuccessfulShards = &successfulshards
	return rb
}

// TotalShards number of total shards

func (rb *SnapshotsRecordBuilder) TotalShards(totalshards string) *SnapshotsRecordBuilder {
	rb.v.TotalShards = &totalshards
	return rb
}
