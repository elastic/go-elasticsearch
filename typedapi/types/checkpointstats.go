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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// CheckpointStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/transform/get_transform_stats/types.ts#L62-L69
type CheckpointStats struct {
	Checkpoint           int64              `json:"checkpoint"`
	CheckpointProgress   *TransformProgress `json:"checkpoint_progress,omitempty"`
	TimeUpperBound       *DateString        `json:"time_upper_bound,omitempty"`
	TimeUpperBoundMillis *EpochMillis       `json:"time_upper_bound_millis,omitempty"`
	Timestamp            *DateString        `json:"timestamp,omitempty"`
	TimestampMillis      *EpochMillis       `json:"timestamp_millis,omitempty"`
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

func (rb *CheckpointStatsBuilder) TimeUpperBound(timeupperbound DateString) *CheckpointStatsBuilder {
	rb.v.TimeUpperBound = &timeupperbound
	return rb
}

func (rb *CheckpointStatsBuilder) TimeUpperBoundMillis(timeupperboundmillis *EpochMillisBuilder) *CheckpointStatsBuilder {
	v := timeupperboundmillis.Build()
	rb.v.TimeUpperBoundMillis = &v
	return rb
}

func (rb *CheckpointStatsBuilder) Timestamp(timestamp DateString) *CheckpointStatsBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}

func (rb *CheckpointStatsBuilder) TimestampMillis(timestampmillis *EpochMillisBuilder) *CheckpointStatsBuilder {
	v := timestampmillis.Build()
	rb.v.TimestampMillis = &v
	return rb
}
