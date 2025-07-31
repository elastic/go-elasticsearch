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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _serialDifferencingAggregation struct {
	v *types.SerialDifferencingAggregation
}

// An aggregation that subtracts values in a time series from themselves at
// different time lags or periods.
func NewSerialDifferencingAggregation() *_serialDifferencingAggregation {

	return &_serialDifferencingAggregation{v: types.NewSerialDifferencingAggregation()}

}

func (s *_serialDifferencingAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_serialDifferencingAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_serialDifferencingAggregation) Format(format string) *_serialDifferencingAggregation {

	s.v.Format = &format

	return s
}

func (s *_serialDifferencingAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_serialDifferencingAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_serialDifferencingAggregation) Lag(lag int) *_serialDifferencingAggregation {

	s.v.Lag = &lag

	return s
}

func (s *_serialDifferencingAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.SerialDiff = s.v

	return container
}

func (s *_serialDifferencingAggregation) SerialDifferencingAggregationCaster() *types.SerialDifferencingAggregation {
	return s.v
}
