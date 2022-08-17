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

// DataframeClassificationSummaryPrecision type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/evaluate_data_frame/types.ts#L60-L63
type DataframeClassificationSummaryPrecision struct {
	AvgPrecision float64                    `json:"avg_precision"`
	Classes      []DataframeEvaluationClass `json:"classes"`
}

// DataframeClassificationSummaryPrecisionBuilder holds DataframeClassificationSummaryPrecision struct and provides a builder API.
type DataframeClassificationSummaryPrecisionBuilder struct {
	v *DataframeClassificationSummaryPrecision
}

// NewDataframeClassificationSummaryPrecision provides a builder for the DataframeClassificationSummaryPrecision struct.
func NewDataframeClassificationSummaryPrecisionBuilder() *DataframeClassificationSummaryPrecisionBuilder {
	r := DataframeClassificationSummaryPrecisionBuilder{
		&DataframeClassificationSummaryPrecision{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeClassificationSummaryPrecision struct
func (rb *DataframeClassificationSummaryPrecisionBuilder) Build() DataframeClassificationSummaryPrecision {
	return *rb.v
}

func (rb *DataframeClassificationSummaryPrecisionBuilder) AvgPrecision(avgprecision float64) *DataframeClassificationSummaryPrecisionBuilder {
	rb.v.AvgPrecision = avgprecision
	return rb
}

func (rb *DataframeClassificationSummaryPrecisionBuilder) Classes(classes []DataframeEvaluationClassBuilder) *DataframeClassificationSummaryPrecisionBuilder {
	tmp := make([]DataframeEvaluationClass, len(classes))
	for _, value := range classes {
		tmp = append(tmp, value.Build())
	}
	rb.v.Classes = tmp
	return rb
}
