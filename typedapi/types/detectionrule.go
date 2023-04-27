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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ruleaction"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// DetectionRule type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ml/_types/Rule.ts#L25-L39
type DetectionRule struct {
	// Actions The set of actions to be triggered when the rule applies. If more than one
	// action is specified the effects of all actions are combined.
	Actions []ruleaction.RuleAction `json:"actions,omitempty"`
	// Conditions An array of numeric conditions when the rule applies. A rule must either have
	// a non-empty scope or at least one condition. Multiple conditions are combined
	// together with a logical AND.
	Conditions []RuleCondition `json:"conditions,omitempty"`
	// Scope A scope of series where the rule applies. A rule must either have a non-empty
	// scope or at least one condition. By default, the scope includes all series.
	// Scoping is allowed for any of the fields that are also specified in
	// `by_field_name`, `over_field_name`, or `partition_field_name`.
	Scope map[string]FilterRef `json:"scope,omitempty"`
}

func (s *DetectionRule) UnmarshalJSON(data []byte) error {

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

		case "actions":
			if err := dec.Decode(&s.Actions); err != nil {
				return err
			}

		case "conditions":
			if err := dec.Decode(&s.Conditions); err != nil {
				return err
			}

		case "scope":
			if s.Scope == nil {
				s.Scope = make(map[string]FilterRef, 0)
			}
			if err := dec.Decode(&s.Scope); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDetectionRule returns a DetectionRule.
func NewDetectionRule() *DetectionRule {
	r := &DetectionRule{
		Scope: make(map[string]FilterRef, 0),
	}

	return r
}
