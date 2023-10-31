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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/appliesto"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conditionoperator"
)

// RuleCondition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Rule.ts#L52-L65
type RuleCondition struct {
	// AppliesTo Specifies the result property to which the condition applies. If your
	// detector uses `lat_long`, `metric`, `rare`, or `freq_rare` functions, you can
	// only specify conditions that apply to time.
	AppliesTo appliesto.AppliesTo `json:"applies_to"`
	// Operator Specifies the condition operator. The available options are greater than,
	// greater than or equals, less than, and less than or equals.
	Operator conditionoperator.ConditionOperator `json:"operator"`
	// Value The value that is compared against the `applies_to` field using the operator.
	Value Float64 `json:"value"`
}

func (s *RuleCondition) UnmarshalJSON(data []byte) error {

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

		case "applies_to":
			if err := dec.Decode(&s.AppliesTo); err != nil {
				return err
			}

		case "operator":
			if err := dec.Decode(&s.Operator); err != nil {
				return err
			}

		case "value":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Value = f
			case float64:
				f := Float64(v)
				s.Value = f
			}

		}
	}
	return nil
}

// NewRuleCondition returns a RuleCondition.
func NewRuleCondition() *RuleCondition {
	r := &RuleCondition{}

	return r
}
