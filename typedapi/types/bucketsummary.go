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

// BucketSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Bucket.ts#L31-L78
type BucketSummary struct {
	// AnomalyScore The maximum anomaly score, between 0-100, for any of the bucket influencers.
	// This is an overall, rate-limited
	// score for the job. All the anomaly records in the bucket contribute to this
	// score. This value might be updated as
	// new data is analyzed.
	AnomalyScore      float64            `json:"anomaly_score"`
	BucketInfluencers []BucketInfluencer `json:"bucket_influencers"`
	// BucketSpan The length of the bucket in seconds. This value matches the bucket span that
	// is specified in the job.
	BucketSpan DurationValueUnitSeconds `json:"bucket_span"`
	// EventCount The number of input data records processed in this bucket.
	EventCount int64 `json:"event_count"`
	// InitialAnomalyScore The maximum anomaly score for any of the bucket influencers. This is the
	// initial value that was calculated at the
	// time the bucket was processed.
	InitialAnomalyScore float64 `json:"initial_anomaly_score"`
	// IsInterim If true, this is an interim result. In other words, the results are
	// calculated based on partial input data.
	IsInterim bool `json:"is_interim"`
	// JobId Identifier for the anomaly detection job.
	JobId Id `json:"job_id"`
	// ProcessingTimeMs The amount of time, in milliseconds, that it took to analyze the bucket
	// contents and calculate results.
	ProcessingTimeMs DurationValueUnitMillis `json:"processing_time_ms"`
	// ResultType Internal. This value is always set to bucket.
	ResultType string `json:"result_type"`
	// Timestamp The start time of the bucket. This timestamp uniquely identifies the bucket.
	// Events that occur exactly at the
	// timestamp of the bucket are included in the results for the bucket.
	Timestamp EpochTimeUnitMillis `json:"timestamp"`
	// TimestampString The start time of the bucket. This timestamp uniquely identifies the bucket.
	// Events that occur exactly at the
	// timestamp of the bucket are included in the results for the bucket.
	TimestampString *DateTime `json:"timestamp_string,omitempty"`
}

// BucketSummaryBuilder holds BucketSummary struct and provides a builder API.
type BucketSummaryBuilder struct {
	v *BucketSummary
}

// NewBucketSummary provides a builder for the BucketSummary struct.
func NewBucketSummaryBuilder() *BucketSummaryBuilder {
	r := BucketSummaryBuilder{
		&BucketSummary{},
	}

	return &r
}

// Build finalize the chain and returns the BucketSummary struct
func (rb *BucketSummaryBuilder) Build() BucketSummary {
	return *rb.v
}

// AnomalyScore The maximum anomaly score, between 0-100, for any of the bucket influencers.
// This is an overall, rate-limited
// score for the job. All the anomaly records in the bucket contribute to this
// score. This value might be updated as
// new data is analyzed.

func (rb *BucketSummaryBuilder) AnomalyScore(anomalyscore float64) *BucketSummaryBuilder {
	rb.v.AnomalyScore = anomalyscore
	return rb
}

func (rb *BucketSummaryBuilder) BucketInfluencers(bucket_influencers []BucketInfluencerBuilder) *BucketSummaryBuilder {
	tmp := make([]BucketInfluencer, len(bucket_influencers))
	for _, value := range bucket_influencers {
		tmp = append(tmp, value.Build())
	}
	rb.v.BucketInfluencers = tmp
	return rb
}

// BucketSpan The length of the bucket in seconds. This value matches the bucket span that
// is specified in the job.

func (rb *BucketSummaryBuilder) BucketSpan(bucketspan *DurationValueUnitSecondsBuilder) *BucketSummaryBuilder {
	v := bucketspan.Build()
	rb.v.BucketSpan = v
	return rb
}

// EventCount The number of input data records processed in this bucket.

func (rb *BucketSummaryBuilder) EventCount(eventcount int64) *BucketSummaryBuilder {
	rb.v.EventCount = eventcount
	return rb
}

// InitialAnomalyScore The maximum anomaly score for any of the bucket influencers. This is the
// initial value that was calculated at the
// time the bucket was processed.

func (rb *BucketSummaryBuilder) InitialAnomalyScore(initialanomalyscore float64) *BucketSummaryBuilder {
	rb.v.InitialAnomalyScore = initialanomalyscore
	return rb
}

// IsInterim If true, this is an interim result. In other words, the results are
// calculated based on partial input data.

func (rb *BucketSummaryBuilder) IsInterim(isinterim bool) *BucketSummaryBuilder {
	rb.v.IsInterim = isinterim
	return rb
}

// JobId Identifier for the anomaly detection job.

func (rb *BucketSummaryBuilder) JobId(jobid Id) *BucketSummaryBuilder {
	rb.v.JobId = jobid
	return rb
}

// ProcessingTimeMs The amount of time, in milliseconds, that it took to analyze the bucket
// contents and calculate results.

func (rb *BucketSummaryBuilder) ProcessingTimeMs(processingtimems *DurationValueUnitMillisBuilder) *BucketSummaryBuilder {
	v := processingtimems.Build()
	rb.v.ProcessingTimeMs = v
	return rb
}

// ResultType Internal. This value is always set to bucket.

func (rb *BucketSummaryBuilder) ResultType(resulttype string) *BucketSummaryBuilder {
	rb.v.ResultType = resulttype
	return rb
}

// Timestamp The start time of the bucket. This timestamp uniquely identifies the bucket.
// Events that occur exactly at the
// timestamp of the bucket are included in the results for the bucket.

func (rb *BucketSummaryBuilder) Timestamp(timestamp *EpochTimeUnitMillisBuilder) *BucketSummaryBuilder {
	v := timestamp.Build()
	rb.v.Timestamp = v
	return rb
}

// TimestampString The start time of the bucket. This timestamp uniquely identifies the bucket.
// Events that occur exactly at the
// timestamp of the bucket are included in the results for the bucket.

func (rb *BucketSummaryBuilder) TimestampString(timestampstring *DateTimeBuilder) *BucketSummaryBuilder {
	v := timestampstring.Build()
	rb.v.TimestampString = &v
	return rb
}
