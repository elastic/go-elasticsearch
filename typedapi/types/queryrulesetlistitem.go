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
// https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// QueryRulesetListItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64/specification/query_rules/list_rulesets/types.ts#L23-L42
type QueryRulesetListItem struct {
	// RuleCriteriaTypesCounts A map of criteria type (e.g. exact) to the number of rules of that type
	RuleCriteriaTypesCounts map[string]int `json:"rule_criteria_types_counts"`
	// RuleTotalCount The number of rules associated with this ruleset
	RuleTotalCount int `json:"rule_total_count"`
	// RuleTypeCounts A map of rule type (e.g. pinned) to the number of rules of that type
	RuleTypeCounts map[string]int `json:"rule_type_counts"`
	// RulesetId Ruleset unique identifier
	RulesetId string `json:"ruleset_id"`
}

func (s *QueryRulesetListItem) UnmarshalJSON(data []byte) error {

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

		case "rule_criteria_types_counts":
			if s.RuleCriteriaTypesCounts == nil {
				s.RuleCriteriaTypesCounts = make(map[string]int, 0)
			}
			if err := dec.Decode(&s.RuleCriteriaTypesCounts); err != nil {
				return fmt.Errorf("%s | %w", "RuleCriteriaTypesCounts", err)
			}

		case "rule_total_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RuleTotalCount", err)
				}
				s.RuleTotalCount = value
			case float64:
				f := int(v)
				s.RuleTotalCount = f
			}

		case "rule_type_counts":
			if s.RuleTypeCounts == nil {
				s.RuleTypeCounts = make(map[string]int, 0)
			}
			if err := dec.Decode(&s.RuleTypeCounts); err != nil {
				return fmt.Errorf("%s | %w", "RuleTypeCounts", err)
			}

		case "ruleset_id":
			if err := dec.Decode(&s.RulesetId); err != nil {
				return fmt.Errorf("%s | %w", "RulesetId", err)
			}

		}
	}
	return nil
}

// NewQueryRulesetListItem returns a QueryRulesetListItem.
func NewQueryRulesetListItem() *QueryRulesetListItem {
	r := &QueryRulesetListItem{
		RuleCriteriaTypesCounts: make(map[string]int, 0),
		RuleTypeCounts:          make(map[string]int, 0),
	}

	return r
}
