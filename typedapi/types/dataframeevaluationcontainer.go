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

// DataframeEvaluationContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeEvaluation.ts#L25-L33
type DataframeEvaluationContainer struct {
	// Classification Classification evaluation evaluates the results of a classification analysis
	// which outputs a prediction that identifies to which of the classes each
	// document belongs.
	Classification *DataframeEvaluationClassification `json:"classification,omitempty"`
	// OutlierDetection Outlier detection evaluates the results of an outlier detection analysis
	// which outputs the probability that each document is an outlier.
	OutlierDetection *DataframeEvaluationOutlierDetection `json:"outlier_detection,omitempty"`
	// Regression Regression evaluation evaluates the results of a regression analysis which
	// outputs a prediction of values.
	Regression *DataframeEvaluationRegression `json:"regression,omitempty"`
}

// DataframeEvaluationContainerBuilder holds DataframeEvaluationContainer struct and provides a builder API.
type DataframeEvaluationContainerBuilder struct {
	v *DataframeEvaluationContainer
}

// NewDataframeEvaluationContainer provides a builder for the DataframeEvaluationContainer struct.
func NewDataframeEvaluationContainerBuilder() *DataframeEvaluationContainerBuilder {
	r := DataframeEvaluationContainerBuilder{
		&DataframeEvaluationContainer{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeEvaluationContainer struct
func (rb *DataframeEvaluationContainerBuilder) Build() DataframeEvaluationContainer {
	return *rb.v
}

// Classification Classification evaluation evaluates the results of a classification analysis
// which outputs a prediction that identifies to which of the classes each
// document belongs.

func (rb *DataframeEvaluationContainerBuilder) Classification(classification *DataframeEvaluationClassificationBuilder) *DataframeEvaluationContainerBuilder {
	v := classification.Build()
	rb.v.Classification = &v
	return rb
}

// OutlierDetection Outlier detection evaluates the results of an outlier detection analysis
// which outputs the probability that each document is an outlier.

func (rb *DataframeEvaluationContainerBuilder) OutlierDetection(outlierdetection *DataframeEvaluationOutlierDetectionBuilder) *DataframeEvaluationContainerBuilder {
	v := outlierdetection.Build()
	rb.v.OutlierDetection = &v
	return rb
}

// Regression Regression evaluation evaluates the results of a regression analysis which
// outputs a prediction of values.

func (rb *DataframeEvaluationContainerBuilder) Regression(regression *DataframeEvaluationRegressionBuilder) *DataframeEvaluationContainerBuilder {
	v := regression.Build()
	rb.v.Regression = &v
	return rb
}
