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

type _dataframeAnalysisClassification struct {
	v *types.DataframeAnalysisClassification
}

// The configuration information necessary to perform classification.
func NewDataframeAnalysisClassification(dependentvariable string) *_dataframeAnalysisClassification {

	tmp := &_dataframeAnalysisClassification{v: types.NewDataframeAnalysisClassification()}

	tmp.DependentVariable(dependentvariable)

	return tmp

}

// Advanced configuration option. Machine learning uses loss guided tree
// growing, which means that the decision trees grow where the regularized loss
// decreases most quickly. This parameter affects loss calculations by acting as
// a multiplier of the tree depth. Higher alpha values result in shallower trees
// and faster training times. By default, this value is calculated during
// hyperparameter optimization. It must be greater than or equal to zero.
func (s *_dataframeAnalysisClassification) Alpha(alpha types.Float64) *_dataframeAnalysisClassification {

	s.v.Alpha = &alpha

	return s
}

func (s *_dataframeAnalysisClassification) ClassAssignmentObjective(classassignmentobjective string) *_dataframeAnalysisClassification {

	s.v.ClassAssignmentObjective = &classassignmentobjective

	return s
}

// Defines which field of the document is to be predicted. It must match one of
// the fields in the index being used to train. If this field is missing from a
// document, then that document will not be used for training, but a prediction
// with the trained model will be generated for it. It is also known as
// continuous target variable.
// For classification analysis, the data type of the field must be numeric
// (`integer`, `short`, `long`, `byte`), categorical (`ip` or `keyword`), or
// `boolean`. There must be no more than 30 different values in this field.
// For regression analysis, the data type of the field must be numeric.
func (s *_dataframeAnalysisClassification) DependentVariable(dependentvariable string) *_dataframeAnalysisClassification {

	s.v.DependentVariable = dependentvariable

	return s
}

// Advanced configuration option. Controls the fraction of data that is used to
// compute the derivatives of the loss function for tree training. A small value
// results in the use of a small fraction of the data. If this value is set to
// be less than 1, accuracy typically improves. However, too small a value may
// result in poor convergence for the ensemble and so require more trees. By
// default, this value is calculated during hyperparameter optimization. It must
// be greater than zero and less than or equal to 1.
func (s *_dataframeAnalysisClassification) DownsampleFactor(downsamplefactor types.Float64) *_dataframeAnalysisClassification {

	s.v.DownsampleFactor = &downsamplefactor

	return s
}

// Advanced configuration option. Specifies whether the training process should
// finish if it is not finding any better performing models. If disabled, the
// training process can take significantly longer and the chance of finding a
// better performing model is unremarkable.
func (s *_dataframeAnalysisClassification) EarlyStoppingEnabled(earlystoppingenabled bool) *_dataframeAnalysisClassification {

	s.v.EarlyStoppingEnabled = &earlystoppingenabled

	return s
}

// Advanced configuration option. The shrinkage applied to the weights. Smaller
// values result in larger forests which have a better generalization error.
// However, larger forests cause slower training. By default, this value is
// calculated during hyperparameter optimization. It must be a value between
// 0.001 and 1.
func (s *_dataframeAnalysisClassification) Eta(eta types.Float64) *_dataframeAnalysisClassification {

	s.v.Eta = &eta

	return s
}

// Advanced configuration option. Specifies the rate at which `eta` increases
// for each new tree that is added to the forest. For example, a rate of 1.05
// increases `eta` by 5% for each extra tree. By default, this value is
// calculated during hyperparameter optimization. It must be between 0.5 and 2.
func (s *_dataframeAnalysisClassification) EtaGrowthRatePerTree(etagrowthratepertree types.Float64) *_dataframeAnalysisClassification {

	s.v.EtaGrowthRatePerTree = &etagrowthratepertree

	return s
}

// Advanced configuration option. Defines the fraction of features that will be
// used when selecting a random bag for each candidate split. By default, this
// value is calculated during hyperparameter optimization.
func (s *_dataframeAnalysisClassification) FeatureBagFraction(featurebagfraction types.Float64) *_dataframeAnalysisClassification {

	s.v.FeatureBagFraction = &featurebagfraction

	return s
}

// Advanced configuration option. A collection of feature preprocessors that
// modify one or more included fields. The analysis uses the resulting one or
// more features instead of the original document field. However, these features
// are ephemeral; they are not stored in the destination index. Multiple
// `feature_processors` entries can refer to the same document fields. Automatic
// categorical feature encoding still occurs for the fields that are unprocessed
// by a custom processor or that have categorical values. Use this property only
// if you want to override the automatic feature encoding of the specified
// fields.
func (s *_dataframeAnalysisClassification) FeatureProcessors(featureprocessors ...types.DataframeAnalysisFeatureProcessorVariant) *_dataframeAnalysisClassification {

	for _, v := range featureprocessors {

		s.v.FeatureProcessors = append(s.v.FeatureProcessors, *v.DataframeAnalysisFeatureProcessorCaster())

	}
	return s
}

// Advanced configuration option. Regularization parameter to prevent
// overfitting on the training data set. Multiplies a linear penalty associated
// with the size of individual trees in the forest. A high gamma value causes
// training to prefer small trees. A small gamma value results in larger
// individual trees and slower training. By default, this value is calculated
// during hyperparameter optimization. It must be a nonnegative value.
func (s *_dataframeAnalysisClassification) Gamma(gamma types.Float64) *_dataframeAnalysisClassification {

	s.v.Gamma = &gamma

	return s
}

// Advanced configuration option. Regularization parameter to prevent
// overfitting on the training data set. Multiplies an L2 regularization term
// which applies to leaf weights of the individual trees in the forest. A high
// lambda value causes training to favor small leaf weights. This behavior makes
// the prediction function smoother at the expense of potentially not being able
// to capture relevant relationships between the features and the dependent
// variable. A small lambda value results in large individual trees and slower
// training. By default, this value is calculated during hyperparameter
// optimization. It must be a nonnegative value.
func (s *_dataframeAnalysisClassification) Lambda(lambda types.Float64) *_dataframeAnalysisClassification {

	s.v.Lambda = &lambda

	return s
}

// Advanced configuration option. A multiplier responsible for determining the
// maximum number of hyperparameter optimization steps in the Bayesian
// optimization procedure. The maximum number of steps is determined based on
// the number of undefined hyperparameters times the maximum optimization rounds
// per hyperparameter. By default, this value is calculated during
// hyperparameter optimization.
func (s *_dataframeAnalysisClassification) MaxOptimizationRoundsPerHyperparameter(maxoptimizationroundsperhyperparameter int) *_dataframeAnalysisClassification {

	s.v.MaxOptimizationRoundsPerHyperparameter = &maxoptimizationroundsperhyperparameter

	return s
}

// Advanced configuration option. Defines the maximum number of decision trees
// in the forest. The maximum value is 2000. By default, this value is
// calculated during hyperparameter optimization.
func (s *_dataframeAnalysisClassification) MaxTrees(maxtrees int) *_dataframeAnalysisClassification {

	s.v.MaxTrees = &maxtrees

	return s
}

// Defines the number of categories for which the predicted probabilities are
// reported. It must be non-negative or -1. If it is -1 or greater than the
// total number of categories, probabilities are reported for all categories; if
// you have a large number of categories, there could be a significant effect on
// the size of your destination index. NOTE: To use the AUC ROC evaluation
// method, `num_top_classes` must be set to -1 or a value greater than or equal
// to the total number of categories.
func (s *_dataframeAnalysisClassification) NumTopClasses(numtopclasses int) *_dataframeAnalysisClassification {

	s.v.NumTopClasses = &numtopclasses

	return s
}

// Advanced configuration option. Specifies the maximum number of feature
// importance values per document to return. By default, no feature importance
// calculation occurs.
func (s *_dataframeAnalysisClassification) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_dataframeAnalysisClassification {

	s.v.NumTopFeatureImportanceValues = &numtopfeatureimportancevalues

	return s
}

// Defines the name of the prediction field in the results. Defaults to
// `<dependent_variable>_prediction`.
func (s *_dataframeAnalysisClassification) PredictionFieldName(field string) *_dataframeAnalysisClassification {

	s.v.PredictionFieldName = &field

	return s
}

// Defines the seed for the random generator that is used to pick training data.
// By default, it is randomly generated. Set it to a specific value to use the
// same training data each time you start a job (assuming other related
// parameters such as `source` and `analyzed_fields` are the same).
func (s *_dataframeAnalysisClassification) RandomizeSeed(randomizeseed types.Float64) *_dataframeAnalysisClassification {

	s.v.RandomizeSeed = &randomizeseed

	return s
}

// Advanced configuration option. Machine learning uses loss guided tree
// growing, which means that the decision trees grow where the regularized loss
// decreases most quickly. This soft limit combines with the
// `soft_tree_depth_tolerance` to penalize trees that exceed the specified
// depth; the regularized loss increases quickly beyond this depth. By default,
// this value is calculated during hyperparameter optimization. It must be
// greater than or equal to 0.
func (s *_dataframeAnalysisClassification) SoftTreeDepthLimit(softtreedepthlimit int) *_dataframeAnalysisClassification {

	s.v.SoftTreeDepthLimit = &softtreedepthlimit

	return s
}

// Advanced configuration option. This option controls how quickly the
// regularized loss increases when the tree depth exceeds
// `soft_tree_depth_limit`. By default, this value is calculated during
// hyperparameter optimization. It must be greater than or equal to 0.01.
func (s *_dataframeAnalysisClassification) SoftTreeDepthTolerance(softtreedepthtolerance types.Float64) *_dataframeAnalysisClassification {

	s.v.SoftTreeDepthTolerance = &softtreedepthtolerance

	return s
}

// Defines what percentage of the eligible documents that will be used for
// training. Documents that are ignored by the analysis (for example those that
// contain arrays with more than one value) won’t be included in the calculation
// for used percentage.
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
