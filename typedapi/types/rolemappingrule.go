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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// RoleMappingRule type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/security/_types/RoleMappingRule.ts#L23-L31
type RoleMappingRule struct {
	All    []RoleMappingRule       `json:"all,omitempty"`
	Any    []RoleMappingRule       `json:"any,omitempty"`
	Except *RoleMappingRule        `json:"except,omitempty"`
	Field  map[string][]FieldValue `json:"field,omitempty"`
}

func (s *RoleMappingRule) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "all":
			if err := dec.Decode(&s.All); err != nil {
				return fmt.Errorf("%s | %w", "All", err)
			}

		case "any":
			if err := dec.Decode(&s.Any); err != nil {
				return fmt.Errorf("%s | %w", "Any", err)
			}

		case "except":
			if err := dec.Decode(&s.Except); err != nil {
				return fmt.Errorf("%s | %w", "Except", err)
			}

		case "field":
			if s.Field == nil {
				s.Field = make(map[string][]FieldValue, 0)
			}
			rawMsg := make(map[string]json.RawMessage, 0)
			dec.Decode(&rawMsg)
			for key, value := range rawMsg {
				v := bytes.TrimSpace(value)
				if len(v) > 0 && v[0] == '[' {
					var o []FieldValue
					if err := json.NewDecoder(bytes.NewReader(v)).Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Field", err)
					}
					s.Field[key] = o
					continue
				}

				var o FieldValue
				if err := json.NewDecoder(bytes.NewReader(v)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Field", err)
				}
				s.Field[key] = append(s.Field[key], o)
			}

		}
	}
	return nil
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
