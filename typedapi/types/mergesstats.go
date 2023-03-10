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

// MergesStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/Stats.ts#L119-L136
type MergesStats struct {
	Current                    int64    `json:"current"`
	CurrentDocs                int64    `json:"current_docs"`
	CurrentSize                *string  `json:"current_size,omitempty"`
	CurrentSizeInBytes         int64    `json:"current_size_in_bytes"`
	Total                      int64    `json:"total"`
	TotalAutoThrottle          *string  `json:"total_auto_throttle,omitempty"`
	TotalAutoThrottleInBytes   int64    `json:"total_auto_throttle_in_bytes"`
	TotalDocs                  int64    `json:"total_docs"`
	TotalSize                  *string  `json:"total_size,omitempty"`
	TotalSizeInBytes           int64    `json:"total_size_in_bytes"`
	TotalStoppedTime           Duration `json:"total_stopped_time,omitempty"`
	TotalStoppedTimeInMillis   int64    `json:"total_stopped_time_in_millis"`
	TotalThrottledTime         Duration `json:"total_throttled_time,omitempty"`
	TotalThrottledTimeInMillis int64    `json:"total_throttled_time_in_millis"`
	TotalTime                  Duration `json:"total_time,omitempty"`
	TotalTimeInMillis          int64    `json:"total_time_in_millis"`
}

// NewMergesStats returns a MergesStats.
func NewMergesStats() *MergesStats {
	r := &MergesStats{}

	return r
}
