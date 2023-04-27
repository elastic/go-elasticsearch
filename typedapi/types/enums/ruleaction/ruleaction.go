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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

// Package ruleaction
package ruleaction

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ml/_types/Rule.ts#L41-L50
type RuleAction struct {
	Name string
}

var (
	Skipresult = RuleAction{"skip_result"}

	Skipmodelupdate = RuleAction{"skip_model_update"}
)

func (r RuleAction) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RuleAction) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "skip_result":
		*r = Skipresult
	case "skip_model_update":
		*r = Skipmodelupdate
	default:
		*r = RuleAction{string(text)}
	}

	return nil
}

func (r RuleAction) String() string {
	return r.Name
}
