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

// CheckpointStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/get_transform_stats/types.ts#L68-L75
type CheckpointStats struct {
	Checkpoint           int64                `json:"checkpoint"`
	CheckpointProgress   *TransformProgress   `json:"checkpoint_progress,omitempty"`
	TimeUpperBound       *DateTime            `json:"time_upper_bound,omitempty"`
	TimeUpperBoundMillis *EpochTimeUnitMillis `json:"time_upper_bound_millis,omitempty"`
	Timestamp            *DateTime            `json:"timestamp,omitempty"`
	TimestampMillis      *EpochTimeUnitMillis `json:"timestamp_millis,omitempty"`
}

// CheckpointStatsBuilder holds CheckpointStats struct and provides a builder API.
type CheckpointStatsBuilder struct {
	v *CheckpointStats
}

// NewCheckpointStats provides a builder for the CheckpointStats struct.
func NewCheckpointStatsBuilder() *CheckpointStatsBuilder {
	r := CheckpointStatsBuilder{
		&CheckpointStats{},
	}

	return &r
}

// Build finalize the chain and returns the CheckpointStats struct
func (rb *CheckpointStatsBuilder) Build() CheckpointStats {
	return *rb.v
}

func (rb *CheckpointStatsBuilder) Checkpoint(checkpoint int64) *CheckpointStatsBuilder {
	rb.v.Checkpoint = checkpoint
	return rb
}

func (rb *CheckpointStatsBuilder) CheckpointProgress(checkpointprogress *TransformProgressBuilder) *CheckpointStatsBuilder {
	v := checkpointprogress.Build()
	rb.v.CheckpointProgress = &v
	return rb
}

func (rb *CheckpointStatsBuilder) TimeUpperBound(timeupperbound *DateTimeBuilder) *CheckpointStatsBuilder {
	v := timeupperbound.Build()
	rb.v.TimeUpperBound = &v
	return rb
}

func (rb *CheckpointStatsBuilder) TimeUpperBoundMillis(timeupperboundmillis *EpochTimeUnitMillisBuilder) *CheckpointStatsBuilder {
	v := timeupperboundmillis.Build()
	rb.v.TimeUpperBoundMillis = &v
	return rb
}

func (rb *CheckpointStatsBuilder) Timestamp(timestamp *DateTimeBuilder) *CheckpointStatsBuilder {
	v := timestamp.Build()
	rb.v.Timestamp = &v
	return rb
}

func (rb *CheckpointStatsBuilder) TimestampMillis(timestampmillis *EpochTimeUnitMillisBuilder) *CheckpointStatsBuilder {
	v := timestampmillis.Build()
	rb.v.TimestampMillis = &v
	return rb
}
