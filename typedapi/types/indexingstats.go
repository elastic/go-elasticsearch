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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

// IndexingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_types/Stats.ts#L101-L116
type IndexingStats struct {
	DeleteCurrent        int64                    `json:"delete_current"`
	DeleteTime           *Duration                `json:"delete_time,omitempty"`
	DeleteTimeInMillis   int64                    `json:"delete_time_in_millis"`
	DeleteTotal          int64                    `json:"delete_total"`
	IndexCurrent         int64                    `json:"index_current"`
	IndexFailed          int64                    `json:"index_failed"`
	IndexTime            *Duration                `json:"index_time,omitempty"`
	IndexTimeInMillis    int64                    `json:"index_time_in_millis"`
	IndexTotal           int64                    `json:"index_total"`
	IsThrottled          bool                     `json:"is_throttled"`
	NoopUpdateTotal      int64                    `json:"noop_update_total"`
	ThrottleTime         *Duration                `json:"throttle_time,omitempty"`
	ThrottleTimeInMillis int64                    `json:"throttle_time_in_millis"`
	Types                map[string]IndexingStats `json:"types,omitempty"`
}

// NewIndexingStats returns a IndexingStats.
func NewIndexingStats() *IndexingStats {
	r := &IndexingStats{
		Types: make(map[string]IndexingStats, 0),
	}

	return r
}
