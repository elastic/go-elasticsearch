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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _percentilesAggregation struct {
	v *types.PercentilesAggregation
}

// A multi-value metrics aggregation that calculates one or more percentiles
// over numeric values extracted from the aggregated documents.
func NewPercentilesAggregation() *_percentilesAggregation {

	return &_percentilesAggregation{v: types.NewPercentilesAggregation()}

}

func (s *_percentilesAggregation) Field(field string) *_percentilesAggregation {

	s.v.Field = &field

	return s
}

func (s *_percentilesAggregation) Format(format string) *_percentilesAggregation {

	s.v.Format = &format

	return s
}

func (s *_percentilesAggregation) Hdr(hdr types.HdrMethodVariant) *_percentilesAggregation {

	s.v.Hdr = hdr.HdrMethodCaster()

	return s
}

func (s *_percentilesAggregation) Keyed(keyed bool) *_percentilesAggregation {

	s.v.Keyed = &keyed

	return s
}

func (s *_percentilesAggregation) Missing(missing types.MissingVariant) *_percentilesAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

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
