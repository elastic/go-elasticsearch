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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _sortOptions struct {
	v *types.SortOptions
}

func NewSortOptions() *_sortOptions {
	return &_sortOptions{v: types.NewSortOptions()}
}

func (s *_sortOptions) Doc_(doc_ types.ScoreSortVariant) *_sortOptions {

	s.v.Doc_ = doc_.ScoreSortCaster()

	return s
}

func (s *_sortOptions) GeoDistance_(geodistance_ types.GeoDistanceSortVariant) *_sortOptions {

	s.v.GeoDistance_ = geodistance_.GeoDistanceSortCaster()

	return s
}

func (s *_sortOptions) Score_(score_ types.ScoreSortVariant) *_sortOptions {

	s.v.Score_ = score_.ScoreSortCaster()

	return s
}

func (s *_sortOptions) Script_(script_ types.ScriptSortVariant) *_sortOptions {

	s.v.Script_ = script_.ScriptSortCaster()

	return s
}

func (s *_sortOptions) SortOptions(sortoptions map[string]types.FieldSort) *_sortOptions {

	s.v.SortOptions = sortoptions
	return s
}

func (s *_sortOptions) AddSortOption(key string, value types.FieldSortVariant) *_sortOptions {

	var tmp map[string]types.FieldSort
	if s.v.SortOptions == nil {
		s.v.SortOptions = make(map[string]types.FieldSort)
	} else {
		tmp = s.v.SortOptions
	}

	tmp[key] = *value.FieldSortCaster()

	s.v.SortOptions = tmp
	return s
}

func (s *_sortOptions) SortOptionsCaster() *types.SortOptions {
	return s.v
}
