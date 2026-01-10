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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conditionop"
)

// WatcherCondition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27/specification/watcher/_types/Conditions.ts#L51-L63
type WatcherCondition struct {
	Always       *AlwaysCondition                                  `json:"always,omitempty"`
	ArrayCompare map[string]ArrayCompareCondition                  `json:"array_compare,omitempty"`
	Compare      map[string]map[conditionop.ConditionOp]FieldValue `json:"compare,omitempty"`
	Never        *NeverCondition                                   `json:"never,omitempty"`
	Script       *ScriptCondition                                  `json:"script,omitempty"`
}

// NewWatcherCondition returns a WatcherCondition.
func NewWatcherCondition() *WatcherCondition {
	r := &WatcherCondition{
		ArrayCompare: make(map[string]ArrayCompareCondition),
		Compare:      make(map[string]map[conditionop.ConditionOp]FieldValue),
	}

	return r
}

type WatcherConditionVariant interface {
	WatcherConditionCaster() *WatcherCondition
}

func (s *WatcherCondition) WatcherConditionCaster() *WatcherCondition {
	return s
}
