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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/calendarinterval"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ratemode"
)

type _rateAggregation struct {
	v *types.RateAggregation
}

// Calculates a rate of documents or a field in each bucket.
// Can only be used inside a `date_histogram` or `composite` aggregation.
func NewRateAggregation() *_rateAggregation {

	return &_rateAggregation{v: types.NewRateAggregation()}

}

func (s *_rateAggregation) Mode(mode ratemode.RateMode) *_rateAggregation {

	s.v.Mode = &mode
	return s
}

func (s *_rateAggregation) Unit(unit calendarinterval.CalendarInterval) *_rateAggregation {

	s.v.Unit = &unit
	return s
}

func (s *_rateAggregation) Field(field string) *_rateAggregation {

	s.v.Field = &field

	return s
}

func (s *_rateAggregation) Format(format string) *_rateAggregation {

	s.v.Format = &format

	return s
}

func (s *_rateAggregation) Missing(missing types.MissingVariant) *_rateAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_rateAggregation) Script(script types.ScriptVariant) *_rateAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_rateAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Rate = s.v

	return container
}

func (s *_rateAggregation) RateAggregationCaster() *types.RateAggregation {
	return s.v
}
