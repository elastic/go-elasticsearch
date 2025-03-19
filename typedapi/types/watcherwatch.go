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
// https://github.com/elastic/elasticsearch-specification/tree/3ea9ce260df22d3244bff5bace485dd97ff4046d

package types

// WatcherWatch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3ea9ce260df22d3244bff5bace485dd97ff4046d/specification/xpack/usage/types.ts#L410-L415
type WatcherWatch struct {
	Action    map[string]Counter  `json:"action,omitempty"`
	Condition map[string]Counter  `json:"condition,omitempty"`
	Input     map[string]Counter  `json:"input"`
	Trigger   WatcherWatchTrigger `json:"trigger"`
}

// NewWatcherWatch returns a WatcherWatch.
func NewWatcherWatch() *WatcherWatch {
	r := &WatcherWatch{
		Action:    make(map[string]Counter),
		Condition: make(map[string]Counter),
		Input:     make(map[string]Counter),
	}

	return r
}

// false
