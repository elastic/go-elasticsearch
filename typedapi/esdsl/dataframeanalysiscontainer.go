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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
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

func (s *_dataframeAnalysisContainer) Classification(classification types.DataframeAnalysisClassificationVariant) *_dataframeAnalysisContainer {

	s.v.Classification = classification.DataframeAnalysisClassificationCaster()

	return s
}

func (s *_dataframeAnalysisContainer) OutlierDetection(outlierdetection types.DataframeAnalysisOutlierDetectionVariant) *_dataframeAnalysisContainer {

	s.v.OutlierDetection = outlierdetection.DataframeAnalysisOutlierDetectionCaster()

	return s
}

func (s *_dataframeAnalysisContainer) Regression(regression types.DataframeAnalysisRegressionVariant) *_dataframeAnalysisContainer {

	s.v.Regression = regression.DataframeAnalysisRegressionCaster()

	return s
}

func (s *_dataframeAnalysisContainer) DataframeAnalysisContainerCaster() *types.DataframeAnalysisContainer {
	return s.v
}
