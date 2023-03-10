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

package types

import (
	"encoding/json"
)

// BoxPlotAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/aggregations/Aggregate.ts#L697-L713
type BoxPlotAggregate struct {
	Lower         Float64                    `json:"lower"`
	LowerAsString *string                    `json:"lower_as_string,omitempty"`
	Max           Float64                    `json:"max"`
	MaxAsString   *string                    `json:"max_as_string,omitempty"`
	Meta          map[string]json.RawMessage `json:"meta,omitempty"`
	Min           Float64                    `json:"min"`
	MinAsString   *string                    `json:"min_as_string,omitempty"`
	Q1            Float64                    `json:"q1"`
	Q1AsString    *string                    `json:"q1_as_string,omitempty"`
	Q2            Float64                    `json:"q2"`
	Q2AsString    *string                    `json:"q2_as_string,omitempty"`
	Q3            Float64                    `json:"q3"`
	Q3AsString    *string                    `json:"q3_as_string,omitempty"`
	Upper         Float64                    `json:"upper"`
	UpperAsString *string                    `json:"upper_as_string,omitempty"`
}

// NewBoxPlotAggregate returns a BoxPlotAggregate.
func NewBoxPlotAggregate() *BoxPlotAggregate {
	r := &BoxPlotAggregate{}

	return r
}
