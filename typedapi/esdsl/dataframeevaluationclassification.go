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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _dataframeEvaluationClassification struct {
	v *types.DataframeEvaluationClassification
}

// Classification evaluation evaluates the results of a classification analysis
// which outputs a prediction that identifies to which of the classes each
// document belongs.
func NewDataframeEvaluationClassification() *_dataframeEvaluationClassification {

	return &_dataframeEvaluationClassification{v: types.NewDataframeEvaluationClassification()}

}

// The field of the index which contains the ground truth. The data type of this
// field can be boolean or integer. If the data type is integer, the value has
// to be either 0 (false) or 1 (true).
func (s *_dataframeEvaluationClassification) ActualField(field string) *_dataframeEvaluationClassification {

	s.v.ActualField = field

	return s
}

// Specifies the metrics that are used for the evaluation.
func (s *_dataframeEvaluationClassification) Metrics(metrics types.DataframeEvaluationClassificationMetricsVariant) *_dataframeEvaluationClassification {

	s.v.Metrics = metrics.DataframeEvaluationClassificationMetricsCaster()

	return s
}

// The field in the index which contains the predicted value, in other words the
// results of the classification analysis.
func (s *_dataframeEvaluationClassification) PredictedField(field string) *_dataframeEvaluationClassification {

	s.v.PredictedField = &field

	return s
}

// The field of the index which is an array of documents of the form {
// "class_name": XXX, "class_probability": YYY }. This field must be defined as
// nested in the mappings.
func (s *_dataframeEvaluationClassification) TopClassesField(field string) *_dataframeEvaluationClassification {

	s.v.TopClassesField = &field

	return s
}

func (s *_dataframeEvaluationClassification) DataframeEvaluationContainerCaster() *types.DataframeEvaluationContainer {
	container := types.NewDataframeEvaluationContainer()

	container.Classification = s.v

	return container
}

func (s *_dataframeEvaluationClassification) DataframeEvaluationClassificationCaster() *types.DataframeEvaluationClassification {
	return s.v
}
