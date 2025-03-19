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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _queryRuleActions struct {
	v *types.QueryRuleActions
}

func NewQueryRuleActions() *_queryRuleActions {

	return &_queryRuleActions{v: types.NewQueryRuleActions()}

}

// The documents to apply the rule to.
// Only one of `ids` or `docs` may be specified and at least one must be
// specified.
// There is a maximum value of 100 documents in a rule.
// You can specify the following attributes for each document:
//
// * `_index`: The index of the document to pin.
// * `_id`: The unique document ID.
func (s *_queryRuleActions) Docs(docs ...types.PinnedDocVariant) *_queryRuleActions {

	for _, v := range docs {

		s.v.Docs = append(s.v.Docs, *v.PinnedDocCaster())

	}
	return s
}

// The unique document IDs of the documents to apply the rule to.
// Only one of `ids` or `docs` may be specified and at least one must be
// specified.
func (s *_queryRuleActions) Ids(ids ...string) *_queryRuleActions {

	for _, v := range ids {

		s.v.Ids = append(s.v.Ids, v)

	}
	return s
}

func (s *_queryRuleActions) QueryRuleActionsCaster() *types.QueryRuleActions {
	return s.v
}
