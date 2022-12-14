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

// TranslogStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/indices/recovery/types.ts#L102-L109
type TranslogStatus struct {
	Percent           Percentage `json:"percent"`
	Recovered         int64      `json:"recovered"`
	Total             int64      `json:"total"`
	TotalOnStart      int64      `json:"total_on_start"`
	TotalTime         *Duration  `json:"total_time,omitempty"`
	TotalTimeInMillis int64      `json:"total_time_in_millis"`
}

// NewTranslogStatus returns a TranslogStatus.
func NewTranslogStatus() *TranslogStatus {
	r := &TranslogStatus{}

	return r
}
