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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package types

// RoleMappingRule type.
//
// https://github.com/elastic/elasticsearch-specification/blob/aa1459fbdcaf57c653729142b3b6e9982373bb1c/specification/security/_types/RoleMappingRule.ts#L23-L31
type RoleMappingRule struct {
	All    []RoleMappingRule       `json:"all,omitempty"`
	Any    []RoleMappingRule       `json:"any,omitempty"`
	Except *RoleMappingRule        `json:"except,omitempty"`
	Field  map[string][]FieldValue `json:"field,omitempty"`
}

// NewRoleMappingRule returns a RoleMappingRule.
func NewRoleMappingRule() *RoleMappingRule {
	r := &RoleMappingRule{
		Field: make(map[string][]FieldValue),
	}

	return r
}

type RoleMappingRuleVariant interface {
	RoleMappingRuleCaster() *RoleMappingRule
}

func (s *RoleMappingRule) RoleMappingRuleCaster() *RoleMappingRule {
	return s
}
