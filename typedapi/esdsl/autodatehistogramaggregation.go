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
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/minimuminterval"
)

type _autoDateHistogramAggregation struct {
	v *types.AutoDateHistogramAggregation
}

// A multi-bucket aggregation similar to the date histogram, except instead of
// providing an interval to use as the width of each bucket, a target number of
// buckets is provided.
func NewAutoDateHistogramAggregation() *_autoDateHistogramAggregation {

	return &_autoDateHistogramAggregation{v: types.NewAutoDateHistogramAggregation()}

}

// The target number of buckets.
func (s *_autoDateHistogramAggregation) Buckets(buckets int) *_autoDateHistogramAggregation {

	s.v.Buckets = &buckets

	return s
}

// The field on which to run the aggregation.
func (s *_autoDateHistogramAggregation) Field(field string) *_autoDateHistogramAggregation {

	s.v.Field = &field

	return s
}

// The date format used to format `key_as_string` in the response.
// If no `format` is specified, the first date format specified in the field
// mapping is used.
func (s *_autoDateHistogramAggregation) Format(format string) *_autoDateHistogramAggregation {

	s.v.Format = &format

	return s
}

// The minimum rounding interval.
// This can make the collection process more efficient, as the aggregation will
// not attempt to round at any interval lower than `minimum_interval`.
func (s *_autoDateHistogramAggregation) MinimumInterval(minimuminterval minimuminterval.MinimumInterval) *_autoDateHistogramAggregation {

	s.v.MinimumInterval = &minimuminterval
	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_autoDateHistogramAggregation) Missing(datetime types.DateTimeVariant) *_autoDateHistogramAggregation {

	s.v.Missing = *datetime.DateTimeCaster()

	return s
}

// Time zone specified as a ISO 8601 UTC offset.
func (s *_autoDateHistogramAggregation) Offset(offset string) *_autoDateHistogramAggregation {

	s.v.Offset = &offset

	return s
}

func (s *_autoDateHistogramAggregation) Params(params map[string]json.RawMessage) *_autoDateHistogramAggregation {

	s.v.Params = params
	return s
}

func (s *_autoDateHistogramAggregation) AddParam(key string, value json.RawMessage) *_autoDateHistogramAggregation {

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

func (s *_autoDateHistogramAggregation) Script(script types.ScriptVariant) *_autoDateHistogramAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

// Time zone ID.
func (s *_autoDateHistogramAggregation) TimeZone(timezone string) *_autoDateHistogramAggregation {

	s.v.TimeZone = &timezone

	return s
}

func (s *_autoDateHistogramAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.AutoDateHistogram = s.v

	return container
}

func (s *_autoDateHistogramAggregation) AutoDateHistogramAggregationCaster() *types.AutoDateHistogramAggregation {
	return s.v
}
