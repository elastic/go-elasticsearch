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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type _script struct {
	v *types.Script
}

// Script used to return matching documents.
// This script must return a boolean value: `true` or `false`.
func NewScript() *_script {

	return &_script{v: types.NewScript()}

}

func (s *_script) Id(id string) *_script {

	s.v.Id = &id

	return s
}

func (s *_script) Lang(lang scriptlanguage.ScriptLanguage) *_script {

	s.v.Lang = &lang
	return s
}

func (s *_script) Options(options map[string]string) *_script {

	s.v.Options = options
	return s
}

func (s *_script) AddOption(key string, value string) *_script {

	var tmp map[string]string
	if s.v.Options == nil {
		s.v.Options = make(map[string]string)
	} else {
		tmp = s.v.Options
	}

	tmp[key] = value

	s.v.Options = tmp
	return s
}

func (s *_script) Params(params map[string]json.RawMessage) *_script {

	s.v.Params = params
	return s
}

func (s *_script) AddParam(key string, value json.RawMessage) *_script {

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

func (s *_script) Source(scriptsource types.ScriptSourceVariant) *_script {

	s.v.Source = *scriptsource.ScriptSourceCaster()

	return s
}

func (s *_script) IntervalsFilterCaster() *types.IntervalsFilter {
	container := types.NewIntervalsFilter()

	container.Script = s.v

	return container
}

func (s *_script) ScriptCaster() *types.Script {
	return s.v
}
