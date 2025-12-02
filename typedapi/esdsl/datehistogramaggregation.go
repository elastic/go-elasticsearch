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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/calendarinterval"
)

type _dateHistogramAggregation struct {
	v *types.DateHistogramAggregation
}

// A multi-bucket values source based aggregation that can be applied on date
// values or date range values extracted from the documents.
// It dynamically builds fixed size (interval) buckets over the values.
func NewDateHistogramAggregation() *_dateHistogramAggregation {

	return &_dateHistogramAggregation{v: types.NewDateHistogramAggregation()}

}

func (s *_dateHistogramAggregation) CalendarInterval(calendarinterval calendarinterval.CalendarInterval) *_dateHistogramAggregation {

	s.v.CalendarInterval = &calendarinterval
	return s
}

func (s *_dateHistogramAggregation) ExtendedBounds(extendedbounds types.ExtendedBoundsFieldDateMathVariant) *_dateHistogramAggregation {

	s.v.ExtendedBounds = extendedbounds.ExtendedBoundsFieldDateMathCaster()

	return s
}

func (s *_dateHistogramAggregation) Field(field string) *_dateHistogramAggregation {

	s.v.Field = &field

	return s
}

func (s *_dateHistogramAggregation) FixedInterval(duration types.DurationVariant) *_dateHistogramAggregation {

	s.v.FixedInterval = *duration.DurationCaster()

	return s
}

func (s *_dateHistogramAggregation) Format(format string) *_dateHistogramAggregation {

	s.v.Format = &format

	return s
}

func (s *_dateHistogramAggregation) HardBounds(hardbounds types.ExtendedBoundsFieldDateMathVariant) *_dateHistogramAggregation {

	s.v.HardBounds = hardbounds.ExtendedBoundsFieldDateMathCaster()

	return s
}

func (s *_dateHistogramAggregation) Interval(duration types.DurationVariant) *_dateHistogramAggregation {

	s.v.Interval = *duration.DurationCaster()

	return s
}

func (s *_dateHistogramAggregation) Keyed(keyed bool) *_dateHistogramAggregation {

	s.v.Keyed = &keyed

	return s
}

func (s *_dateHistogramAggregation) MinDocCount(mindoccount int) *_dateHistogramAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

func (s *_dateHistogramAggregation) Missing(datetime types.DateTimeVariant) *_dateHistogramAggregation {

	s.v.Missing = *datetime.DateTimeCaster()

	return s
}

func (s *_dateHistogramAggregation) Offset(duration types.DurationVariant) *_dateHistogramAggregation {

	s.v.Offset = *duration.DurationCaster()

	return s
}

func (s *_dateHistogramAggregation) Order(aggregateorder types.AggregateOrderVariant) *_dateHistogramAggregation {

	s.v.Order = *aggregateorder.AggregateOrderCaster()

	return s
}

func (s *_dateHistogramAggregation) Params(params map[string]json.RawMessage) *_dateHistogramAggregation {

	s.v.Params = params
	return s
}

func (s *_dateHistogramAggregation) AddParam(key string, value json.RawMessage) *_dateHistogramAggregation {

	var tmp map[string]json.RawMessage
	if s.v.Params == nil {
		s.v.Params = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Params
	}

	tmp[key] = value

	s.v.Params = tmp
	return s
}

func (s *_dateHistogramAggregation) Script(script types.ScriptVariant) *_dateHistogramAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_dateHistogramAggregation) TimeZone(timezone string) *_dateHistogramAggregation {

	s.v.TimeZone = &timezone

	return s
}

func (s *_dateHistogramAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.DateHistogram = s.v

	return container
}

func (s *_dateHistogramAggregation) PivotGroupByContainerCaster() *types.PivotGroupByContainer {
	container := types.NewPivotGroupByContainer()

	container.DateHistogram = s.v

	return container
}

func (s *_dateHistogramAggregation) DateHistogramAggregationCaster() *types.DateHistogramAggregation {
	return s.v
}
