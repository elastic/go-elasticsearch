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

// RecoveryStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/Stats.ts#L161-L166
type RecoveryStats struct {
	CurrentAsSource      int64    `json:"current_as_source"`
	CurrentAsTarget      int64    `json:"current_as_target"`
	ThrottleTime         Duration `json:"throttle_time,omitempty"`
	ThrottleTimeInMillis int64    `json:"throttle_time_in_millis"`
}

// NewRecoveryStats returns a RecoveryStats.
func NewRecoveryStats() *RecoveryStats {
	r := &RecoveryStats{}

	return r
}
