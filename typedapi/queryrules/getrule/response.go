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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package getrule

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/queryruletype"
)

// Response holds the response body struct for the package getrule
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/query_rules/get_rule/QueryRuleGetResponse.ts#L22-L25
type Response struct {

	// Actions The actions to take when the rule is matched.
	// The format of this action depends on the rule type.
	Actions types.QueryRuleActions `json:"actions"`
	// Criteria The criteria that must be met for the rule to be applied.
	// If multiple criteria are specified for a rule, all criteria must be met for
	// the rule to be applied.
	Criteria []types.QueryRuleCriteria `json:"criteria"`
	Priority *int                      `json:"priority,omitempty"`
	// RuleId A unique identifier for the rule.
	RuleId string `json:"rule_id"`
	// Type The type of rule.
	// `pinned` will identify and pin specific documents to the top of search
	// results.
	// `exclude` will exclude specific documents from search results.
	Type queryruletype.QueryRuleType `json:"type"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
