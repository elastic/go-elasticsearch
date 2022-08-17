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

// DataframeRegressionSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/evaluate_data_frame/types.ts#L39-L44
type DataframeRegressionSummary struct {
	Huber    *DataframeEvaluationValue `json:"huber,omitempty"`
	Mse      *DataframeEvaluationValue `json:"mse,omitempty"`
	Msle     *DataframeEvaluationValue `json:"msle,omitempty"`
	RSquared *DataframeEvaluationValue `json:"r_squared,omitempty"`
}

// DataframeRegressionSummaryBuilder holds DataframeRegressionSummary struct and provides a builder API.
type DataframeRegressionSummaryBuilder struct {
	v *DataframeRegressionSummary
}

// NewDataframeRegressionSummary provides a builder for the DataframeRegressionSummary struct.
func NewDataframeRegressionSummaryBuilder() *DataframeRegressionSummaryBuilder {
	r := DataframeRegressionSummaryBuilder{
		&DataframeRegressionSummary{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeRegressionSummary struct
func (rb *DataframeRegressionSummaryBuilder) Build() DataframeRegressionSummary {
	return *rb.v
}

func (rb *DataframeRegressionSummaryBuilder) Huber(huber *DataframeEvaluationValueBuilder) *DataframeRegressionSummaryBuilder {
	v := huber.Build()
	rb.v.Huber = &v
	return rb
}

func (rb *DataframeRegressionSummaryBuilder) Mse(mse *DataframeEvaluationValueBuilder) *DataframeRegressionSummaryBuilder {
	v := mse.Build()
	rb.v.Mse = &v
	return rb
}

func (rb *DataframeRegressionSummaryBuilder) Msle(msle *DataframeEvaluationValueBuilder) *DataframeRegressionSummaryBuilder {
	v := msle.Build()
	rb.v.Msle = &v
	return rb
}

func (rb *DataframeRegressionSummaryBuilder) RSquared(rsquared *DataframeEvaluationValueBuilder) *DataframeRegressionSummaryBuilder {
	v := rsquared.Build()
	rb.v.RSquared = &v
	return rb
}
