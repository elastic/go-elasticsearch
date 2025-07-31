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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _elisionTokenFilter struct {
	v *types.ElisionTokenFilter
}

func NewElisionTokenFilter() *_elisionTokenFilter {

	return &_elisionTokenFilter{v: types.NewElisionTokenFilter()}

}

func (s *_elisionTokenFilter) Articles(articles ...string) *_elisionTokenFilter {

	for _, v := range articles {

		s.v.Articles = append(s.v.Articles, v)

	}
	return s
}

func (s *_elisionTokenFilter) ArticlesCase(stringifiedboolean types.StringifiedbooleanVariant) *_elisionTokenFilter {

	s.v.ArticlesCase = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_elisionTokenFilter) ArticlesPath(articlespath string) *_elisionTokenFilter {

	s.v.ArticlesPath = &articlespath

	return s
}

func (s *_elisionTokenFilter) Version(versionstring string) *_elisionTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_elisionTokenFilter) ElisionTokenFilterCaster() *types.ElisionTokenFilter {
	return s.v
}
