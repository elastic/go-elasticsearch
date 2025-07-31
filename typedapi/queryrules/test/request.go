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

package test

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package test
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/query_rules/test/QueryRulesetTestRequest.ts#L24-L57
type Request struct {

	// MatchCriteria The match criteria to apply to rules in the given query ruleset.
	// Match criteria should match the keys defined in the `criteria.metadata` field
	// of the rule.
	MatchCriteria map[string]json.RawMessage `json:"match_criteria"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		MatchCriteria: make(map[string]json.RawMessage, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Test request: %w", err)
	}

	return &req, nil
}
