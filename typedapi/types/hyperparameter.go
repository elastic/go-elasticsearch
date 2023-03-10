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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// Hyperparameter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/TrainedModel.ts#L206-L220
type Hyperparameter struct {
	// AbsoluteImportance A positive number showing how much the parameter influences the variation of
	// the loss function. For hyperparameters with values that are not specified by
	// the user but tuned during hyperparameter optimization.
	AbsoluteImportance *Float64 `json:"absolute_importance,omitempty"`
	// Name Name of the hyperparameter.
	Name string `json:"name"`
	// RelativeImportance A number between 0 and 1 showing the proportion of influence on the variation
	// of the loss function among all tuned hyperparameters. For hyperparameters
	// with values that are not specified by the user but tuned during
	// hyperparameter optimization.
	RelativeImportance *Float64 `json:"relative_importance,omitempty"`
	// Supplied Indicates if the hyperparameter is specified by the user (true) or optimized
	// (false).
	Supplied bool `json:"supplied"`
	// Value The value of the hyperparameter, either optimized or specified by the user.
	Value Float64 `json:"value"`
}

// NewHyperparameter returns a Hyperparameter.
func NewHyperparameter() *Hyperparameter {
	r := &Hyperparameter{}

	return r
}
