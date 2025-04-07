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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldsortnumerictype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

type _fieldSort struct {
	v *types.FieldSort
}

func NewFieldSort(order sortorder.SortOrder) *_fieldSort {

	tmp := &_fieldSort{v: types.NewFieldSort()}

	tmp.Order(order)

	return tmp

}

func (s *_fieldSort) Format(format string) *_fieldSort {

	s.v.Format = &format

	return s
}

func (s *_fieldSort) Missing(missing types.MissingVariant) *_fieldSort {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_fieldSort) Mode(mode sortmode.SortMode) *_fieldSort {

	s.v.Mode = &mode
	return s
}

func (s *_fieldSort) Nested(nested types.NestedSortValueVariant) *_fieldSort {

	s.v.Nested = nested.NestedSortValueCaster()

	return s
}

func (s *_fieldSort) NumericType(numerictype fieldsortnumerictype.FieldSortNumericType) *_fieldSort {

	s.v.NumericType = &numerictype
	return s
}

func (s *_fieldSort) Order(order sortorder.SortOrder) *_fieldSort {

	s.v.Order = &order
	return s
}

func (s *_fieldSort) UnmappedType(unmappedtype fieldtype.FieldType) *_fieldSort {

	s.v.UnmappedType = &unmappedtype
	return s
}

func (s *_fieldSort) FieldSortCaster() *types.FieldSort {
	return s.v
}
