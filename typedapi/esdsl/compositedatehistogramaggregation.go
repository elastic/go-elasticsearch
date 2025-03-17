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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/valuetype"
)

type _compositeDateHistogramAggregation struct {
	v *types.CompositeDateHistogramAggregation
}

func NewCompositeDateHistogramAggregation() *_compositeDateHistogramAggregation {

	return &_compositeDateHistogramAggregation{v: types.NewCompositeDateHistogramAggregation()}

}

// Either `calendar_interval` or `fixed_interval` must be present
func (s *_compositeDateHistogramAggregation) CalendarInterval(durationlarge string) *_compositeDateHistogramAggregation {

	s.v.CalendarInterval = &durationlarge

	return s
}

// Either `field` or `script` must be present
func (s *_compositeDateHistogramAggregation) Field(field string) *_compositeDateHistogramAggregation {

	s.v.Field = &field

	return s
}

// Either `calendar_interval` or `fixed_interval` must be present
func (s *_compositeDateHistogramAggregation) FixedInterval(durationlarge string) *_compositeDateHistogramAggregation {

	s.v.FixedInterval = &durationlarge

	return s
}

func (s *_compositeDateHistogramAggregation) Format(format string) *_compositeDateHistogramAggregation {

	s.v.Format = &format

	return s
}

func (s *_compositeDateHistogramAggregation) MissingBucket(missingbucket bool) *_compositeDateHistogramAggregation {

	s.v.MissingBucket = &missingbucket

	return s
}

func (s *_compositeDateHistogramAggregation) MissingOrder(missingorder missingorder.MissingOrder) *_compositeDateHistogramAggregation {

	s.v.MissingOrder = &missingorder
	return s
}

func (s *_compositeDateHistogramAggregation) Offset(duration types.DurationVariant) *_compositeDateHistogramAggregation {

	s.v.Offset = *duration.DurationCaster()

	return s
}

func (s *_compositeDateHistogramAggregation) Order(order sortorder.SortOrder) *_compositeDateHistogramAggregation {

	s.v.Order = &order
	return s
}

// Either `field` or `script` must be present
func (s *_compositeDateHistogramAggregation) Script(script types.ScriptVariant) *_compositeDateHistogramAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_compositeDateHistogramAggregation) TimeZone(timezone string) *_compositeDateHistogramAggregation {

	s.v.TimeZone = &timezone

	return s
}

func (s *_compositeDateHistogramAggregation) ValueType(valuetype valuetype.ValueType) *_compositeDateHistogramAggregation {

	s.v.ValueType = &valuetype
	return s
}

func (s *_compositeDateHistogramAggregation) CompositeDateHistogramAggregationCaster() *types.CompositeDateHistogramAggregation {
	return s.v
}
