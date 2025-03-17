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

// Specifies whether the feature influence calculation is enabled.
func (s *_dataframeAnalysisOutlierDetection) ComputeFeatureInfluence(computefeatureinfluence bool) *_dataframeAnalysisOutlierDetection {

	s.v.ComputeFeatureInfluence = &computefeatureinfluence

	return s
}

// The minimum outlier score that a document needs to have in order to calculate
// its feature influence score. Value range: 0-1.
func (s *_dataframeAnalysisOutlierDetection) FeatureInfluenceThreshold(featureinfluencethreshold types.Float64) *_dataframeAnalysisOutlierDetection {

	s.v.FeatureInfluenceThreshold = &featureinfluencethreshold

	return s
}

// The method that outlier detection uses. Available methods are `lof`, `ldof`,
// `distance_kth_nn`, `distance_knn`, and `ensemble`. The default value is
// ensemble, which means that outlier detection uses an ensemble of different
// methods and normalises and combines their individual outlier scores to obtain
// the overall outlier score.
func (s *_dataframeAnalysisOutlierDetection) Method(method string) *_dataframeAnalysisOutlierDetection {

	s.v.Method = &method

	return s
}

// Defines the value for how many nearest neighbors each method of outlier
// detection uses to calculate its outlier score. When the value is not set,
// different values are used for different ensemble members. This default
// behavior helps improve the diversity in the ensemble; only override it if you
// are confident that the value you choose is appropriate for the data set.
func (s *_dataframeAnalysisOutlierDetection) NNeighbors(nneighbors int) *_dataframeAnalysisOutlierDetection {

	s.v.NNeighbors = &nneighbors

	return s
}

// The proportion of the data set that is assumed to be outlying prior to
// outlier detection. For example, 0.05 means it is assumed that 5% of values
// are real outliers and 95% are inliers.
func (s *_dataframeAnalysisOutlierDetection) OutlierFraction(outlierfraction types.Float64) *_dataframeAnalysisOutlierDetection {

	s.v.OutlierFraction = &outlierfraction

	return s
}

// If true, the following operation is performed on the columns before computing
// outlier scores: `(x_i - mean(x_i)) / sd(x_i)`.
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
