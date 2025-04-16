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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _watcherCondition struct {
	v *types.WatcherCondition
}

func NewWatcherCondition() *_watcherCondition {
	return &_watcherCondition{v: types.NewWatcherCondition()}
}

// AdditionalWatcherConditionProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_watcherCondition) AdditionalWatcherConditionProperty(key string, value json.RawMessage) *_watcherCondition {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalWatcherConditionProperty = tmp
	return s
}

func (s *_watcherCondition) Always(always types.AlwaysConditionVariant) *_watcherCondition {

	s.v.Always = always.AlwaysConditionCaster()

	return s
}

// ArrayCompare is a single key dictionnary.
// It will replace the current value on each call.
func (s *_watcherCondition) ArrayCompare(key string, value types.ArrayCompareConditionVariant) *_watcherCondition {

	tmp := make(map[string]types.ArrayCompareCondition)

	tmp[key] = *value.ArrayCompareConditionCaster()

	s.v.ArrayCompare = tmp
	return s
}

func (s *_watcherCondition) Never(never types.NeverConditionVariant) *_watcherCondition {

	s.v.Never = never.NeverConditionCaster()

	return s
}

func (s *_watcherCondition) Script(script types.ScriptConditionVariant) *_watcherCondition {

	s.v.Script = script.ScriptConditionCaster()

	return s
}

func (s *_watcherCondition) WatcherConditionCaster() *types.WatcherCondition {
	return s.v
}
