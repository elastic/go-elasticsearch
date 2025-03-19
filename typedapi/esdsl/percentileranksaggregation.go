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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _percentileRanksAggregation struct {
	v *types.PercentileRanksAggregation
}

// A multi-value metrics aggregation that calculates one or more percentile
// ranks over numeric values extracted from the aggregated documents.
func NewPercentileRanksAggregation() *_percentileRanksAggregation {

	return &_percentileRanksAggregation{v: types.NewPercentileRanksAggregation()}

}

// The field on which to run the aggregation.
func (s *_percentileRanksAggregation) Field(field string) *_percentileRanksAggregation {

	s.v.Field = &field

	return s
}

func (s *_percentileRanksAggregation) Format(format string) *_percentileRanksAggregation {

	s.v.Format = &format

	return s
}

// Uses the alternative High Dynamic Range Histogram algorithm to calculate
// percentile ranks.
func (s *_percentileRanksAggregation) Hdr(hdr types.HdrMethodVariant) *_percentileRanksAggregation {

	s.v.Hdr = hdr.HdrMethodCaster()

	return s
}

// By default, the aggregation associates a unique string key with each bucket
// and returns the ranges as a hash rather than an array.
// Set to `false` to disable this behavior.
func (s *_percentileRanksAggregation) Keyed(keyed bool) *_percentileRanksAggregation {

	s.v.Keyed = &keyed

	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_percentileRanksAggregation) Missing(missing types.MissingVariant) *_percentileRanksAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_percentileRanksAggregation) Script(script types.ScriptVariant) *_percentileRanksAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

// Sets parameters for the default TDigest algorithm used to calculate
// percentile ranks.
func (s *_percentileRanksAggregation) Tdigest(tdigest types.TDigestVariant) *_percentileRanksAggregation {

	s.v.Tdigest = tdigest.TDigestCaster()

	return s
}

// An array of values for which to calculate the percentile ranks.
func (s *_percentileRanksAggregation) Values(values []types.Float64) *_percentileRanksAggregation {

	s.v.Values = &values

	return s
}

func (s *_percentileRanksAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.PercentileRanks = s.v

	return container
}

func (s *_percentileRanksAggregation) PercentileRanksAggregationCaster() *types.PercentileRanksAggregation {
	return s.v
}
