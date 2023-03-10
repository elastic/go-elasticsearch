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

// AnomalyExplanation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/Anomaly.ts#L156-L197
type AnomalyExplanation struct {
	// AnomalyCharacteristicsImpact Impact from the duration and magnitude of the detected anomaly relative to
	// the historical average.
	AnomalyCharacteristicsImpact *int `json:"anomaly_characteristics_impact,omitempty"`
	// AnomalyLength Length of the detected anomaly in the number of buckets.
	AnomalyLength *int `json:"anomaly_length,omitempty"`
	// AnomalyType Type of the detected anomaly: `spike` or `dip`.
	AnomalyType *string `json:"anomaly_type,omitempty"`
	// HighVariancePenalty Indicates reduction of anomaly score for the bucket with large confidence
	// intervals. If a bucket has large confidence intervals, the score is reduced.
	HighVariancePenalty *bool `json:"high_variance_penalty,omitempty"`
	// IncompleteBucketPenalty If the bucket contains fewer samples than expected, the score is reduced.
	IncompleteBucketPenalty *bool `json:"incomplete_bucket_penalty,omitempty"`
	// LowerConfidenceBound Lower bound of the 95% confidence interval.
	LowerConfidenceBound *Float64 `json:"lower_confidence_bound,omitempty"`
	// MultiBucketImpact Impact of the deviation between actual and typical values in the past 12
	// buckets.
	MultiBucketImpact *int `json:"multi_bucket_impact,omitempty"`
	// SingleBucketImpact Impact of the deviation between actual and typical values in the current
	// bucket.
	SingleBucketImpact *int `json:"single_bucket_impact,omitempty"`
	// TypicalValue Typical (expected) value for this bucket.
	TypicalValue *Float64 `json:"typical_value,omitempty"`
	// UpperConfidenceBound Upper bound of the 95% confidence interval.
	UpperConfidenceBound *Float64 `json:"upper_confidence_bound,omitempty"`
}

// NewAnomalyExplanation returns a AnomalyExplanation.
func NewAnomalyExplanation() *AnomalyExplanation {
	r := &AnomalyExplanation{}

	return r
}
