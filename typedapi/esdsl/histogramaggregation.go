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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _histogramAggregation struct {
	v *types.HistogramAggregation
}

// A multi-bucket values source based aggregation that can be applied on numeric
// values or numeric range values extracted from the documents.
// It dynamically builds fixed size (interval) buckets over the values.
func NewHistogramAggregation() *_histogramAggregation {

	return &_histogramAggregation{v: types.NewHistogramAggregation()}

}

// Enables extending the bounds of the histogram beyond the data itself.
func (s *_histogramAggregation) ExtendedBounds(extendedbounds types.ExtendedBoundsdoubleVariant) *_histogramAggregation {

	s.v.ExtendedBounds = extendedbounds.ExtendedBoundsdoubleCaster()

	return s
}

// The name of the field to aggregate on.
func (s *_histogramAggregation) Field(field string) *_histogramAggregation {

	s.v.Field = &field

	return s
}

func (s *_histogramAggregation) Format(format string) *_histogramAggregation {

	s.v.Format = &format

	return s
}

// Limits the range of buckets in the histogram.
// It is particularly useful in the case of open data ranges that can result in
// a very large number of buckets.
func (s *_histogramAggregation) HardBounds(hardbounds types.ExtendedBoundsdoubleVariant) *_histogramAggregation {

	s.v.HardBounds = hardbounds.ExtendedBoundsdoubleCaster()

	return s
}

// The interval for the buckets.
// Must be a positive decimal.
func (s *_histogramAggregation) Interval(interval types.Float64) *_histogramAggregation {

	s.v.Interval = &interval

	return s
}

// If `true`, returns buckets as a hash instead of an array, keyed by the bucket
// keys.
func (s *_histogramAggregation) Keyed(keyed bool) *_histogramAggregation {

	s.v.Keyed = &keyed

	return s
}

// Only returns buckets that have `min_doc_count` number of documents.
// By default, the response will fill gaps in the histogram with empty buckets.
func (s *_histogramAggregation) MinDocCount(mindoccount int) *_histogramAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_histogramAggregation) Missing(missing types.Float64) *_histogramAggregation {

	s.v.Missing = &missing

	return s
}

// By default, the bucket keys start with 0 and then continue in even spaced
// steps of `interval`.
// The bucket boundaries can be shifted by using the `offset` option.
func (s *_histogramAggregation) Offset(offset types.Float64) *_histogramAggregation {

	s.v.Offset = &offset

	return s
}

// The sort order of the returned buckets.
// By default, the returned buckets are sorted by their key ascending.
func (s *_histogramAggregation) Order(aggregateorder types.AggregateOrderVariant) *_histogramAggregation {

	s.v.Order = *aggregateorder.AggregateOrderCaster()

	return s
}

func (s *_histogramAggregation) Script(script types.ScriptVariant) *_histogramAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_histogramAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Histogram = s.v

	return container
}

func (s *_histogramAggregation) PivotGroupByContainerCaster() *types.PivotGroupByContainer {
	container := types.NewPivotGroupByContainer()

	container.Histogram = s.v

	return container
}

func (s *_histogramAggregation) HistogramAggregationCaster() *types.HistogramAggregation {
	return s.v
}
