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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _histogramAggregation struct {
	v *types.HistogramAggregation
}

// A multi-bucket values source based aggregation that can be applied on numeric
// values or numeric range values extracted from the documents.
// It dynamically builds fixed size (interval) buckets over the values.
func NewHistogramAggregation() *_histogramAggregation {

	return &_histogramAggregation{v: types.NewHistogramAggregation()}

}

func (s *_histogramAggregation) ExtendedBounds(extendedbounds types.ExtendedBoundsdoubleVariant) *_histogramAggregation {

	s.v.ExtendedBounds = extendedbounds.ExtendedBoundsdoubleCaster()

	return s
}

func (s *_histogramAggregation) Field(field string) *_histogramAggregation {

	s.v.Field = &field

	return s
}

func (s *_histogramAggregation) Format(format string) *_histogramAggregation {

	s.v.Format = &format

	return s
}

func (s *_histogramAggregation) HardBounds(hardbounds types.ExtendedBoundsdoubleVariant) *_histogramAggregation {

	s.v.HardBounds = hardbounds.ExtendedBoundsdoubleCaster()

	return s
}

func (s *_histogramAggregation) Interval(interval types.Float64) *_histogramAggregation {

	s.v.Interval = &interval

	return s
}

func (s *_histogramAggregation) Keyed(keyed bool) *_histogramAggregation {

	s.v.Keyed = &keyed

	return s
}

func (s *_histogramAggregation) MinDocCount(mindoccount int) *_histogramAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

func (s *_histogramAggregation) Missing(missing types.Float64) *_histogramAggregation {

	s.v.Missing = &missing

	return s
}

func (s *_histogramAggregation) Offset(offset types.Float64) *_histogramAggregation {

	s.v.Offset = &offset

	return s
}

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
