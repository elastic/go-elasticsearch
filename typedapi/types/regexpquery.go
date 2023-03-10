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

// RegexpQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/query_dsl/term.ts#L102-L114
type RegexpQuery struct {
	Boost                 *float32 `json:"boost,omitempty"`
	CaseInsensitive       *bool    `json:"case_insensitive,omitempty"`
	Flags                 *string  `json:"flags,omitempty"`
	MaxDeterminizedStates *int     `json:"max_determinized_states,omitempty"`
	QueryName_            *string  `json:"_name,omitempty"`
	Rewrite               *string  `json:"rewrite,omitempty"`
	Value                 string   `json:"value"`
}

// NewRegexpQuery returns a RegexpQuery.
func NewRegexpQuery() *RegexpQuery {
	r := &RegexpQuery{}

	return r
}
