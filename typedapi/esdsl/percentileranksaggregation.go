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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _percentileRanksAggregation struct {
	v *types.PercentileRanksAggregation
}

// A multi-value metrics aggregation that calculates one or more percentile
// ranks over numeric values extracted from the aggregated documents.
func NewPercentileRanksAggregation() *_percentileRanksAggregation {

	return &_percentileRanksAggregation{v: types.NewPercentileRanksAggregation()}

}

func (s *_percentileRanksAggregation) Field(field string) *_percentileRanksAggregation {

	s.v.Field = &field

	return s
}

func (s *_percentileRanksAggregation) Format(format string) *_percentileRanksAggregation {

	s.v.Format = &format

	return s
}

func (s *_percentileRanksAggregation) Hdr(hdr types.HdrMethodVariant) *_percentileRanksAggregation {

	s.v.Hdr = hdr.HdrMethodCaster()

	return s
}

func (s *_percentileRanksAggregation) Keyed(keyed bool) *_percentileRanksAggregation {

	s.v.Keyed = &keyed

	return s
}

func (s *_percentileRanksAggregation) Missing(missing types.MissingVariant) *_percentileRanksAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_percentileRanksAggregation) Script(script types.ScriptVariant) *_percentileRanksAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_percentileRanksAggregation) Tdigest(tdigest types.TDigestVariant) *_percentileRanksAggregation {

	s.v.Tdigest = tdigest.TDigestCaster()

	return s
}

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
