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

type _percentilesAggregation struct {
	v *types.PercentilesAggregation
}

// A multi-value metrics aggregation that calculates one or more percentiles
// over numeric values extracted from the aggregated documents.
func NewPercentilesAggregation() *_percentilesAggregation {

	return &_percentilesAggregation{v: types.NewPercentilesAggregation()}

}

// The field on which to run the aggregation.
func (s *_percentilesAggregation) Field(field string) *_percentilesAggregation {

	s.v.Field = &field

	return s
}

func (s *_percentilesAggregation) Format(format string) *_percentilesAggregation {

	s.v.Format = &format

	return s
}

// Uses the alternative High Dynamic Range Histogram algorithm to calculate
// percentiles.
func (s *_percentilesAggregation) Hdr(hdr types.HdrMethodVariant) *_percentilesAggregation {

	s.v.Hdr = hdr.HdrMethodCaster()

	return s
}

// By default, the aggregation associates a unique string key with each bucket
// and returns the ranges as a hash rather than an array.
// Set to `false` to disable this behavior.
func (s *_percentilesAggregation) Keyed(keyed bool) *_percentilesAggregation {

	s.v.Keyed = &keyed

	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_percentilesAggregation) Missing(missing types.MissingVariant) *_percentilesAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

// The percentiles to calculate.
func (s *_percentilesAggregation) Percents(percents ...types.Float64) *_percentilesAggregation {

	for _, v := range percents {

		s.v.Percents = append(s.v.Percents, v)

	}
	return s
}

func (s *_percentilesAggregation) Script(script types.ScriptVariant) *_percentilesAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

// Sets parameters for the default TDigest algorithm used to calculate
// percentiles.
func (s *_percentilesAggregation) Tdigest(tdigest types.TDigestVariant) *_percentilesAggregation {

	s.v.Tdigest = tdigest.TDigestCaster()

	return s
}

func (s *_percentilesAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Percentiles = s.v

	return container
}

func (s *_percentilesAggregation) PercentilesAggregationCaster() *types.PercentilesAggregation {
	return s.v
}
