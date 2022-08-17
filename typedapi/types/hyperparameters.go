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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// Hyperparameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L395-L410
type Hyperparameters struct {
	Alpha                                  *float64 `json:"alpha,omitempty"`
	DownsampleFactor                       *float64 `json:"downsample_factor,omitempty"`
	Eta                                    *float64 `json:"eta,omitempty"`
	EtaGrowthRatePerTree                   *float64 `json:"eta_growth_rate_per_tree,omitempty"`
	FeatureBagFraction                     *float64 `json:"feature_bag_fraction,omitempty"`
	Gamma                                  *float64 `json:"gamma,omitempty"`
	Lambda                                 *float64 `json:"lambda,omitempty"`
	MaxAttemptsToAddTree                   *int     `json:"max_attempts_to_add_tree,omitempty"`
	MaxOptimizationRoundsPerHyperparameter *int     `json:"max_optimization_rounds_per_hyperparameter,omitempty"`
	MaxTrees                               *int     `json:"max_trees,omitempty"`
	NumFolds                               *int     `json:"num_folds,omitempty"`
	NumSplitsPerFeature                    *int     `json:"num_splits_per_feature,omitempty"`
	SoftTreeDepthLimit                     *int     `json:"soft_tree_depth_limit,omitempty"`
	SoftTreeDepthTolerance                 *float64 `json:"soft_tree_depth_tolerance,omitempty"`
}

// HyperparametersBuilder holds Hyperparameters struct and provides a builder API.
type HyperparametersBuilder struct {
	v *Hyperparameters
}

// NewHyperparameters provides a builder for the Hyperparameters struct.
func NewHyperparametersBuilder() *HyperparametersBuilder {
	r := HyperparametersBuilder{
		&Hyperparameters{},
	}

	return &r
}

// Build finalize the chain and returns the Hyperparameters struct
func (rb *HyperparametersBuilder) Build() Hyperparameters {
	return *rb.v
}

func (rb *HyperparametersBuilder) Alpha(alpha float64) *HyperparametersBuilder {
	rb.v.Alpha = &alpha
	return rb
}

func (rb *HyperparametersBuilder) DownsampleFactor(downsamplefactor float64) *HyperparametersBuilder {
	rb.v.DownsampleFactor = &downsamplefactor
	return rb
}

func (rb *HyperparametersBuilder) Eta(eta float64) *HyperparametersBuilder {
	rb.v.Eta = &eta
	return rb
}

func (rb *HyperparametersBuilder) EtaGrowthRatePerTree(etagrowthratepertree float64) *HyperparametersBuilder {
	rb.v.EtaGrowthRatePerTree = &etagrowthratepertree
	return rb
}

func (rb *HyperparametersBuilder) FeatureBagFraction(featurebagfraction float64) *HyperparametersBuilder {
	rb.v.FeatureBagFraction = &featurebagfraction
	return rb
}

func (rb *HyperparametersBuilder) Gamma(gamma float64) *HyperparametersBuilder {
	rb.v.Gamma = &gamma
	return rb
}

func (rb *HyperparametersBuilder) Lambda(lambda float64) *HyperparametersBuilder {
	rb.v.Lambda = &lambda
	return rb
}

func (rb *HyperparametersBuilder) MaxAttemptsToAddTree(maxattemptstoaddtree int) *HyperparametersBuilder {
	rb.v.MaxAttemptsToAddTree = &maxattemptstoaddtree
	return rb
}

func (rb *HyperparametersBuilder) MaxOptimizationRoundsPerHyperparameter(maxoptimizationroundsperhyperparameter int) *HyperparametersBuilder {
	rb.v.MaxOptimizationRoundsPerHyperparameter = &maxoptimizationroundsperhyperparameter
	return rb
}

func (rb *HyperparametersBuilder) MaxTrees(maxtrees int) *HyperparametersBuilder {
	rb.v.MaxTrees = &maxtrees
	return rb
}

func (rb *HyperparametersBuilder) NumFolds(numfolds int) *HyperparametersBuilder {
	rb.v.NumFolds = &numfolds
	return rb
}

func (rb *HyperparametersBuilder) NumSplitsPerFeature(numsplitsperfeature int) *HyperparametersBuilder {
	rb.v.NumSplitsPerFeature = &numsplitsperfeature
	return rb
}

func (rb *HyperparametersBuilder) SoftTreeDepthLimit(softtreedepthlimit int) *HyperparametersBuilder {
	rb.v.SoftTreeDepthLimit = &softtreedepthlimit
	return rb
}

func (rb *HyperparametersBuilder) SoftTreeDepthTolerance(softtreedepthtolerance float64) *HyperparametersBuilder {
	rb.v.SoftTreeDepthTolerance = &softtreedepthtolerance
	return rb
}
