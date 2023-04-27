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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conditionop"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// WatcherCondition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/watcher/_types/Conditions.ts#L47-L59
type WatcherCondition struct {
	Always       *AlwaysCondition                                  `json:"always,omitempty"`
	ArrayCompare map[string]ArrayCompareCondition                  `json:"array_compare,omitempty"`
	Compare      map[string]map[conditionop.ConditionOp]FieldValue `json:"compare,omitempty"`
	Never        *NeverCondition                                   `json:"never,omitempty"`
	Script       *ScriptCondition                                  `json:"script,omitempty"`
}

func (s *WatcherCondition) UnmarshalJSON(data []byte) error {

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

		case "always":
			if err := dec.Decode(&s.Always); err != nil {
				return err
			}

		case "array_compare":
			if s.ArrayCompare == nil {
				s.ArrayCompare = make(map[string]ArrayCompareCondition, 0)
			}
			if err := dec.Decode(&s.ArrayCompare); err != nil {
				return err
			}

		case "compare":
			if s.Compare == nil {
				s.Compare = make(map[string]map[conditionop.ConditionOp]FieldValue, 0)
			}
			if err := dec.Decode(&s.Compare); err != nil {
				return err
			}

		case "never":
			if err := dec.Decode(&s.Never); err != nil {
				return err
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewWatcherCondition returns a WatcherCondition.
func NewWatcherCondition() *WatcherCondition {
	r := &WatcherCondition{
		ArrayCompare: make(map[string]ArrayCompareCondition, 0),
		Compare:      make(map[string]map[conditionop.ConditionOp]FieldValue, 0),
	}

	return r
}
