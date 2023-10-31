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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Package sampleraggregationexecutionhint
package sampleraggregationexecutionhint

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/bucket.ts#L343-L356
type SamplerAggregationExecutionHint struct {
	Name string
}

var (
	Map = SamplerAggregationExecutionHint{"map"}

	Globalordinals = SamplerAggregationExecutionHint{"global_ordinals"}

	Byteshash = SamplerAggregationExecutionHint{"bytes_hash"}
)

func (s SamplerAggregationExecutionHint) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SamplerAggregationExecutionHint) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "map":
		*s = Map
	case "global_ordinals":
		*s = Globalordinals
	case "bytes_hash":
		*s = Byteshash
	default:
		*s = SamplerAggregationExecutionHint{string(text)}
	}

	return nil
}

func (s SamplerAggregationExecutionHint) String() string {
	return s.Name
}
