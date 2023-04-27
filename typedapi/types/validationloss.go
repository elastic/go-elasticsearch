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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

// ValidationLoss type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ml/_types/DataframeAnalytics.ts#L428-L433
type ValidationLoss struct {
	// FoldValues Validation loss values for every added decision tree during the forest
	// growing procedure.
	FoldValues []string `json:"fold_values"`
	// LossType The type of the loss metric. For example, binomial_logistic.
	LossType string `json:"loss_type"`
}

// NewValidationLoss returns a ValidationLoss.
func NewValidationLoss() *ValidationLoss {
	r := &ValidationLoss{}

	return r
}
