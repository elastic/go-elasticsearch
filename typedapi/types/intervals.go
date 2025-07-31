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
	"encoding/json"
	"fmt"
)

// Intervals type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/query_dsl/fulltext.ts#L83-L112
type Intervals struct {
	AdditionalIntervalsProperty map[string]json.RawMessage `json:"-"`
	// AllOf Returns matches that span a combination of other rules.
	AllOf *IntervalsAllOf `json:"all_of,omitempty"`
	// AnyOf Returns intervals produced by any of its sub-rules.
	AnyOf *IntervalsAnyOf `json:"any_of,omitempty"`
	// Fuzzy Matches analyzed text.
	Fuzzy *IntervalsFuzzy `json:"fuzzy,omitempty"`
	// Match Matches analyzed text.
	Match *IntervalsMatch `json:"match,omitempty"`
	// Prefix Matches terms that start with a specified set of characters.
	Prefix *IntervalsPrefix `json:"prefix,omitempty"`
	Range  *IntervalsRange  `json:"range,omitempty"`
	Regexp *IntervalsRegexp `json:"regexp,omitempty"`
	// Wildcard Matches terms using a wildcard pattern.
	Wildcard *IntervalsWildcard `json:"wildcard,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s Intervals) MarshalJSON() ([]byte, error) {
	type opt Intervals
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalIntervalsProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalIntervalsProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewIntervals returns a Intervals.
func NewIntervals() *Intervals {
	r := &Intervals{
		AdditionalIntervalsProperty: make(map[string]json.RawMessage),
	}

	return r
}
