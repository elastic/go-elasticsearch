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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _scriptTransform struct {
	v *types.ScriptTransform
}

func NewScriptTransform() *_scriptTransform {

	return &_scriptTransform{v: types.NewScriptTransform()}

}

func (s *_scriptTransform) Id(id string) *_scriptTransform {

	s.v.Id = &id

	return s
}

func (s *_scriptTransform) Lang(lang string) *_scriptTransform {

	s.v.Lang = &lang

	return s
}

func (s *_scriptTransform) Params(params map[string]json.RawMessage) *_scriptTransform {

	s.v.Params = params
	return s
}

func (s *_scriptTransform) AddParam(key string, value json.RawMessage) *_scriptTransform {

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

func (s *_scriptTransform) Source(scriptsource types.ScriptSourceVariant) *_scriptTransform {

	s.v.Source = *scriptsource.ScriptSourceCaster()

	return s
}

func (s *_scriptTransform) TransformContainerCaster() *types.TransformContainer {
	container := types.NewTransformContainer()

	container.Script = s.v

	return container
}

func (s *_scriptTransform) ScriptTransformCaster() *types.ScriptTransform {
	return s.v
}
