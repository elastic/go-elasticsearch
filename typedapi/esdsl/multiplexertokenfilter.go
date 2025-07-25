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

type _multiplexerTokenFilter struct {
	v *types.MultiplexerTokenFilter
}

func NewMultiplexerTokenFilter() *_multiplexerTokenFilter {

	return &_multiplexerTokenFilter{v: types.NewMultiplexerTokenFilter()}

}

func (s *_multiplexerTokenFilter) Filters(filters ...string) *_multiplexerTokenFilter {

	for _, v := range filters {

		s.v.Filters = append(s.v.Filters, v)

	}
	return s
}

func (s *_multiplexerTokenFilter) PreserveOriginal(stringifiedboolean types.StringifiedbooleanVariant) *_multiplexerTokenFilter {

	s.v.PreserveOriginal = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_multiplexerTokenFilter) Version(versionstring string) *_multiplexerTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_multiplexerTokenFilter) MultiplexerTokenFilterCaster() *types.MultiplexerTokenFilter {
	return s.v
}
