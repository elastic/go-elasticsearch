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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/queryrulecriteriatype"
)

// QueryRuleCriteria type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/query_rules/_types/QueryRuleset.ts#L65-L93
type QueryRuleCriteria struct {
	// Metadata The metadata field to match against.
	// This metadata will be used to match against `match_criteria` sent in the
	// rule.
	// It is required for all criteria types except `always`.
	Metadata *string `json:"metadata,omitempty"`
	// Type The type of criteria. The following criteria types are supported:
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
	Type queryrulecriteriatype.QueryRuleCriteriaType `json:"type"`
	// Values The values to match against the `metadata` field.
	// Only one value must match for the criteria to be met.
	// It is required for all criteria types except `always`.
	Values []json.RawMessage `json:"values,omitempty"`
}

func (s *QueryRuleCriteria) UnmarshalJSON(data []byte) error {

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

		case "metadata":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Metadata = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "values":
			if err := dec.Decode(&s.Values); err != nil {
				return fmt.Errorf("%s | %w", "Values", err)
			}

		}
	}
	return nil
}

// NewQueryRuleCriteria returns a QueryRuleCriteria.
func NewQueryRuleCriteria() *QueryRuleCriteria {
	r := &QueryRuleCriteria{}

	return r
}
