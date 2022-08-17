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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/watcherstate"
)

// WatcherNodeStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/stats/types.ts#L33-L40
type WatcherNodeStats struct {
	CurrentWatches      []WatchRecordStats        `json:"current_watches,omitempty"`
	ExecutionThreadPool ExecutionThreadPool       `json:"execution_thread_pool"`
	NodeId              Id                        `json:"node_id"`
	QueuedWatches       []WatchRecordQueuedStats  `json:"queued_watches,omitempty"`
	WatchCount          int64                     `json:"watch_count"`
	WatcherState        watcherstate.WatcherState `json:"watcher_state"`
}

// WatcherNodeStatsBuilder holds WatcherNodeStats struct and provides a builder API.
type WatcherNodeStatsBuilder struct {
	v *WatcherNodeStats
}

// NewWatcherNodeStats provides a builder for the WatcherNodeStats struct.
func NewWatcherNodeStatsBuilder() *WatcherNodeStatsBuilder {
	r := WatcherNodeStatsBuilder{
		&WatcherNodeStats{},
	}

	return &r
}

// Build finalize the chain and returns the WatcherNodeStats struct
func (rb *WatcherNodeStatsBuilder) Build() WatcherNodeStats {
	return *rb.v
}

func (rb *WatcherNodeStatsBuilder) CurrentWatches(current_watches []WatchRecordStatsBuilder) *WatcherNodeStatsBuilder {
	tmp := make([]WatchRecordStats, len(current_watches))
	for _, value := range current_watches {
		tmp = append(tmp, value.Build())
	}
	rb.v.CurrentWatches = tmp
	return rb
}

func (rb *WatcherNodeStatsBuilder) ExecutionThreadPool(executionthreadpool *ExecutionThreadPoolBuilder) *WatcherNodeStatsBuilder {
	v := executionthreadpool.Build()
	rb.v.ExecutionThreadPool = v
	return rb
}

func (rb *WatcherNodeStatsBuilder) NodeId(nodeid Id) *WatcherNodeStatsBuilder {
	rb.v.NodeId = nodeid
	return rb
}

func (rb *WatcherNodeStatsBuilder) QueuedWatches(queued_watches []WatchRecordQueuedStatsBuilder) *WatcherNodeStatsBuilder {
	tmp := make([]WatchRecordQueuedStats, len(queued_watches))
	for _, value := range queued_watches {
		tmp = append(tmp, value.Build())
	}
	rb.v.QueuedWatches = tmp
	return rb
}

func (rb *WatcherNodeStatsBuilder) WatchCount(watchcount int64) *WatcherNodeStatsBuilder {
	rb.v.WatchCount = watchcount
	return rb
}

func (rb *WatcherNodeStatsBuilder) WatcherState(watcherstate watcherstate.WatcherState) *WatcherNodeStatsBuilder {
	rb.v.WatcherState = watcherstate
	return rb
}
