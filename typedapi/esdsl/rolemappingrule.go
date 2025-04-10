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
)

type _roleMappingRule struct {
	v *types.RoleMappingRule
}

func NewRoleMappingRule() *_roleMappingRule {
	return &_roleMappingRule{v: types.NewRoleMappingRule()}
}

// AdditionalRoleMappingRuleProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_roleMappingRule) AdditionalRoleMappingRuleProperty(key string, value json.RawMessage) *_roleMappingRule {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalRoleMappingRuleProperty = tmp
	return s
}

func (s *_roleMappingRule) All(alls ...types.RoleMappingRuleVariant) *_roleMappingRule {

	for _, v := range alls {

		s.v.All = append(s.v.All, *v.RoleMappingRuleCaster())

	}
	return s
}

func (s *_roleMappingRule) Any(anies ...types.RoleMappingRuleVariant) *_roleMappingRule {

	for _, v := range anies {

		s.v.Any = append(s.v.Any, *v.RoleMappingRuleCaster())

	}
	return s
}

func (s *_roleMappingRule) Except(except types.RoleMappingRuleVariant) *_roleMappingRule {

	s.v.Except = except.RoleMappingRuleCaster()

	return s
}

func (s *_roleMappingRule) RoleMappingRuleCaster() *types.RoleMappingRule {
	return s.v
}
