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

type _histogramGrouping struct {
	v *types.HistogramGrouping
}

func NewHistogramGrouping(interval int64) *_histogramGrouping {

	tmp := &_histogramGrouping{v: types.NewHistogramGrouping()}

	tmp.Interval(interval)

	return tmp

}

// The set of fields that you wish to build histograms for.
// All fields specified must be some kind of numeric.
// Order does not matter.
func (s *_histogramGrouping) Fields(fields ...string) *_histogramGrouping {

	s.v.Fields = fields

	return s
}

// The interval of histogram buckets to be generated when rolling up.
// For example, a value of `5` creates buckets that are five units wide (`0-5`,
// `5-10`, etc).
// Note that only one interval can be specified in the histogram group, meaning
// that all fields being grouped via the histogram must share the same interval.
func (s *_histogramGrouping) Interval(interval int64) *_histogramGrouping {

	s.v.Interval = interval

	return s
}

func (s *_histogramGrouping) HistogramGroupingCaster() *types.HistogramGrouping {
	return s.v
}
