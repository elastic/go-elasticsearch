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

// BucketSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/Bucket.ts#L31-L78
type BucketSummary struct {
	// AnomalyScore The maximum anomaly score, between 0-100, for any of the bucket influencers.
	// This is an overall, rate-limited
	// score for the job. All the anomaly records in the bucket contribute to this
	// score. This value might be updated as
	// new data is analyzed.
	AnomalyScore      Float64            `json:"anomaly_score"`
	BucketInfluencers []BucketInfluencer `json:"bucket_influencers"`
	// BucketSpan The length of the bucket in seconds. This value matches the bucket span that
	// is specified in the job.
	BucketSpan int64 `json:"bucket_span"`
	// EventCount The number of input data records processed in this bucket.
	EventCount int64 `json:"event_count"`
	// InitialAnomalyScore The maximum anomaly score for any of the bucket influencers. This is the
	// initial value that was calculated at the
	// time the bucket was processed.
	InitialAnomalyScore Float64 `json:"initial_anomaly_score"`
	// IsInterim If true, this is an interim result. In other words, the results are
	// calculated based on partial input data.
	IsInterim bool `json:"is_interim"`
	// JobId Identifier for the anomaly detection job.
	JobId string `json:"job_id"`
	// ProcessingTimeMs The amount of time, in milliseconds, that it took to analyze the bucket
	// contents and calculate results.
	ProcessingTimeMs int64 `json:"processing_time_ms"`
	// ResultType Internal. This value is always set to bucket.
	ResultType string `json:"result_type"`
	// Timestamp The start time of the bucket. This timestamp uniquely identifies the bucket.
	// Events that occur exactly at the
	// timestamp of the bucket are included in the results for the bucket.
	Timestamp int64 `json:"timestamp"`
	// TimestampString The start time of the bucket. This timestamp uniquely identifies the bucket.
	// Events that occur exactly at the
	// timestamp of the bucket are included in the results for the bucket.
	TimestampString DateTime `json:"timestamp_string,omitempty"`
}

// NewBucketSummary returns a BucketSummary.
func NewBucketSummary() *BucketSummary {
	r := &BucketSummary{}

	return r
}
