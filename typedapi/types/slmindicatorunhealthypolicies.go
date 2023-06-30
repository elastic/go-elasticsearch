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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// SlmIndicatorUnhealthyPolicies type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_global/health_report/types.ts#L165-L168
type SlmIndicatorUnhealthyPolicies struct {
	Count                       int64            `json:"count"`
	InvocationsSinceLastSuccess map[string]int64 `json:"invocations_since_last_success,omitempty"`
}

func (s *SlmIndicatorUnhealthyPolicies) UnmarshalJSON(data []byte) error {

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

		case "count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "invocations_since_last_success":
			if s.InvocationsSinceLastSuccess == nil {
				s.InvocationsSinceLastSuccess = make(map[string]int64, 0)
			}
			if err := dec.Decode(&s.InvocationsSinceLastSuccess); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSlmIndicatorUnhealthyPolicies returns a SlmIndicatorUnhealthyPolicies.
func NewSlmIndicatorUnhealthyPolicies() *SlmIndicatorUnhealthyPolicies {
	r := &SlmIndicatorUnhealthyPolicies{
		InvocationsSinceLastSuccess: make(map[string]int64, 0),
	}

	return r
}
