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

type _dataframeAnalysisContainer struct {
	v *types.DataframeAnalysisContainer
}

func NewDataframeAnalysisContainer() *_dataframeAnalysisContainer {
	return &_dataframeAnalysisContainer{v: types.NewDataframeAnalysisContainer()}
}

// AdditionalDataframeAnalysisContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_dataframeAnalysisContainer) AdditionalDataframeAnalysisContainerProperty(key string, value json.RawMessage) *_dataframeAnalysisContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalDataframeAnalysisContainerProperty = tmp
	return s
}

// The configuration information necessary to perform classification.
func (s *_dataframeAnalysisContainer) Classification(classification types.DataframeAnalysisClassificationVariant) *_dataframeAnalysisContainer {

	s.v.Classification = classification.DataframeAnalysisClassificationCaster()

	return s
}

// The configuration information necessary to perform outlier detection. NOTE:
// Advanced parameters are for fine-tuning classification analysis. They are set
// automatically by hyperparameter optimization to give the minimum validation
// error. It is highly recommended to use the default values unless you fully
// understand the function of these parameters.
func (s *_dataframeAnalysisContainer) OutlierDetection(outlierdetection types.DataframeAnalysisOutlierDetectionVariant) *_dataframeAnalysisContainer {

	s.v.OutlierDetection = outlierdetection.DataframeAnalysisOutlierDetectionCaster()

	return s
}

// The configuration information necessary to perform regression. NOTE: Advanced
// parameters are for fine-tuning regression analysis. They are set
// automatically by hyperparameter optimization to give the minimum validation
// error. It is highly recommended to use the default values unless you fully
// understand the function of these parameters.
func (s *_dataframeAnalysisContainer) Regression(regression types.DataframeAnalysisRegressionVariant) *_dataframeAnalysisContainer {

	s.v.Regression = regression.DataframeAnalysisRegressionCaster()

	return s
}

func (s *_dataframeAnalysisContainer) DataframeAnalysisContainerCaster() *types.DataframeAnalysisContainer {
	return s.v
}
