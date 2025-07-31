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

type _fielddataFrequencyFilter struct {
	v *types.FielddataFrequencyFilter
}

func NewFielddataFrequencyFilter(max types.Float64, min types.Float64, minsegmentsize int) *_fielddataFrequencyFilter {

	tmp := &_fielddataFrequencyFilter{v: types.NewFielddataFrequencyFilter()}

	tmp.Max(max)

	tmp.Min(min)

	tmp.MinSegmentSize(minsegmentsize)

	return tmp

}

func (s *_fielddataFrequencyFilter) Max(max types.Float64) *_fielddataFrequencyFilter {

	s.v.Max = max

	return s
}

func (s *_fielddataFrequencyFilter) Min(min types.Float64) *_fielddataFrequencyFilter {

	s.v.Min = min

	return s
}

func (s *_fielddataFrequencyFilter) MinSegmentSize(minsegmentsize int) *_fielddataFrequencyFilter {

	s.v.MinSegmentSize = minsegmentsize

	return s
}

func (s *_fielddataFrequencyFilter) FielddataFrequencyFilterCaster() *types.FielddataFrequencyFilter {
	return s.v
}
