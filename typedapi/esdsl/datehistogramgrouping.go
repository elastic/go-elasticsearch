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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _dateHistogramGrouping struct {
	v *types.DateHistogramGrouping
}

func NewDateHistogramGrouping() *_dateHistogramGrouping {

	return &_dateHistogramGrouping{v: types.NewDateHistogramGrouping()}

}

// The interval of time buckets to be generated when rolling up.
func (s *_dateHistogramGrouping) CalendarInterval(duration types.DurationVariant) *_dateHistogramGrouping {

	s.v.CalendarInterval = *duration.DurationCaster()

	return s
}

// How long to wait before rolling up new documents.
// By default, the indexer attempts to roll up all data that is available.
// However, it is not uncommon for data to arrive out of order.
// The indexer is unable to deal with data that arrives after a time-span has
// been rolled up.
// You need to specify a delay that matches the longest period of time you
// expect out-of-order data to arrive.
func (s *_dateHistogramGrouping) Delay(duration types.DurationVariant) *_dateHistogramGrouping {

	s.v.Delay = *duration.DurationCaster()

	return s
}

// The date field that is to be rolled up.
func (s *_dateHistogramGrouping) Field(field string) *_dateHistogramGrouping {

	s.v.Field = field

	return s
}

// The interval of time buckets to be generated when rolling up.
func (s *_dateHistogramGrouping) FixedInterval(duration types.DurationVariant) *_dateHistogramGrouping {

	s.v.FixedInterval = *duration.DurationCaster()

	return s
}

func (s *_dateHistogramGrouping) Format(format string) *_dateHistogramGrouping {

	s.v.Format = &format

	return s
}

func (s *_dateHistogramGrouping) Interval(duration types.DurationVariant) *_dateHistogramGrouping {

	s.v.Interval = *duration.DurationCaster()

	return s
}

// Defines what `time_zone` the rollup documents are stored as.
// Unlike raw data, which can shift timezones on the fly, rolled documents have
// to be stored with a specific timezone.
// By default, rollup documents are stored in `UTC`.
func (s *_dateHistogramGrouping) TimeZone(timezone string) *_dateHistogramGrouping {

	s.v.TimeZone = &timezone

	return s
}

func (s *_dateHistogramGrouping) DateHistogramGroupingCaster() *types.DateHistogramGrouping {
	return s.v
}
