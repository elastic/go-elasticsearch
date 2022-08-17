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

// OverallBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Bucket.ts#L130-L145
type OverallBucket struct {
	// BucketSpan The length of the bucket in seconds. Matches the job with the longest
	// bucket_span value.
	BucketSpan DurationValueUnitSeconds `json:"bucket_span"`
	// IsInterim If true, this is an interim result. In other words, the results are
	// calculated based on partial input data.
	IsInterim bool `json:"is_interim"`
	// Jobs An array of objects that contain the max_anomaly_score per job_id.
	Jobs []OverallBucketJob `json:"jobs"`
	// OverallScore The top_n average of the maximum bucket anomaly_score per job.
	OverallScore float64 `json:"overall_score"`
	// ResultType Internal. This is always set to overall_bucket.
	ResultType string `json:"result_type"`
	// Timestamp The start time of the bucket for which these results were calculated.
	Timestamp EpochTimeUnitMillis `json:"timestamp"`
	// TimestampString The start time of the bucket for which these results were calculated.
	TimestampString DateTime `json:"timestamp_string"`
}

// OverallBucketBuilder holds OverallBucket struct and provides a builder API.
type OverallBucketBuilder struct {
	v *OverallBucket
}

// NewOverallBucket provides a builder for the OverallBucket struct.
func NewOverallBucketBuilder() *OverallBucketBuilder {
	r := OverallBucketBuilder{
		&OverallBucket{},
	}

	return &r
}

// Build finalize the chain and returns the OverallBucket struct
func (rb *OverallBucketBuilder) Build() OverallBucket {
	return *rb.v
}

// BucketSpan The length of the bucket in seconds. Matches the job with the longest
// bucket_span value.

func (rb *OverallBucketBuilder) BucketSpan(bucketspan *DurationValueUnitSecondsBuilder) *OverallBucketBuilder {
	v := bucketspan.Build()
	rb.v.BucketSpan = v
	return rb
}

// IsInterim If true, this is an interim result. In other words, the results are
// calculated based on partial input data.

func (rb *OverallBucketBuilder) IsInterim(isinterim bool) *OverallBucketBuilder {
	rb.v.IsInterim = isinterim
	return rb
}

// Jobs An array of objects that contain the max_anomaly_score per job_id.

func (rb *OverallBucketBuilder) Jobs(jobs []OverallBucketJobBuilder) *OverallBucketBuilder {
	tmp := make([]OverallBucketJob, len(jobs))
	for _, value := range jobs {
		tmp = append(tmp, value.Build())
	}
	rb.v.Jobs = tmp
	return rb
}

// OverallScore The top_n average of the maximum bucket anomaly_score per job.

func (rb *OverallBucketBuilder) OverallScore(overallscore float64) *OverallBucketBuilder {
	rb.v.OverallScore = overallscore
	return rb
}

// ResultType Internal. This is always set to overall_bucket.

func (rb *OverallBucketBuilder) ResultType(resulttype string) *OverallBucketBuilder {
	rb.v.ResultType = resulttype
	return rb
}

// Timestamp The start time of the bucket for which these results were calculated.

func (rb *OverallBucketBuilder) Timestamp(timestamp *EpochTimeUnitMillisBuilder) *OverallBucketBuilder {
	v := timestamp.Build()
	rb.v.Timestamp = v
	return rb
}

// TimestampString The start time of the bucket for which these results were calculated.

func (rb *OverallBucketBuilder) TimestampString(timestampstring *DateTimeBuilder) *OverallBucketBuilder {
	v := timestampstring.Build()
	rb.v.TimestampString = v
	return rb
}
