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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type _roleTemplateScript struct {
	v *types.RoleTemplateScript
}

func NewRoleTemplateScript() *_roleTemplateScript {

	return &_roleTemplateScript{v: types.NewRoleTemplateScript()}

}

func (s *_roleTemplateScript) Id(id string) *_roleTemplateScript {

	s.v.Id = &id

	return s
}

func (s *_roleTemplateScript) Lang(lang scriptlanguage.ScriptLanguage) *_roleTemplateScript {

	s.v.Lang = &lang
	return s
}

func (s *_roleTemplateScript) Options(options map[string]string) *_roleTemplateScript {

	s.v.Options = options
	return s
}

func (s *_roleTemplateScript) AddOption(key string, value string) *_roleTemplateScript {

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

func (s *_roleTemplateScript) Params(params map[string]json.RawMessage) *_roleTemplateScript {

	s.v.Params = params
	return s
}

func (s *_roleTemplateScript) AddParam(key string, value json.RawMessage) *_roleTemplateScript {

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

func (s *_roleTemplateScript) Source(roletemplateinlinequery types.RoleTemplateInlineQueryVariant) *_roleTemplateScript {

	s.v.Source = *roletemplateinlinequery.RoleTemplateInlineQueryCaster()

	return s
}

func (s *_roleTemplateScript) RoleTemplateScriptCaster() *types.RoleTemplateScript {
	return s.v
}
