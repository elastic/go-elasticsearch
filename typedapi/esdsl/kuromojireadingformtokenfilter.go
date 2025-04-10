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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _kuromojiReadingFormTokenFilter struct {
	v *types.KuromojiReadingFormTokenFilter
}

func NewKuromojiReadingFormTokenFilter(useromaji bool) *_kuromojiReadingFormTokenFilter {

	tmp := &_kuromojiReadingFormTokenFilter{v: types.NewKuromojiReadingFormTokenFilter()}

	tmp.UseRomaji(useromaji)

	return tmp

}

func (s *_kuromojiReadingFormTokenFilter) UseRomaji(useromaji bool) *_kuromojiReadingFormTokenFilter {

	s.v.UseRomaji = useromaji

	return s
}

func (s *_kuromojiReadingFormTokenFilter) Version(versionstring string) *_kuromojiReadingFormTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_kuromojiReadingFormTokenFilter) KuromojiReadingFormTokenFilterCaster() *types.KuromojiReadingFormTokenFilter {
	return s.v
}
