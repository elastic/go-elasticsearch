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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/keeptypesmode"
)

type _keepTypesTokenFilter struct {
	v *types.KeepTypesTokenFilter
}

func NewKeepTypesTokenFilter() *_keepTypesTokenFilter {

	return &_keepTypesTokenFilter{v: types.NewKeepTypesTokenFilter()}

}

func (s *_keepTypesTokenFilter) Mode(mode keeptypesmode.KeepTypesMode) *_keepTypesTokenFilter {

	s.v.Mode = &mode
	return s
}

func (s *_keepTypesTokenFilter) Types(types ...string) *_keepTypesTokenFilter {

	for _, v := range types {

		s.v.Types = append(s.v.Types, v)

	}
	return s
}

func (s *_keepTypesTokenFilter) Version(versionstring string) *_keepTypesTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_keepTypesTokenFilter) KeepTypesTokenFilterCaster() *types.KeepTypesTokenFilter {
	return s.v
}
