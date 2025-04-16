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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeAnalysisOutlierDetection struct {
	v *types.DataframeAnalysisOutlierDetection
}

// The configuration information necessary to perform outlier detection. NOTE:
// Advanced parameters are for fine-tuning classification analysis. They are set
// automatically by hyperparameter optimization to give the minimum validation
// error. It is highly recommended to use the default values unless you fully
// understand the function of these parameters.
func NewDataframeAnalysisOutlierDetection() *_dataframeAnalysisOutlierDetection {

	return &_dataframeAnalysisOutlierDetection{v: types.NewDataframeAnalysisOutlierDetection()}

}

func (s *_dataframeAnalysisOutlierDetection) ComputeFeatureInfluence(computefeatureinfluence bool) *_dataframeAnalysisOutlierDetection {

	s.v.ComputeFeatureInfluence = &computefeatureinfluence

	return s
}

func (s *_dataframeAnalysisOutlierDetection) FeatureInfluenceThreshold(featureinfluencethreshold types.Float64) *_dataframeAnalysisOutlierDetection {

	s.v.FeatureInfluenceThreshold = &featureinfluencethreshold

	return s
}

func (s *_dataframeAnalysisOutlierDetection) Method(method string) *_dataframeAnalysisOutlierDetection {

	s.v.Method = &method

	return s
}

func (s *_dataframeAnalysisOutlierDetection) NNeighbors(nneighbors int) *_dataframeAnalysisOutlierDetection {

	s.v.NNeighbors = &nneighbors

	return s
}

func (s *_dataframeAnalysisOutlierDetection) OutlierFraction(outlierfraction types.Float64) *_dataframeAnalysisOutlierDetection {

	s.v.OutlierFraction = &outlierfraction

	return s
}

func (s *_dataframeAnalysisOutlierDetection) StandardizationEnabled(standardizationenabled bool) *_dataframeAnalysisOutlierDetection {

	s.v.StandardizationEnabled = &standardizationenabled

	return s
}

func (s *_dataframeAnalysisOutlierDetection) DataframeAnalysisContainerCaster() *types.DataframeAnalysisContainer {
	container := types.NewDataframeAnalysisContainer()

	container.OutlierDetection = s.v

	return container
}

func (s *_dataframeAnalysisOutlierDetection) DataframeAnalysisOutlierDetectionCaster() *types.DataframeAnalysisOutlierDetection {
	return s.v
}
