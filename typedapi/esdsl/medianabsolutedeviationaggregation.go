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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _medianAbsoluteDeviationAggregation struct {
	v *types.MedianAbsoluteDeviationAggregation
}

// A single-value aggregation that approximates the median absolute deviation of
// its search results.
func NewMedianAbsoluteDeviationAggregation() *_medianAbsoluteDeviationAggregation {

	return &_medianAbsoluteDeviationAggregation{v: types.NewMedianAbsoluteDeviationAggregation()}

}

func (s *_medianAbsoluteDeviationAggregation) Compression(compression types.Float64) *_medianAbsoluteDeviationAggregation {

	s.v.Compression = &compression

	return s
}

func (s *_medianAbsoluteDeviationAggregation) Field(field string) *_medianAbsoluteDeviationAggregation {

	s.v.Field = &field

	return s
}

func (s *_medianAbsoluteDeviationAggregation) Format(format string) *_medianAbsoluteDeviationAggregation {

	s.v.Format = &format

	return s
}

func (s *_medianAbsoluteDeviationAggregation) Missing(missing types.MissingVariant) *_medianAbsoluteDeviationAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_medianAbsoluteDeviationAggregation) Script(script types.ScriptVariant) *_medianAbsoluteDeviationAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_medianAbsoluteDeviationAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.MedianAbsoluteDeviation = s.v

	return container
}

func (s *_medianAbsoluteDeviationAggregation) MedianAbsoluteDeviationAggregationCaster() *types.MedianAbsoluteDeviationAggregation {
	return s.v
}
