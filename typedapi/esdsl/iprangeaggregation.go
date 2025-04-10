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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ipRangeAggregation struct {
	v *types.IpRangeAggregation
}

// A multi-bucket value source based aggregation that enables the user to define
// a set of IP ranges - each representing a bucket.
func NewIpRangeAggregation() *_ipRangeAggregation {

	return &_ipRangeAggregation{v: types.NewIpRangeAggregation()}

}

func (s *_ipRangeAggregation) Field(field string) *_ipRangeAggregation {

	s.v.Field = &field

	return s
}

func (s *_ipRangeAggregation) Ranges(ranges ...types.IpRangeAggregationRangeVariant) *_ipRangeAggregation {

	for _, v := range ranges {

		s.v.Ranges = append(s.v.Ranges, *v.IpRangeAggregationRangeCaster())

	}
	return s
}

func (s *_ipRangeAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.IpRange = s.v

	return container
}

func (s *_ipRangeAggregation) IpRangeAggregationCaster() *types.IpRangeAggregation {
	return s.v
}
