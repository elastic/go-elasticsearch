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
// https://github.com/elastic/elasticsearch-specification/tree/f6a370d0fba975752c644fc730f7c45610e28f36

// Package shardroutingstate
package shardroutingstate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/f6a370d0fba975752c644fc730f7c45610e28f36/specification/indices/stats/types.ts#L169-L174
type ShardRoutingState struct {
	Name string
}

var (
	UNASSIGNED = ShardRoutingState{"UNASSIGNED"}

	INITIALIZING = ShardRoutingState{"INITIALIZING"}

	STARTED = ShardRoutingState{"STARTED"}

	RELOCATING = ShardRoutingState{"RELOCATING"}
)

func (s ShardRoutingState) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *ShardRoutingState) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "unassigned":
		*s = UNASSIGNED
	case "initializing":
		*s = INITIALIZING
	case "started":
		*s = STARTED
	case "relocating":
		*s = RELOCATING
	default:
		*s = ShardRoutingState{string(text)}
	}

	return nil
}

func (s ShardRoutingState) String() string {
	return s.Name
}
