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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/executionphase"
)

// WatchRecordStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/stats/types.ts#L54-L60
type WatchRecordStats struct {
	ExecutedActions []string                      `json:"executed_actions,omitempty"`
	ExecutionPhase  executionphase.ExecutionPhase `json:"execution_phase"`
	ExecutionTime   DateTime                      `json:"execution_time"`
	TriggeredTime   DateTime                      `json:"triggered_time"`
	WatchId         Id                            `json:"watch_id"`
	WatchRecordId   Id                            `json:"watch_record_id"`
}

// WatchRecordStatsBuilder holds WatchRecordStats struct and provides a builder API.
type WatchRecordStatsBuilder struct {
	v *WatchRecordStats
}

// NewWatchRecordStats provides a builder for the WatchRecordStats struct.
func NewWatchRecordStatsBuilder() *WatchRecordStatsBuilder {
	r := WatchRecordStatsBuilder{
		&WatchRecordStats{},
	}

	return &r
}

// Build finalize the chain and returns the WatchRecordStats struct
func (rb *WatchRecordStatsBuilder) Build() WatchRecordStats {
	return *rb.v
}

func (rb *WatchRecordStatsBuilder) ExecutedActions(executed_actions ...string) *WatchRecordStatsBuilder {
	rb.v.ExecutedActions = executed_actions
	return rb
}

func (rb *WatchRecordStatsBuilder) ExecutionPhase(executionphase executionphase.ExecutionPhase) *WatchRecordStatsBuilder {
	rb.v.ExecutionPhase = executionphase
	return rb
}

func (rb *WatchRecordStatsBuilder) ExecutionTime(executiontime *DateTimeBuilder) *WatchRecordStatsBuilder {
	v := executiontime.Build()
	rb.v.ExecutionTime = v
	return rb
}

func (rb *WatchRecordStatsBuilder) TriggeredTime(triggeredtime *DateTimeBuilder) *WatchRecordStatsBuilder {
	v := triggeredtime.Build()
	rb.v.TriggeredTime = v
	return rb
}

func (rb *WatchRecordStatsBuilder) WatchId(watchid Id) *WatchRecordStatsBuilder {
	rb.v.WatchId = watchid
	return rb
}

func (rb *WatchRecordStatsBuilder) WatchRecordId(watchrecordid Id) *WatchRecordStatsBuilder {
	rb.v.WatchRecordId = watchrecordid
	return rb
}
