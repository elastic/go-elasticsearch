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

// DataframeClassificationSummaryAccuracy type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/evaluate_data_frame/types.ts#L70-L73
type DataframeClassificationSummaryAccuracy struct {
	Classes         []DataframeEvaluationClass `json:"classes"`
	OverallAccuracy float64                    `json:"overall_accuracy"`
}

// DataframeClassificationSummaryAccuracyBuilder holds DataframeClassificationSummaryAccuracy struct and provides a builder API.
type DataframeClassificationSummaryAccuracyBuilder struct {
	v *DataframeClassificationSummaryAccuracy
}

// NewDataframeClassificationSummaryAccuracy provides a builder for the DataframeClassificationSummaryAccuracy struct.
func NewDataframeClassificationSummaryAccuracyBuilder() *DataframeClassificationSummaryAccuracyBuilder {
	r := DataframeClassificationSummaryAccuracyBuilder{
		&DataframeClassificationSummaryAccuracy{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeClassificationSummaryAccuracy struct
func (rb *DataframeClassificationSummaryAccuracyBuilder) Build() DataframeClassificationSummaryAccuracy {
	return *rb.v
}

func (rb *DataframeClassificationSummaryAccuracyBuilder) Classes(classes []DataframeEvaluationClassBuilder) *DataframeClassificationSummaryAccuracyBuilder {
	tmp := make([]DataframeEvaluationClass, len(classes))
	for _, value := range classes {
		tmp = append(tmp, value.Build())
	}
	rb.v.Classes = tmp
	return rb
}

func (rb *DataframeClassificationSummaryAccuracyBuilder) OverallAccuracy(overallaccuracy float64) *DataframeClassificationSummaryAccuracyBuilder {
	rb.v.OverallAccuracy = overallaccuracy
	return rb
}
