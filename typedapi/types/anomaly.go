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

// Anomaly type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/Anomaly.ts#L24-L51
type Anomaly struct {
	Actual []float64 `json:"actual,omitempty"`
	// AnomalyScoreExplanation Information about the factors impacting the initial anomaly score.
	AnomalyScoreExplanation *AnomalyExplanation `json:"anomaly_score_explanation,omitempty"`
	BucketSpan              int64               `json:"bucket_span"`
	ByFieldName             *string             `json:"by_field_name,omitempty"`
	ByFieldValue            *string             `json:"by_field_value,omitempty"`
	Causes                  []AnomalyCause      `json:"causes,omitempty"`
	DetectorIndex           int                 `json:"detector_index"`
	FieldName               *string             `json:"field_name,omitempty"`
	Function                *string             `json:"function,omitempty"`
	FunctionDescription     *string             `json:"function_description,omitempty"`
	Influencers             []Influence         `json:"influencers,omitempty"`
	InitialRecordScore      float64             `json:"initial_record_score"`
	IsInterim               bool                `json:"is_interim"`
	JobId                   string              `json:"job_id"`
	OverFieldName           *string             `json:"over_field_name,omitempty"`
	OverFieldValue          *string             `json:"over_field_value,omitempty"`
	PartitionFieldName      *string             `json:"partition_field_name,omitempty"`
	PartitionFieldValue     *string             `json:"partition_field_value,omitempty"`
	Probability             float64             `json:"probability"`
	RecordScore             float64             `json:"record_score"`
	ResultType              string              `json:"result_type"`
	Timestamp               int64               `json:"timestamp"`
	Typical                 []float64           `json:"typical,omitempty"`
}

// NewAnomaly returns a Anomaly.
func NewAnomaly() *Anomaly {
	r := &Anomaly{}

	return r
}
