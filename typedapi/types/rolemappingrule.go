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

package types

import (
	"encoding/json"
	"fmt"
)

// RoleMappingRule type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/security/_types/RoleMappingRule.ts#L23-L31
type RoleMappingRule struct {
	AdditionalRoleMappingRuleProperty map[string]json.RawMessage `json:"-"`
	All                               []RoleMappingRule          `json:"all,omitempty"`
	Any                               []RoleMappingRule          `json:"any,omitempty"`
	Except                            *RoleMappingRule           `json:"except,omitempty"`
	Field                             map[string][]FieldValue    `json:"field,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s RoleMappingRule) MarshalJSON() ([]byte, error) {
	type opt RoleMappingRule
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalRoleMappingRuleProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalRoleMappingRuleProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewRoleMappingRule returns a RoleMappingRule.
func NewRoleMappingRule() *RoleMappingRule {
	r := &RoleMappingRule{
		AdditionalRoleMappingRuleProperty: make(map[string]json.RawMessage),
		Field:                             make(map[string][]FieldValue),
	}

	return r
}

// true

type RoleMappingRuleVariant interface {
	RoleMappingRuleCaster() *RoleMappingRule
}

func (s *RoleMappingRule) RoleMappingRuleCaster() *RoleMappingRule {
	return s
}
