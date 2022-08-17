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

// JobTimingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Job.ts#L109-L118
type JobTimingStats struct {
	AverageBucketProcessingTimeMs                   *DurationValueUnitFloatMillis `json:"average_bucket_processing_time_ms,omitempty"`
	BucketCount                                     int64                         `json:"bucket_count"`
	ExponentialAverageBucketProcessingTimeMs        *DurationValueUnitFloatMillis `json:"exponential_average_bucket_processing_time_ms,omitempty"`
	ExponentialAverageBucketProcessingTimePerHourMs DurationValueUnitFloatMillis  `json:"exponential_average_bucket_processing_time_per_hour_ms"`
	JobId                                           Id                            `json:"job_id"`
	MaximumBucketProcessingTimeMs                   *DurationValueUnitFloatMillis `json:"maximum_bucket_processing_time_ms,omitempty"`
	MinimumBucketProcessingTimeMs                   *DurationValueUnitFloatMillis `json:"minimum_bucket_processing_time_ms,omitempty"`
	TotalBucketProcessingTimeMs                     DurationValueUnitFloatMillis  `json:"total_bucket_processing_time_ms"`
}

// JobTimingStatsBuilder holds JobTimingStats struct and provides a builder API.
type JobTimingStatsBuilder struct {
	v *JobTimingStats
}

// NewJobTimingStats provides a builder for the JobTimingStats struct.
func NewJobTimingStatsBuilder() *JobTimingStatsBuilder {
	r := JobTimingStatsBuilder{
		&JobTimingStats{},
	}

	return &r
}

// Build finalize the chain and returns the JobTimingStats struct
func (rb *JobTimingStatsBuilder) Build() JobTimingStats {
	return *rb.v
}

func (rb *JobTimingStatsBuilder) AverageBucketProcessingTimeMs(averagebucketprocessingtimems *DurationValueUnitFloatMillisBuilder) *JobTimingStatsBuilder {
	v := averagebucketprocessingtimems.Build()
	rb.v.AverageBucketProcessingTimeMs = &v
	return rb
}

func (rb *JobTimingStatsBuilder) BucketCount(bucketcount int64) *JobTimingStatsBuilder {
	rb.v.BucketCount = bucketcount
	return rb
}

func (rb *JobTimingStatsBuilder) ExponentialAverageBucketProcessingTimeMs(exponentialaveragebucketprocessingtimems *DurationValueUnitFloatMillisBuilder) *JobTimingStatsBuilder {
	v := exponentialaveragebucketprocessingtimems.Build()
	rb.v.ExponentialAverageBucketProcessingTimeMs = &v
	return rb
}

func (rb *JobTimingStatsBuilder) ExponentialAverageBucketProcessingTimePerHourMs(exponentialaveragebucketprocessingtimeperhourms *DurationValueUnitFloatMillisBuilder) *JobTimingStatsBuilder {
	v := exponentialaveragebucketprocessingtimeperhourms.Build()
	rb.v.ExponentialAverageBucketProcessingTimePerHourMs = v
	return rb
}

func (rb *JobTimingStatsBuilder) JobId(jobid Id) *JobTimingStatsBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *JobTimingStatsBuilder) MaximumBucketProcessingTimeMs(maximumbucketprocessingtimems *DurationValueUnitFloatMillisBuilder) *JobTimingStatsBuilder {
	v := maximumbucketprocessingtimems.Build()
	rb.v.MaximumBucketProcessingTimeMs = &v
	return rb
}

func (rb *JobTimingStatsBuilder) MinimumBucketProcessingTimeMs(minimumbucketprocessingtimems *DurationValueUnitFloatMillisBuilder) *JobTimingStatsBuilder {
	v := minimumbucketprocessingtimems.Build()
	rb.v.MinimumBucketProcessingTimeMs = &v
	return rb
}

func (rb *JobTimingStatsBuilder) TotalBucketProcessingTimeMs(totalbucketprocessingtimems *DurationValueUnitFloatMillisBuilder) *JobTimingStatsBuilder {
	v := totalbucketprocessingtimems.Build()
	rb.v.TotalBucketProcessingTimeMs = v
	return rb
}
