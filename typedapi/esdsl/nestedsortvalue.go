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

type _nestedSortValue struct {
	v *types.NestedSortValue
}

func NewNestedSortValue() *_nestedSortValue {

	return &_nestedSortValue{v: types.NewNestedSortValue()}

}

func (s *_nestedSortValue) Filter(filter types.QueryVariant) *_nestedSortValue {

	s.v.Filter = filter.QueryCaster()

	return s
}

func (s *_nestedSortValue) MaxChildren(maxchildren int) *_nestedSortValue {

	s.v.MaxChildren = &maxchildren

	return s
}

func (s *_nestedSortValue) Nested(nested types.NestedSortValueVariant) *_nestedSortValue {

	s.v.Nested = nested.NestedSortValueCaster()

	return s
}

func (s *_nestedSortValue) Path(field string) *_nestedSortValue {

	s.v.Path = field

	return s
}

func (s *_nestedSortValue) NestedSortValueCaster() *types.NestedSortValue {
	return s.v
}
