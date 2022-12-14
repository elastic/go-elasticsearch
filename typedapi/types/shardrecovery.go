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

// ShardRecovery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/indices/recovery/types.ts#L118-L135
type ShardRecovery struct {
	Id                int64                `json:"id"`
	Index             RecoveryIndexStatus  `json:"index"`
	Primary           bool                 `json:"primary"`
	Source            RecoveryOrigin       `json:"source"`
	Stage             string               `json:"stage"`
	Start             *RecoveryStartStatus `json:"start,omitempty"`
	StartTime         *DateTime            `json:"start_time,omitempty"`
	StartTimeInMillis int64                `json:"start_time_in_millis"`
	StopTime          *DateTime            `json:"stop_time,omitempty"`
	StopTimeInMillis  *int64               `json:"stop_time_in_millis,omitempty"`
	Target            RecoveryOrigin       `json:"target"`
	TotalTime         *Duration            `json:"total_time,omitempty"`
	TotalTimeInMillis int64                `json:"total_time_in_millis"`
	Translog          TranslogStatus       `json:"translog"`
	Type              string               `json:"type"`
	VerifyIndex       VerifyIndex          `json:"verify_index"`
}

// NewShardRecovery returns a ShardRecovery.
func NewShardRecovery() *ShardRecovery {
	r := &ShardRecovery{}

	return r
}
