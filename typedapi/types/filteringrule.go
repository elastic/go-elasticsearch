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
// https://github.com/elastic/elasticsearch-specification/tree/19027dbdd366978ccae41842a040a636730e7c10

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/filteringpolicy"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/filteringrulerule"
)

// FilteringRule type.
//
// https://github.com/elastic/elasticsearch-specification/blob/19027dbdd366978ccae41842a040a636730e7c10/specification/connector/_types/Connector.ts#L170-L179
type FilteringRule struct {
	CreatedAt DateTime                            `json:"created_at,omitempty"`
	Field     string                              `json:"field"`
	Id        string                              `json:"id"`
	Order     int                                 `json:"order"`
	Policy    filteringpolicy.FilteringPolicy     `json:"policy"`
	Rule      filteringrulerule.FilteringRuleRule `json:"rule"`
	UpdatedAt DateTime                            `json:"updated_at,omitempty"`
	Value     string                              `json:"value"`
}

func (s *FilteringRule) UnmarshalJSON(data []byte) error {

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

		case "created_at":
			if err := dec.Decode(&s.CreatedAt); err != nil {
				return fmt.Errorf("%s | %w", "CreatedAt", err)
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "order":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Order", err)
				}
				s.Order = value
			case float64:
				f := int(v)
				s.Order = f
			}

		case "policy":
			if err := dec.Decode(&s.Policy); err != nil {
				return fmt.Errorf("%s | %w", "Policy", err)
			}

		case "rule":
			if err := dec.Decode(&s.Rule); err != nil {
				return fmt.Errorf("%s | %w", "Rule", err)
			}

		case "updated_at":
			if err := dec.Decode(&s.UpdatedAt); err != nil {
				return fmt.Errorf("%s | %w", "UpdatedAt", err)
			}

		case "value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Value", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Value = o

		}
	}
	return nil
}

// NewFilteringRule returns a FilteringRule.
func NewFilteringRule() *FilteringRule {
	r := &FilteringRule{}

	return r
}
