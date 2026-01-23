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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _sortCombinations struct {
	v types.SortCombinations
}

func NewSortCombinations() *_sortCombinations {
	return &_sortCombinations{v: nil}
}

func (u *_sortCombinations) Field(field string) *_sortCombinations {

	u.v = &field

	return u
}

func (u *_sortCombinations) SortOptions(sortoptions types.SortOptionsVariant) *_sortCombinations {

	u.v = sortoptions.SortOptionsCaster()

	return u
}

// Interface implementation for SortOptions in SortCombinations union
func (u *_sortOptions) SortCombinationsCaster() *types.SortCombinations {
	t := types.SortCombinations(u.v)
	return &t
}

func (u *_sortCombinations) SortCombinationsCaster() *types.SortCombinations {
	return &u.v
}
