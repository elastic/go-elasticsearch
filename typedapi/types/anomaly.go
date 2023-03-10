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

// Anomaly type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/Anomaly.ts#L24-L121
type Anomaly struct {
	// Actual The actual value for the bucket.
	Actual []Float64 `json:"actual,omitempty"`
	// AnomalyScoreExplanation Information about the factors impacting the initial anomaly score.
	AnomalyScoreExplanation *AnomalyExplanation `json:"anomaly_score_explanation,omitempty"`
	// BucketSpan The length of the bucket in seconds. This value matches the `bucket_span`
	// that is specified in the job.
	BucketSpan int64 `json:"bucket_span"`
	// ByFieldName The field used to split the data. In particular, this property is used for
	// analyzing the splits with respect to their own history. It is used for
	// finding unusual values in the context of the split.
	ByFieldName *string `json:"by_field_name,omitempty"`
	// ByFieldValue The value of `by_field_name`.
	ByFieldValue *string `json:"by_field_value,omitempty"`
	// Causes For population analysis, an over field must be specified in the detector.
	// This property contains an array of anomaly records that are the causes for
	// the anomaly that has been identified for the over field. This sub-resource
	// contains the most anomalous records for the `over_field_name`. For
	// scalability reasons, a maximum of the 10 most significant causes of the
	// anomaly are returned. As part of the core analytical modeling, these
	// low-level anomaly records are aggregated for their parent over field record.
	// The `causes` resource contains similar elements to the record resource,
	// namely `actual`, `typical`, `geo_results.actual_point`,
	// `geo_results.typical_point`, `*_field_name` and `*_field_value`. Probability
	// and scores are not applicable to causes.
	Causes []AnomalyCause `json:"causes,omitempty"`
	// DetectorIndex A unique identifier for the detector.
	DetectorIndex int `json:"detector_index"`
	// FieldName Certain functions require a field to operate on, for example, `sum()`. For
	// those functions, this value is the name of the field to be analyzed.
	FieldName *string `json:"field_name,omitempty"`
	// Function The function in which the anomaly occurs, as specified in the detector
	// configuration. For example, `max`.
	Function *string `json:"function,omitempty"`
	// FunctionDescription The description of the function in which the anomaly occurs, as specified in
	// the detector configuration.
	FunctionDescription *string `json:"function_description,omitempty"`
	// GeoResults If the detector function is `lat_long`, this object contains comma delimited
	// strings for the latitude and longitude of the actual and typical values.
	GeoResults *GeoResults `json:"geo_results,omitempty"`
	// Influencers If influencers were specified in the detector configuration, this array
	// contains influencers that contributed to or were to blame for an anomaly.
	Influencers []Influence `json:"influencers,omitempty"`
	// InitialRecordScore A normalized score between 0-100, which is based on the probability of the
	// anomalousness of this record. This is the initial value that was calculated
	// at the time the bucket was processed.
	InitialRecordScore Float64 `json:"initial_record_score"`
	// IsInterim If true, this is an interim result. In other words, the results are
	// calculated based on partial input data.
	IsInterim bool `json:"is_interim"`
	// JobId Identifier for the anomaly detection job.
	JobId string `json:"job_id"`
	// OverFieldName The field used to split the data. In particular, this property is used for
	// analyzing the splits with respect to the history of all splits. It is used
	// for finding unusual values in the population of all splits.
	OverFieldName *string `json:"over_field_name,omitempty"`
	// OverFieldValue The value of `over_field_name`.
	OverFieldValue *string `json:"over_field_value,omitempty"`
	// PartitionFieldName The field used to segment the analysis. When you use this property, you have
	// completely independent baselines for each value of this field.
	PartitionFieldName *string `json:"partition_field_name,omitempty"`
	// PartitionFieldValue The value of `partition_field_name`.
	PartitionFieldValue *string `json:"partition_field_value,omitempty"`
	// Probability The probability of the individual anomaly occurring, in the range 0 to 1. For
	// example, `0.0000772031`. This value can be held to a high precision of over
	// 300 decimal places, so the `record_score` is provided as a human-readable and
	// friendly interpretation of this.
	Probability Float64 `json:"probability"`
	// RecordScore A normalized score between 0-100, which is based on the probability of the
	// anomalousness of this record. Unlike `initial_record_score`, this value will
	// be updated by a re-normalization process as new data is analyzed.
	RecordScore Float64 `json:"record_score"`
	// ResultType Internal. This is always set to `record`.
	ResultType string `json:"result_type"`
	// Timestamp The start time of the bucket for which these results were calculated.
	Timestamp int64 `json:"timestamp"`
	// Typical The typical value for the bucket, according to analytical modeling.
	Typical []Float64 `json:"typical,omitempty"`
}

// NewAnomaly returns a Anomaly.
func NewAnomaly() *Anomaly {
	r := &Anomaly{}

	return r
}
