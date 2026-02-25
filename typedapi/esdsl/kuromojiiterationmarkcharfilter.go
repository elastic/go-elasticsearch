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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _kuromojiIterationMarkCharFilter struct {
	v *types.KuromojiIterationMarkCharFilter
}

func NewKuromojiIterationMarkCharFilter(normalizekana bool, normalizekanji bool) *_kuromojiIterationMarkCharFilter {

	tmp := &_kuromojiIterationMarkCharFilter{v: types.NewKuromojiIterationMarkCharFilter()}

	tmp.NormalizeKana(normalizekana)

	tmp.NormalizeKanji(normalizekanji)

	return tmp

}

func (s *_kuromojiIterationMarkCharFilter) NormalizeKana(normalizekana bool) *_kuromojiIterationMarkCharFilter {

	s.v.NormalizeKana = normalizekana

	return s
}

func (s *_kuromojiIterationMarkCharFilter) NormalizeKanji(normalizekanji bool) *_kuromojiIterationMarkCharFilter {

	s.v.NormalizeKanji = normalizekanji

	return s
}

func (s *_kuromojiIterationMarkCharFilter) Version(versionstring string) *_kuromojiIterationMarkCharFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_kuromojiIterationMarkCharFilter) KuromojiIterationMarkCharFilterCaster() *types.KuromojiIterationMarkCharFilter {
	return s.v
}
