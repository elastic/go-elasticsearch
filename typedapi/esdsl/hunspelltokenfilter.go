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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hunspellTokenFilter struct {
	v *types.HunspellTokenFilter
}

func NewHunspellTokenFilter(locale string) *_hunspellTokenFilter {

	tmp := &_hunspellTokenFilter{v: types.NewHunspellTokenFilter()}

	tmp.Locale(locale)

	return tmp

}

func (s *_hunspellTokenFilter) Dedup(dedup bool) *_hunspellTokenFilter {

	s.v.Dedup = &dedup

	return s
}

func (s *_hunspellTokenFilter) Dictionary(dictionary string) *_hunspellTokenFilter {

	s.v.Dictionary = &dictionary

	return s
}

func (s *_hunspellTokenFilter) Locale(locale string) *_hunspellTokenFilter {

	s.v.Locale = locale

	return s
}

func (s *_hunspellTokenFilter) LongestOnly(longestonly bool) *_hunspellTokenFilter {

	s.v.LongestOnly = &longestonly

	return s
}

func (s *_hunspellTokenFilter) Version(versionstring string) *_hunspellTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_hunspellTokenFilter) HunspellTokenFilterCaster() *types.HunspellTokenFilter {
	return s.v
}
