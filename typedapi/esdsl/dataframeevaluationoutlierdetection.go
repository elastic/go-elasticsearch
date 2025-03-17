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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _dataframeEvaluationOutlierDetection struct {
	v *types.DataframeEvaluationOutlierDetection
}

// Outlier detection evaluates the results of an outlier detection analysis
// which outputs the probability that each document is an outlier.
func NewDataframeEvaluationOutlierDetection() *_dataframeEvaluationOutlierDetection {

	return &_dataframeEvaluationOutlierDetection{v: types.NewDataframeEvaluationOutlierDetection()}

}

// The field of the index which contains the ground truth. The data type of this
// field can be boolean or integer. If the data type is integer, the value has
// to be either 0 (false) or 1 (true).
func (s *_dataframeEvaluationOutlierDetection) ActualField(field string) *_dataframeEvaluationOutlierDetection {

	s.v.ActualField = field

	return s
}

// Specifies the metrics that are used for the evaluation.
func (s *_dataframeEvaluationOutlierDetection) Metrics(metrics types.DataframeEvaluationOutlierDetectionMetricsVariant) *_dataframeEvaluationOutlierDetection {

	s.v.Metrics = metrics.DataframeEvaluationOutlierDetectionMetricsCaster()

	return s
}

// The field of the index that defines the probability of whether the item
// belongs to the class in question or not. It’s the field that contains the
// results of the analysis.
func (s *_dataframeEvaluationOutlierDetection) PredictedProbabilityField(field string) *_dataframeEvaluationOutlierDetection {

	s.v.PredictedProbabilityField = field

	return s
}

func (s *_dataframeEvaluationOutlierDetection) DataframeEvaluationContainerCaster() *types.DataframeEvaluationContainer {
	container := types.NewDataframeEvaluationContainer()

	container.OutlierDetection = s.v

	return container
}

func (s *_dataframeEvaluationOutlierDetection) DataframeEvaluationOutlierDetectionCaster() *types.DataframeEvaluationOutlierDetection {
	return s.v
}
