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

// BucketInfluencer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/Bucket.ts#L80-L128
type BucketInfluencer struct {
	// AnomalyScore A normalized score between 0-100, which is calculated for each bucket
	// influencer. This score might be updated as
	// newer data is analyzed.
	AnomalyScore float64 `json:"anomaly_score"`
	// BucketSpan The length of the bucket in seconds. This value matches the bucket span that
	// is specified in the job.
	BucketSpan int64 `json:"bucket_span"`
	// InfluencerFieldName The field name of the influencer.
	InfluencerFieldName string `json:"influencer_field_name"`
	// InitialAnomalyScore The score between 0-100 for each bucket influencer. This score is the initial
	// value that was calculated at the
	// time the bucket was processed.
	InitialAnomalyScore float64 `json:"initial_anomaly_score"`
	// IsInterim If true, this is an interim result. In other words, the results are
	// calculated based on partial input data.
	IsInterim bool `json:"is_interim"`
	// JobId Identifier for the anomaly detection job.
	JobId string `json:"job_id"`
	// Probability The probability that the bucket has this behavior, in the range 0 to 1. This
	// value can be held to a high precision
	// of over 300 decimal places, so the `anomaly_score` is provided as a
	// human-readable and friendly interpretation of
	// this.
	Probability float64 `json:"probability"`
	// RawAnomalyScore Internal.
	RawAnomalyScore float64 `json:"raw_anomaly_score"`
	// ResultType Internal. This value is always set to `bucket_influencer`.
	ResultType string `json:"result_type"`
	// Timestamp The start time of the bucket for which these results were calculated.
	Timestamp int64 `json:"timestamp"`
	// TimestampString The start time of the bucket for which these results were calculated.
	TimestampString *DateTime `json:"timestamp_string,omitempty"`
}

// NewBucketInfluencer returns a BucketInfluencer.
func NewBucketInfluencer() *BucketInfluencer {
	r := &BucketInfluencer{}

	return r
}
