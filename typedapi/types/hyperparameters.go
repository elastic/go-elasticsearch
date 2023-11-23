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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Hyperparameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/DataframeAnalytics.ts#L419-L525
type Hyperparameters struct {
	// Alpha Advanced configuration option.
	// Machine learning uses loss guided tree growing, which means that the decision
	// trees grow where the regularized loss decreases most quickly.
	// This parameter affects loss calculations by acting as a multiplier of the
	// tree depth.
	// Higher alpha values result in shallower trees and faster training times.
	// By default, this value is calculated during hyperparameter optimization.
	// It must be greater than or equal to zero.
	Alpha *Float64 `json:"alpha,omitempty"`
	// DownsampleFactor Advanced configuration option.
	// Controls the fraction of data that is used to compute the derivatives of the
	// loss function for tree training.
	// A small value results in the use of a small fraction of the data.
	// If this value is set to be less than 1, accuracy typically improves.
	// However, too small a value may result in poor convergence for the ensemble
	// and so require more trees.
	// By default, this value is calculated during hyperparameter optimization.
	// It must be greater than zero and less than or equal to 1.
	DownsampleFactor *Float64 `json:"downsample_factor,omitempty"`
	// Eta Advanced configuration option.
	// The shrinkage applied to the weights.
	// Smaller values result in larger forests which have a better generalization
	// error.
	// However, larger forests cause slower training.
	// By default, this value is calculated during hyperparameter optimization.
	// It must be a value between `0.001` and `1`.
	Eta *Float64 `json:"eta,omitempty"`
	// EtaGrowthRatePerTree Advanced configuration option.
	// Specifies the rate at which `eta` increases for each new tree that is added
	// to the forest.
	// For example, a rate of 1.05 increases `eta` by 5% for each extra tree.
	// By default, this value is calculated during hyperparameter optimization.
	// It must be between `0.5` and `2`.
	EtaGrowthRatePerTree *Float64 `json:"eta_growth_rate_per_tree,omitempty"`
	// FeatureBagFraction Advanced configuration option.
	// Defines the fraction of features that will be used when selecting a random
	// bag for each candidate split.
	// By default, this value is calculated during hyperparameter optimization.
	FeatureBagFraction *Float64 `json:"feature_bag_fraction,omitempty"`
	// Gamma Advanced configuration option.
	// Regularization parameter to prevent overfitting on the training data set.
	// Multiplies a linear penalty associated with the size of individual trees in
	// the forest.
	// A high gamma value causes training to prefer small trees.
	// A small gamma value results in larger individual trees and slower training.
	// By default, this value is calculated during hyperparameter optimization.
	// It must be a nonnegative value.
	Gamma *Float64 `json:"gamma,omitempty"`
	// Lambda Advanced configuration option.
	// Regularization parameter to prevent overfitting on the training data set.
	// Multiplies an L2 regularization term which applies to leaf weights of the
	// individual trees in the forest.
	// A high lambda value causes training to favor small leaf weights.
	// This behavior makes the prediction function smoother at the expense of
	// potentially not being able to capture relevant relationships between the
	// features and the dependent variable.
	// A small lambda value results in large individual trees and slower training.
	// By default, this value is calculated during hyperparameter optimization.
	// It must be a nonnegative value.
	Lambda *Float64 `json:"lambda,omitempty"`
	// MaxAttemptsToAddTree If the algorithm fails to determine a non-trivial tree (more than a single
	// leaf), this parameter determines how many of such consecutive failures are
	// tolerated.
	// Once the number of attempts exceeds the threshold, the forest training stops.
	MaxAttemptsToAddTree *int `json:"max_attempts_to_add_tree,omitempty"`
	// MaxOptimizationRoundsPerHyperparameter Advanced configuration option.
	// A multiplier responsible for determining the maximum number of hyperparameter
	// optimization steps in the Bayesian optimization procedure.
	// The maximum number of steps is determined based on the number of undefined
	// hyperparameters times the maximum optimization rounds per hyperparameter.
	// By default, this value is calculated during hyperparameter optimization.
	MaxOptimizationRoundsPerHyperparameter *int `json:"max_optimization_rounds_per_hyperparameter,omitempty"`
	// MaxTrees Advanced configuration option.
	// Defines the maximum number of decision trees in the forest.
	// The maximum value is 2000.
	// By default, this value is calculated during hyperparameter optimization.
	MaxTrees *int `json:"max_trees,omitempty"`
	// NumFolds The maximum number of folds for the cross-validation procedure.
	NumFolds *int `json:"num_folds,omitempty"`
	// NumSplitsPerFeature Determines the maximum number of splits for every feature that can occur in a
	// decision tree when the tree is trained.
	NumSplitsPerFeature *int `json:"num_splits_per_feature,omitempty"`
	// SoftTreeDepthLimit Advanced configuration option.
	// Machine learning uses loss guided tree growing, which means that the decision
	// trees grow where the regularized loss decreases most quickly.
	// This soft limit combines with the `soft_tree_depth_tolerance` to penalize
	// trees that exceed the specified depth; the regularized loss increases quickly
	// beyond this depth.
	// By default, this value is calculated during hyperparameter optimization.
	// It must be greater than or equal to 0.
	SoftTreeDepthLimit *int `json:"soft_tree_depth_limit,omitempty"`
	// SoftTreeDepthTolerance Advanced configuration option.
	// This option controls how quickly the regularized loss increases when the tree
	// depth exceeds `soft_tree_depth_limit`.
	// By default, this value is calculated during hyperparameter optimization.
	// It must be greater than or equal to 0.01.
	SoftTreeDepthTolerance *Float64 `json:"soft_tree_depth_tolerance,omitempty"`
}

func (s *Hyperparameters) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "alpha":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Alpha = &f
			case float64:
				f := Float64(v)
				s.Alpha = &f
			}

		case "downsample_factor":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.DownsampleFactor = &f
			case float64:
				f := Float64(v)
				s.DownsampleFactor = &f
			}

		case "eta":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Eta = &f
			case float64:
				f := Float64(v)
				s.Eta = &f
			}

		case "eta_growth_rate_per_tree":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.EtaGrowthRatePerTree = &f
			case float64:
				f := Float64(v)
				s.EtaGrowthRatePerTree = &f
			}

		case "feature_bag_fraction":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.FeatureBagFraction = &f
			case float64:
				f := Float64(v)
				s.FeatureBagFraction = &f
			}

		case "gamma":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Gamma = &f
			case float64:
				f := Float64(v)
				s.Gamma = &f
			}

		case "lambda":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Lambda = &f
			case float64:
				f := Float64(v)
				s.Lambda = &f
			}

		case "max_attempts_to_add_tree":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxAttemptsToAddTree = &value
			case float64:
				f := int(v)
				s.MaxAttemptsToAddTree = &f
			}

		case "max_optimization_rounds_per_hyperparameter":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxOptimizationRoundsPerHyperparameter = &value
			case float64:
				f := int(v)
				s.MaxOptimizationRoundsPerHyperparameter = &f
			}

		case "max_trees":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxTrees = &value
			case float64:
				f := int(v)
				s.MaxTrees = &f
			}

		case "num_folds":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NumFolds = &value
			case float64:
				f := int(v)
				s.NumFolds = &f
			}

		case "num_splits_per_feature":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NumSplitsPerFeature = &value
			case float64:
				f := int(v)
				s.NumSplitsPerFeature = &f
			}

		case "soft_tree_depth_limit":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.SoftTreeDepthLimit = &value
			case float64:
				f := int(v)
				s.SoftTreeDepthLimit = &f
			}

		case "soft_tree_depth_tolerance":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.SoftTreeDepthTolerance = &f
			case float64:
				f := Float64(v)
				s.SoftTreeDepthTolerance = &f
			}

		}
	}
	return nil
}

// NewHyperparameters returns a Hyperparameters.
func NewHyperparameters() *Hyperparameters {
	r := &Hyperparameters{}

	return r
}
