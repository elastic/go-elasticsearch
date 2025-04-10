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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _ruleRetriever struct {
	v *types.RuleRetriever
}

// A retriever that replaces the functionality of a rule query.
func NewRuleRetriever(matchcriteria json.RawMessage, retriever types.RetrieverContainerVariant) *_ruleRetriever {

	tmp := &_ruleRetriever{v: types.NewRuleRetriever()}

	tmp.MatchCriteria(matchcriteria)

	tmp.Retriever(retriever)

	return tmp

}

func (s *_ruleRetriever) Filter(filters ...types.QueryVariant) *_ruleRetriever {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_ruleRetriever) MatchCriteria(matchcriteria json.RawMessage) *_ruleRetriever {

	s.v.MatchCriteria = matchcriteria

	return s
}

func (s *_ruleRetriever) MinScore(minscore float32) *_ruleRetriever {

	s.v.MinScore = &minscore

	return s
}

func (s *_ruleRetriever) RankWindowSize(rankwindowsize int) *_ruleRetriever {

	s.v.RankWindowSize = &rankwindowsize

	return s
}

func (s *_ruleRetriever) Retriever(retriever types.RetrieverContainerVariant) *_ruleRetriever {

	s.v.Retriever = *retriever.RetrieverContainerCaster()

	return s
}

func (s *_ruleRetriever) RulesetIds(rulesetids ...string) *_ruleRetriever {

	for _, v := range rulesetids {

		s.v.RulesetIds = append(s.v.RulesetIds, v)

	}
	return s
}

func (s *_ruleRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	container := types.NewRetrieverContainer()

	container.Rule = s.v

	return container
}

func (s *_ruleRetriever) RuleRetrieverCaster() *types.RuleRetriever {
	return s.v
}
