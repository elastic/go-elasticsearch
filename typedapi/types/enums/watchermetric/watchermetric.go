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
// https://github.com/elastic/elasticsearch-specification/tree/17ac39c7f9266bc303baa029f90194aecb1c3b7c

// Package watchermetric
package watchermetric

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/17ac39c7f9266bc303baa029f90194aecb1c3b7c/specification/watcher/stats/types.ts#L42-L48
type WatcherMetric struct {
	Name string
}

var (
	All = WatcherMetric{"_all"}

	Queuedwatches = WatcherMetric{"queued_watches"}

	Currentwatches = WatcherMetric{"current_watches"}

	Pendingwatches = WatcherMetric{"pending_watches"}
)

func (w WatcherMetric) MarshalText() (text []byte, err error) {
	return []byte(w.String()), nil
}

func (w *WatcherMetric) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "_all":
		*w = All
	case "queued_watches":
		*w = Queuedwatches
	case "current_watches":
		*w = Currentwatches
	case "pending_watches":
		*w = Pendingwatches
	default:
		*w = WatcherMetric{string(text)}
	}

	return nil
}

func (w WatcherMetric) String() string {
	return w.Name
}
