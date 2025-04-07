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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
)

type _scriptCondition struct {
	v *types.ScriptCondition
}

func NewScriptCondition() *_scriptCondition {

	return &_scriptCondition{v: types.NewScriptCondition()}

}

func (s *_scriptCondition) Id(id string) *_scriptCondition {

	s.v.Id = &id

	return s
}

func (s *_scriptCondition) Lang(lang scriptlanguage.ScriptLanguage) *_scriptCondition {

	s.v.Lang = &lang
	return s
}

func (s *_scriptCondition) Params(params map[string]json.RawMessage) *_scriptCondition {

	s.v.Params = params
	return s
}

func (s *_scriptCondition) AddParam(key string, value json.RawMessage) *_scriptCondition {

	var tmp map[string]json.RawMessage
	if s.v.Params == nil {
		s.v.Params = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Params
	}

	tmp[key] = value

	s.v.Params = tmp
	return s
}

func (s *_scriptCondition) Source(scriptsource types.ScriptSourceVariant) *_scriptCondition {

	s.v.Source = *scriptsource.ScriptSourceCaster()

	return s
}

func (s *_scriptCondition) WatcherConditionCaster() *types.WatcherCondition {
	container := types.NewWatcherCondition()

	container.Script = s.v

	return container
}

func (s *_scriptCondition) ScriptConditionCaster() *types.ScriptCondition {
	return s.v
}
