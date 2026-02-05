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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package types

// ChangeType type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/_types/aggregations/Aggregate.ts#L398-L410
type ChangeType struct {
	Dip                *Dip                `json:"dip,omitempty"`
	DistributionChange *DistributionChange `json:"distribution_change,omitempty"`
	Indeterminable     *Indeterminable     `json:"indeterminable,omitempty"`
	NonStationary      *NonStationary      `json:"non_stationary,omitempty"`
	Spike              *Spike              `json:"spike,omitempty"`
	Stationary         *Stationary         `json:"stationary,omitempty"`
	StepChange         *StepChange         `json:"step_change,omitempty"`
	TrendChange        *TrendChange        `json:"trend_change,omitempty"`
}

// NewChangeType returns a ChangeType.
func NewChangeType() *ChangeType {
	r := &ChangeType{}

	return r
}
