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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationtype"
)

type _icuNormalizationTokenFilter struct {
	v *types.IcuNormalizationTokenFilter
}

func NewIcuNormalizationTokenFilter(name icunormalizationtype.IcuNormalizationType) *_icuNormalizationTokenFilter {

	tmp := &_icuNormalizationTokenFilter{v: types.NewIcuNormalizationTokenFilter()}

	tmp.Name(name)

	return tmp

}

func (s *_icuNormalizationTokenFilter) Name(name icunormalizationtype.IcuNormalizationType) *_icuNormalizationTokenFilter {

	s.v.Name = name
	return s
}

func (s *_icuNormalizationTokenFilter) Version(versionstring string) *_icuNormalizationTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_icuNormalizationTokenFilter) IcuNormalizationTokenFilterCaster() *types.IcuNormalizationTokenFilter {
	return s.v
}
