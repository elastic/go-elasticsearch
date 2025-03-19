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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
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

// Query to filter the documents that can match.
func (s *_ruleRetriever) Filter(filters ...types.QueryVariant) *_ruleRetriever {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

// The match criteria that will determine if a rule in the provided rulesets
// should be applied.
func (s *_ruleRetriever) MatchCriteria(matchcriteria json.RawMessage) *_ruleRetriever {

	s.v.MatchCriteria = matchcriteria

	return s
}

// Minimum _score for matching documents. Documents with a lower _score are not
// included in the top documents.
func (s *_ruleRetriever) MinScore(minscore float32) *_ruleRetriever {

	s.v.MinScore = &minscore

	return s
}

// This value determines the size of the individual result set.
func (s *_ruleRetriever) RankWindowSize(rankwindowsize int) *_ruleRetriever {

	s.v.RankWindowSize = &rankwindowsize

	return s
}

// The retriever whose results rules should be applied to.
func (s *_ruleRetriever) Retriever(retriever types.RetrieverContainerVariant) *_ruleRetriever {

	s.v.Retriever = *retriever.RetrieverContainerCaster()

	return s
}

// The ruleset IDs containing the rules this retriever is evaluating against.
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
