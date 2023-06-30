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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Package cardinalityexecutionmode
package cardinalityexecutionmode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/aggregations/metric.ts#L54-L60
type CardinalityExecutionMode struct {
	Name string
}

var (
	Globalordinals = CardinalityExecutionMode{"global_ordinals"}

	Segmentordinals = CardinalityExecutionMode{"segment_ordinals"}

	Direct = CardinalityExecutionMode{"direct"}

	Savememoryheuristic = CardinalityExecutionMode{"save_memory_heuristic"}

	Savetimeheuristic = CardinalityExecutionMode{"save_time_heuristic"}
)

func (c CardinalityExecutionMode) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CardinalityExecutionMode) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "global_ordinals":
		*c = Globalordinals
	case "segment_ordinals":
		*c = Segmentordinals
	case "direct":
		*c = Direct
	case "save_memory_heuristic":
		*c = Savememoryheuristic
	case "save_time_heuristic":
		*c = Savetimeheuristic
	default:
		*c = CardinalityExecutionMode{string(text)}
	}

	return nil
}

func (c CardinalityExecutionMode) String() string {
	return c.Name
}
