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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _dataframeEvaluationContainer struct {
	v *types.DataframeEvaluationContainer
}

func NewDataframeEvaluationContainer() *_dataframeEvaluationContainer {
	return &_dataframeEvaluationContainer{v: types.NewDataframeEvaluationContainer()}
}

// AdditionalDataframeEvaluationContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_dataframeEvaluationContainer) AdditionalDataframeEvaluationContainerProperty(key string, value json.RawMessage) *_dataframeEvaluationContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalDataframeEvaluationContainerProperty = tmp
	return s
}

// Classification evaluation evaluates the results of a classification analysis
// which outputs a prediction that identifies to which of the classes each
// document belongs.
func (s *_dataframeEvaluationContainer) Classification(classification types.DataframeEvaluationClassificationVariant) *_dataframeEvaluationContainer {

	s.v.Classification = classification.DataframeEvaluationClassificationCaster()

	return s
}

// Outlier detection evaluates the results of an outlier detection analysis
// which outputs the probability that each document is an outlier.
func (s *_dataframeEvaluationContainer) OutlierDetection(outlierdetection types.DataframeEvaluationOutlierDetectionVariant) *_dataframeEvaluationContainer {

	s.v.OutlierDetection = outlierdetection.DataframeEvaluationOutlierDetectionCaster()

	return s
}

// Regression evaluation evaluates the results of a regression analysis which
// outputs a prediction of values.
func (s *_dataframeEvaluationContainer) Regression(regression types.DataframeEvaluationRegressionVariant) *_dataframeEvaluationContainer {

	s.v.Regression = regression.DataframeEvaluationRegressionCaster()

	return s
}

func (s *_dataframeEvaluationContainer) DataframeEvaluationContainerCaster() *types.DataframeEvaluationContainer {
	return s.v
}
