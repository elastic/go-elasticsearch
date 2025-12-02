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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeAnalysisRegression struct {
	v *types.DataframeAnalysisRegression
}

// The configuration information necessary to perform regression. NOTE: Advanced
// parameters are for fine-tuning regression analysis. They are set
// automatically by hyperparameter optimization to give the minimum validation
// error. It is highly recommended to use the default values unless you fully
// understand the function of these parameters.
func NewDataframeAnalysisRegression(dependentvariable string) *_dataframeAnalysisRegression {

	tmp := &_dataframeAnalysisRegression{v: types.NewDataframeAnalysisRegression()}

	tmp.DependentVariable(dependentvariable)

	return tmp

}

func (s *_dataframeAnalysisRegression) Alpha(alpha types.Float64) *_dataframeAnalysisRegression {

	s.v.Alpha = &alpha

	return s
}

func (s *_dataframeAnalysisRegression) DependentVariable(dependentvariable string) *_dataframeAnalysisRegression {

	s.v.DependentVariable = dependentvariable

	return s
}

func (s *_dataframeAnalysisRegression) DownsampleFactor(downsamplefactor types.Float64) *_dataframeAnalysisRegression {

	s.v.DownsampleFactor = &downsamplefactor

	return s
}

func (s *_dataframeAnalysisRegression) EarlyStoppingEnabled(earlystoppingenabled bool) *_dataframeAnalysisRegression {

	s.v.EarlyStoppingEnabled = &earlystoppingenabled

	return s
}

func (s *_dataframeAnalysisRegression) Eta(eta types.Float64) *_dataframeAnalysisRegression {

	s.v.Eta = &eta

	return s
}

func (s *_dataframeAnalysisRegression) EtaGrowthRatePerTree(etagrowthratepertree types.Float64) *_dataframeAnalysisRegression {

	s.v.EtaGrowthRatePerTree = &etagrowthratepertree

	return s
}

func (s *_dataframeAnalysisRegression) FeatureBagFraction(featurebagfraction types.Float64) *_dataframeAnalysisRegression {

	s.v.FeatureBagFraction = &featurebagfraction

	return s
}

func (s *_dataframeAnalysisRegression) FeatureProcessors(featureprocessors ...types.DataframeAnalysisFeatureProcessorVariant) *_dataframeAnalysisRegression {

	for _, v := range featureprocessors {

		s.v.FeatureProcessors = append(s.v.FeatureProcessors, *v.DataframeAnalysisFeatureProcessorCaster())

	}
	return s
}

func (s *_dataframeAnalysisRegression) Gamma(gamma types.Float64) *_dataframeAnalysisRegression {

	s.v.Gamma = &gamma

	return s
}

func (s *_dataframeAnalysisRegression) Lambda(lambda types.Float64) *_dataframeAnalysisRegression {

	s.v.Lambda = &lambda

	return s
}

func (s *_dataframeAnalysisRegression) LossFunction(lossfunction string) *_dataframeAnalysisRegression {

	s.v.LossFunction = &lossfunction

	return s
}

func (s *_dataframeAnalysisRegression) LossFunctionParameter(lossfunctionparameter types.Float64) *_dataframeAnalysisRegression {

	s.v.LossFunctionParameter = &lossfunctionparameter

	return s
}

func (s *_dataframeAnalysisRegression) MaxOptimizationRoundsPerHyperparameter(maxoptimizationroundsperhyperparameter int) *_dataframeAnalysisRegression {

	s.v.MaxOptimizationRoundsPerHyperparameter = &maxoptimizationroundsperhyperparameter

	return s
}

func (s *_dataframeAnalysisRegression) MaxTrees(maxtrees int) *_dataframeAnalysisRegression {

	s.v.MaxTrees = &maxtrees

	return s
}

func (s *_dataframeAnalysisRegression) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_dataframeAnalysisRegression {

	s.v.NumTopFeatureImportanceValues = &numtopfeatureimportancevalues

	return s
}

func (s *_dataframeAnalysisRegression) PredictionFieldName(field string) *_dataframeAnalysisRegression {

	s.v.PredictionFieldName = &field

	return s
}

func (s *_dataframeAnalysisRegression) RandomizeSeed(randomizeseed types.Float64) *_dataframeAnalysisRegression {

	s.v.RandomizeSeed = &randomizeseed

	return s
}

func (s *_dataframeAnalysisRegression) SoftTreeDepthLimit(softtreedepthlimit int) *_dataframeAnalysisRegression {

	s.v.SoftTreeDepthLimit = &softtreedepthlimit

	return s
}

func (s *_dataframeAnalysisRegression) SoftTreeDepthTolerance(softtreedepthtolerance types.Float64) *_dataframeAnalysisRegression {

	s.v.SoftTreeDepthTolerance = &softtreedepthtolerance

	return s
}

func (s *_dataframeAnalysisRegression) TrainingPercent(percentage types.PercentageVariant) *_dataframeAnalysisRegression {

	s.v.TrainingPercent = *percentage.PercentageCaster()

	return s
}

func (s *_dataframeAnalysisRegression) DataframeAnalysisContainerCaster() *types.DataframeAnalysisContainer {
	container := types.NewDataframeAnalysisContainer()

	container.Regression = s.v

	return container
}

func (s *_dataframeAnalysisRegression) DataframeAnalysisRegressionCaster() *types.DataframeAnalysisRegression {
	return s.v
}
