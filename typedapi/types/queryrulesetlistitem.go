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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// QueryRulesetListItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/query_ruleset/list/types.ts#L22-L31
type QueryRulesetListItem struct {
	// RulesCount The number of rules associated with this ruleset
	RulesCount int `json:"rules_count"`
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

		case "rules_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.RulesCount = value
			case float64:
				f := int(v)
				s.RulesCount = f
			}

		case "ruleset_id":
			if err := dec.Decode(&s.RulesetId); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewQueryRulesetListItem returns a QueryRulesetListItem.
func NewQueryRulesetListItem() *QueryRulesetListItem {
	r := &QueryRulesetListItem{}

	return r
}
