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

package types

// Intervals type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/fulltext.ts#L83-L110
type Intervals struct {
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
	// Wildcard Matches terms using a wildcard pattern.
	Wildcard *IntervalsWildcard `json:"wildcard,omitempty"`
}

// NewIntervals returns a Intervals.
func NewIntervals() *Intervals {
	r := &Intervals{}

	return r
}
