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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// Hyperparameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/DataframeAnalytics.ts#L395-L410
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

// NewHyperparameters returns a Hyperparameters.
func NewHyperparameters() *Hyperparameters {
	r := &Hyperparameters{}

	return r
}
