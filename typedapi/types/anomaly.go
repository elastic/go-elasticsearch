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

// Anomaly type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Anomaly.ts#L24-L47
type Anomaly struct {
	Actual              []float64                `json:"actual,omitempty"`
	BucketSpan          DurationValueUnitSeconds `json:"bucket_span"`
	ByFieldName         *string                  `json:"by_field_name,omitempty"`
	ByFieldValue        *string                  `json:"by_field_value,omitempty"`
	Causes              []AnomalyCause           `json:"causes,omitempty"`
	DetectorIndex       int                      `json:"detector_index"`
	FieldName           *string                  `json:"field_name,omitempty"`
	Function            *string                  `json:"function,omitempty"`
	FunctionDescription *string                  `json:"function_description,omitempty"`
	Influencers         []Influence              `json:"influencers,omitempty"`
	InitialRecordScore  float64                  `json:"initial_record_score"`
	IsInterim           bool                     `json:"is_interim"`
	JobId               string                   `json:"job_id"`
	OverFieldName       *string                  `json:"over_field_name,omitempty"`
	OverFieldValue      *string                  `json:"over_field_value,omitempty"`
	PartitionFieldName  *string                  `json:"partition_field_name,omitempty"`
	PartitionFieldValue *string                  `json:"partition_field_value,omitempty"`
	Probability         float64                  `json:"probability"`
	RecordScore         float64                  `json:"record_score"`
	ResultType          string                   `json:"result_type"`
	Timestamp           EpochTimeUnitMillis      `json:"timestamp"`
	Typical             []float64                `json:"typical,omitempty"`
}

// AnomalyBuilder holds Anomaly struct and provides a builder API.
type AnomalyBuilder struct {
	v *Anomaly
}

// NewAnomaly provides a builder for the Anomaly struct.
func NewAnomalyBuilder() *AnomalyBuilder {
	r := AnomalyBuilder{
		&Anomaly{},
	}

	return &r
}

// Build finalize the chain and returns the Anomaly struct
func (rb *AnomalyBuilder) Build() Anomaly {
	return *rb.v
}

func (rb *AnomalyBuilder) Actual(actual ...float64) *AnomalyBuilder {
	rb.v.Actual = actual
	return rb
}

func (rb *AnomalyBuilder) BucketSpan(bucketspan *DurationValueUnitSecondsBuilder) *AnomalyBuilder {
	v := bucketspan.Build()
	rb.v.BucketSpan = v
	return rb
}

func (rb *AnomalyBuilder) ByFieldName(byfieldname string) *AnomalyBuilder {
	rb.v.ByFieldName = &byfieldname
	return rb
}

func (rb *AnomalyBuilder) ByFieldValue(byfieldvalue string) *AnomalyBuilder {
	rb.v.ByFieldValue = &byfieldvalue
	return rb
}

func (rb *AnomalyBuilder) Causes(causes []AnomalyCauseBuilder) *AnomalyBuilder {
	tmp := make([]AnomalyCause, len(causes))
	for _, value := range causes {
		tmp = append(tmp, value.Build())
	}
	rb.v.Causes = tmp
	return rb
}

func (rb *AnomalyBuilder) DetectorIndex(detectorindex int) *AnomalyBuilder {
	rb.v.DetectorIndex = detectorindex
	return rb
}

func (rb *AnomalyBuilder) FieldName(fieldname string) *AnomalyBuilder {
	rb.v.FieldName = &fieldname
	return rb
}

func (rb *AnomalyBuilder) Function(function string) *AnomalyBuilder {
	rb.v.Function = &function
	return rb
}

func (rb *AnomalyBuilder) FunctionDescription(functiondescription string) *AnomalyBuilder {
	rb.v.FunctionDescription = &functiondescription
	return rb
}

func (rb *AnomalyBuilder) Influencers(influencers []InfluenceBuilder) *AnomalyBuilder {
	tmp := make([]Influence, len(influencers))
	for _, value := range influencers {
		tmp = append(tmp, value.Build())
	}
	rb.v.Influencers = tmp
	return rb
}

func (rb *AnomalyBuilder) InitialRecordScore(initialrecordscore float64) *AnomalyBuilder {
	rb.v.InitialRecordScore = initialrecordscore
	return rb
}

func (rb *AnomalyBuilder) IsInterim(isinterim bool) *AnomalyBuilder {
	rb.v.IsInterim = isinterim
	return rb
}

func (rb *AnomalyBuilder) JobId(jobid string) *AnomalyBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *AnomalyBuilder) OverFieldName(overfieldname string) *AnomalyBuilder {
	rb.v.OverFieldName = &overfieldname
	return rb
}

func (rb *AnomalyBuilder) OverFieldValue(overfieldvalue string) *AnomalyBuilder {
	rb.v.OverFieldValue = &overfieldvalue
	return rb
}

func (rb *AnomalyBuilder) PartitionFieldName(partitionfieldname string) *AnomalyBuilder {
	rb.v.PartitionFieldName = &partitionfieldname
	return rb
}

func (rb *AnomalyBuilder) PartitionFieldValue(partitionfieldvalue string) *AnomalyBuilder {
	rb.v.PartitionFieldValue = &partitionfieldvalue
	return rb
}

func (rb *AnomalyBuilder) Probability(probability float64) *AnomalyBuilder {
	rb.v.Probability = probability
	return rb
}

func (rb *AnomalyBuilder) RecordScore(recordscore float64) *AnomalyBuilder {
	rb.v.RecordScore = recordscore
	return rb
}

func (rb *AnomalyBuilder) ResultType(resulttype string) *AnomalyBuilder {
	rb.v.ResultType = resulttype
	return rb
}

func (rb *AnomalyBuilder) Timestamp(timestamp *EpochTimeUnitMillisBuilder) *AnomalyBuilder {
	v := timestamp.Build()
	rb.v.Timestamp = v
	return rb
}

func (rb *AnomalyBuilder) Typical(typical ...float64) *AnomalyBuilder {
	rb.v.Typical = typical
	return rb
}
