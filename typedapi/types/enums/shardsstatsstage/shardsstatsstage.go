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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Package shardsstatsstage
package shardsstatsstage

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/snapshot/_types/SnapshotShardsStatsStage.ts#L20-L31
type ShardsStatsStage struct {
	Name string
}

var (
	DONE = ShardsStatsStage{"DONE"}

	FAILURE = ShardsStatsStage{"FAILURE"}

	FINALIZE = ShardsStatsStage{"FINALIZE"}

	INIT = ShardsStatsStage{"INIT"}

	STARTED = ShardsStatsStage{"STARTED"}
)

func (s ShardsStatsStage) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *ShardsStatsStage) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "DONE":
		*s = DONE
	case "FAILURE":
		*s = FAILURE
	case "FINALIZE":
		*s = FINALIZE
	case "INIT":
		*s = INIT
	case "STARTED":
		*s = STARTED
	default:
		*s = ShardsStatsStage{string(text)}
	}

	return nil
}

func (s ShardsStatsStage) String() string {
	return s.Name
}
