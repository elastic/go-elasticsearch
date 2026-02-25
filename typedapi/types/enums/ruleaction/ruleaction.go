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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package ruleaction
package ruleaction

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/ml/_types/Rule.ts#L41-L50
type RuleAction struct {
	Name string
}

var (

	// Skipresult The result will not be created. Unless you also specify `skip_model_update`,
	// the model will be updated as usual with the corresponding series value.
	Skipresult = RuleAction{"skip_result"}

	// Skipmodelupdate The value for that series will not be used to update the model. Unless you
	// also specify `skip_result`, the results will be created as usual. This action
	// is suitable when certain values are expected to be consistently anomalous and
	// they affect the model in a way that negatively impacts the rest of the
	// results.
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
