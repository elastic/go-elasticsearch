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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// JobTimingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/Job.ts#L109-L118
type JobTimingStats struct {
	AverageBucketProcessingTimeMs                   Float64 `json:"average_bucket_processing_time_ms,omitempty"`
	BucketCount                                     int64   `json:"bucket_count"`
	ExponentialAverageBucketProcessingTimeMs        Float64 `json:"exponential_average_bucket_processing_time_ms,omitempty"`
	ExponentialAverageBucketProcessingTimePerHourMs Float64 `json:"exponential_average_bucket_processing_time_per_hour_ms"`
	JobId                                           string  `json:"job_id"`
	MaximumBucketProcessingTimeMs                   Float64 `json:"maximum_bucket_processing_time_ms,omitempty"`
	MinimumBucketProcessingTimeMs                   Float64 `json:"minimum_bucket_processing_time_ms,omitempty"`
	TotalBucketProcessingTimeMs                     Float64 `json:"total_bucket_processing_time_ms"`
}

// NewJobTimingStats returns a JobTimingStats.
func NewJobTimingStats() *JobTimingStats {
	r := &JobTimingStats{}

	return r
}
