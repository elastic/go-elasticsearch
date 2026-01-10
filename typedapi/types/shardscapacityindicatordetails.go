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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package types

// ShardsCapacityIndicatorDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27/specification/_global/health_report/types.ts#L198-L201
type ShardsCapacityIndicatorDetails struct {
	Data   ShardsCapacityIndicatorTierDetail `json:"data"`
	Frozen ShardsCapacityIndicatorTierDetail `json:"frozen"`
}

// NewShardsCapacityIndicatorDetails returns a ShardsCapacityIndicatorDetails.
func NewShardsCapacityIndicatorDetails() *ShardsCapacityIndicatorDetails {
	r := &ShardsCapacityIndicatorDetails{}

	return r
}
