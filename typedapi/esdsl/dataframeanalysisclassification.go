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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeAnalysisClassification struct {
	v *types.DataframeAnalysisClassification
}

// The configuration information necessary to perform classification.
func NewDataframeAnalysisClassification(dependentvariable string) *_dataframeAnalysisClassification {

	tmp := &_dataframeAnalysisClassification{v: types.NewDataframeAnalysisClassification()}

	tmp.DependentVariable(dependentvariable)

	return tmp

}

func (s *_dataframeAnalysisClassification) ClassAssignmentObjective(classassignmentobjective string) *_dataframeAnalysisClassification {

	s.v.ClassAssignmentObjective = &classassignmentobjective

	return s
}

func (s *_dataframeAnalysisClassification) NumTopClasses(numtopclasses int) *_dataframeAnalysisClassification {

	s.v.NumTopClasses = &numtopclasses

	return s
}

func (s *_dataframeAnalysisClassification) Alpha(alpha types.Float64) *_dataframeAnalysisClassification {

	s.v.Alpha = &alpha

	return s
}

func (s *_dataframeAnalysisClassification) DependentVariable(dependentvariable string) *_dataframeAnalysisClassification {

	s.v.DependentVariable = dependentvariable

	return s
}

func (s *_dataframeAnalysisClassification) DownsampleFactor(downsamplefactor types.Float64) *_dataframeAnalysisClassification {

	s.v.DownsampleFactor = &downsamplefactor

	return s
}

func (s *_dataframeAnalysisClassification) EarlyStoppingEnabled(earlystoppingenabled bool) *_dataframeAnalysisClassification {

	s.v.EarlyStoppingEnabled = &earlystoppingenabled

	return s
}

func (s *_dataframeAnalysisClassification) Eta(eta types.Float64) *_dataframeAnalysisClassification {

	s.v.Eta = &eta

	return s
}

func (s *_dataframeAnalysisClassification) EtaGrowthRatePerTree(etagrowthratepertree types.Float64) *_dataframeAnalysisClassification {

	s.v.EtaGrowthRatePerTree = &etagrowthratepertree

	return s
}

func (s *_dataframeAnalysisClassification) FeatureBagFraction(featurebagfraction types.Float64) *_dataframeAnalysisClassification {

	s.v.FeatureBagFraction = &featurebagfraction

	return s
}

func (s *_dataframeAnalysisClassification) FeatureProcessors(featureprocessors ...types.DataframeAnalysisFeatureProcessorVariant) *_dataframeAnalysisClassification {

	for _, v := range featureprocessors {

		s.v.FeatureProcessors = append(s.v.FeatureProcessors, *v.DataframeAnalysisFeatureProcessorCaster())

	}
	return s
}

func (s *_dataframeAnalysisClassification) Gamma(gamma types.Float64) *_dataframeAnalysisClassification {

	s.v.Gamma = &gamma

	return s
}

func (s *_dataframeAnalysisClassification) Lambda(lambda types.Float64) *_dataframeAnalysisClassification {

	s.v.Lambda = &lambda

	return s
}

func (s *_dataframeAnalysisClassification) MaxOptimizationRoundsPerHyperparameter(maxoptimizationroundsperhyperparameter int) *_dataframeAnalysisClassification {

	s.v.MaxOptimizationRoundsPerHyperparameter = &maxoptimizationroundsperhyperparameter

	return s
}

func (s *_dataframeAnalysisClassification) MaxTrees(maxtrees int) *_dataframeAnalysisClassification {

	s.v.MaxTrees = &maxtrees

	return s
}

func (s *_dataframeAnalysisClassification) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_dataframeAnalysisClassification {

	s.v.NumTopFeatureImportanceValues = &numtopfeatureimportancevalues

	return s
}

func (s *_dataframeAnalysisClassification) PredictionFieldName(field string) *_dataframeAnalysisClassification {

	s.v.PredictionFieldName = &field

	return s
}

func (s *_dataframeAnalysisClassification) RandomizeSeed(randomizeseed types.Float64) *_dataframeAnalysisClassification {

	s.v.RandomizeSeed = &randomizeseed

	return s
}

func (s *_dataframeAnalysisClassification) SoftTreeDepthLimit(softtreedepthlimit int) *_dataframeAnalysisClassification {

	s.v.SoftTreeDepthLimit = &softtreedepthlimit

	return s
}

func (s *_dataframeAnalysisClassification) SoftTreeDepthTolerance(softtreedepthtolerance types.Float64) *_dataframeAnalysisClassification {

	s.v.SoftTreeDepthTolerance = &softtreedepthtolerance

	return s
}

func (s *_dataframeAnalysisClassification) TrainingPercent(percentage types.PercentageVariant) *_dataframeAnalysisClassification {

	s.v.TrainingPercent = *percentage.PercentageCaster()

	return s
}

func (s *_dataframeAnalysisClassification) DataframeAnalysisContainerCaster() *types.DataframeAnalysisContainer {
	container := types.NewDataframeAnalysisContainer()

	container.Classification = s.v

	return container
}

func (s *_dataframeAnalysisClassification) DataframeAnalysisClassificationCaster() *types.DataframeAnalysisClassification {
	return s.v
}
