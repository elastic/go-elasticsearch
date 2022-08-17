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

// OverallBucketJob type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Bucket.ts#L146-L149
type OverallBucketJob struct {
	JobId           Id      `json:"job_id"`
	MaxAnomalyScore float64 `json:"max_anomaly_score"`
}

// OverallBucketJobBuilder holds OverallBucketJob struct and provides a builder API.
type OverallBucketJobBuilder struct {
	v *OverallBucketJob
}

// NewOverallBucketJob provides a builder for the OverallBucketJob struct.
func NewOverallBucketJobBuilder() *OverallBucketJobBuilder {
	r := OverallBucketJobBuilder{
		&OverallBucketJob{},
	}

	return &r
}

// Build finalize the chain and returns the OverallBucketJob struct
func (rb *OverallBucketJobBuilder) Build() OverallBucketJob {
	return *rb.v
}

func (rb *OverallBucketJobBuilder) JobId(jobid Id) *OverallBucketJobBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *OverallBucketJobBuilder) MaxAnomalyScore(maxanomalyscore float64) *OverallBucketJobBuilder {
	rb.v.MaxAnomalyScore = maxanomalyscore
	return rb
}
