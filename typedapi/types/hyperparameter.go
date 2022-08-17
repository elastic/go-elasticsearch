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

// Hyperparameter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L205-L219
type Hyperparameter struct {
	// AbsoluteImportance A positive number showing how much the parameter influences the variation of
	// the loss function. For hyperparameters with values that are not specified by
	// the user but tuned during hyperparameter optimization.
	AbsoluteImportance *float64 `json:"absolute_importance,omitempty"`
	// Name Name of the hyperparameter.
	Name Name `json:"name"`
	// RelativeImportance A number between 0 and 1 showing the proportion of influence on the variation
	// of the loss function among all tuned hyperparameters. For hyperparameters
	// with values that are not specified by the user but tuned during
	// hyperparameter optimization.
	RelativeImportance *float64 `json:"relative_importance,omitempty"`
	// Supplied Indicates if the hyperparameter is specified by the user (true) or optimized
	// (false).
	Supplied bool `json:"supplied"`
	// Value The value of the hyperparameter, either optimized or specified by the user.
	Value float64 `json:"value"`
}

// HyperparameterBuilder holds Hyperparameter struct and provides a builder API.
type HyperparameterBuilder struct {
	v *Hyperparameter
}

// NewHyperparameter provides a builder for the Hyperparameter struct.
func NewHyperparameterBuilder() *HyperparameterBuilder {
	r := HyperparameterBuilder{
		&Hyperparameter{},
	}

	return &r
}

// Build finalize the chain and returns the Hyperparameter struct
func (rb *HyperparameterBuilder) Build() Hyperparameter {
	return *rb.v
}

// AbsoluteImportance A positive number showing how much the parameter influences the variation of
// the loss function. For hyperparameters with values that are not specified by
// the user but tuned during hyperparameter optimization.

func (rb *HyperparameterBuilder) AbsoluteImportance(absoluteimportance float64) *HyperparameterBuilder {
	rb.v.AbsoluteImportance = &absoluteimportance
	return rb
}

// Name Name of the hyperparameter.

func (rb *HyperparameterBuilder) Name(name Name) *HyperparameterBuilder {
	rb.v.Name = name
	return rb
}

// RelativeImportance A number between 0 and 1 showing the proportion of influence on the variation
// of the loss function among all tuned hyperparameters. For hyperparameters
// with values that are not specified by the user but tuned during
// hyperparameter optimization.

func (rb *HyperparameterBuilder) RelativeImportance(relativeimportance float64) *HyperparameterBuilder {
	rb.v.RelativeImportance = &relativeimportance
	return rb
}

// Supplied Indicates if the hyperparameter is specified by the user (true) or optimized
// (false).

func (rb *HyperparameterBuilder) Supplied(supplied bool) *HyperparameterBuilder {
	rb.v.Supplied = supplied
	return rb
}

// Value The value of the hyperparameter, either optimized or specified by the user.

func (rb *HyperparameterBuilder) Value(value float64) *HyperparameterBuilder {
	rb.v.Value = value
	return rb
}
