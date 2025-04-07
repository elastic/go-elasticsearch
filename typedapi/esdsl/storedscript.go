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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type _storedScript struct {
	v *types.StoredScript
}

func NewStoredScript(lang scriptlanguage.ScriptLanguage) *_storedScript {

	tmp := &_storedScript{v: types.NewStoredScript()}

	tmp.Lang(lang)

	return tmp

}

func (s *_storedScript) Lang(lang scriptlanguage.ScriptLanguage) *_storedScript {

	s.v.Lang = lang
	return s
}

func (s *_storedScript) Options(options map[string]string) *_storedScript {

	s.v.Options = options
	return s
}

func (s *_storedScript) AddOption(key string, value string) *_storedScript {

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

func (s *_storedScript) Source(scriptsource types.ScriptSourceVariant) *_storedScript {

	s.v.Source = *scriptsource.ScriptSourceCaster()

	return s
}

func (s *_storedScript) StoredScriptCaster() *types.StoredScript {
	return s.v
}
