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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/minimuminterval"
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

func (s *_autoDateHistogramAggregation) Buckets(buckets int) *_autoDateHistogramAggregation {

	s.v.Buckets = &buckets

	return s
}

func (s *_autoDateHistogramAggregation) Field(field string) *_autoDateHistogramAggregation {

	s.v.Field = &field

	return s
}

func (s *_autoDateHistogramAggregation) Format(format string) *_autoDateHistogramAggregation {

	s.v.Format = &format

	return s
}

func (s *_autoDateHistogramAggregation) MinimumInterval(minimuminterval minimuminterval.MinimumInterval) *_autoDateHistogramAggregation {

	s.v.MinimumInterval = &minimuminterval
	return s
}

func (s *_autoDateHistogramAggregation) Missing(datetime types.DateTimeVariant) *_autoDateHistogramAggregation {

	s.v.Missing = *datetime.DateTimeCaster()

	return s
}

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
