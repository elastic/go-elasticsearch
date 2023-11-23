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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Package watcherstate
package watcherstate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/watcher/stats/types.ts#L26-L31
type WatcherState struct {
	Name string
}

var (
	Stopped = WatcherState{"stopped"}

	Starting = WatcherState{"starting"}

	Started = WatcherState{"started"}

	Stopping = WatcherState{"stopping"}
)

func (w WatcherState) MarshalText() (text []byte, err error) {
	return []byte(w.String()), nil
}

func (w *WatcherState) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "stopped":
		*w = Stopped
	case "starting":
		*w = Starting
	case "started":
		*w = Started
	case "stopping":
		*w = Stopping
	default:
		*w = WatcherState{string(text)}
	}

	return nil
}

func (w WatcherState) String() string {
	return w.Name
}
