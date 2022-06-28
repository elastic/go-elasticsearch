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

// JobTimingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/ml/_types/Job.ts#L104-L113
type JobTimingStats struct {
	AverageBucketProcessingTimeMs                   *float64 `json:"average_bucket_processing_time_ms,omitempty"`
	BucketCount                                     int64    `json:"bucket_count"`
	ExponentialAverageBucketProcessingTimeMs        *float64 `json:"exponential_average_bucket_processing_time_ms,omitempty"`
	ExponentialAverageBucketProcessingTimePerHourMs float64  `json:"exponential_average_bucket_processing_time_per_hour_ms"`
	JobId                                           Id       `json:"job_id"`
	MaximumBucketProcessingTimeMs                   *float64 `json:"maximum_bucket_processing_time_ms,omitempty"`
	MinimumBucketProcessingTimeMs                   *float64 `json:"minimum_bucket_processing_time_ms,omitempty"`
	TotalBucketProcessingTimeMs                     float64  `json:"total_bucket_processing_time_ms"`
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

func (rb *JobTimingStatsBuilder) AverageBucketProcessingTimeMs(averagebucketprocessingtimems float64) *JobTimingStatsBuilder {
	rb.v.AverageBucketProcessingTimeMs = &averagebucketprocessingtimems
	return rb
}

func (rb *JobTimingStatsBuilder) BucketCount(bucketcount int64) *JobTimingStatsBuilder {
	rb.v.BucketCount = bucketcount
	return rb
}

func (rb *JobTimingStatsBuilder) ExponentialAverageBucketProcessingTimeMs(exponentialaveragebucketprocessingtimems float64) *JobTimingStatsBuilder {
	rb.v.ExponentialAverageBucketProcessingTimeMs = &exponentialaveragebucketprocessingtimems
	return rb
}

func (rb *JobTimingStatsBuilder) ExponentialAverageBucketProcessingTimePerHourMs(exponentialaveragebucketprocessingtimeperhourms float64) *JobTimingStatsBuilder {
	rb.v.ExponentialAverageBucketProcessingTimePerHourMs = exponentialaveragebucketprocessingtimeperhourms
	return rb
}

func (rb *JobTimingStatsBuilder) JobId(jobid Id) *JobTimingStatsBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *JobTimingStatsBuilder) MaximumBucketProcessingTimeMs(maximumbucketprocessingtimems float64) *JobTimingStatsBuilder {
	rb.v.MaximumBucketProcessingTimeMs = &maximumbucketprocessingtimems
	return rb
}

func (rb *JobTimingStatsBuilder) MinimumBucketProcessingTimeMs(minimumbucketprocessingtimems float64) *JobTimingStatsBuilder {
	rb.v.MinimumBucketProcessingTimeMs = &minimumbucketprocessingtimems
	return rb
}

func (rb *JobTimingStatsBuilder) TotalBucketProcessingTimeMs(totalbucketprocessingtimems float64) *JobTimingStatsBuilder {
	rb.v.TotalBucketProcessingTimeMs = totalbucketprocessingtimems
	return rb
}
