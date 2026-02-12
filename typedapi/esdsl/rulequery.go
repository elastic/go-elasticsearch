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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _ruleQuery struct {
	v *types.RuleQuery
}

func NewRuleQuery(matchcriteria json.RawMessage, organic types.QueryVariant) *_ruleQuery {

	tmp := &_ruleQuery{v: types.NewRuleQuery()}

	tmp.MatchCriteria(matchcriteria)

	tmp.Organic(organic)

	return tmp

}

func (s *_ruleQuery) MatchCriteria(matchcriteria json.RawMessage) *_ruleQuery {

	s.v.MatchCriteria = matchcriteria

	return s
}

func (s *_ruleQuery) Organic(organic types.QueryVariant) *_ruleQuery {

	s.v.Organic = *organic.QueryCaster()

	return s
}

func (s *_ruleQuery) RulesetId(rulesetid string) *_ruleQuery {

	s.v.RulesetId = &rulesetid

	return s
}

func (s *_ruleQuery) RulesetIds(rulesetids ...string) *_ruleQuery {

	s.v.RulesetIds = make([]string, len(rulesetids))
	s.v.RulesetIds = rulesetids

	return s
}

func (s *_ruleQuery) Boost(boost float32) *_ruleQuery {

	s.v.Boost = &boost

	return s
}

func (s *_ruleQuery) QueryName_(queryname_ string) *_ruleQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_ruleQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Rule = s.v

	return container
}

func (s *_ruleQuery) RuleQueryCaster() *types.RuleQuery {
	return s.v
}
