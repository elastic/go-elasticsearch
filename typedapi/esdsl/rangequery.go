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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _rangeQuery struct {
	v types.RangeQuery
}

func NewRangeQuery() *_rangeQuery {
	return &_rangeQuery{v: nil}
}

func (u *_rangeQuery) UntypedRangeQuery(untypedrangequery types.UntypedRangeQueryVariant) *_rangeQuery {

	u.v = untypedrangequery.UntypedRangeQueryCaster()

	return u
}

// Interface implementation for UntypedRangeQuery in RangeQuery union
func (u *_untypedRangeQuery) RangeQueryCaster() *types.RangeQuery {
	t := types.RangeQuery(u.v)
	return &t
}

func (u *_rangeQuery) DateRangeQuery(daterangequery types.DateRangeQueryVariant) *_rangeQuery {

	u.v = daterangequery.DateRangeQueryCaster()

	return u
}

// Interface implementation for DateRangeQuery in RangeQuery union
func (u *_dateRangeQuery) RangeQueryCaster() *types.RangeQuery {
	t := types.RangeQuery(u.v)
	return &t
}

func (u *_rangeQuery) NumberRangeQuery(numberrangequery types.NumberRangeQueryVariant) *_rangeQuery {

	u.v = numberrangequery.NumberRangeQueryCaster()

	return u
}

// Interface implementation for NumberRangeQuery in RangeQuery union
func (u *_numberRangeQuery) RangeQueryCaster() *types.RangeQuery {
	t := types.RangeQuery(u.v)
	return &t
}

func (u *_rangeQuery) TermRangeQuery(termrangequery types.TermRangeQueryVariant) *_rangeQuery {

	u.v = termrangequery.TermRangeQueryCaster()

	return u
}

// Interface implementation for TermRangeQuery in RangeQuery union
func (u *_termRangeQuery) RangeQueryCaster() *types.RangeQuery {
	t := types.RangeQuery(u.v)
	return &t
}

func (u *_rangeQuery) RangeQueryCaster() *types.RangeQuery {
	return &u.v
}
