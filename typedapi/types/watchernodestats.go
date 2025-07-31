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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/watcherstate"
)

// WatcherNodeStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/watcher/stats/types.ts#L33-L61
type WatcherNodeStats struct {
	// CurrentWatches The current executing watches metric gives insight into the watches that are
	// currently being executed by Watcher.
	// Additional information is shared per watch that is currently executing.
	// This information includes the `watch_id`, the time its execution started and
	// its current execution phase.
	// To include this metric, the `metric` option should be set to
	// `current_watches` or `_all`.
	// In addition you can also specify the `emit_stacktraces=true` parameter, which
	// adds stack traces for each watch that is being run.
	// These stack traces can give you more insight into an execution of a watch.
	CurrentWatches      []WatchRecordStats  `json:"current_watches,omitempty"`
	ExecutionThreadPool ExecutionThreadPool `json:"execution_thread_pool"`
	NodeId              string              `json:"node_id"`
	// QueuedWatches Watcher moderates the execution of watches such that their execution won't
	// put too much pressure on the node and its resources.
	// If too many watches trigger concurrently and there isn't enough capacity to
	// run them all, some of the watches are queued, waiting for the current running
	// watches to finish.s
	// The queued watches metric gives insight on these queued watches.
	//
	// To include this metric, the `metric` option should include `queued_watches`
	// or `_all`.
	QueuedWatches []WatchRecordQueuedStats `json:"queued_watches,omitempty"`
	// WatchCount The number of watches currently registered.
	WatchCount int64 `json:"watch_count"`
	// WatcherState The current state of Watcher.
	WatcherState watcherstate.WatcherState `json:"watcher_state"`
}

func (s *WatcherNodeStats) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "current_watches":
			if err := dec.Decode(&s.CurrentWatches); err != nil {
				return fmt.Errorf("%s | %w", "CurrentWatches", err)
			}

		case "execution_thread_pool":
			if err := dec.Decode(&s.ExecutionThreadPool); err != nil {
				return fmt.Errorf("%s | %w", "ExecutionThreadPool", err)
			}

		case "node_id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return fmt.Errorf("%s | %w", "NodeId", err)
			}

		case "queued_watches":
			if err := dec.Decode(&s.QueuedWatches); err != nil {
				return fmt.Errorf("%s | %w", "QueuedWatches", err)
			}

		case "watch_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "WatchCount", err)
				}
				s.WatchCount = value
			case float64:
				f := int64(v)
				s.WatchCount = f
			}

		case "watcher_state":
			if err := dec.Decode(&s.WatcherState); err != nil {
				return fmt.Errorf("%s | %w", "WatcherState", err)
			}

		}
	}
	return nil
}

// NewWatcherNodeStats returns a WatcherNodeStats.
func NewWatcherNodeStats() *WatcherNodeStats {
	r := &WatcherNodeStats{}

	return r
}
