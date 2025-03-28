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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/queryrulecriteriatype"
)

type _queryRuleCriteria struct {
	v *types.QueryRuleCriteria
}

func NewQueryRuleCriteria(type_ queryrulecriteriatype.QueryRuleCriteriaType) *_queryRuleCriteria {

	tmp := &_queryRuleCriteria{v: types.NewQueryRuleCriteria()}

	tmp.Type(type_)

	return tmp

}

// The metadata field to match against.
// This metadata will be used to match against `match_criteria` sent in the
// rule.
// It is required for all criteria types except `always`.
func (s *_queryRuleCriteria) Metadata(metadata string) *_queryRuleCriteria {

	s.v.Metadata = &metadata

	return s
}

// The type of criteria. The following criteria types are supported:
//
// * `always`: Matches all queries, regardless of input.
// * `contains`: Matches that contain this value anywhere in the field meet the
// criteria defined by the rule. Only applicable for string values.
// * `exact`: Only exact matches meet the criteria defined by the rule.
// Applicable for string or numerical values.
// * `fuzzy`: Exact matches or matches within the allowed Levenshtein Edit
// Distance meet the criteria defined by the rule. Only applicable for string
// values.
// * `gt`: Matches with a value greater than this value meet the criteria
// defined by the rule. Only applicable for numerical values.
// * `gte`: Matches with a value greater than or equal to this value meet the
// criteria defined by the rule. Only applicable for numerical values.
// * `lt`: Matches with a value less than this value meet the criteria defined
// by the rule. Only applicable for numerical values.
// * `lte`: Matches with a value less than or equal to this value meet the
// criteria defined by the rule. Only applicable for numerical values.
// * `prefix`: Matches that start with this value meet the criteria defined by
// the rule. Only applicable for string values.
// * `suffix`: Matches that end with this value meet the criteria defined by the
// rule. Only applicable for string values.
func (s *_queryRuleCriteria) Type(type_ queryrulecriteriatype.QueryRuleCriteriaType) *_queryRuleCriteria {

	s.v.Type = type_
	return s
}

// The values to match against the `metadata` field.
// Only one value must match for the criteria to be met.
// It is required for all criteria types except `always`.
func (s *_queryRuleCriteria) Values(values ...json.RawMessage) *_queryRuleCriteria {

	for _, v := range values {

		s.v.Values = append(s.v.Values, v)

	}
	return s
}

func (s *_queryRuleCriteria) QueryRuleCriteriaCaster() *types.QueryRuleCriteria {
	return s.v
}
