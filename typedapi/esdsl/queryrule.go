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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/queryruletype"
)

type _queryRule struct {
	v *types.QueryRule
}

func NewQueryRule(actions types.QueryRuleActionsVariant, type_ queryruletype.QueryRuleType) *_queryRule {

	tmp := &_queryRule{v: types.NewQueryRule()}

	tmp.Actions(actions)

	tmp.Type(type_)

	return tmp

}

// The actions to take when the rule is matched.
// The format of this action depends on the rule type.
func (s *_queryRule) Actions(actions types.QueryRuleActionsVariant) *_queryRule {

	s.v.Actions = *actions.QueryRuleActionsCaster()

	return s
}

// The criteria that must be met for the rule to be applied.
// If multiple criteria are specified for a rule, all criteria must be met for
// the rule to be applied.
func (s *_queryRule) Criteria(criteria ...types.QueryRuleCriteriaVariant) *_queryRule {

	s.v.Criteria = make([]types.QueryRuleCriteria, len(criteria))
	for i, v := range criteria {
		s.v.Criteria[i] = *v.QueryRuleCriteriaCaster()
	}

	return s
}

func (s *_queryRule) Priority(priority int) *_queryRule {

	s.v.Priority = &priority

	return s
}

// A unique identifier for the rule.
func (s *_queryRule) RuleId(id string) *_queryRule {

	s.v.RuleId = id

	return s
}

// The type of rule.
// `pinned` will identify and pin specific documents to the top of search
// results.
// `exclude` will exclude specific documents from search results.
func (s *_queryRule) Type(type_ queryruletype.QueryRuleType) *_queryRule {

	s.v.Type = type_
	return s
}

func (s *_queryRule) QueryRuleCaster() *types.QueryRule {
	return s.v
}
