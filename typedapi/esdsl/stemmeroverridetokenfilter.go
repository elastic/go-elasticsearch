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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _stemmerOverrideTokenFilter struct {
	v *types.StemmerOverrideTokenFilter
}

func NewStemmerOverrideTokenFilter() *_stemmerOverrideTokenFilter {

	return &_stemmerOverrideTokenFilter{v: types.NewStemmerOverrideTokenFilter()}

}

func (s *_stemmerOverrideTokenFilter) Rules(rules ...string) *_stemmerOverrideTokenFilter {

	for _, v := range rules {

		s.v.Rules = append(s.v.Rules, v)

	}
	return s
}

func (s *_stemmerOverrideTokenFilter) RulesPath(rulespath string) *_stemmerOverrideTokenFilter {

	s.v.RulesPath = &rulespath

	return s
}

func (s *_stemmerOverrideTokenFilter) Version(versionstring string) *_stemmerOverrideTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_stemmerOverrideTokenFilter) StemmerOverrideTokenFilterCaster() *types.StemmerOverrideTokenFilter {
	return s.v
}
