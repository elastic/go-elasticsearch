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
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conditionop"
)

// WatcherCondition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/watcher/_types/Conditions.ts#L50-L62
type WatcherCondition struct {
	AdditionalWatcherConditionProperty map[string]json.RawMessage                        `json:"-"`
	Always                             *AlwaysCondition                                  `json:"always,omitempty"`
	ArrayCompare                       map[string]ArrayCompareCondition                  `json:"array_compare,omitempty"`
	Compare                            map[string]map[conditionop.ConditionOp]FieldValue `json:"compare,omitempty"`
	Never                              *NeverCondition                                   `json:"never,omitempty"`
	Script                             *ScriptCondition                                  `json:"script,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s WatcherCondition) MarshalJSON() ([]byte, error) {
	type opt WatcherCondition
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
	for key, value := range s.AdditionalWatcherConditionProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalWatcherConditionProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewWatcherCondition returns a WatcherCondition.
func NewWatcherCondition() *WatcherCondition {
	r := &WatcherCondition{
		AdditionalWatcherConditionProperty: make(map[string]json.RawMessage),
		ArrayCompare:                       make(map[string]ArrayCompareCondition),
		Compare:                            make(map[string]map[conditionop.ConditionOp]FieldValue),
	}

	return r
}
