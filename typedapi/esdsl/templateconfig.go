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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _templateConfig struct {
	v *types.TemplateConfig
}

func NewTemplateConfig() *_templateConfig {

	return &_templateConfig{v: types.NewTemplateConfig()}

}

// If `true`, returns detailed information about score calculation as part of
// each hit.
func (s *_templateConfig) Explain(explain bool) *_templateConfig {

	s.v.Explain = &explain

	return s
}

// The ID of the search template to use. If no `source` is specified,
// this parameter is required.
func (s *_templateConfig) Id(id string) *_templateConfig {

	s.v.Id = &id

	return s
}

// Key-value pairs used to replace Mustache variables in the template.
// The key is the variable name.
// The value is the variable value.
func (s *_templateConfig) Params(params map[string]json.RawMessage) *_templateConfig {

	s.v.Params = params
	return s
}

func (s *_templateConfig) AddParam(key string, value json.RawMessage) *_templateConfig {

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

// If `true`, the query execution is profiled.
func (s *_templateConfig) Profile(profile bool) *_templateConfig {

	s.v.Profile = &profile

	return s
}

// An inline search template. Supports the same parameters as the search API's
// request body. It also supports Mustache variables. If no `id` is specified,
// this
// parameter is required.
func (s *_templateConfig) Source(source string) *_templateConfig {

	s.v.Source = &source

	return s
}

func (s *_templateConfig) TemplateConfigCaster() *types.TemplateConfig {
	return s.v
}
