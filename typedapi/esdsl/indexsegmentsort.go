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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortmissing"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortorder"
)

type _indexSegmentSort struct {
	v *types.IndexSegmentSort
}

func NewIndexSegmentSort() *_indexSegmentSort {

	return &_indexSegmentSort{v: types.NewIndexSegmentSort()}

}

func (s *_indexSegmentSort) Field(fields ...string) *_indexSegmentSort {

	s.v.Field = fields

	return s
}

func (s *_indexSegmentSort) Missing(missings ...segmentsortmissing.SegmentSortMissing) *_indexSegmentSort {

	s.v.Missing = make([]segmentsortmissing.SegmentSortMissing, len(missings))
	s.v.Missing = missings

	return s
}

func (s *_indexSegmentSort) Mode(modes ...segmentsortmode.SegmentSortMode) *_indexSegmentSort {

	s.v.Mode = make([]segmentsortmode.SegmentSortMode, len(modes))
	s.v.Mode = modes

	return s
}

func (s *_indexSegmentSort) Order(orders ...segmentsortorder.SegmentSortOrder) *_indexSegmentSort {

	s.v.Order = make([]segmentsortorder.SegmentSortOrder, len(orders))
	s.v.Order = orders

	return s
}

func (s *_indexSegmentSort) IndexSegmentSortCaster() *types.IndexSegmentSort {
	return s.v
}
