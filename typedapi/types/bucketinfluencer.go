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

// BucketInfluencer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Bucket.ts#L80-L128
type BucketInfluencer struct {
	// AnomalyScore A normalized score between 0-100, which is calculated for each bucket
	// influencer. This score might be updated as
	// newer data is analyzed.
	AnomalyScore float64 `json:"anomaly_score"`
	// BucketSpan The length of the bucket in seconds. This value matches the bucket span that
	// is specified in the job.
	BucketSpan DurationValueUnitSeconds `json:"bucket_span"`
	// InfluencerFieldName The field name of the influencer.
	InfluencerFieldName Field `json:"influencer_field_name"`
	// InitialAnomalyScore The score between 0-100 for each bucket influencer. This score is the initial
	// value that was calculated at the
	// time the bucket was processed.
	InitialAnomalyScore float64 `json:"initial_anomaly_score"`
	// IsInterim If true, this is an interim result. In other words, the results are
	// calculated based on partial input data.
	IsInterim bool `json:"is_interim"`
	// JobId Identifier for the anomaly detection job.
	JobId Id `json:"job_id"`
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
	Timestamp EpochTimeUnitMillis `json:"timestamp"`
	// TimestampString The start time of the bucket for which these results were calculated.
	TimestampString *DateTime `json:"timestamp_string,omitempty"`
}

// BucketInfluencerBuilder holds BucketInfluencer struct and provides a builder API.
type BucketInfluencerBuilder struct {
	v *BucketInfluencer
}

// NewBucketInfluencer provides a builder for the BucketInfluencer struct.
func NewBucketInfluencerBuilder() *BucketInfluencerBuilder {
	r := BucketInfluencerBuilder{
		&BucketInfluencer{},
	}

	return &r
}

// Build finalize the chain and returns the BucketInfluencer struct
func (rb *BucketInfluencerBuilder) Build() BucketInfluencer {
	return *rb.v
}

// AnomalyScore A normalized score between 0-100, which is calculated for each bucket
// influencer. This score might be updated as
// newer data is analyzed.

func (rb *BucketInfluencerBuilder) AnomalyScore(anomalyscore float64) *BucketInfluencerBuilder {
	rb.v.AnomalyScore = anomalyscore
	return rb
}

// BucketSpan The length of the bucket in seconds. This value matches the bucket span that
// is specified in the job.

func (rb *BucketInfluencerBuilder) BucketSpan(bucketspan *DurationValueUnitSecondsBuilder) *BucketInfluencerBuilder {
	v := bucketspan.Build()
	rb.v.BucketSpan = v
	return rb
}

// InfluencerFieldName The field name of the influencer.

func (rb *BucketInfluencerBuilder) InfluencerFieldName(influencerfieldname Field) *BucketInfluencerBuilder {
	rb.v.InfluencerFieldName = influencerfieldname
	return rb
}

// InitialAnomalyScore The score between 0-100 for each bucket influencer. This score is the initial
// value that was calculated at the
// time the bucket was processed.

func (rb *BucketInfluencerBuilder) InitialAnomalyScore(initialanomalyscore float64) *BucketInfluencerBuilder {
	rb.v.InitialAnomalyScore = initialanomalyscore
	return rb
}

// IsInterim If true, this is an interim result. In other words, the results are
// calculated based on partial input data.

func (rb *BucketInfluencerBuilder) IsInterim(isinterim bool) *BucketInfluencerBuilder {
	rb.v.IsInterim = isinterim
	return rb
}

// JobId Identifier for the anomaly detection job.

func (rb *BucketInfluencerBuilder) JobId(jobid Id) *BucketInfluencerBuilder {
	rb.v.JobId = jobid
	return rb
}

// Probability The probability that the bucket has this behavior, in the range 0 to 1. This
// value can be held to a high precision
// of over 300 decimal places, so the `anomaly_score` is provided as a
// human-readable and friendly interpretation of
// this.

func (rb *BucketInfluencerBuilder) Probability(probability float64) *BucketInfluencerBuilder {
	rb.v.Probability = probability
	return rb
}

// RawAnomalyScore Internal.

func (rb *BucketInfluencerBuilder) RawAnomalyScore(rawanomalyscore float64) *BucketInfluencerBuilder {
	rb.v.RawAnomalyScore = rawanomalyscore
	return rb
}

// ResultType Internal. This value is always set to `bucket_influencer`.

func (rb *BucketInfluencerBuilder) ResultType(resulttype string) *BucketInfluencerBuilder {
	rb.v.ResultType = resulttype
	return rb
}

// Timestamp The start time of the bucket for which these results were calculated.

func (rb *BucketInfluencerBuilder) Timestamp(timestamp *EpochTimeUnitMillisBuilder) *BucketInfluencerBuilder {
	v := timestamp.Build()
	rb.v.Timestamp = v
	return rb
}

// TimestampString The start time of the bucket for which these results were calculated.

func (rb *BucketInfluencerBuilder) TimestampString(timestampstring *DateTimeBuilder) *BucketInfluencerBuilder {
	v := timestampstring.Build()
	rb.v.TimestampString = &v
	return rb
}
