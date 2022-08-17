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

// DataframeEvaluationSummaryAucRoc type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/evaluate_data_frame/types.ts#L50-L52
type DataframeEvaluationSummaryAucRoc struct {
	Curve []DataframeEvaluationSummaryAucRocCurveItem `json:"curve,omitempty"`
	Value float64                                     `json:"value"`
}

// DataframeEvaluationSummaryAucRocBuilder holds DataframeEvaluationSummaryAucRoc struct and provides a builder API.
type DataframeEvaluationSummaryAucRocBuilder struct {
	v *DataframeEvaluationSummaryAucRoc
}

// NewDataframeEvaluationSummaryAucRoc provides a builder for the DataframeEvaluationSummaryAucRoc struct.
func NewDataframeEvaluationSummaryAucRocBuilder() *DataframeEvaluationSummaryAucRocBuilder {
	r := DataframeEvaluationSummaryAucRocBuilder{
		&DataframeEvaluationSummaryAucRoc{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeEvaluationSummaryAucRoc struct
func (rb *DataframeEvaluationSummaryAucRocBuilder) Build() DataframeEvaluationSummaryAucRoc {
	return *rb.v
}

func (rb *DataframeEvaluationSummaryAucRocBuilder) Curve(curve []DataframeEvaluationSummaryAucRocCurveItemBuilder) *DataframeEvaluationSummaryAucRocBuilder {
	tmp := make([]DataframeEvaluationSummaryAucRocCurveItem, len(curve))
	for _, value := range curve {
		tmp = append(tmp, value.Build())
	}
	rb.v.Curve = tmp
	return rb
}

func (rb *DataframeEvaluationSummaryAucRocBuilder) Value(value float64) *DataframeEvaluationSummaryAucRocBuilder {
	rb.v.Value = value
	return rb
}
