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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _uniqueTokenFilter struct {
	v *types.UniqueTokenFilter
}

func NewUniqueTokenFilter() *_uniqueTokenFilter {

	return &_uniqueTokenFilter{v: types.NewUniqueTokenFilter()}

}

func (s *_uniqueTokenFilter) OnlyOnSamePosition(onlyonsameposition bool) *_uniqueTokenFilter {

	s.v.OnlyOnSamePosition = &onlyonsameposition

	return s
}

func (s *_uniqueTokenFilter) Version(versionstring string) *_uniqueTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_uniqueTokenFilter) UniqueTokenFilterCaster() *types.UniqueTokenFilter {
	return s.v
}
