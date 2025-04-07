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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// This is provide all the types that are part of the union.
type _aggregationRange struct {
	v types.AggregationRange
}

func NewAggregationRange() *_aggregationRange {
	return &_aggregationRange{v: nil}
}

func (u *_aggregationRange) UntypedAggregationRange(untypedaggregationrange types.UntypedAggregationRangeVariant) *_aggregationRange {

	u.v = &untypedaggregationrange

	return u
}

// Interface implementation for UntypedAggregationRange in AggregationRange union
func (u *_untypedAggregationRange) AggregationRangeCaster() *types.AggregationRange {
	t := types.AggregationRange(u.v)
	return &t
}

func (u *_aggregationRange) DateAggregationRange(dateaggregationrange types.DateAggregationRangeVariant) *_aggregationRange {

	u.v = &dateaggregationrange

	return u
}

// Interface implementation for DateAggregationRange in AggregationRange union
func (u *_dateAggregationRange) AggregationRangeCaster() *types.AggregationRange {
	t := types.AggregationRange(u.v)
	return &t
}

func (u *_aggregationRange) NumberAggregationRange(numberaggregationrange types.NumberAggregationRangeVariant) *_aggregationRange {

	u.v = &numberaggregationrange

	return u
}

// Interface implementation for NumberAggregationRange in AggregationRange union
func (u *_numberAggregationRange) AggregationRangeCaster() *types.AggregationRange {
	t := types.AggregationRange(u.v)
	return &t
}

func (u *_aggregationRange) TermAggregationRange(termaggregationrange types.TermAggregationRangeVariant) *_aggregationRange {

	u.v = &termaggregationrange

	return u
}

// Interface implementation for TermAggregationRange in AggregationRange union
func (u *_termAggregationRange) AggregationRangeCaster() *types.AggregationRange {
	t := types.AggregationRange(u.v)
	return &t
}

func (u *_aggregationRange) AggregationRangeCaster() *types.AggregationRange {
	return &u.v
}
