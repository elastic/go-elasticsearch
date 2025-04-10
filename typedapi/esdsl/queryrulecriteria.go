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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/queryrulecriteriatype"
)

type _queryRuleCriteria struct {
	v *types.QueryRuleCriteria
}

func NewQueryRuleCriteria(type_ queryrulecriteriatype.QueryRuleCriteriaType) *_queryRuleCriteria {

	tmp := &_queryRuleCriteria{v: types.NewQueryRuleCriteria()}

	tmp.Type(type_)

	return tmp

}

func (s *_queryRuleCriteria) Metadata(metadata string) *_queryRuleCriteria {

	s.v.Metadata = &metadata

	return s
}

func (s *_queryRuleCriteria) Type(type_ queryrulecriteriatype.QueryRuleCriteriaType) *_queryRuleCriteria {

	s.v.Type = type_
	return s
}

func (s *_queryRuleCriteria) Values(values ...json.RawMessage) *_queryRuleCriteria {

	for _, v := range values {

		s.v.Values = append(s.v.Values, v)

	}
	return s
}

func (s *_queryRuleCriteria) QueryRuleCriteriaCaster() *types.QueryRuleCriteria {
	return s.v
}
