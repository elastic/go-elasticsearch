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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

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
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/DataframeAnalytics.ts#L395-L410
type Hyperparameters struct {
	Alpha                                  *Float64 `json:"alpha,omitempty"`
	DownsampleFactor                       *Float64 `json:"downsample_factor,omitempty"`
	Eta                                    *Float64 `json:"eta,omitempty"`
	EtaGrowthRatePerTree                   *Float64 `json:"eta_growth_rate_per_tree,omitempty"`
	FeatureBagFraction                     *Float64 `json:"feature_bag_fraction,omitempty"`
	Gamma                                  *Float64 `json:"gamma,omitempty"`
	Lambda                                 *Float64 `json:"lambda,omitempty"`
	MaxAttemptsToAddTree                   *int     `json:"max_attempts_to_add_tree,omitempty"`
	MaxOptimizationRoundsPerHyperparameter *int     `json:"max_optimization_rounds_per_hyperparameter,omitempty"`
	MaxTrees                               *int     `json:"max_trees,omitempty"`
	NumFolds                               *int     `json:"num_folds,omitempty"`
	NumSplitsPerFeature                    *int     `json:"num_splits_per_feature,omitempty"`
	SoftTreeDepthLimit                     *int     `json:"soft_tree_depth_limit,omitempty"`
	SoftTreeDepthTolerance                 *Float64 `json:"soft_tree_depth_tolerance,omitempty"`
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
